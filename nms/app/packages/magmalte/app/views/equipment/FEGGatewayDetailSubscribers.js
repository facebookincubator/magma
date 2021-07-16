/**
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

import type {FEGGatewayDetailType} from './FEGGatewayDetailMain';
import type {subscriber} from '@fbcnms/magma-api';

import ActionTable from '../../components/ActionTable';
import FEGSubscriberContext from '../../components/context/FEGSubscriberContext';
import Link from '@material-ui/core/Link';
import LoadingFiller from '@fbcnms/ui/components/LoadingFiller';
import React from 'react';
import nullthrows from '@fbcnms/util/nullthrows';

import {
  FEG_SUBSCRIBER,
  REFRESH_INTERVAL,
  useRefreshingContext,
} from '../../components/context/RefreshContext';
import {FetchSubscribers} from '../../state/lte/SubscriberState';
import {useEffect, useState} from 'react';
import {useRouter} from '@fbcnms/ui/hooks';

type SubscriberRowType = {
  name: string,
  id: string,
  service: string,
};

/**
 * Returns a table of subscribers serviced by the federation gateway.
 *
 * @param {federation_gateway} gwInfo The Federation Gateway
 * @param {boolean} refresh Boolean Value telling to refresh or not.
 */
export default function GatewayDetailSubscribers(props: FEGGatewayDetailType) {
  const {history, match} = useRouter();
  const networkId: string = nullthrows(match.params.networkId);
  const [subscriberRows, setSubscriberRows] = useState<
    Array<SubscriberRowType>,
  >([]);
  const [isLoading, setIsLoading] = useState(true);
  // Auto refresh  every 30 seconds
  const ctx = useRefreshingContext({
    context: FEGSubscriberContext,
    networkId: networkId,
    type: FEG_SUBSCRIBER,
    interval: REFRESH_INTERVAL,
    refresh: props.refresh,
  });
  const sessionState = ctx?.sessionState || {};
  const subscriberToNetworkIdMap = {};

  Object.keys(sessionState).map(servicedNetworkId => {
    // $FlowIgnore
    const servicedNetworkSessionState = sessionState[servicedNetworkId] || {};
    Object?.keys(servicedNetworkSessionState).map(subscriberImsi => {
      subscriberToNetworkIdMap[subscriberImsi] = servicedNetworkId;
    });
  });
  // get all the subscribers IMSI number serviced by the federation network
  const subscribersImsi = JSON.stringify(Object.keys(subscriberToNetworkIdMap));

  useEffect(() => {
    const fetchSubscribersInfo = async () => {
      const newSubscriberRows = [];
      await Promise.all(
        Object.keys(subscriberToNetworkIdMap).map(async subscriberImsi => {
          // $FlowIgnore
          const subscriberInfo: subscriber = await FetchSubscribers({
            networkId: subscriberToNetworkIdMap[subscriberImsi],
            id: subscriberImsi,
          });
          newSubscriberRows.push({
            name: subscriberInfo?.name || subscriberImsi,
            id: subscriberImsi,
            service: subscriberInfo?.lte?.state || '-',
          });
        }),
      );
      setSubscriberRows(newSubscriberRows);
      setIsLoading(false);
    };
    fetchSubscribersInfo();
    // rerun only when a new subscriber session has been added
  }, [subscribersImsi]); // eslint-disable-line react-hooks/exhaustive-deps

  if (isLoading) {
    return <LoadingFiller />;
  }

  return (
    <ActionTable
      title=""
      data={subscriberRows}
      columns={[
        {title: 'Name', field: 'name'},
        {
          title: 'Subscriber ID',
          field: 'id',
          render: currRow => (
            <Link
              variant="body2"
              component="button"
              onClick={() => {
                history.push(
                  match.url.replace(
                    `equipment/overview/gateway/${props.gwInfo.id}`,
                    `subscribers/overview/${currRow.id}`,
                  ),
                );
              }}>
              {currRow.id}
            </Link>
          ),
        },
        {title: 'Service', field: 'service'},
      ]}
      options={{
        actionsColumnIndex: -1,
        pageSizeOptions: [10],
        toolbar: false,
      }}
    />
  );
}
