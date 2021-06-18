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

// ENODEB Representation of an enodeB
// swagger:model enodeb
type ENODEB struct {

	// attached gateway id
	// Read Only: true
	AttachedGatewayID string `json:"attached_gateway_id,omitempty"`

	// config
	// Required: true
	Config *ENODEBConfiguration `json:"config"`

	// description
	Description string `json:"description,omitempty"`

	// enodeb config
	ENODEBConfig *ENODEBConfig `json:"enodeb_config,omitempty"`

	// name
	// Required: true
	// Min Length: 1
	Name string `json:"name"`

	// serial
	// Required: true
	// Min Length: 1
	Serial string `json:"serial"`
}

// Validate validates this enodeb
func (m *ENODEB) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateENODEBConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSerial(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ENODEB) validateConfig(formats strfmt.Registry) error {

	if err := validate.Required("config", "body", m.Config); err != nil {
		return err
	}

	if m.Config != nil {
		if err := m.Config.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("config")
			}
			return err
		}
	}

	return nil
}

func (m *ENODEB) validateENODEBConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.ENODEBConfig) { // not required
		return nil
	}

	if m.ENODEBConfig != nil {
		if err := m.ENODEBConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("enodeb_config")
			}
			return err
		}
	}

	return nil
}

func (m *ENODEB) validateName(formats strfmt.Registry) error {

	if err := validate.RequiredString("name", "body", string(m.Name)); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(m.Name), 1); err != nil {
		return err
	}

	return nil
}

func (m *ENODEB) validateSerial(formats strfmt.Registry) error {

	if err := validate.RequiredString("serial", "body", string(m.Serial)); err != nil {
		return err
	}

	if err := validate.MinLength("serial", "body", string(m.Serial), 1); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ENODEB) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ENODEB) UnmarshalBinary(b []byte) error {
	var res ENODEB
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}