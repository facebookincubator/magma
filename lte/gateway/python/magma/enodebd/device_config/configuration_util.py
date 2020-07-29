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

from lte.protos.mconfig import mconfigs_pb2


def get_enb_rf_tx_desired(
    mconfig: mconfigs_pb2.EnodebD,
    enb_serial: str,
) -> bool:
    """ True if the mconfig specifies to enable transmit on the eNB """
    if mconfig.enb_configs_by_serial is not None and \
            len(mconfig.enb_configs_by_serial) > 0:
        if enb_serial in mconfig.enb_configs_by_serial:
            enb_config = mconfig.enb_configs_by_serial[enb_serial]
            return enb_config.transmit_enabled
        else:
            raise KeyError('Missing eNB from mconfig: %s' % enb_serial)
    return mconfig.allow_enodeb_transmit


def is_enb_registered(mconfig: mconfigs_pb2.EnodebD, enb_serial: str) -> bool:
    """
    True if either:
        - the eNodeB is registered by serial to the Access Gateway
        or
        - the Access Gateway accepts all eNodeB devices
    """
    if mconfig.enb_configs_by_serial is not None and \
            len(mconfig.enb_configs_by_serial) > 0:
        if enb_serial in mconfig.enb_configs_by_serial:
            return True
        else:
            return False
    return True
