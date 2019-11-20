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
  RemoveEquipmentTypeMutationMutationResponse,
  RemoveEquipmentTypeMutationMutationVariables,
} from './__generated__/RemoveEquipmentTypeMutation.graphql';

const mutation = graphql`
  mutation RemoveEquipmentTypeMutation($id: ID!) {
    removeEquipmentType(id: $id)
  }
`;

export default (
  variables: RemoveEquipmentTypeMutationMutationVariables,
  callbacks?: MutationCallbacks<RemoveEquipmentTypeMutationMutationResponse>,
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
