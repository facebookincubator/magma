/**
 * Copyright (c) 2016-present, Facebook, Inc.
 * All rights reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree. An additional grant
 * of patent rights can be found in the PATENTS file in the same directory.
 */
#include "SessionRules.h"

namespace magma {

SessionRules::SessionRules(StaticRuleStore &static_rule_ref):
  static_rules_(static_rule_ref)
{
}

bool SessionRules::get_charging_key_for_rule_id(
  const std::string &rule_id,
  uint32_t *charging_key)
{
  // first check dynamic rules and then static rules
  if (dynamic_rules_.get_charging_key_for_rule_id(rule_id, charging_key)) {
    return true;
  }
  if (static_rules_.get_charging_key_for_rule_id(rule_id, charging_key)) {
    return true;
  }
  return false;
}

bool SessionRules::get_monitoring_key_for_rule_id(
  const std::string &rule_id,
  std::string *monitoring_key)
{
  // first check dynamic rules and then static rules
  if (dynamic_rules_.get_monitoring_key_for_rule_id(rule_id, monitoring_key)) {
    return true;
  }
  if (static_rules_.get_monitoring_key_for_rule_id(rule_id, monitoring_key)) {
    return true;
  }
  return false;
}

void SessionRules::insert_dynamic_rule(const PolicyRule &rule)
{
  dynamic_rules_.insert_rule(rule);
}

bool SessionRules::remove_dynamic_rule(
  const std::string &rule_id,
  PolicyRule *rule_out)
{
  return dynamic_rules_.remove_rule(rule_id, rule_out);
}

/**
 * For the charging key, get any applicable rules from the static rule set
 * and the dynamic rule set
 */
void SessionRules::add_rules_to_action(
  ServiceAction &action,
  uint32_t charging_key)
{
  static_rules_.get_rule_ids_for_charging_key(
    charging_key, *action.get_mutable_rule_ids());
  dynamic_rules_.get_rule_definitions_for_charging_key(
    charging_key, *action.get_mutable_rule_definitions());
}

void SessionRules::add_rules_to_action(
  ServiceAction &action,
  std::string monitoring_key)
{
  static_rules_.get_rule_ids_for_monitoring_key(
    monitoring_key, *action.get_mutable_rule_ids());
  dynamic_rules_.get_rule_definitions_for_monitoring_key(
    monitoring_key, *action.get_mutable_rule_definitions());
}

} // namespace magma
