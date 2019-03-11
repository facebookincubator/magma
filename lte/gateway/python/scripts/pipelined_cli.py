#!/usr/bin/env python3

"""
Copyright (c) 2016-present, Facebook, Inc.
All rights reserved.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree. An additional grant
of patent rights can be found in the PATENTS file in the same directory.
"""

import argparse
import binascii
import errno
from pprint import pprint
import subprocess

from magma.common.rpc_utils import grpc_wrapper
from magma.subscriberdb.sid import SIDUtils
from magma.configuration.service_configs import load_service_config
from magma.pipelined.bridge_util import BridgeTools
from orc8r.protos.common_pb2 import Void
from lte.protos.pipelined_pb2 import (
    ActivateFlowsRequest,
    DeactivateFlowsRequest,
    RuleModResult,
)
from lte.protos.pipelined_pb2_grpc import PipelinedStub
from lte.protos.policydb_pb2 import FlowMatch, FlowDescription, PolicyRule


# --------------------------
# Metering App
# --------------------------

@grpc_wrapper
def get_subscriber_metering_flows(client, _):
    flow_table = client.GetSubscriberMeteringFlows(Void())
    print(flow_table)


def create_metering_parser(apps):
    """
    Creates the argparse subparser for the metering app
    """
    app = apps.add_parser('meter')
    subparsers = app.add_subparsers(title='subcommands', dest='cmd')

    # Add subcommands
    subcmd = subparsers.add_parser('dump_flows',
                                   help='Prints all subscriber metering flows')
    subcmd.set_defaults(func=get_subscriber_metering_flows)


# --------------------------
# Enforcement App
# --------------------------

@grpc_wrapper
def activate_flows(client, args):
    request = ActivateFlowsRequest(
        sid=SIDUtils.to_pb(args.imsi),
        rule_ids=args.rule_ids.split(','))
    response = client.ActivateFlows(request)
    _print_rule_mod_results(response.static_rule_results)


@grpc_wrapper
def deactivate_flows(client, args):
    request = DeactivateFlowsRequest(
        sid=SIDUtils.to_pb(args.imsi),
        rule_ids=args.rule_ids.split(','))
    client.DeactivateFlows(request)


@grpc_wrapper
def activate_dynamic_rule(client, args):
    request = ActivateFlowsRequest(
        sid=SIDUtils.to_pb(args.imsi),
        dynamic_rules=[PolicyRule(
            id=args.rule_id,
            priority=args.priority,
            hard_timeout=args.hard_timeout,
            flow_list=[
                FlowDescription(match=FlowMatch(
                    ipv4_dst=args.ipv4_dst, direction=FlowMatch.UPLINK)),
                FlowDescription(match=FlowMatch(
                    ipv4_src=args.ipv4_dst, direction=FlowMatch.DOWNLINK)),
            ],
        )])
    response = client.ActivateFlows(request)
    _print_rule_mod_results(response.dynamic_rule_results)


def _print_rule_mod_results(results):
    # The message cannot be directly printed because SUCCESS is mapped to 0,
    # which is ignored in the printing by default.
    for result in results:
        print(result.rule_id,
              RuleModResult.Result.Name(result.result))


@grpc_wrapper
def display_flows(_unused, args):
    pipelined_config = load_service_config('pipelined')
    bridge_name = pipelined_config['bridge_name']
    flows = []
    try:
        flows = BridgeTools.get_flows_for_bridge(bridge_name, args.table_num)
    except subprocess.CalledProcessError as e:
        if (e.returncode == errno.EPERM):
            print("Need to run as root to dump flows")
        return

    # Parse the flows and print it decoding note
    for flow in flows[1:]:
        flow = flow.replace('00', '').replace('.', '')
        # If there is a note, decode it otherwise just print the flow
        if 'note:' in flow:
            prefix = flow.split('note:')
            print(prefix[0] + "note:" + str(binascii.unhexlify(prefix[1])))
        else:
            print(flow)


