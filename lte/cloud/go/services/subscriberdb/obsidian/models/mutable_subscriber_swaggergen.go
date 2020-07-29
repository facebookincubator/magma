// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	models1 "magma/lte/cloud/go/services/policydb/obsidian/models"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// MutableSubscriber Subset of subscriber field which are mutable
// swagger:model mutable_subscriber
type MutableSubscriber struct {

	// active apns
	ActiveApns ApnList `json:"active_apns,omitempty"`

	// Base names which are active for this subscriber
	ActiveBaseNames []models1.BaseName `json:"active_base_names,omitempty"`

	// Policies which are active for this subscriber
	ActivePolicies []models1.PolicyID `json:"active_policies,omitempty"`

	// id
	// Required: true
	ID models1.SubscriberID `json:"id"`

	// lte
	// Required: true
	Lte *LteSubscription `json:"lte"`

	// Name for the subscriber
	Name string `json:"name,omitempty"`

	// static ips
	StaticIps SubscriberStaticIps `json:"static_ips,omitempty"`
}

// Validate validates this mutable subscriber
func (m *MutableSubscriber) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActiveApns(formats); err != nil {
		res = append(res, err)
	}

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

	if err := m.validateStaticIps(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MutableSubscriber) validateActiveApns(formats strfmt.Registry) error {

	if swag.IsZero(m.ActiveApns) { // not required
		return nil
	}

	if err := m.ActiveApns.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("active_apns")
		}
		return err
	}

	return nil
}

func (m *MutableSubscriber) validateActiveBaseNames(formats strfmt.Registry) error {

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

func (m *MutableSubscriber) validateActivePolicies(formats strfmt.Registry) error {

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

func (m *MutableSubscriber) validateID(formats strfmt.Registry) error {

	if err := m.ID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("id")
		}
		return err
	}

	return nil
}

func (m *MutableSubscriber) validateLte(formats strfmt.Registry) error {

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

func (m *MutableSubscriber) validateStaticIps(formats strfmt.Registry) error {

	if swag.IsZero(m.StaticIps) { // not required
		return nil
	}

	if err := m.StaticIps.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("static_ips")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MutableSubscriber) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MutableSubscriber) UnmarshalBinary(b []byte) error {
	var res MutableSubscriber
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
