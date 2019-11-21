/**
 * Copyright 2004-present Facebook. All Rights Reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 *
 * @flow
 * @format
 */

import RelayEnvironment from '../../../common/RelayEnvironment';
import {buildLocationTypeFilterConfigs, getLocationTypes} from '../FilterUtils';
import {graphql} from 'relay-runtime';
import {useGraphQL} from '@fbcnms/ui/hooks';

const locationTypesQuery = graphql`
  query locationTypesHookLocationTypesQuery {
    locationTypes(first: 20) {
      edges {
        node {
          id
          name
        }
      }
    }
  }
`;

const useLocationTypes = () => {
  const locationTypesResponse = useGraphQL(
    RelayEnvironment,
    locationTypesQuery,
    {},
  );
  return buildLocationTypeFilterConfigs(
    getLocationTypes(locationTypesResponse.response),
  );
};

export default useLocationTypes;
