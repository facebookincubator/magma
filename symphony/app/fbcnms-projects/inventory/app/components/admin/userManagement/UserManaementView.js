/**
 * Copyright 2004-present Facebook. All Rights Reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 *
 * @flow
 * @format
 */
import type {ContextRouter} from 'react-router';
import type {NavigatableView} from '@fbcnms/ui/components/design-system/View/NavigatableViews';

import * as React from 'react';
import AppContext from '@fbcnms/ui/context/AppContext';
import NavigatableViews from '@fbcnms/ui/components/design-system/View/NavigatableViews';
import NewUserDialog from './users/NewUserDialog';
import PermissionsGroupCard from './groups/PermissionsGroupCard';
import PermissionsGroupsView, {
  PERMISSION_GROUPS_VIEW_NAME,
} from './groups/PermissionsGroupsView';
import PermissionsPoliciesView, {
  PERMISSION_POLICIES_VIEW_NAME,
} from './policies/PermissionsPoliciesView';
import PermissionsPolicyCard from './policies/PermissionsPolicyCard';
import Strings from '@fbcnms/strings/Strings';
import UsersView from './users/UsersView';
import fbt from 'fbt';
import {
  ButtonAction,
  OptionsAction,
} from '@fbcnms/ui/components/design-system/View/ViewHeaderActions';
import {FormContextProvider} from '../../../common/FormContext';
import {NEW_DIALOG_PARAM, POLICY_TYPES} from './utils/UserManagementUtils';
import {UserManagementContextProvider} from './UserManagementContext';
import {useCallback, useContext, useMemo, useState} from 'react';
import {useHistory, withRouter} from 'react-router-dom';

const USERS_HEADER = fbt(
  'Users & Roles',
  'Header for view showing system users settings',
);

type Props = ContextRouter;

const UserManaementView = ({match}: Props) => {
  const history = useHistory();
  const basePath = match.path;
  const [addingNewUser, setAddingNewUser] = useState(false);
  const gotoGroupsPage = useCallback(() => history.push(`${basePath}/groups`), [
    history,
    basePath,
  ]);
  const gotoPoliciesPage = useCallback(
    () => history.push(`${basePath}/policies`),
    [history, basePath],
  );

  const {isFeatureEnabled} = useContext(AppContext);
  const userManagementDevMode = isFeatureEnabled('user_management_dev');

  const usersComponent = useMemo(() => {
    return {
      header: {
        title: `${USERS_HEADER}`,
        subtitle:
          'Add and manage your organization users, and set their role to control their global settings',
        actionButtons: [
          <ButtonAction action={() => setAddingNewUser(true)}>
            <fbt desc="">Add User</fbt>
          </ButtonAction>,
        ],
      },
      children: <UsersView />,
    };
  }, []);

  const VIEWS: Array<NavigatableView> = useMemo(() => {
    const views = [
      {
        routingPath: 'users/:id',
        targetPath: 'users/all',
        menuItem: {
          label: USERS_HEADER,
          tooltip: `${USERS_HEADER}`,
        },
        component: usersComponent,
      },
      {
        routingPath: 'groups',
        menuItem: {
          label: PERMISSION_GROUPS_VIEW_NAME,
          tooltip: `${PERMISSION_GROUPS_VIEW_NAME}`,
        },
        component: {
          header: {
            title: `${PERMISSION_GROUPS_VIEW_NAME}`,
            subtitle:
              'Create groups with different rules and add users to apply permissions',
            actionButtons: userManagementDevMode
              ? [
                  <ButtonAction
                    action={() => history.push(`group/${NEW_DIALOG_PARAM}`)}>
                    <fbt desc="">Create Group</fbt>
                  </ButtonAction>,
                ]
              : [],
          },
          children: <PermissionsGroupsView />,
        },
      },
      {
        routingPath: 'group/:id',
        component: {
          children: (
            <PermissionsGroupCard
              redirectToGroupsView={gotoGroupsPage}
              onClose={gotoGroupsPage}
            />
          ),
        },
        relatedMenuItemIndex: 1,
      },
    ];

    if (userManagementDevMode) {
      views.push(
        {
          routingPath: 'policies',
          menuItem: {
            label: PERMISSION_POLICIES_VIEW_NAME,
            tooltip: `${PERMISSION_POLICIES_VIEW_NAME}`,
          },
          component: {
            header: {
              title: `${PERMISSION_POLICIES_VIEW_NAME}`,
              subtitle: 'Manage policies and apply them to groups.',
              actionButtons: [
                <OptionsAction
                  options={[
                    POLICY_TYPES.InventoryPolicy,
                    POLICY_TYPES.WorkforcePolicy,
                  ].map(type => ({
                    key: type.key,
                    value: type.key,
                    label: fbt(
                      fbt.param('policy type', type.value) + ' Policy',
                      'create policy of given type',
                    ),
                  }))}
                  optionAction={typeKey => {
                    history.push(`policy/${NEW_DIALOG_PARAM}?type=${typeKey}`);
                  }}>
                  <fbt desc="">Create Policy</fbt>
                </OptionsAction>,
              ],
            },
            children: <PermissionsPoliciesView />,
          },
        },
        {
          routingPath: 'policy/:id',
          component: {
            children: (
              <PermissionsPolicyCard
                redirectToPoliciesView={gotoPoliciesPage}
                onClose={gotoPoliciesPage}
              />
            ),
          },
          relatedMenuItemIndex: 3,
        },
      );
    }

    return views;
  }, [gotoGroupsPage, gotoPoliciesPage, history, userManagementDevMode]);

  return (
    <UserManagementContextProvider>
      <FormContextProvider permissions={{adminRightsRequired: true}}>
        <NavigatableViews
          header={Strings.admin.users.viewHeader}
          views={VIEWS}
          routingBasePath={basePath}
        />
      </FormContextProvider>
      {addingNewUser && (
        <NewUserDialog onClose={() => setAddingNewUser(false)} />
      )}
    </UserManagementContextProvider>
  );
};

export default withRouter(UserManaementView);
