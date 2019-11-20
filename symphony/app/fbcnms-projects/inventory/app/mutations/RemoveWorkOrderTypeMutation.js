/**
 * Copyright 2004-present Facebook. All Rights Reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 *
 * @flow
 * @format
 */

import RelayEnvironment from '../common/RelayEnvironment.js';
import {commitMutation, graphql} from 'react-relay';
import type {MutationCallbacks} from './MutationCallbacks.js';
import type {
  RemoveWorkOrderTypeMutationMutationResponse,
  RemoveWorkOrderTypeMutationVariables,
} from './__generated__/RemoveWorkOrderTypeMutation.graphql';

const mutation = graphql`
  mutation RemoveWorkOrderTypeMutation($id: ID!) {
    removeWorkOrderType(id: $id)
  }
`;

export default (
  variables: RemoveWorkOrderTypeMutationVariables,
  callbacks?: MutationCallbacks<RemoveWorkOrderTypeMutationMutationResponse>,
  updater?: (store: any) => void,
) => {
  const {onCompleted, onError} = callbacks ? callbacks : {};
  commitMutation(RelayEnvironment, {
    mutation,
    variables,
    updater,
    onCompleted,
    onError,
  });
};
