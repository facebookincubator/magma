// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// NetworkCellularConfigs Cellular configuration for a network
// swagger:model network_cellular_configs
type NetworkCellularConfigs struct {

	// epc
	Epc *NetworkEpcConfigs `json:"epc,omitempty"`

	// feg network id
	FegNetworkID string `json:"feg_network_id,omitempty"`

	// ran
	Ran *NetworkRanConfigs `json:"ran,omitempty"`
}

// Validate validates this network cellular configs
func (m *NetworkCellularConfigs) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEpc(formats); err != nil {
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

func (m *NetworkCellularConfigs) validateEpc(formats strfmt.Registry) error {

	if swag.IsZero(m.Epc) { // not required
		return nil
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

func (m *NetworkCellularConfigs) validateRan(formats strfmt.Registry) error {

	if swag.IsZero(m.Ran) { // not required
		return nil
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
func (m *NetworkCellularConfigs) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetworkCellularConfigs) UnmarshalBinary(b []byte) error {
	var res NetworkCellularConfigs
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
