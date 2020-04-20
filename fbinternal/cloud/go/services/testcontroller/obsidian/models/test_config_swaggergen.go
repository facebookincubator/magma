// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TestConfig test config
// swagger:model test_config
type TestConfig struct {

	// Minimum time to wait between upgrade script executions, in minutes
	// Required: true
	MinWaitTime int64 `json:"min_wait_time"`

	// URL of debian package repo
	// Required: true
	// Min Length: 1
	PackageRepo string `json:"package_repo"`

	// Release channel for package repo (stretch-beta, stretch-dev, etc)
	// Required: true
	// Min Length: 1
	ReleaseChannel string `json:"release_channel"`

	// Gateway ID of the target gateway
	// Required: true
	// Min Length: 1
	TargetGatewayID string `json:"target_gateway_id"`

	// ID of upgrade tier to bump when a new version is found in the package repo
	// Required: true
	// Min Length: 1
	TargetUpgradeTier string `json:"target_upgrade_tier"`

	// Command to execute test script
	// Required: true
	TestScriptCommand string `json:"test_script_command"`

	// Path of executable test script on testcontroller gateway
	// Required: true
	TestScriptPath string `json:"test_script_path"`

	// Gateway ID of the testcontroller gateway
	// Required: true
	// Min Length: 1
	TestcontrollerGatewayID string `json:"testcontroller_gateway_id"`
}

// Validate validates this test config
func (m *TestConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMinWaitTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePackageRepo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReleaseChannel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTargetGatewayID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTargetUpgradeTier(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTestScriptCommand(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTestScriptPath(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTestcontrollerGatewayID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TestConfig) validateMinWaitTime(formats strfmt.Registry) error {

	if err := validate.Required("min_wait_time", "body", int64(m.MinWaitTime)); err != nil {
		return err
	}

	return nil
}

func (m *TestConfig) validatePackageRepo(formats strfmt.Registry) error {

	if err := validate.RequiredString("package_repo", "body", string(m.PackageRepo)); err != nil {
		return err
	}

	if err := validate.MinLength("package_repo", "body", string(m.PackageRepo), 1); err != nil {
		return err
	}

	return nil
}

func (m *TestConfig) validateReleaseChannel(formats strfmt.Registry) error {

	if err := validate.RequiredString("release_channel", "body", string(m.ReleaseChannel)); err != nil {
		return err
	}

	if err := validate.MinLength("release_channel", "body", string(m.ReleaseChannel), 1); err != nil {
		return err
	}

	return nil
}

func (m *TestConfig) validateTargetGatewayID(formats strfmt.Registry) error {

	if err := validate.RequiredString("target_gateway_id", "body", string(m.TargetGatewayID)); err != nil {
		return err
	}

	if err := validate.MinLength("target_gateway_id", "body", string(m.TargetGatewayID), 1); err != nil {
		return err
	}

	return nil
}

func (m *TestConfig) validateTargetUpgradeTier(formats strfmt.Registry) error {

	if err := validate.RequiredString("target_upgrade_tier", "body", string(m.TargetUpgradeTier)); err != nil {
		return err
	}

	if err := validate.MinLength("target_upgrade_tier", "body", string(m.TargetUpgradeTier), 1); err != nil {
		return err
	}

	return nil
}

func (m *TestConfig) validateTestScriptCommand(formats strfmt.Registry) error {

	if err := validate.RequiredString("test_script_command", "body", string(m.TestScriptCommand)); err != nil {
		return err
	}

	return nil
}

func (m *TestConfig) validateTestScriptPath(formats strfmt.Registry) error {

	if err := validate.RequiredString("test_script_path", "body", string(m.TestScriptPath)); err != nil {
		return err
	}

	return nil
}

func (m *TestConfig) validateTestcontrollerGatewayID(formats strfmt.Registry) error {

	if err := validate.RequiredString("testcontroller_gateway_id", "body", string(m.TestcontrollerGatewayID)); err != nil {
		return err
	}

	if err := validate.MinLength("testcontroller_gateway_id", "body", string(m.TestcontrollerGatewayID), 1); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TestConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TestConfig) UnmarshalBinary(b []byte) error {
	var res TestConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
