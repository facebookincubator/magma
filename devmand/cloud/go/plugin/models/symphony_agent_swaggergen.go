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

// SymphonyAgent Full description of a Symphony agent
// swagger:model symphony_agent
type SymphonyAgent struct {

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

	// status
	Status *models3.GatewayStatus `json:"status,omitempty"`

	// tier
	// Required: true
	Tier models3.TierID `json:"tier"`
}

// Validate validates this symphony agent
func (m *SymphonyAgent) Validate(formats strfmt.Registry) error {
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

	if err := m.validateStatus(formats); err != nil {
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

func (m *SymphonyAgent) validateDescription(formats strfmt.Registry) error {

	if err := m.Description.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("description")
		}
		return err
	}

	return nil
}

func (m *SymphonyAgent) validateDevice(formats strfmt.Registry) error {

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

func (m *SymphonyAgent) validateID(formats strfmt.Registry) error {

	if err := m.ID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("id")
		}
		return err
	}

	return nil
}

func (m *SymphonyAgent) validateMagmad(formats strfmt.Registry) error {

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

func (m *SymphonyAgent) validateManagedDevices(formats strfmt.Registry) error {

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

func (m *SymphonyAgent) validateName(formats strfmt.Registry) error {

	if err := m.Name.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("name")
		}
		return err
	}

	return nil
}

func (m *SymphonyAgent) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if m.Status != nil {
		if err := m.Status.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

func (m *SymphonyAgent) validateTier(formats strfmt.Registry) error {

	if err := m.Tier.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("tier")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SymphonyAgent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SymphonyAgent) UnmarshalBinary(b []byte) error {
	var res SymphonyAgent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