@grpc_wrapper
def get_policy_usage(client, _):
    rule_table = client.GetPolicyUsage(Void())
    pprint(rule_table)


def create_enforcement_parser(apps):
    """
    Creates the argparse subparser for the enforcement app
    """
    app = apps.add_parser('enforcement')
    subparsers = app.add_subparsers(title='subcommands', dest='cmd')

    # Add subcommands
    subcmd = subparsers.add_parser('activate_flows', help='Activate flows')
    subcmd.add_argument('--imsi', help='Subscriber ID', default='IMSI12345')
    subcmd.add_argument('--rule_ids',
                        help='Comma separated rule ids', default='rule1,rule2')
    subcmd.set_defaults(func=activate_flows)

    subcmd = subparsers.add_parser('deactivate_flows', help='Deactivate flows')
    subcmd.add_argument('--imsi', help='Subscriber ID', default='IMSI12345')
    subcmd.add_argument('--rule_ids',
                        help='Comma separated rule ids', default='rule1,rule2')
    subcmd.set_defaults(func=deactivate_flows)

    subcmd = subparsers.add_parser('activate_dynamic_rule',
                                   help='Activate dynamic flows')
    subcmd.add_argument('--imsi', help='Subscriber ID', default='IMSI12345')
    subcmd.add_argument('--rule_id', help='rule id to add', default='rule1')
    subcmd.add_argument('--ipv4_dst', help='ipv4 dst for rule', default='')
    subcmd.add_argument('--priority', help='priority for rule',
                        type=int, default=0)
    subcmd.add_argument('--hard_timeout', help='hard timeout for rule',
                        type=int, default=0)
    subcmd.set_defaults(func=activate_dynamic_rule)

    subcmd = subparsers.add_parser('display_flows', help='Display flows')
    subcmd.add_argument('--table_num', help='table number to filter')
    subcmd.set_defaults(func=display_flows)

    subcmd = subparsers.add_parser('get_policy_usage',
                                   help='Get policy usage stats')
    subcmd.set_defaults(func=get_policy_usage)


# --------------------------
# Debugging
# --------------------------

@grpc_wrapper
def get_table_assignment(client, args):
    response = client.GetAllTableAssignments(Void())
    table_assignments = response.table_assignments
    if args.apps:
        app_filter = args.apps.split(',')
        table_assignments = [table_assignment for table_assignment in
                             table_assignments if
                             table_assignment.app_name in app_filter]

    table_template = '{:<25}{:<20}{:<25}'
    print(table_template.format('App', 'Main Table', 'Scratch Tables'))
    print('-' * 70)
    for table_assignment in table_assignments:
        print(table_template.format(
            table_assignment.app_name,
            table_assignment.main_table,
            str([table for table in table_assignment.scratch_tables])))


def create_debug_parser(apps):
    """
    Creates the argparse subparser for the debugging commands
    """
    app = apps.add_parser('debug')
    subparsers = app.add_subparsers(title='subcommands', dest='cmd')

    # Add subcommands
    subcmd = subparsers.add_parser('table_assignment',
                                   help='Get the table assignment for apps.')
    subcmd.add_argument('--apps',
                        help='Comma separated list of app names. If not set, '
                             'all table assignments will be printed.')
    subcmd.set_defaults(func=get_table_assignment)


# --------------------------
# Pipelined base CLI
# --------------------------

def create_parser():
    """
    Creates the argparse parser with all the arguments.
    """
    parser = argparse.ArgumentParser(
        description='Management CLI for pipelined',
        formatter_class=argparse.ArgumentDefaultsHelpFormatter)
    apps = parser.add_subparsers(title='apps', dest='cmd')
    create_metering_parser(apps)
    create_enforcement_parser(apps)
    create_debug_parser(apps)
    return parser


def main():
    parser = create_parser()

    # Parse the args
    args = parser.parse_args()
    if not args.cmd:
        parser.print_usage()
        exit(1)

    # Execute the subcommand function
    args.func(args, PipelinedStub, 'pipelined')


if __name__ == "__main__":
    main()
