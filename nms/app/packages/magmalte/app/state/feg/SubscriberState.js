/*
 * Copyright 2020 The Magma Authors.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * @flow strict-local
 * @format
 */

import type {EnqueueSnackbarOptions} from 'notistack';
import type {network_id, subscriber_state} from '@fbcnms/magma-api';

import {FetchSubscriberState} from '../lte/SubscriberState';
import {getServicedAccessNetworks} from '../../components/FEGServicingAccessGatewayKPIs';

type InitSubscriberStateProps = {
  networkId: network_id,
  setSessionState: ({[network_id]: {[string]: subscriber_state}}) => void,
  enqueueSnackbar?: (
    msg: string,
    cfg: EnqueueSnackbarOptions,
  ) => ?(string | number),
};

/**
 * Initalizes the Subscriber State with all the sessions that
 * this federation network services.
 *
 * @param {network_id} networkId Id of the federation network.
 * @param {({[string]: federation_gateway}) => void} setSessionState Sets the subscriber session state.
 * @param {(msg, cfg,) => ?(string | number),} enqueueSnackbar Snackbar to display error.
 */
export async function InitSubscriberState(props: InitSubscriberStateProps) {
  const {networkId, setSessionState, enqueueSnackbar} = props;
  const sessionState = await FetchFegSubscriberState({
    networkId,
    enqueueSnackbar,
  });
  setSessionState(sessionState);
}

type FetchProps = {
  networkId: string,
  enqueueSnackbar?: (
    msg: string,
    cfg: EnqueueSnackbarOptions,
  ) => ?(string | number),
};

/**
 * Fetches and returns the subscriber session state of all the serviced
 * federated lte networks under by this federation network and whose
 * subscriber session is not managed by the HSS.
 *
 * @param {network_id} networkId Id of the federation network.
 * @param {(msg, cfg,) => ?(string | number),} enqueueSnackbar Snackbar to display error.
 */
export async function FetchFegSubscriberState(props: FetchProps) {
  const {networkId, enqueueSnackbar} = props;
  const servicedAccessNetworks = await getServicedAccessNetworks(
    networkId,
    enqueueSnackbar,
  );
  const sessionState = {};
  for (const servicedAccessNetwork of servicedAccessNetworks) {
    // only save subscriber state of networks with hss disabled
    if (!servicedAccessNetwork?.cellular?.epc?.hss_relay_enabled) {
      const servicedAccessNetworkId = servicedAccessNetwork.id;
      const state = await FetchSubscriberState({
        networkId: servicedAccessNetworkId,
        enqueueSnackbar,
      });
      // group session states under their network id
      sessionState[servicedAccessNetworkId] = state;
    }
  }
  return sessionState;
}
