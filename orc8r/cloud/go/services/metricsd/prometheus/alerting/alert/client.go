/*
 * Copyright (c) Facebook, Inc. and its affiliates.
 * All rights reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 */

package alert

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/prometheus/prometheus/pkg/rulefmt"
	"gopkg.in/yaml.v2"
)

const (
	rulesFilePostfix = "_rules.yml"
)

// PrometheusAlertClient provides thread-safe methods for writing, reading,
// and modifying alert configuration files
type PrometheusAlertClient interface {
	ValidateRule(rule rulefmt.Rule) error
	RuleExists(rulename, networkID string) bool
	WriteRule(rule rulefmt.Rule, networkID string) error
	UpdateRule(rule rulefmt.Rule, networkID string) error
	ReadRules(ruleName string, networkID string) ([]rulefmt.Rule, error)
	DeleteRule(ruleName string, networkID string) error
	BulkUpdateRules(rules []rulefmt.Rule, networkID string) (BulkUpdateResults, error)
}

type client struct {
	fileLocks *FileLocker
	rulesDir  string
}

func NewClient(rulesDir string) (PrometheusAlertClient, error) {
	fileLocks, err := NewFileLocker(rulesDir)
	if err != nil {
		return nil, err
	}
	return &client{
		fileLocks: fileLocks,
		rulesDir:  rulesDir,
	}, nil
}

// ValidateRule checks that a new alert rule is a valid specification
func (c *client) ValidateRule(rule rulefmt.Rule) error {
	errs := rule.Validate()
	if len(errs) != 0 {
		return fmt.Errorf("invalid rule: %v", errs)
	}
	return nil
}

func (c *client) RuleExists(rulename, networkID string) bool {
	filename := makeFilename(networkID, c.rulesDir)

	c.fileLocks.Lock(filename)
	defer c.fileLocks.Unlock(filename)

	ruleFile, err := c.initializeRuleFile(filename, networkID)
	if err != nil {
		return false
	}
	return ruleFile.GetRule(rulename) != nil
}

// WriteRule takes an alerting rule and writes it to the rules file for the
// given networkID
func (c *client) WriteRule(rule rulefmt.Rule, networkID string) error {
	filename := makeFilename(networkID, c.rulesDir)

	c.fileLocks.Lock(filename)
	defer c.fileLocks.Unlock(filename)

	ruleFile, err := c.initializeRuleFile(filename, networkID)
	if err != nil {
		return err
	}
	ruleFile.AddRule(rule)

	err = c.writeRuleFile(ruleFile, filename)
	if err != nil {
		return err
	}
	return nil
}

func (c *client) UpdateRule(rule rulefmt.Rule, networkID string) error {
	filename := makeFilename(networkID, c.rulesDir)

	c.fileLocks.Lock(filename)
	defer c.fileLocks.Unlock(filename)

	ruleFile, err := c.initializeRuleFile(filename, networkID)
	if err != nil {
		return err
	}

	err = SecureRule(&rule, networkID)
	if err != nil {
		return err
	}

	err = ruleFile.ReplaceRule(rule)
	if err != nil {
		return err
	}

	err = c.writeRuleFile(ruleFile, filename)
	if err != nil {
		return err
	}
	return nil
}

func (c *client) ReadRules(ruleName string, networkID string) ([]rulefmt.Rule, error) {
	filename := makeFilename(networkID, c.rulesDir)
	c.fileLocks.RLock(filename)
	defer c.fileLocks.RUnlock(filename)

	ruleFile, err := c.readRuleFile(makeFilename(networkID, c.rulesDir))
	if err != nil {
		return []rulefmt.Rule{}, err
	}
	if ruleName == "" {
		return ruleFile.Rules(), nil
	}
	foundRule := ruleFile.GetRule(ruleName)
	if foundRule == nil {
		return nil, fmt.Errorf("rule %s not found", ruleName)
	}
	return []rulefmt.Rule{*foundRule}, nil
}

func (c *client) DeleteRule(ruleName string, networkID string) error {
	filename := makeFilename(networkID, c.rulesDir)
	c.fileLocks.Lock(filename)
	defer c.fileLocks.Unlock(filename)

	ruleFile, err := c.readRuleFile(filename)
	if err != nil {
		return err
	}

	err = ruleFile.DeleteRule(ruleName)
	if err != nil {
		return err
	}

	err = c.writeRuleFile(ruleFile, filename)
	if err != nil {
		return err
	}
	return nil
}

func (c *client) BulkUpdateRules(rules []rulefmt.Rule, networkID string) (BulkUpdateResults, error) {
	filename := makeFilename(networkID, c.rulesDir)
	c.fileLocks.Lock(filename)
	defer c.fileLocks.Unlock(filename)

	ruleFile, err := c.readRuleFile(filename)
	if err != nil {
		return BulkUpdateResults{}, err
	}

	results := NewBulkUpdateResults()
	for _, newRule := range rules {
		ruleName := newRule.Alert
		err := SecureRule(&newRule, networkID)
		if err != nil {
			results.Errors[ruleName] = err
			continue
		}
		if ruleFile.GetRule(ruleName) != nil {
			err := ruleFile.ReplaceRule(newRule)
			if err != nil {
				results.Errors[ruleName] = err
			} else {
				results.Statuses[ruleName] = "updated"
			}
		} else {
			ruleFile.AddRule(newRule)
			results.Statuses[ruleName] = "created"
		}
	}

	err = c.writeRuleFile(ruleFile, filename)
	if err != nil {
		return results, err
	}
	return results, nil
}

func (c *client) writeRuleFile(ruleFile *File, filename string) error {
	yamlFile, err := yaml.Marshal(ruleFile)
	err = ioutil.WriteFile(filename, yamlFile, 0666)
	if err != nil {
		return fmt.Errorf("error writing rules file: %v\n", yamlFile)
	}
	return nil
}

func (c *client) initializeRuleFile(filename, networkID string) (*File, error) {
	if _, err := os.Stat(filename); err == nil {
		file, err := c.readRuleFile(filename)
		if err != nil {
			return nil, err
		}
		return file, nil
	}
	return NewFile(networkID), nil
}

func (c *client) readRuleFile(requestedFile string) (*File, error) {
	ruleFile := File{}
	file, err := ioutil.ReadFile(requestedFile)
	if err != nil {
		return &File{}, fmt.Errorf("error reading rules files: %v", err)
	}
	err = yaml.Unmarshal(file, &ruleFile)
	return &ruleFile, err
}

type BulkUpdateResults struct {
	Errors   map[string]error
	Statuses map[string]string
}

func NewBulkUpdateResults() BulkUpdateResults {
	return BulkUpdateResults{
		Errors:   make(map[string]error, 0),
		Statuses: make(map[string]string, 0),
	}
}

func (r BulkUpdateResults) String() string {
	str := strings.Builder{}
	if len(r.Errors) > 0 {
		str.WriteString("Errors: \n")
		for name, err := range r.Errors {
			str.WriteString(fmt.Sprintf("\t%s: %s\n", name, err))
		}
	}
	if len(r.Statuses) > 0 {
		str.WriteString("Statuses: \n")
		for name, status := range r.Statuses {
			str.WriteString(fmt.Sprintf("\t%s: %s\n", name, status))
		}
	}
	return str.String()
}

func makeFilename(networkID, path string) string {
	return path + "/" + networkID + rulesFilePostfix
}
