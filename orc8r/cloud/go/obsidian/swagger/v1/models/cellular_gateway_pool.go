// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CellularGatewayPool Highly available gateway pool in an LTE network
// swagger:model cellular_gateway_pool
type CellularGatewayPool struct {

	// config
	// Required: true
	Config *CellularGatewayPoolConfigs `json:"config"`

	// gateway ids
	// Required: true
	// Unique: true
	GatewayIds []GatewayID `json:"gateway_ids"`

	// gateway pool id
	// Required: true
	GatewayPoolID GatewayPoolID `json:"gateway_pool_id"`

	// gateway pool name
	GatewayPoolName string `json:"gateway_pool_name,omitempty"`
}

// Validate validates this cellular gateway pool
func (m *CellularGatewayPool) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGatewayIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGatewayPoolID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CellularGatewayPool) validateConfig(formats strfmt.Registry) error {

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

func (m *CellularGatewayPool) validateGatewayIds(formats strfmt.Registry) error {

	if err := validate.Required("gateway_ids", "body", m.GatewayIds); err != nil {
		return err
	}

	if err := validate.UniqueItems("gateway_ids", "body", m.GatewayIds); err != nil {
		return err
	}

	for i := 0; i < len(m.GatewayIds); i++ {

		if err := m.GatewayIds[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gateway_ids" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *CellularGatewayPool) validateGatewayPoolID(formats strfmt.Registry) error {

	if err := m.GatewayPoolID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("gateway_pool_id")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CellularGatewayPool) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CellularGatewayPool) UnmarshalBinary(b []byte) error {
	var res CellularGatewayPool
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
