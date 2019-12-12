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
  AggregationOperator,
  BinaryOperator,
  FunctionName,
  LabelOperator,
} from './PromQLTypes';

export interface Expression {
  toPromQL(): string;
}

export class Function implements Expression {
  name: FunctionName;
  arguments: Array<Expression>;

  constructor(name: FunctionName, args: Array<Expression>) {
    this.name = name;
    this.arguments = args;
  }

  toPromQL(): string {
    return (
      `${this.name}(` +
      this.arguments.map(arg => arg.toPromQL()).join(',') +
      ')'
    );
  }
}

export class InstantSelector implements Expression {
  selectorName: ?string;
  labels: ?Labels;

  constructor(selectorName: ?string, labels: ?Labels) {
    this.selectorName = selectorName;
    this.labels = labels || new Labels();
  }

  toPromQL(): string {
    return (
      (this.selectorName || '') + (this.labels ? this.labels.toPromQL() : '')
    );
  }
}

export class RangeSelector extends InstantSelector {
  range: Range;

  constructor(selector: InstantSelector, range: Range) {
    super(selector.selectorName, selector.labels);
    this.range = range;
  }

  toPromQL(): string {
    return `${super.toPromQL()}[${this.range.toString()}]`;
  }
}

export class Range {
  unit: string;
  value: number;

  constructor(value: number, unit: string) {
    this.unit = unit;
    this.value = value;
  }

  toString(): string {
    return this.value + this.unit;
  }
}

/**
 * The modifier methods of Labels mutate the underlying data
 * and return `this` to enable chaining on constructors.
 */
export class Labels {
  labels: Array<Label>;
  constructor(labels: ?Array<Label>) {
    this.labels = labels || [];
  }

  toPromQL(): string {
    if (this.labels.length === 0) {
      return '';
    }
    return '{' + this.labels.map(label => label.toString()).join(',') + '}';
  }

  addLabel(name: string, value: string, operator: LabelOperator) {
    this.labels.push(new Label(name, value, operator));
  }

  addEqual(name: string, value: string): Labels {
    this.labels.push(new Label(name, value, '='));
    return this;
  }

  addNotEqual(name: string, value: string): Labels {
    this.labels.push(new Label(name, value, '!='));
    return this;
  }

  addRegex(name: string, value: string): Labels {
    this.labels.push(new Label(name, value, '=~'));
    return this;
  }

  addNotRegex(name: string, value: string): Labels {
    this.labels.push(new Label(name, value, '!~'));
    return this;
  }

  setIndex(
    i: number,
    name: string,
    value: string,
    operator: ?LabelOperator,
  ): Labels {
    if (i >= 0 && i < this.len()) {
      this.labels[i].name = name;
      this.labels[i].value = value;
      this.labels[i].operator = operator || this.labels[i].operator;
    }
    return this;
  }

  remove(i: number): Labels {
    if (i >= 0 && i < this.len()) {
      this.labels.splice(i, 1);
    }
    return this;
  }

  len(): number {
    return this.labels.length;
  }

  copy(): Labels {
    const ret = new Labels();
    this.labels.forEach(label => {
      ret.addLabel(label.name, label.value, label.operator);
    });
    return ret;
  }
}

export class Label {
  name: string;
  value: string;
  operator: LabelOperator;

  constructor(name: string, value: string, operator: LabelOperator) {
    this.name = name;
    this.value = value;
    this.operator = operator;
  }

  toString(): string {
    return `${this.name}${this.operator}"${this.value}"`;
  }
}

export class Scalar implements Expression {
  value: number;

  constructor(value: number) {
    this.value = value;
  }

  toPromQL(): string {
    return this.value.toString();
  }
}

export class BinaryOperation implements Expression {
  lh: Expression;
  rh: Expression;
  operator: BinaryOperator;
  clause: ?Clause;

  constructor(
    lh: Expression,
    rh: Expression,
    operator: BinaryOperator,
    clause: ?Clause,
  ) {
    this.lh = lh;
    this.rh = rh;
    this.operator = operator;
    this.clause = clause;
  }

  toPromQL(): string {
    return (
      `${this.lh.toPromQL()} ${this.operator} ` +
      (this.clause ? this.clause.toString() + ' ' : '') +
      `${this.rh.toPromQL()}`
    );
  }
}

export class Clause {
  operator: string;
  labelList: Array<string>;

  constructor(operator: string, labelList: Array<string>) {
    this.operator = operator;
    this.labelList = labelList;
  }

  toString(): string {
    return `${this.operator} (` + this.labelList.join(',') + ')';
  }
}

export class AggregationOperation implements Expression {
  name: AggregationOperator;
  parameters: Array<Expression>;
  clause: ?Clause;

  constructor(
    name: AggregationOperator,
    parameters: Array<Expression>,
    clause: ?Clause,
  ) {
    this.name = name;
    (this.parameters = parameters), (this.clause = clause);
  }

  toPromQL(): string {
    return (
      `${this.name}(` +
      this.parameters.map(param => param.toPromQL()).join(',') +
      ')' +
      (this.clause ? ' ' + this.clause.toString() : '')
    );
  }
}
