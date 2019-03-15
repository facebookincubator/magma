"""
Copyright (c) 2016-present, Facebook, Inc.
All rights reserved.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree. An additional grant
of patent rights can be found in the PATENTS file in the same directory.
"""
import grpc

import os
import asyncio
import logging
from typing import List

import snowflake
from google.protobuf import json_format
from google.protobuf.struct_pb2 import Struct
from orc8r.protos import magmad_pb2, magmad_pb2_grpc

from magma.common.rpc_utils import return_void
from magma.common.service import MagmaService
from magma.configuration.mconfig_managers import MconfigManager
from magma.magmad.generic_command.command_executor import \
    CommandExecutor
from magma.magmad.service_manager import ServiceManager

from .network_check import ping, traceroute


class MagmadRpcServicer(magmad_pb2_grpc.MagmadServicer):
    """
    gRPC based server for Magmad.
    """

    def __init__(self,
                 magma_service: MagmaService,
                 services: List[str],
                 service_manager: ServiceManager,
                 mconfig_manager: MconfigManager,
                 command_executor: CommandExecutor,
                 loop):
        """
        Constructor for the magmad RPC servicer

        Args:
            magma_service:
                MagmaService servicer (service303) which runs on the same
                server as this servicer
            services:
                List of services that magmad manages

            service_manager: ServiceManger instance
            mconfig_manager: MconfigManager instance
            loop: event loop
        """
        self._service_manager = service_manager
        self._services = services
        self._mconfig_manager = mconfig_manager
        self._magma_service = magma_service
        self._command_executor = command_executor
        self._loop = loop

    def add_to_server(self, server):
        """
        Add the servicer to a gRPC server
        """
        magmad_pb2_grpc.add_MagmadServicer_to_server(self, server)

    @return_void
    def StartServices(self, _, context):
        """
        Start all magma services
        """
        self._loop.create_task(self._service_manager.start_services())

    @return_void
    def StopServices(self, _, context):
        """
        Stop all magma services
        """
        self._loop.create_task(self._service_manager.stop_services())

    @return_void
    def Reboot(self, _, context):
        """
        Reboot the gateway device
        """
        async def run_reboot():
            await asyncio.sleep(1)
            os.system('reboot')

        self._loop.create_task(run_reboot())

    @return_void
    def RestartServices(self, request, context):
        """
        Restart specified magma services.
        If no services specified, restart all services.
        """
        async def run_restart():
            await asyncio.sleep(1)
            await self._service_manager.restart_services(request.services)

        logging.info("Restarting following services: %s", request.services)
        self._loop.create_task(run_restart())

    @return_void
    def SetConfigs(self, request, context):
        """
        Set the stored mconfig to a specific value. Restarts all services
        other than magmad and reloads the magmad mconfig in-place.

        If the gateway configs provided are empty, this will delete the
        managed configurations, reverting the stored mconfigs to the image
        defaults.
        """
        if request.configs_by_key is None or \
            len(request.configs_by_key) == 0:
            self._mconfig_manager.delete_stored_mconfig()
        else:
            # TODO: support streaming mconfig manager impl
            self._mconfig_manager.update_stored_mconfig(
                json_format.MessageToJson(request),
            )

        self._loop.create_task(
            self._service_manager.restart_services(self._services)
        )
        self._magma_service.reload_mconfig()

    def GetConfigs(self, _, context):
        # TODO: support streaming mconfig manager impl
        return self._mconfig_manager.load_mconfig()

    def RunNetworkTests(self, request, context):
        """
        Execute some specified network commands to check gateway network health
        """
        ping_results = self.__ping_specified_hosts(request.pings)
        traceroute_results = self.__traceroute_specified_hosts(
            request.traceroutes)

        return magmad_pb2.NetworkTestResponse(pings=ping_results,
                                              traceroutes=traceroute_results)

    def GetGatewayId(self, _, context):
        """
        Get gateway hardware ID
        """
        return magmad_pb2.GetGatewayIdResponse(gateway_id=snowflake.snowflake())

    def GenericCommand(self, request, context):
        """
        Execute generic command. This method will run the command with params
        as specified in the command executor's command table, then return
        the response of the command.
        """
        if 'generic_command_config' not in self._magma_service.config:
            context.set_code(grpc.StatusCode.NOT_FOUND)
            context.set_details('Generic command config not found')
            return magmad_pb2.GenericCommandResponse()

        params = json_format.MessageToDict(request.params)

        # Run the execute command coroutine. Return an error if it times out or
        # if an exception occurs.
        logging.info('Running generic command %s with parameters %s',
                     request.command, params)
        future = asyncio.run_coroutine_threadsafe(
            self._command_executor.execute_command(request.command, params),
            self._loop)

        timeout = self._magma_service.config['generic_command_config']\
            .get('timeout_secs', 15)

        response = magmad_pb2.GenericCommandResponse()
        try:
            result = future.result(timeout=timeout)
            logging.debug('Command was successful')
            response.response.MergeFrom(
                json_format.ParseDict(result, Struct()))
        except asyncio.TimeoutError:
            logging.error('Error running command %s! Command timed out',
                          request.command)
            future.cancel()
            context.set_code(grpc.StatusCode.DEADLINE_EXCEEDED)
            context.set_details('Command timed out')
        except Exception as e:  # pylint: disable=broad-except
            logging.error('Error running command %s! %s: %s',
                          request.command, e.__class__.__name__, e)
            context.set_code(grpc.StatusCode.UNKNOWN)
            context.set_details('{}: {}'.format(e.__class__.__name__, str(e)))

        return response

    @staticmethod
    def __ping_specified_hosts(ping_param_protos):
        def create_ping_result_proto(ping_result):
            if ping_result.error:
                return magmad_pb2.PingResult(
                    error=ping_result.error,
                    host_or_ip=ping_result.host_or_ip,
                    num_packets=ping_result.num_packets,
                )
            else:
                return magmad_pb2.PingResult(
                    host_or_ip=ping_result.host_or_ip,
                    num_packets=ping_result.num_packets,
                    packets_transmitted=ping_result.stats.packets_transmitted,
                    packets_received=ping_result.stats.packets_received,
                    avg_response_ms=ping_result.stats.rtt_avg,
                )

        pings_to_exec = [ping.PingCommandParams(
            host_or_ip=p.host_or_ip,
            num_packets=p.num_packets,
            timeout_secs=None,
        ) for p in ping_param_protos]
        ping_results = ping.ping(pings_to_exec)
        return map(create_ping_result_proto, ping_results)

    @staticmethod
    def __traceroute_specified_hosts(traceroute_param_protos):
        def create_result_proto(result):
            if result.error:
                return magmad_pb2.TracerouteResult(
                    error=result.error,
                    host_or_ip=result.host_or_ip,
                )
            else:
                return magmad_pb2.TracerouteResult(
                    host_or_ip=result.host_or_ip,
                    hops=create_hop_protos(result.stats.hops),
                )

        def create_hop_protos(hops):
            hop_protos = []
            for hop in hops:
                hop_protos.append(magmad_pb2.TracerouteHop(
                    idx=hop.idx,
                    probes=create_probe_protos(hop.probes),
                ))
            return hop_protos

        def create_probe_protos(probes):
            return [magmad_pb2.TracerouteProbe(
                hostname=probe.hostname,
                ip=probe.ip_addr,
                rtt_ms=probe.rtt_ms,
            ) for probe in probes]

        traceroutes_to_exec = [traceroute.TracerouteParams(
            host_or_ip=param.host_or_ip,
            max_hops=param.max_hops,
            bytes_per_packet=param.bytes_per_packet
        ) for param in traceroute_param_protos]
        traceroute_results = traceroute.traceroute(traceroutes_to_exec)
        return map(create_result_proto, traceroute_results)
