/**
 * Copyright 2004-present Facebook. All Rights Reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 *
 *
 * @flow strict-local
 * @format
 */

// flowlint untyped-import:off

import type {AddUsersGroupMutationResponse} from '../../../mutations/__generated__/AddUsersGroupMutation.graphql';
import type {EditUserMutationResponse} from '../../../mutations/__generated__/EditUserMutation.graphql';
import type {EditUsersGroupMutationResponse} from '../../../mutations/__generated__/EditUsersGroupMutation.graphql';
import type {MutationCallbacks} from '../../../mutations/MutationCallbacks.js';
import type {StoreUpdater} from '../../../common/RelayEnvironment';
import type {User, UserPermissionsGroup} from './TempTypes';
import type {
  UserManagementContextQuery,
  UserRole,
} from './__generated__/UserManagementContextQuery.graphql';
import type {UserManagementContext_UserQuery} from './__generated__/UserManagementContext_UserQuery.graphql';

import * as React from 'react';
import AddUsersGroupMutation from '../../../mutations/AddUsersGroupMutation';
import EditUserMutation from '../../../mutations/EditUserMutation';
import EditUsersGroupMutation from '../../../mutations/EditUsersGroupMutation';
import LoadingIndicator from '../../../common/LoadingIndicator';
import RelayEnvironment from '../../../common/RelayEnvironment';
import axios from 'axios';
import nullthrows from 'nullthrows';
import {ConnectionHandler, fetchQuery, graphql} from 'relay-runtime';
import {LogEvents, ServerLogger} from '../../../common/LoggingUtils';
import {RelayEnvironmentProvider} from 'react-relay/hooks';
import {Suspense} from 'react';
import {USER_ROLES} from './TempTypes';
import {getGraphError} from '../../../common/EntUtils';
import {useContext} from 'react';
import {useLazyLoadQuery} from 'react-relay/hooks';

export type UserManagementContextValue = {
  groups: Array<UserPermissionsGroup>,
  users: Array<User>,
  addUser: (user: User, password: string) => Promise<User>,
  editUser: (newUserValue: User, updater?: StoreUpdater) => Promise<User>,
  changeUserPassword: (user: User, password: string) => Promise<User>,
  addGroup: UserPermissionsGroup => Promise<UserPermissionsGroup>,
  editGroup: UserPermissionsGroup => Promise<UserPermissionsGroup>,
};

const userQuery = graphql`
  query UserManagementContext_UserQuery($authID: String!) {
    user(authID: $authID) {
      id
    }
  }
`;

const roleToNodeRole = (role: UserRole) =>
  role === USER_ROLES.USER.key ? 0 : 3;

const createNewUserInNode = (newUserValue: User, password: string) => {
  const newUserPayload = {
    email: newUserValue.authID,
    password,
    role: roleToNodeRole(newUserValue.role),
    networkIDs: [],
  };
  return axios.post('/user/async/', newUserPayload);
};
const getUserEntIdByAuthID = authID => {
  return fetchQuery<UserManagementContext_UserQuery>(
    RelayEnvironment,
    userQuery,
    {
      authID,
    },
  ).then(response => nullthrows(response.user?.id));
};
const setNewUserEntValues = (userEntId: string, userValues: User) => {
  userValues.id = userEntId;
  const addNewUserToStore = store => {
    const rootQuery = store.getRoot();
    const newNode = store.get(userValues.id);
    if (newNode == null) {
      return;
    }
    const users = ConnectionHandler.getConnection(
      rootQuery,
      'UserManagementContext_users',
    );
    if (users == null) {
      return;
    }
    const edge = ConnectionHandler.createEdge(
      store,
      users,
      newNode,
      'UserEdge',
    );
    ConnectionHandler.insertEdgeAfter(users, edge);
  };
  return editUser(userValues, addNewUserToStore);
};
const addUser = (newUserValue: User, password: string) => {
  return createNewUserInNode(newUserValue, password)
    .then(() => getUserEntIdByAuthID(newUserValue.authID))
    .then(userId => setNewUserEntValues(userId, newUserValue));
};

const updateUserInNode = (
  email: string,
  role?: UserRole,
  password?: string,
) => {
  const updateUserPayload = {};
  if (password != null) {
    updateUserPayload.password = password;
  }
  if (role != null) {
    updateUserPayload.role = roleToNodeRole(role);
  }
  if (Object.keys(updateUserPayload).length === 0) {
    return Promise.resolve();
  }
  return axios.put(`/user/set/${email}`, updateUserPayload);
};

const changeUserPassword = (user: User, password: string) => {
  return updateUserInNode(user.authID, undefined, password).then(() => user);
};

const editUser = (newUserValue: User, updater?: StoreUpdater) => {
  return new Promise<User>((resolve, reject) => {
    const callbacks: MutationCallbacks<EditUserMutationResponse> = {
      onCompleted: (response, errors) => {
        if (errors && errors[0]) {
          reject(errors[0].message);
        }
        resolve(userResponse2User(response.editUser));
        // TEMP: Need to update Node with the new role.
        // (Once Node is changed to take the role from graph,
        //  we can remove this)
        updateUserInNode(newUserValue.authID, newUserValue.role).catch(
          error => {
            ServerLogger.error(LogEvents.CLIENT_FATAL_ERROR, {
              message: error.message,
              stack: error.stack,
            });
          },
        );
      },
      onError: e => {
        reject(getGraphError(e));
      },
    };
    EditUserMutation(
      {
        input: {
          id: newUserValue.id,
          firstName: newUserValue.firstName,
          lastName: newUserValue.lastName,
          role: newUserValue.role,
          status: newUserValue.status,
        },
      },
      callbacks,
      updater,
    );
  });
};

