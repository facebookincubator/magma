/**
 * Copyright 2004-present Facebook. All Rights Reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 *
 * @flow
 * @format
 */

import {useEffect, useState} from 'react';
import {Environment, fetchQuery} from 'relay-runtime';

export default function(
  env: Environment,
  query: any,
  variables: {[string]: mixed},
) {
  const [error, setError] = useState(null);
  const [response, setResponse] = useState(null);
  const [isLoading, setIsLoading] = useState(false);

  useEffect(() => {
    if (!Object.keys(variables).length) {
      return;
    }

    setError(null);
    setIsLoading(true);
    fetchQuery(env, query, variables)
      .then(response => {
        setResponse(response);
        setIsLoading(false);
      })
      .catch(error => {
        setError(error);
        setIsLoading(false);
      });
  }, [query().text, JSON.stringify(variables)]);

  return {error, response, isLoading};
}
