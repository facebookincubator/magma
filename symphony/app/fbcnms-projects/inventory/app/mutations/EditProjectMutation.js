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
  EditProjectMutationResponse,
  EditProjectMutationVariables,
} from './__generated__/EditProjectMutation.graphql';
import type {MutationCallbacks} from './MutationCallbacks.js';

import RelayEnvironemnt from '../common/RelayEnvironment.js';
import {commitMutation, graphql} from 'react-relay';

const mutation = graphql`
  mutation EditProjectMutation($input: EditProjectInput!) {
    editProject(input: $input) {
      id
      name
      description
      creator
      properties {
        stringValue
        intValue
        floatValue
        booleanValue
        latitudeValue
        longitudeValue
        rangeFromValue
        rangeToValue
        propertyType {
          id
          name
          type
          isEditable
          isInstanceProperty
          stringValue
          intValue
          floatValue
          booleanValue
          latitudeValue
          longitudeValue
          rangeFromValue
          rangeToValue
        }
      }
    }
  }
`;

export default (
  variables: EditProjectMutationVariables,
  callbacks?: MutationCallbacks<EditProjectMutationResponse>,
  // eslint-disable-next-line flowtype/no-weak-types
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
