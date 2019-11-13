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

// PolicyRule policy rule
// swagger:model policy_rule
type PolicyRule struct {

	// Subscribers which have been assigned this policy not as part of a base name
	AssignedSubscribers []SubscriberID `json:"assigned_subscribers,omitempty"`

	// id
	// Required: true
	ID PolicyID `json:"id"`

	// rule
	// Required: true
	Rule *PolicyRuleConfig `json:"rule"`
}

// Validate validates this policy rule
func (m *PolicyRule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAssignedSubscribers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRule(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PolicyRule) validateAssignedSubscribers(formats strfmt.Registry) error {

	if swag.IsZero(m.AssignedSubscribers) { // not required
		return nil
	}

	for i := 0; i < len(m.AssignedSubscribers); i++ {

		if err := m.AssignedSubscribers[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("assigned_subscribers" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *PolicyRule) validateID(formats strfmt.Registry) error {

	if err := m.ID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("id")
		}
		return err
	}

	return nil
}

func (m *PolicyRule) validateRule(formats strfmt.Registry) error {

	if err := validate.Required("rule", "body", m.Rule); err != nil {
		return err
	}

	if m.Rule != nil {
		if err := m.Rule.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rule")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PolicyRule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PolicyRule) UnmarshalBinary(b []byte) error {
	var res PolicyRule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
