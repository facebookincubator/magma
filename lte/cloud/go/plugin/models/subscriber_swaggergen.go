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

// Subscriber subscriber
// swagger:model subscriber
type Subscriber struct {

	// Base names which are active for this subscriber
	ActiveBaseNames []BaseName `json:"active_base_names,omitempty"`

	// Policies which are active for this subscriber
	ActivePolicies []PolicyID `json:"active_policies,omitempty"`

	// id
	// Required: true
	ID SubscriberID `json:"id"`

	// lte
	// Required: true
	Lte *LteSubscription `json:"lte"`
}

// Validate validates this subscriber
func (m *Subscriber) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActiveBaseNames(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateActivePolicies(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLte(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Subscriber) validateActiveBaseNames(formats strfmt.Registry) error {

	if swag.IsZero(m.ActiveBaseNames) { // not required
		return nil
	}

	for i := 0; i < len(m.ActiveBaseNames); i++ {

		if err := m.ActiveBaseNames[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("active_base_names" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *Subscriber) validateActivePolicies(formats strfmt.Registry) error {

	if swag.IsZero(m.ActivePolicies) { // not required
		return nil
	}

	for i := 0; i < len(m.ActivePolicies); i++ {

		if err := m.ActivePolicies[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("active_policies" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *Subscriber) validateID(formats strfmt.Registry) error {

	if err := m.ID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("id")
		}
		return err
	}

	return nil
}

func (m *Subscriber) validateLte(formats strfmt.Registry) error {

	if err := validate.Required("lte", "body", m.Lte); err != nil {
		return err
	}

	if m.Lte != nil {
		if err := m.Lte.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lte")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Subscriber) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Subscriber) UnmarshalBinary(b []byte) error {
	var res Subscriber
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
