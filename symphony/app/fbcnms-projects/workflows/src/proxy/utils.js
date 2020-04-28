/**
 * Copyright 2004-present Facebook. All Rights Reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 *
 * @flow
 * @format
 */

import logging from '@fbcnms/logging';
import streamify from 'stream-array';
import {JSONPath} from 'jsonpath-plus';

const logger = logging.getLogger(module);

import type {ProxyRequest, Task} from '../types';

// Global prefix for taskdefs which can be used by all tenants.
export const GLOBAL_PREFIX: string = 'GLOBAL';

// This is used to separate tenant id from name in workflowdefs and taskdefs
export const INFIX_SEPARATOR: string = '___';

const SUB_WORKFLOW: string = 'SUB_WORKFLOW';
const DECISION: string = 'DECISION';
const FORK: string = 'FORK';
const SYSTEM_TASK_TYPES: Array<string> = [
  SUB_WORKFLOW,
  DECISION,
  'EVENT',
  'HTTP',
  FORK,
  'FORK_JOIN',
  'FORK_JOIN_DYNAMIC',
  'JOIN',
  'EXCLUSIVE_JOIN',
  'WAIT',
  'DYNAMIC',
  'LAMBDA',
  'TERMINATE',
  'KAFKA_PUBLISH',
  'DO_WHILE',
];

function isAllowedSystemTask(task: Task): boolean {
  return SYSTEM_TASK_TYPES.includes(task.type);
}

export function isSubworkflowTask(task: Task): boolean {
  return SUB_WORKFLOW === task.type;
}

export function isDecisionTask(task: Task): boolean {
  return DECISION === task.type;
}

export function isForkTask(task: Task): boolean {
  return FORK === task.type;
}

export function assertAllowedSystemTask(task: Task): void {
  if (!isAllowedSystemTask(task)) {
    logger.error(
      `Task type is not allowed: ` + ` in '${JSON.stringify(task)}'`,
    );
    // TODO create Exception class
    throw 'Task type is not allowed';
  }

  // assert decisions recursively
  if (isDecisionTask(task)) {
    const defaultCaseTasks = task.defaultCase ? task.defaultCase : [];
    for (const task of defaultCaseTasks) {
      assertAllowedSystemTask(task);
    }

    const decisionCaseIdToTasks: {[string]: Array<Task>} = task.decisionCases
      ? task.decisionCases
      : {};
    const values: Array<Array<Task>> = objectToValues(decisionCaseIdToTasks);
    for (const tasks of values) {
      for (const task of tasks) {
        assertAllowedSystemTask(task);
      }
    }
  }
}

export function objectToValues<A, B>(obj: {[key: A]: B}): Array<B> {
  return ((Object.values(obj): any): Array<B>);
}

export function withInfixSeparator(s: string): string {
  return s + INFIX_SEPARATOR;
}

export function addTenantIdPrefix(
  tenantId: string,
  objectWithName: {name: string},
): void {
  assertNameIsWithoutInfixSeparator(objectWithName);
  objectWithName.name = withInfixSeparator(tenantId) + objectWithName.name;
}

export function assertNameIsWithoutInfixSeparator(objectWithName: {
  name: string,
}): void {
  assertValueIsWithoutInfixSeparator(objectWithName.name);
}

// TODO: disallow ':'
export function assertValueIsWithoutInfixSeparator(value: string): void {
  if (value.indexOf(INFIX_SEPARATOR) > -1) {
    logger.error(`Value must not contain '${INFIX_SEPARATOR}' in '${value}'`);
    // TODO create Exception class
    throw 'Value must not contain INFIX_SEPARATOR';
  }
}

export function getTenantId(req: ProxyRequest): string {
  const tenantId: ?string = req.headers['x-auth-organization'];
  if (tenantId == null) {
    logger.error('x-auth-organization header not found');
    throw 'x-auth-organization header not found';
  }
  if (tenantId == GLOBAL_PREFIX) {
    logger.error(`Illegal name for TenantId: '${tenantId}'`);
    throw 'Illegal TenantId';
  }
  return tenantId;
}

export function createProxyOptionsBuffer(
  modifiedBody: any,
  req: ProxyRequest,
): any {
  // if request transformer returned modified body,
  // serialize it to new request stream. Original
  // request stream was already consumed. See `buffer` option
  // in node-http-proxy.
  if (typeof modifiedBody === 'object') {
    modifiedBody = JSON.stringify(modifiedBody);
  }
  if (typeof modifiedBody === 'string') {
    req.headers['content-length'] = modifiedBody.length;
    // create an array
    modifiedBody = [modifiedBody];
  } else {
    logger.error(`Unknown type: '${modifiedBody}'`);
    throw 'Unknown type';
  }
  return streamify(modifiedBody);
}

// Mass remove tenant prefix from json object.
// Setting allowGlobal to true implies that tasks are being processed,
// those starting with global prefix will not be touched.
export function removeTenantPrefix(
  tenantId: string,
  json: any,
  jsonPath: string,
  allowGlobal: boolean,
): void {
  const tenantWithInfixSeparator = withInfixSeparator(tenantId);
  const globalPrefix = withInfixSeparator(GLOBAL_PREFIX);
  const result = findValuesByJsonPath(json, jsonPath);
  for (const idx in result) {
    const item = result[idx];
    const prop = item.parent[item.parentProperty];
    if (allowGlobal && prop.indexOf(globalPrefix) == 0) {
      continue;
    }
    // expect tenantId prefix
    if (prop.indexOf(tenantWithInfixSeparator) != 0) {
      if (jsonPath.indexOf('taskDefName') != -1) {
        // Skipping tenant removal in taskDefName
        //  This is expected as some tasks do not require task def
        //  and might contain just some default
        continue;
      }

      logger.error(
        `Name must start with tenantId prefix` +
          `tenantId:'${tenantId}',json:'${json}',jsonPath:'${jsonPath}'` +
          `,item:'${item}'`,
      );
      throw 'Name must start with tenantId prefix'; // TODO create Exception class
    }
    // remove prefix
    item.parent[item.parentProperty] = prop.substr(
      tenantWithInfixSeparator.length,
    );
  }
}

// See removeTenantPrefix
export function removeTenantPrefixes(
  tenantId: string,
  json: any,
  jsonPathToAllowGlobal: {[string]: boolean},
): void {
  for (const key in jsonPathToAllowGlobal) {
    removeTenantPrefix(tenantId, json, key, jsonPathToAllowGlobal[key]);
  }
}

export function findValuesByJsonPath(
  json: any,
  path: string,
  resultType: string = 'all',
) {
  const result = JSONPath({json, path, resultType});
  logger.debug(`For path '${path}' found ${result.length} items`);
  return result;
}

export function anythingTo<T>(anything: any): T {
  if (anything != null) {
    return (anything: T);
  } else {
    throw 'Unexpected: value does not exist';
  }
}
