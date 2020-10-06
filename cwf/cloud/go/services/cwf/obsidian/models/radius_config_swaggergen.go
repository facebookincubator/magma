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

// RadiusConfig built-in radius server configuration
// swagger:model radiusConfig
type RadiusConfig struct {

	// d a e addr
	// Pattern: [0-9a-f\:\.]*(:[0-9]{1,5})?
	DAEAddr string `json:"DAE_addr,omitempty"`

	// acct addr
	// Pattern: [0-9a-f\:\.]*(:[0-9]{1,5})?
	AcctAddr string `json:"acct_addr,omitempty"`

	// auth addr
	// Pattern: [0-9a-f\:\.]*(:[0-9]{1,5})?
	AuthAddr string `json:"auth_addr,omitempty"`

	// network
	Network string `json:"network,omitempty"`

	// secret
	// Format: byte
	Secret strfmt.Base64 `json:"secret,omitempty"`
}

// Validate validates this radius config
func (m *RadiusConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDAEAddr(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAcctAddr(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAuthAddr(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecret(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RadiusConfig) validateDAEAddr(formats strfmt.Registry) error {

	if swag.IsZero(m.DAEAddr) { // not required
		return nil
	}

	if err := validate.Pattern("DAE_addr", "body", string(m.DAEAddr), `[0-9a-f\:\.]*(:[0-9]{1,5})?`); err != nil {
		return err
	}

	return nil
}

func (m *RadiusConfig) validateAcctAddr(formats strfmt.Registry) error {

	if swag.IsZero(m.AcctAddr) { // not required
		return nil
	}

	if err := validate.Pattern("acct_addr", "body", string(m.AcctAddr), `[0-9a-f\:\.]*(:[0-9]{1,5})?`); err != nil {
		return err
	}

	return nil
}

func (m *RadiusConfig) validateAuthAddr(formats strfmt.Registry) error {

	if swag.IsZero(m.AuthAddr) { // not required
		return nil
	}

	if err := validate.Pattern("auth_addr", "body", string(m.AuthAddr), `[0-9a-f\:\.]*(:[0-9]{1,5})?`); err != nil {
		return err
	}

	return nil
}

func (m *RadiusConfig) validateSecret(formats strfmt.Registry) error {

	if swag.IsZero(m.Secret) { // not required
		return nil
	}

	// Format "byte" (base64 string) is already validated when unmarshalled

	return nil
}

// MarshalBinary interface implementation
func (m *RadiusConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RadiusConfig) UnmarshalBinary(b []byte) error {
	var res RadiusConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
