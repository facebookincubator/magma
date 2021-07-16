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
import type {DataRows} from '../../components/DataGrid';

import DataGrid from '../../components/DataGrid';
import FEGGatewayContext from '../../components/context/FEGGatewayContext';
import LoadingFiller from '@fbcnms/ui/components/LoadingFiller';
import MagmaV1API from '@fbcnms/magma-api/client/WebClient';
import React from 'react';
import nullthrows from '@fbcnms/util/nullthrows';
import useMagmaAPI from '@fbcnms/ui/magma/useMagmaAPI';

import {
  DynamicServices,
  HEALTHY_GATEWAY,
  HEALTHY_STATUS,
  UNAVAILABLE_SERVICE,
  UNHEALTHY_GATEWAY,
} from '../../components/GatewayUtils';
import {
  FEG_GATEWAY,
  REFRESH_INTERVAL,
  useRefreshingContext,
} from '../../components/context/RefreshContext';
import {useRouter} from '@fbcnms/ui/hooks';

export default function GatewayDetailStatus({refresh}: {refresh: boolean}) {
  const {match} = useRouter();
  const networkId: string = nullthrows(match.params.networkId);
  const gatewayId: string = nullthrows(match.params.gatewayId);
  // Auto refresh gateways every 30 seconds
  const refreshCtx = useRefreshingContext({
    context: FEGGatewayContext,
    networkId: networkId,
    type: FEG_GATEWAY,
    interval: REFRESH_INTERVAL,
    id: gatewayId,
    refresh: refresh,
  });
  const fegGateways = refreshCtx.fegGateways || {};
  const health = refreshCtx.health || {};
  const gwInfo = fegGateways[gatewayId] || {};
  let checkInTime;

  if (
    gwInfo.status &&
    gwInfo.status.checkin_time !== undefined &&
    gwInfo.status.checkin_time !== null
  ) {
    checkInTime = new Date(gwInfo.status.checkin_time);
  }

  const startTime = Math.floor(Date.now() / 1000);
  const {response: cpuPercent, isLoading: isCpuPercentLoading} = useMagmaAPI(
    MagmaV1API.getNetworksByNetworkIdPrometheusQueryRange,
    {
      networkId: networkId,
      query: `cpu_percent{gatewayID="${gwInfo.id}", service="magmad"}`,
      start: startTime.toString(),
    },
  );

  const logAggregation =
    !!gwInfo.magmad.dynamic_services &&
    gwInfo.magmad.dynamic_services.includes(DynamicServices.TD_AGENT_BIT);

  const eventAggregation =
    !!gwInfo.magmad.dynamic_services &&
    gwInfo.magmad.dynamic_services.includes(DynamicServices.EVENTD);

  const cpeMonitoring =
    !!gwInfo.magmad.dynamic_services &&
    gwInfo.magmad.dynamic_services.includes(DynamicServices.MONITORD);

  const gwHealth = health[gwInfo?.id].status
    ? health[gwInfo?.id].status === HEALTHY_STATUS
      ? HEALTHY_GATEWAY
      : UNHEALTHY_GATEWAY
    : UNAVAILABLE_SERVICE;

  if (isCpuPercentLoading) {
    return <LoadingFiller />;
  }

  const data: DataRows[] = [
    [
      {
        category: 'Health',
        value: gwHealth,
        statusCircle: true,
        // make kpi inactive if health status had error (health service not enabled)
        statusInactive: health[gwInfo?.id]?.status ? false : true,
        status: gwHealth === HEALTHY_GATEWAY,
        tooltip:
          "Federation gateway's health as reported by the health service",
      },
      {
        category: 'Last Check in',
        value: checkInTime?.toLocaleString() ?? '-',
        statusCircle: false,
        tooltip: 'The last Time the gateway checked in',
      },
      {
        category: 'CPU Usage',
        value: cpuPercent?.data?.result?.[0]?.values?.[0]?.[1] ?? 'Unknown',
        unit:
          cpuPercent?.data?.result?.[0]?.values?.[0]?.[1] ?? false ? '%' : '',
        statusCircle: false,
        tooltip: 'Current Gateway CPU %',
      },
    ],
    [
      {
        category: 'Event Aggregation',
        value: eventAggregation ? 'Enabled' : 'Disabled',
        statusCircle: true,
        status: eventAggregation,
      },
      {
        category: 'Log Aggregation',
        value: logAggregation ? 'Enabled' : 'Disabled',
        statusCircle: true,
        status: logAggregation,
      },
      {
        category: 'CPE Monitoring',
        value: cpeMonitoring ? 'Enabled' : 'Disabled',
        statusCircle: true,
        status: cpeMonitoring,
      },
    ],
  ];
  return <DataGrid data={data} />;
}
