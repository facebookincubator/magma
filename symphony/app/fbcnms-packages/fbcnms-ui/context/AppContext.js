/**
 * Copyright 2004-present Facebook. All Rights Reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 *
 * @flow
 * @format
 */
'use strict';

import type {FeatureID} from '@fbcnms/types/features';

import * as React from 'react';
import emptyFunction from '@fbcnms/util/emptyFunction';

export type User = {
  tenant: string,
  email: string,
  isSuperUser: boolean,
};

export type AppContextType = {
  csrfToken: ?string,
  version: ?string,
  networkIds: string[],
  tabs: string[],
  user: User,
  showExpandButton: () => void,
  hideExpandButton: () => void,
  isFeatureEnabled: FeatureID => boolean,
  ssoEnabled: boolean,
};

const AppContext = React.createContext<AppContextType>({
  csrfToken: null,
  version: null,
  networkIds: [],
  tabs: [],
  user: {tenant: '', email: '', isSuperUser: false},
  showExpandButton: emptyFunction,
  hideExpandButton: emptyFunction,
  isFeatureEnabled: () => false,
  ssoEnabled: false,
});

type Props = {|
  children: React.Node,
  networkIDs?: string[],
|};

export function AppContextProvider(props: Props) {
  const {appData} = window.CONFIG;
  const value = {
    ...appData,
    networkIds: props.networkIDs || [],
    isFeatureEnabled: (featureID: FeatureID): boolean => {
      return appData.enabledFeatures.indexOf(featureID) !== -1;
    },
  };

  return (
    <AppContext.Provider value={value}>{props.children}</AppContext.Provider>
  );
}

export default AppContext;
