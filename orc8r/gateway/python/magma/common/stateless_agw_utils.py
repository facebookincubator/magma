"""
Copyright 2020 The Magma Authors.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree.

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
"""

import logging
import subprocess

from magma.configuration.service_configs import (
    load_override_config,
    load_service_config,
    save_override_config,
)
from orc8r.protos import magmad_pb2

STATELESS_SERVICE_CONFIGS = [
    ("mme", "use_stateless", True),
    ("mobilityd", "persist_to_redis", True),
    ("pipelined", "clean_restart", False),
    ("pipelined", "redis_enabled", True),
    ("sessiond", "support_stateless", True),
]


def _get_service_start_args_list(service_name):
    return ["service", service_name, "start"]


def _get_service_stop_args_list(service_name):
    return ["service", service_name, "stop"]


def _check_stateless_service_config(service, config_name, config_value):
    service_config = load_service_config(service)
    if service_config.get(config_name) == config_value:
        logging.info("STATELESS\t%s -> %s" % (service, config_name))
        return magmad_pb2.StatelessAgwMode.STATELESS

    logging.info("STATEFUL\t%s -> %s" % (service, config_name))
    return magmad_pb2.StatelessAgwMode.STATEFUL


def check_stateless_agw():
    num_stateful = 0
    for service, config, value in STATELESS_SERVICE_CONFIGS:
        if (
            _check_stateless_service_config(service, config, value)
            == magmad_pb2.StatelessAgwMode.STATEFUL
        ):
            num_stateful += 1

    if num_stateful == 0:
        res = magmad_pb2.StatelessAgwMode.STATELESS
    elif num_stateful == len(STATELESS_SERVICE_CONFIGS):
        res = magmad_pb2.StatelessAgwMode.STATEFUL
    else:
        res = magmad_pb2.StatelessAgwMode.CORRUPT

    logging.debug(
        "Check returning %s", magmad_pb2.StatelessAgwMode.AgwMode.Name(res)
    )
    return res


def _restart_sctpd():
    logging.info("Restarting sctpd")
    subprocess.call("service sctpd restart".split())


def enable_stateless_agw():
    if check_stateless_agw() == magmad_pb2.StatelessAgwMode.STATELESS:
        logging.info("Nothing to enable, AGW is stateless")
    for service, config, value in STATELESS_SERVICE_CONFIGS:
        cfg = load_override_config(service) or {}
        cfg[config] = value
        save_override_config(service, cfg)

    # restart Sctpd so that eNB connections are reset and local state cleared
    _restart_sctpd()


def disable_stateless_agw():
    if check_stateless_agw() == magmad_pb2.StatelessAgwMode.STATEFUL:
        logging.info("Nothing to disable, AGW is stateful")
    for service, config, value in STATELESS_SERVICE_CONFIGS:
        cfg = load_override_config(service) or {}

        # remove the stateless override
        cfg.pop(config, None)

        save_override_config(service, cfg)

    # restart Sctpd so that eNB connections are reset and local state cleared
    _restart_sctpd()
