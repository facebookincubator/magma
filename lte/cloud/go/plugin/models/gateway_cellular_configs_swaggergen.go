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

// GatewayCellularConfigs Cellular configuration for LTE gateway
// swagger:model gateway_cellular_configs
type GatewayCellularConfigs struct {

	// epc
	// Required: true
	Epc *GatewayEpcConfigs `json:"epc"`

	// non eps service
	NonEpsService *GatewayNonEpsConfigs `json:"non_eps_service,omitempty"`

	// ran
	// Required: true
	Ran *GatewayRanConfigs `json:"ran"`
}

// Validate validates this gateway cellular configs
func (m *GatewayCellularConfigs) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEpc(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNonEpsService(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRan(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GatewayCellularConfigs) validateEpc(formats strfmt.Registry) error {

	if err := validate.Required("epc", "body", m.Epc); err != nil {
		return err
	}

	if m.Epc != nil {
		if err := m.Epc.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("epc")
			}
			return err
		}
	}

	return nil
}

func (m *GatewayCellularConfigs) validateNonEpsService(formats strfmt.Registry) error {

	if swag.IsZero(m.NonEpsService) { // not required
		return nil
	}

	if m.NonEpsService != nil {
		if err := m.NonEpsService.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("non_eps_service")
			}
			return err
		}
	}

	return nil
}

func (m *GatewayCellularConfigs) validateRan(formats strfmt.Registry) error {

	if err := validate.Required("ran", "body", m.Ran); err != nil {
		return err
	}

	if m.Ran != nil {
		if err := m.Ran.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ran")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GatewayCellularConfigs) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GatewayCellularConfigs) UnmarshalBinary(b []byte) error {
	var res GatewayCellularConfigs
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
