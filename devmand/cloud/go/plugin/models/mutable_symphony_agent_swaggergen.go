// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	models2 "magma/orc8r/cloud/go/models"
	models3 "magma/orc8r/cloud/go/pluginimpl/models"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// MutableSymphonyAgent Description of a Symphony agent with the read-only fields omitted
// swagger:model mutable_symphony_agent
type MutableSymphonyAgent struct {

	// description
	// Required: true
	Description models2.GatewayDescription `json:"description"`

	// device
	// Required: true
	Device *models3.GatewayDevice `json:"device"`

	// id
	// Required: true
	ID models2.GatewayID `json:"id"`

	// magmad
	// Required: true
	Magmad *models3.MagmadGatewayConfigs `json:"magmad"`

	// managed devices
	// Required: true
	ManagedDevices ManagedDevices `json:"managed_devices"`

	// name
	// Required: true
	Name models2.GatewayName `json:"name"`

	// tier
	// Required: true
	Tier models3.TierID `json:"tier"`
}

// Validate validates this mutable symphony agent
func (m *MutableSymphonyAgent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDevice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMagmad(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateManagedDevices(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTier(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MutableSymphonyAgent) validateDescription(formats strfmt.Registry) error {

	if err := m.Description.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("description")
		}
		return err
	}

	return nil
}

func (m *MutableSymphonyAgent) validateDevice(formats strfmt.Registry) error {

	if err := validate.Required("device", "body", m.Device); err != nil {
		return err
	}

	if m.Device != nil {
		if err := m.Device.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("device")
			}
			return err
		}
	}

	return nil
}

func (m *MutableSymphonyAgent) validateID(formats strfmt.Registry) error {

	if err := m.ID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("id")
		}
		return err
	}

	return nil
}

func (m *MutableSymphonyAgent) validateMagmad(formats strfmt.Registry) error {

	if err := validate.Required("magmad", "body", m.Magmad); err != nil {
		return err
	}

	if m.Magmad != nil {
		if err := m.Magmad.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("magmad")
			}
			return err
		}
	}

	return nil
}

func (m *MutableSymphonyAgent) validateManagedDevices(formats strfmt.Registry) error {

	if err := validate.Required("managed_devices", "body", m.ManagedDevices); err != nil {
		return err
	}

	if err := m.ManagedDevices.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("managed_devices")
		}
		return err
	}

	return nil
}

func (m *MutableSymphonyAgent) validateName(formats strfmt.Registry) error {

	if err := m.Name.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("name")
		}
		return err
	}

	return nil
}

func (m *MutableSymphonyAgent) validateTier(formats strfmt.Registry) error {

	if err := m.Tier.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("tier")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MutableSymphonyAgent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MutableSymphonyAgent) UnmarshalBinary(b []byte) error {
	var res MutableSymphonyAgent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
