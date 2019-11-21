/**
 * Copyright 2004-present Facebook. All Rights Reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 *
 * @flow
 * @format
 */

import Admin from './admin/Admin';
import AppContent from '@fbcnms/ui/components/layout/AppContent';
import AppContext, {AppContextProvider} from '@fbcnms/ui/context/AppContext';
import AppSideBar from '@fbcnms/ui/components/layout/AppSideBar';
import ApplicationMain from '@fbcnms/ui/components/ApplicationMain';
import Automation from './automation/Automation';
import Configure from '../pages/Configure';
import EquipmentComparisonView from './comparison_view/EquipmentComparisonView';
import ExpandButtonContext from './context/ExpandButtonContext';
import IDToolMain from './id/IDToolMain';
import Inventory from '../pages/Inventory';
import InventoryComparisonView from './comparison_view/InventoryComparisonView';
import LocationsMap from './map/LocationsMap';
import MagmaMain from '@fbcnms/magmalte/app/components/Main';
import MainNavListItems from './MainNavListItems';
import React, {useCallback, useContext, useEffect, useState} from 'react';
import ServiceComparisonView from './services/ServiceComparisonView';
import Settings from './Settings';
import WorkOrdersMain from './work_orders/WorkOrdersMain';

import nullthrows from '@fbcnms/util/nullthrows';
import {Redirect, Route, Switch} from 'react-router-dom';
import {getProjectLinks} from '@fbcnms/magmalte/app/common/projects';
import {makeStyles} from '@material-ui/styles';
import {setLoggerUser} from '../common/LoggingUtils';
import {shouldShowSettings} from '@fbcnms/magmalte/app/components/Settings';
import {useRouter} from '@fbcnms/ui/hooks';

const useStyles = makeStyles({
  root: {
    display: 'flex',
  },
});

function Index() {
  const classes = useStyles();
  const {isExpandButtonShown, expand, collapse, isExpanded} = useContext(
    ExpandButtonContext,
  );

  const multiSubjectReports = useContext(AppContext).isFeatureEnabled(
    'multi_subject_reports',
  );
  const {tabs, user, ssoEnabled} = useContext(AppContext);
  const {location, relativeUrl} = useRouter();

  return (
    <div className={classes.root}>
      <AppSideBar
        useExpandButton={location.pathname.includes('inventory/inventory')}
        expanded={isExpanded}
        showExpandButton={isExpandButtonShown}
        onExpandClicked={() => (isExpanded ? collapse() : expand())}
        mainItems={<MainNavListItems />}
        projects={getProjectLinks(tabs, user)}
        showSettings={shouldShowSettings({
          isSuperUser: user.isSuperUser,
          ssoEnabled,
        })}
        user={nullthrows(user)}
      />
      <AppContent>
        <Switch>
          <Route path={relativeUrl('/configure')} component={Configure} />
          <Route path={relativeUrl('/inventory')} component={Inventory} />
          <Route path={relativeUrl('/map')} component={LocationsMap} />
          <Route
            path={relativeUrl('/search')}
            component={
              multiSubjectReports
                ? InventoryComparisonView
                : EquipmentComparisonView
            }
          />
          <Route
            path={relativeUrl('/services')}
            component={ServiceComparisonView}
          />
          <Route path={relativeUrl('/settings')} component={Settings} />
          <Redirect exact from="/" to={relativeUrl('/inventory')} />
          <Redirect exact from="/inventory" to={relativeUrl('/inventory')} />
        </Switch>
      </AppContent>
    </div>
  );
}

function Main() {
  useEffect(() => setLoggerUser(window.CONFIG.appData.user), []);
  const [isExpanded, setIsExpanded] = useState(true);
  const expand = useCallback(() => setIsExpanded(true), []);
  const collapse = useCallback(() => setIsExpanded(false), []);

  const [isExpandButtonShown, showExpandButton] = useState(false);
  return (
    <ApplicationMain>
      <AppContextProvider>
        <ExpandButtonContext.Provider
          value={{
            showExpandButton: () => showExpandButton(true),
            hideExpandButton: () => showExpandButton(false),
            expand: expand,
            collapse: collapse,
            isExpanded,
            isExpandButtonShown,
          }}>
          <Index />
        </ExpandButtonContext.Provider>
      </AppContextProvider>
    </ApplicationMain>
  );
}

export default () => (
  <Switch>
    <Route path="/nms" component={MagmaMain} />
    <Route path="/inventory" component={Main} />
    <Route path="/workorders" component={WorkOrdersMain} />
    <Route path="/admin" component={Admin} />
    <Route path="/automation" component={Automation} />
    <Route path="/id" component={IDToolMain} />
  </Switch>
);