const editGroup = (newGroupValue: UserPermissionsGroup) => {
  return new Promise<UserPermissionsGroup>((resolve, reject) => {
    const callbacks: MutationCallbacks<EditUsersGroupMutationResponse> = {
      onCompleted: (response, errors) => {
        if (errors && errors[0]) {
          reject(errors[0].message);
        }
        resolve(groupResponse2Group(response.editUsersGroup));
      },
      onError: e => {
        reject(getGraphError(e));
      },
    };
    EditUsersGroupMutation(
      {
        input: {
          id: newGroupValue.id,
          name: newGroupValue.name,
          description: newGroupValue.description,
          status: newGroupValue.status,
        },
      },
      callbacks,
    );
  });
};

const addGroup = (newGroupValue: UserPermissionsGroup) => {
  return new Promise<UserPermissionsGroup>((resolve, reject) => {
    const callbacks: MutationCallbacks<AddUsersGroupMutationResponse> = {
      onCompleted: (response, errors) => {
        if (errors && errors[0]) {
          reject(errors[0].message);
        }
        resolve(groupResponse2Group(response.addUsersGroup));
      },
      onError: e => {
        reject(getGraphError(e));
      },
    };

    const addNewGroupToStore = store => {
      const rootQuery = store.getRoot();
      // eslint-disable-next-line no-warning-comments
      // $FlowFixMe (T62907961) Relay flow types
      const newNode = store.getRootField('addUsersGroup');
      if (newNode == null) {
        return;
      }
      const groups = ConnectionHandler.getConnection(
        rootQuery,
        'UserManagementContext_usersGroups',
      );
      if (groups == null) {
        return;
      }
      const edge = ConnectionHandler.createEdge(
        store,
        groups,
        newNode,
        'UsersGroupEdge',
      );
      ConnectionHandler.insertEdgeAfter(groups, edge);
    };
    AddUsersGroupMutation(
      {
        input: {
          name: newGroupValue.name,
          description: newGroupValue.description,
        },
      },
      callbacks,
      addNewGroupToStore,
    );
  });
};

const UserManagementContext = React.createContext<UserManagementContextValue>({
  groups: [],
  users: [],
  addUser,
  editUser,
  changeUserPassword,
  addGroup,
  editGroup,
});

export function useUserManagement() {
  return useContext(UserManagementContext);
}

type Props = {
  children: React.Node,
};

const usersQuery = graphql`
  query UserManagementContextQuery {
    users(first: 500) @connection(key: "UserManagementContext_users") {
      edges {
        node {
          id
          authID
          firstName
          lastName
          email
          status
          role
          groups {
            id
            name
          }
          profilePhoto {
            id
            fileName
            storeKey
          }
        }
      }
    }
    usersGroups(first: 500)
      @connection(key: "UserManagementContext_usersGroups") {
      edges {
        node {
          id
          name
          description
          status
          members {
            id
            authID
          }
        }
      }
    }
  }
`;

const userResponse2User = userResponse => ({
  id: userResponse.id,
  authID: userResponse.authID,
  firstName: userResponse.firstName,
  lastName: userResponse.lastName,
  role: userResponse.role,
  status: userResponse.status,
});

const usersResponse2Users = usersResponse => {
  const users: Array<User> = [];
  const usersEdges = usersResponse?.edges;
  if (usersEdges == null) {
    return [];
  }
  // using 'for' and not simple 'map' beacuse of flow.
  for (let i = 0; i < usersEdges.length; i++) {
    const userNode = usersEdges[i].node;
    if (userNode == null) {
      continue;
    }
    users.push({
      id: userNode.id,
      authID: userNode.authID,
      firstName: userNode.firstName,
      lastName: userNode.lastName,
      role: userNode.role,
      status: userNode.status,
    });
  }
  return users;
};

const groupResponse2Group = groupResponse => ({
  id: groupResponse.id,
  name: groupResponse.name,
  description: groupResponse.description || '',
  status: groupResponse.status,
  members: groupResponse.members,
});

const groupsResponse2Groups = groupsResponse => {
  const groups: Array<UserPermissionsGroup> = [];
  const groupsEdges = groupsResponse?.edges;
  if (groupsEdges == null) {
    return [];
  }
  // using 'for' and not simple 'map' beacuse of flow.
  for (let i = 0; i < groupsEdges.length; i++) {
    const groupNode = groupsEdges[i].node;
    if (groupNode == null) {
      continue;
    }
    groups.push(groupResponse2Group(groupNode));
  }
  return groups;
};

function ProviderWrap(props: Props) {
  const providerValue = (users, groups) => ({
    groups,
    users,
    addUser,
    editUser,
    changeUserPassword,
    addGroup,
    editGroup,
  });

  const data = useLazyLoadQuery<UserManagementContextQuery>(usersQuery);

  return (
    <UserManagementContext.Provider
      value={providerValue(
        usersResponse2Users(data.users),
        groupsResponse2Groups(data.usersGroups),
      )}>
      {props.children}
    </UserManagementContext.Provider>
  );
}

export function UserManagementContextProvider(props: Props) {
  return (
    <RelayEnvironmentProvider environment={RelayEnvironment}>
      <Suspense fallback={<LoadingIndicator />}>
        <ProviderWrap {...props} />
      </Suspense>
    </RelayEnvironmentProvider>
  );
}

export default UserManagementContext;
