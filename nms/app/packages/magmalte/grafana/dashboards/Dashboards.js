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

//$FlowFixMe TODO: type this package
import * as Grafana from 'grafana-dash-gen';

const netIDVar = 'networkID';
const gwIDVar = 'gatewayID';

const variableSortNumbers: {[VariableSortOption]: number} = {
  none: 0,
  'alpha-asc': 1,
  'alpha-desc': 2,
  'num-asc': 3,
  'num-desc': 4,
  'alpha-insensitive-asc': 5,
  'alpha-insensitive-desc': 6,
};

export const getNetworkTemplate = (
  networkIDs: Array<string>,
): TemplateConfig => {
  return customVariableTemplate({
    name: netIDVar,
    type: 'custom',
    options: networkIDs,
    sort: 'alpha-insensitive-asc',
    includeAll: true,
  });
};

// This templating schema will produce a variable in the dashboard
// named gatewayID which is a multi-selectable option of all the
// gateways associated with this organization that exist for the
// currently selected $networkID. $networkID variable must also
// be configured for this dashboard in order for it to work
export const gatewayTemplate: TemplateConfig = variableTemplate({
  name: gwIDVar,
  query: `label_values({networkID=~"$networkID",gatewayID=~".+"}, ${gwIDVar})`,
  regex: `/.+/`,
  sort: 'alpha-insensitive-asc',
  includeAll: true,
});

export const imsiTemplate = variableTemplate({
  labelName: 'imsi',
  query: `label_values({networkID=~"$networkID",IMSI=~".+"}, IMSI)`,
  regex: `/.+/`,
  sort: 'alpha-insensitive-asc',
  includeAll: true,
});

export const msisdnTemplate = variableTemplate({
  labelName: 'msisdn',
  query: `label_values(msisdn)`,
  regex: `/.+/`,
  sort: 'num-asc',
  includeAll: false,
});

export const apnTemplate = variableTemplate({
  labelName: 'apn',
  query: `label_values({networkID=~"$networkID",apn=~".+"},apn)`,
  regex: `/.+/`,
  sort: 'alpha-insensitive-asc',
  includeAll: true,
});

export const NetworkDBData = (networkIDs: Array<string>): GrafanaDBData => {
  return {
    title: 'Networks',
    description:
      'Metrics relevant to the whole network. Do not edit: edits will be overwritten. Save this dashboard under another name to copy and edit.',
    templates: [getNetworkTemplate(networkIDs)],
    rows: [
      {
        title: '',
        panels: [
          {
            title: 'Number of Connected UEs',
            targets: [
              {
                expr:
                  'sum(ue_connected{networkID=~"$networkID"}) by (networkID)',
                legendFormat: '{{networkID}}',
              },
            ],
            aggregates: {avg: true, max: true},
          },
          {
            title: 'Number of Registered UEs',
            targets: [
              {
                expr:
                  'sum(ue_registered{networkID=~"$networkID"}) by (networkID)',
                legendFormat: '{{networkID}}',
              },
            ],
            aggregates: {avg: true, max: true},
          },
          {
            title: 'Number of Connected eNBs',
            targets: [
              {
                expr:
                  'sum(enb_connected{networkID=~"$networkID"}) by (networkID)',
                legendFormat: '{{networkID}}',
              },
            ],
            aggregates: {avg: true, max: true},
          },
          {
            title: 'S1 Setup',
            targets: [
              {
                expr: 'sum(s1_setup{networkID=~"$networkID"}) by (networkID)',
                legendFormat: 'Total: {{networkID}}',
              },
              {
                expr:
                  'sum(s1_setup{networkID=~"$networkID",result="success"}) by (networkID)',
                legendFormat: 'Success: {{networkID}}',
              },
              {
                expr:
                  'sum(s1_setup{networkID=~"$networkID"})by(networkID)-sum(s1_setup{result="success",networkID=~"$networkID"}) by (networkID)',
                legendFormat: 'Failure: {{networkID}}',
              },
            ],
          },
          {
            title: 'Attach/Reg Attempts',
            targets: [
              {
                expr: 'sum(ue_attach{networkID=~"$networkID"}) by (networkID)',
                legendFormat: 'Total: {{networkID}}',
              },
              {
                expr:
                  'sum(ue_attach{networkID=~"$networkID",result="attach_proc_successful"}) by (networkID)',
                legendFormat: 'Success: {{networkID}}',
              },
              {
                expr:
                  'sum(ue_attach{networkID=~"$networkID"}) by (networkID) -sum(s1_setup{result="attach_proc_successful",networkID=~"$networkID"}) by (networkID)',
                legendFormat: 'Failure: {{networkID}}',
              },
            ],
          },
          {
            title: 'Detach/Dereg Attempts',
            targets: [
              {
                expr: 'sum(ue_detach{networkID=~"$networkID"}) by (networkID)',
                legendFormat: 'Total: {{networkID}}',
              },
              {
                expr:
                  'sum(ue_detach{networkID=~"$networkID",result="attach_proc_successful"}) by (networkID)',
                legendFormat: 'Success: {{networkID}}',
              },
              {
                expr:
                  'sum(ue_detach{networkID=~"$networkID"}) by (networkID) -sum(s1_setup{result="attach_proc_successful",networkID=~"$networkID"}) by (networkID)',
                legendFormat: 'Failure: {{networkID}}',
              },
            ],
          },
          {
            title: 'GPS Connection Uptime',
            targets: [
              {
                expr:
                  'avg(enodeb_gps_connected{networkID=~"$networkID"}) by (networkID)',
                legendFormat: '{{networkID}}',
              },
            ],
            unit: 's',
          },
          {
            title: 'Device Transmitting Status',
            targets: [
              {
                expr:
                  'avg(enodeb_rf_tx_enabled{networkID=~"$networkID"}) by (networkID)',
                legendFormat: '{{networkID}}',
              },
            ],
          },
          {
            title: 'Service Requests',
            targets: [
              {
                expr:
                  'sum(service_request{networkID=~"$networkID"}) by (networkID)',
                legendFormat: 'Total: {{networkID}}',
              },
              {
                expr:
                  'sum(service_request{networkID=~"$networkID",result="success"}) by (networkID)',
                legendFormat: 'Success: {{networkID}}',
              },
              {
                expr:
                  'sum(service_request{networkID=~"$networkID"}) by (networkID)-sum(s1_setup{result="success",networkID=~"$networkID"}) by (networkID)',
                legendFormat: 'Failure: {{networkID}}',
              },
            ],
          },
        ],
      },
    ],
  };
};

