// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	models3 "magma/orc8r/cloud/go/models"
	models4 "magma/orc8r/cloud/go/services/orchestrator/obsidian/models"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// MutableLteGateway LTE gateway object with read-only fields omitted
// swagger:model mutable_lte_gateway
type MutableLteGateway struct {

	// apn resources
	// Required: true
	ApnResources ApnResources `json:"apn_resources"`

	// cellular
	// Required: true
	Cellular *GatewayCellularConfigs `json:"cellular"`

	// connected enodeb serials
	// Required: true
	ConnectedEnodebSerials EnodebSerials `json:"connected_enodeb_serials"`

	// description
	// Required: true
	Description models3.GatewayDescription `json:"description"`

	// device
	// Required: true
	Device *models4.GatewayDevice `json:"device"`

	// id
	// Required: true
	ID models3.GatewayID `json:"id"`

	// magmad
	// Required: true
	Magmad *models4.MagmadGatewayConfigs `json:"magmad"`

	// name
	// Required: true
	Name models3.GatewayName `json:"name"`

	// tier
	// Required: true
	Tier models4.TierID `json:"tier"`
}

// Validate validates this mutable lte gateway
func (m *MutableLteGateway) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateApnResources(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCellular(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConnectedEnodebSerials(formats); err != nil {
		res = append(res, err)
	}

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

func (m *MutableLteGateway) validateApnResources(formats strfmt.Registry) error {

	if err := m.ApnResources.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("apn_resources")
		}
		return err
	}

	return nil
}

func (m *MutableLteGateway) validateCellular(formats strfmt.Registry) error {

	if err := validate.Required("cellular", "body", m.Cellular); err != nil {
		return err
	}

	if m.Cellular != nil {
		if err := m.Cellular.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cellular")
			}
			return err
		}
	}

	return nil
}

func (m *MutableLteGateway) validateConnectedEnodebSerials(formats strfmt.Registry) error {

	if err := validate.Required("connected_enodeb_serials", "body", m.ConnectedEnodebSerials); err != nil {
		return err
	}

	if err := m.ConnectedEnodebSerials.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("connected_enodeb_serials")
		}
		return err
	}

	return nil
}

func (m *MutableLteGateway) validateDescription(formats strfmt.Registry) error {

	if err := m.Description.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("description")
		}
		return err
	}

	return nil
}

func (m *MutableLteGateway) validateDevice(formats strfmt.Registry) error {

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

func (m *MutableLteGateway) validateID(formats strfmt.Registry) error {

	if err := m.ID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("id")
		}
		return err
	}

	return nil
}

func (m *MutableLteGateway) validateMagmad(formats strfmt.Registry) error {

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

func (m *MutableLteGateway) validateName(formats strfmt.Registry) error {

	if err := m.Name.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("name")
		}
		return err
	}

	return nil
}

func (m *MutableLteGateway) validateTier(formats strfmt.Registry) error {

	if err := m.Tier.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("tier")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MutableLteGateway) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MutableLteGateway) UnmarshalBinary(b []byte) error {
	var res MutableLteGateway
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
