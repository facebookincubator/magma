/**
 * Copyright 2004-present Facebook. All Rights Reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 *
 * @flow
 * @format
 */

import React from 'react';
import CircularProgress from '@material-ui/core/CircularProgress';
import {Line} from 'react-chartjs-2';
import moment from 'moment';

import {MagmaAPIUrls} from '../../common/MagmaAPI';
import {useAxios, useSnackbar, useRouter} from '@fbcnms/ui/hooks';
import {useMemo} from 'react';
import {makeStyles} from '@material-ui/styles';

type Props = {
  label: string,
  unit: string,
  query: string,
  timeRange: TimeRange,
};

const useStyles = makeStyles({
  loadingContainer: {
    paddingTop: 100,
    textAlign: 'center',
  },
});

export type TimeRange = '24_hours' | '7_days' | '14_days' | '30_days';
type RangeValue = {
  days: number,
  step: string,
  unit: string,
};

const RANGE_VALUES: {[TimeRange]: RangeValue} = {
  '24_hours': {
    days: 1,
    step: '15m',
    unit: 'hour',
  },
  '7_days': {
    days: 7,
    step: '2h',
    unit: 'day',
  },
  '14_days': {
    days: 14,
    step: '4h',
    unit: 'day',
  },
  '30_days': {
    days: 14,
    step: '8h',
    unit: 'day',
  },
};

const COLORS = ['blue', 'red', 'green', 'yellow', 'purple', 'black'];

function Progress() {
  const classes = useStyles();
  return (
    <div className={classes.loadingContainer}>
      <CircularProgress />
    </div>
  );
}

function getStartEndStep(timeRange: TimeRange) {
  const {days, step} = RANGE_VALUES[timeRange];
  const end = moment().toISOString();
  const start = moment()
    .subtract({days})
    .toISOString();
  return {start, end, step};
}

function getUnit(timeRange: TimeRange) {
  return RANGE_VALUES[timeRange].unit;
}

function getColorForIndex(index: number) {
  return COLORS[index % COLORS.length];
}

export default function AsyncMetric(props: Props) {
  const {match} = useRouter();
  const startEndStep = useMemo(() => getStartEndStep(props.timeRange), [
    props.timeRange,
  ]);

  const {error, isLoading, response} = useAxios({
    method: 'get',
    url: MagmaAPIUrls.metricsQueryRange(match),
    params: {
      query: props.query,
      ...startEndStep,
    },
  });

  useSnackbar('Error getting metric ' + props.label, {variant: 'error'}, error);

  const data = useMemo(() => {
    const result = response?.data.data.result;
    if (!result || result.length === 0) {
      return null;
    }

    return {
      datasets: result.map((it, index) => ({
        label: JSON.stringify(it.metric),
        unit: props.unit || '',
        fill: false,
        lineTension: 0,
        pointRadius: 4,
        borderColor: getColorForIndex(index),
        backgroundColor: 'transparent',
        data: it.values.map(i => ({
          t: i[0] * 1000,
          y: parseInt(i[1], 10),
        })),
      })),
    };
  }, [response, props.unit]);

  if (error || isLoading || !response) {
    return <Progress />;
  }

  if (data === null) {
    return 'No Data';
  }

  return (
    <Line
      options={{
        maintainAspectRatio: false,
        scaleShowValues: true,
        scales: {
          xAxes: [
            {
              type: 'time',
              time: {
                unit: getUnit(props.timeRange),
                round: 'second',
                tooltipFormat: ' YYYY/MM/DD h:mm:ss a',
              },
              scaleLabel: {
                display: true,
                labelString: 'Date',
              },
            },
          ],
          yAxes: [
            {
              position: 'left',
              scaleLabel: {
                display: true,
                labelString: props.unit,
              },
            },
          ],
        },
        tooltips: {
          enabled: true,
          mode: 'nearest',
          callbacks: {
            label: (tooltipItem, data) =>
              tooltipItem.yLabel + data.datasets[tooltipItem.datasetIndex].unit,
          },
        },
      }}
      legend={{display: false}}
      data={data}
    />
  );
}
