/**
 * Copyright 2004-present Facebook. All Rights Reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 *
 * @flow
 * @format
 */

import type {
  EditWorkOrderMutationResponse,
  EditWorkOrderMutationVariables,
} from './__generated__/EditWorkOrderMutation.graphql';
import type {MutationCallbacks} from './MutationCallbacks.js';

import RelayEnvironemnt from '../common/RelayEnvironment.js';
import {commitMutation, graphql} from 'react-relay';

const mutation = graphql`
  mutation EditWorkOrderMutation($input: EditWorkOrderInput!) {
    editWorkOrder(input: $input) {
      id
      name
      description
      ownerName
      creationDate
      installDate
      status
      priority
      assignee
      ...WorkOrderDetails_workOrder
      ...WorkOrdersView_workOrder
    }
  }
`;

export default (
  variables: EditWorkOrderMutationVariables,
  callbacks?: MutationCallbacks<EditWorkOrderMutationResponse>,
  updater?: (store: any) => void,
) => {
  const {onCompleted, onError} = callbacks ? callbacks : {};
  commitMutation(RelayEnvironemnt, {
    mutation,
    variables,
    updater,
    onCompleted,
    onError,
  });
};
