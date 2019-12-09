/**
 * Copyright 2004-present Facebook. All Rights Reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 *
 * @flow
 * @format
 */

import type {ApiUtil} from '../components/AlarmsApi';

/**
 * I don't understand how to properly type these mocks so using any for now.
 * The consuming code is all strongly typed, this shouldn't be much of an issue.
 */
// eslint-disable-next-line flowtype/no-weak-types
export const useMagmaAPIMock = jest.fn<any, any>(() => ({
  isLoading: false,
  response: [],
  error: null,
}));

// eslint-disable-next-line flowtype/no-weak-types
export const apiMock = jest.fn<any, any>();

/**
 * Make sure when adding new functions to ApiUtil to add their mocks here
 */
export function mockApiUtil(merge?: $Shape<ApiUtil>): ApiUtil {
  return Object.assign(
    {
      useAlarmsApi: useMagmaAPIMock,
      viewFiringAlerts: apiMock,
      viewMatchingAlerts: apiMock,
      createAlertRule: apiMock,
      editAlertRule: apiMock,
      getAlertRules: apiMock,
      deleteAlertRule: apiMock,
      getReceivers: apiMock,
      getRoutes: apiMock,
      getSuppressions: apiMock,
      getMetricSeries: apiMock,
    },
    merge || {},
  );
}