export const GatewayDBData = (networkIDs: Array<string>): GrafanaDBData => {
  return {
    title: 'Gateways',
    description:
      'Metrics relevant to the gateways. Do not edit: edits will be overwritten. Save this dashboard under another name to copy and edit.',
    templates: [getNetworkTemplate(networkIDs), gatewayTemplate],
    rows: [
      {
        title: '',
        panels: [
          {
            title: 'E-Node B Status',
            targets: [
              {
                expr:
                  'enodeb_rf_tx_enabled{gatewayID=~"$gatewayID",networkID=~"$networkID"}',
                legendFormat: '{{gatewayID}}',
              },
            ],
          },
          {
            title: 'Connected Subscribers',
            targets: [
              {
                expr:
                  'ue_connected{gatewayID=~"$gatewayID",networkID=~"$networkID"}',
                legendFormat: '{{gatewayID}}',
              },
            ],
          },
          {
            title: 'Download Throughput',
            targets: [
              {
                expr:
                  'pdcp_user_plane_bytes_dl{gatewayID=~"$gatewayID",service="enodebd",networkID=~"$networkID"}/1000',
                legendFormat: '{{gatewayID}}',
              },
            ],
            unit: 'Bps',
          },
          {
            title: 'Upload Throughput',
            targets: [
              {
                expr:
                  'pdcp_user_plane_bytes_ul{gatewayID=~"$gatewayID",service="enodebd",networkID=~"$networkID"}/1000',
                legendFormat: '{{gatewayID}}',
              },
            ],
            unit: 'Bps',
          },
          {
            title: 'Latency',
            targets: [
              {
                expr:
                  'magmad_ping_rtt_ms{gatewayID=~"$gatewayID",networkID=~"$networkID",metric="rtt_ms"}',
                legendFormat: '{{gatewayID}}',
              },
            ],
            unit: 's',
          },
          {
            title: 'Gateway CPU %',
            targets: [
              {
                expr:
                  'cpu_percent{gatewayID=~"$gatewayID",networkID=~"$networkID"}',
                legendFormat: '{{gatewayID}}',
              },
            ],
            unit: 'percent',
          },
          {
            title: 'Temperature',
            targets: [
              {
                expr:
                  'temperature{gatewayID=~"$gatewayID",networkID=~"$networkID"}',
                legendFormat: '{{gatewayID}} - {{sensor}}',
              },
            ],
            yMin: null,
            unit: 'celsius',
          },
          {
            title: 'Disk %',
            targets: [
              {
                expr:
                  'disk_percent{gatewayID=~"$gatewayID",networkID=~"$networkID"}',
                legendFormat: '{{gatewayID}}',
              },
            ],
            unit: 'percent',
          },
          {
            title: 's6a Auth Failure',
            targets: [
              {
                expr:
                  's6a_auth_failure{gatewayID=~"$gatewayID",networkID=~"$networkID"}',
                legendFormat: '{{gatewayID}}',
              },
            ],
          },
        ],
      },
    ],
  };

export const SubscriberDBData: GrafanaDBData = {
  title: 'Subscribers',
  description:
    'Metrics relevant to subscribers. Do not edit: edits will be overwritten. Save this dashboard under another name to copy and edit.',
  templates: [networkTemplate, msisdnTemplate],
  rows: [
    {
      title: '',
      panels: [
        {
          title: 'UE Data Usage In',
          targets: [
            {
              expr:
                'sum(ue_reported_usage{msisdn=~"$msisdn", direction="down"}) by (IMSI, apn, msisdn)',
              legendFormat: '{{IMSI}}, MSISDN: {{msisdn}}, APN: {{apn}}',
            },
          ],
          unit: 'decbytes',
          description: 'Inbound data per subscriber measured in bytes.',
        },
        {
          title: 'UE Data Usage Out',
          targets: [
            {
              expr:
                'sum(ue_reported_usage{msisdn=~"$msisdn", direction="up"}) by (IMSI, apn, msisdn)',
              legendFormat: '{{IMSI}}, MSISDN: {{msisdn}}, APN: {{apn}}',
            },
          ],
          unit: 'decbytes',
          description: 'Outbound data per subscriber measured in bytes.',
        },
        {
          title: 'Throughput In',
          targets: [
            {
              expr:
                'avg(rate(ue_reported_usage{msisdn=~"$msisdn", direction="down"}[5m])) by (IMSI, apn, msisdn)',
              legendFormat: '{{IMSI}}, MSISDN: {{msisdn}}, APN: {{apn}}',
            },
          ],
          unit: 'Bps',
          description:
            'Inbound data rate per subscriber measured in bytes/second.',
        },
        {
          title: 'Throughput Out',
          targets: [
            {
              expr:
                'avg(rate(ue_reported_usage{msisdn=~"$msisdn", direction="up"}[5m])) by (IMSI, apn, msisdn)',
              legendFormat: '{{IMSI}}, MSISDN: {{msisdn}}, APN: {{apn}}',
            },
          ],
          unit: 'Bps',
          description:
            'Outbound data rate per subscriber measured in bytes/second.',
        },
      ],
    },
  ],
};

export const InternalDBData = (networkIDs: Array<string>): GrafanaDBData => {
  return {
    title: 'Internal',
    description:
      'Metrics relevant to the internals of gateways. Do not edit: edits will be overwritten. Save this dashboard under another name to copy and edit.',
    templates: [getNetworkTemplate(networkIDs), gatewayTemplate],
    rows: [
      {
        title: '',
        panels: [
          {
            title: 'Physical Memory Utilization Percent',
            targets: [
              {
                expr:
                  'mem_free{gatewayID=~"$gatewayID"}/mem_total{gatewayID=~"$gatewayID",networkID=~"$networkID"} * 100',
                legendFormat: '{{gatewayID}}',
              },
            ],
          },
          {
            title: 'Temperature',
            targets: [
              {
                expr:
                  'temperature{gatewayID=~"$gatewayID",sensor="coretemp_0",networkID=~"$networkID"}',
                legendFormat: '{{gatewayID}} - {{sensor}}',
              },
            ],
            unit: 'percent',
          },
          {
            title: 'Virtual Memory Percent',
            targets: [
              {
                expr:
                  'virtual_memory_percent{gatewayID=~"$gatewayID",networkID=~"$networkID"}',
                legendFormat: '{{gatewayID}}',
              },
            ],
            unit: 'percent',
          },
          {
            title: 'Backhaul Latency',
            targets: [
              {
                expr:
                  'magmad_ping_rtt_ms{gatewayID=~"$gatewayID",service="magmad",host="8.8.8.8",metric="rtt_ms",networkID=~"$networkID"}',
                legendFormat: '{{gatewayID}}',
              },
            ],
            unit: 's',
          },
          {
            title: 'System Uptime',
            targets: [
              {
                expr:
                  'process_uptime_seconds{gatewayID=~"$gatewayID",networkID=~"$networkID"}',
                legendFormat: '{{gatewayID}}-{{service}}',
              },
            ],
            unit: 's',
          },
          {
            title: 'Number of Service Restarts',
            targets: [
              {
                expr:
                  'unexpected_service_restarts{gatewayID=~"$gatewayID",networkID=~"$networkID"}',
                legendFormat: '{{gatewayID}}-{{service_name}}',
              },
            ],
          },
        ],
      },
    ],
  };
};

export function createDashboard(dbdata: GrafanaDBData) {
  const rows = dbdata.rows.map(conf => {
    const row = new Grafana.Row({title: conf.title});
    conf.panels.forEach(panel => row.addPanel(newPanel(panel)));
    return row;
  });
  const db = new Grafana.Dashboard({
    schemaVersion: 6,
    title: dbdata.title,
    templating: dbdata.templates,
    description: dbdata.description,
    rows,
  });
  db.state.editable = false;

  // Necessary to make custom templates display "all" option
  for (const template of db.state?.templating?.list) {
    if (template.type === 'custom' && template.includeAll) {
      template.options.unshift({selected: true, text: 'All', value: '$__all'});
    }
    template.current = template.options[0];
  }
  return db;
}

export type GrafanaDBData = {
  title: string,
  description: string,
  rows: Array<GrafanaDBRow>,
  templates: Array<TemplateConfig>,
};

type GrafanaDBRow = {
  title: string,
  panels: PanelParams[],
};

type PanelParams = {
  title: string,
  targets: Array<{expr: string, legendFormat?: string}>,
  unit?: string,
  yMin?: ?number,
  aggregates?: {avg?: boolean, max?: boolean},
  description?: string,
};

function newPanel(params: PanelParams) {
  const pan = new Grafana.Panels.Graph({
    title: params.title,
    span: 6,
    datasource: 'default',
    description: params.description ?? '',
  });
  // Have to add this after to avoid grafana-dash-gen from forcing the target
  // into a Graphite format
  pan.state.targets = params.targets;

  // "short" is the default unit for grafana (no unit)
  pan.state.y_formats[0] = params.unit ?? 'short';

  // yMin should be 0 at minimum unless otherwise specified.
  // null is used to indicate 'auto' in grafana
  if (params.yMin === null) {
    pan.state.grid.leftMin = null;
  } else {
    pan.state.grid.leftMin = params.yMin ?? 0;
  }

  pan.state.legend.avg = params.aggregates?.avg ?? false;
  pan.state.legend.max = params.aggregates?.max ?? false;
  return pan;
}

export type TemplateParams = {
  name: string,
  query?: string,
  options?: Array<string>,
  regex?: string,
  sort?: VariableSortOption,
  includeAll: boolean,
  type?: string,
};

type VariableSortOption =
  | 'none'
  | 'alpha-asc'
  | 'alpha-desc'
  | 'num-asc'
  | 'num-desc'
  | 'alpha-insensitive-asc'
  | 'alpha-insensitive-desc';

export function variableTemplate(params: TemplateParams): TemplateConfig {
  return {
    allValue: '.+',
    definition: params.query,
    hide: 0,
    includeAll: params.includeAll,
    allFormat: 'glob',
    multi: true,
    name: params.name,
    query: params.query ?? '',
    options: params.options ?? [],
    regex: params.regex,
    type: params.type ?? 'query',
    refresh: true,
    useTags: false,
    sort: params.sort ? variableSortNumbers[params.sort] : 0,
  };
}

export function customVariableTemplate(params: TemplateParams): TemplateConfig {
  return {
    options: params.options ?? [],
    includeAll: true,
    name: params.name,
    multi: true,
    allFormat: 'glob',
    allValue: '.+',
  };
}

export type TemplateConfig = {
  allValue: string,
  definition?: string,
  hide?: number,
  includeAll: boolean,
  allFormat: string,
  multi: boolean,
  name: string,
  query?: string,
  options: Array<string>,
  regex?: string,
  type?: string,
  refresh?: boolean,
  useTags?: boolean,
};
