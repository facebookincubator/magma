// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// MutableSubscriber Subset of subscriber field which are mutable
//
// swagger:model mutable_subscriber
type MutableSubscriber struct {

	// active apns
	ActiveAPNS APNList `json:"active_apns,omitempty"`

	// active base names
	ActiveBaseNames BaseNames `json:"active_base_names,omitempty"`

	// active policies
	ActivePolicies PolicyIds `json:"active_policies,omitempty"`

	// active policies by apn
	ActivePoliciesByAPN PolicyIdsByAPN `json:"active_policies_by_apn,omitempty"`

	// id
	// Required: true
	ID SubscriberID `json:"id"`

	// lte
	// Required: true
	LTE *LTESubscription `json:"lte"`

	// Name for the subscriber
	// Example: Jane Doe
	Name string `json:"name,omitempty"`

	// static ips
	StaticIps SubscriberStaticIps `json:"static_ips,omitempty"`
}

// Validate validates this mutable subscriber
func (m *MutableSubscriber) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActiveAPNS(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateActiveBaseNames(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateActivePolicies(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateActivePoliciesByAPN(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLTE(formats); err != nil {
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

func (m *MutableSubscriber) validateActiveAPNS(formats strfmt.Registry) error {
	if swag.IsZero(m.ActiveAPNS) { // not required
		return nil
	}

	if err := m.ActiveAPNS.Validate(formats); err != nil {
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

	if err := m.ActiveBaseNames.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("active_base_names")
		}
		return err
	}

	return nil
}

func (m *MutableSubscriber) validateActivePolicies(formats strfmt.Registry) error {
	if swag.IsZero(m.ActivePolicies) { // not required
		return nil
	}

	if err := m.ActivePolicies.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("active_policies")
		}
		return err
	}

	return nil
}

func (m *MutableSubscriber) validateActivePoliciesByAPN(formats strfmt.Registry) error {
	if swag.IsZero(m.ActivePoliciesByAPN) { // not required
		return nil
	}

	if m.ActivePoliciesByAPN != nil {
		if err := m.ActivePoliciesByAPN.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("active_policies_by_apn")
			}
			return err
		}
	}

	return nil
}

func (m *MutableSubscriber) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", SubscriberID(m.ID)); err != nil {
		return err
	}

	if err := m.ID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("id")
		}
		return err
	}

	return nil
}

func (m *MutableSubscriber) validateLTE(formats strfmt.Registry) error {

	if err := validate.Required("lte", "body", m.LTE); err != nil {
		return err
	}

	if m.LTE != nil {
		if err := m.LTE.Validate(formats); err != nil {
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

	if m.StaticIps != nil {
		if err := m.StaticIps.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("static_ips")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this mutable subscriber based on the context it is used
func (m *MutableSubscriber) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateActiveAPNS(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateActiveBaseNames(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateActivePolicies(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateActivePoliciesByAPN(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLTE(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStaticIps(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MutableSubscriber) contextValidateActiveAPNS(ctx context.Context, formats strfmt.Registry) error {

	if err := m.ActiveAPNS.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("active_apns")
		}
		return err
	}

	return nil
}

func (m *MutableSubscriber) contextValidateActiveBaseNames(ctx context.Context, formats strfmt.Registry) error {

	if err := m.ActiveBaseNames.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("active_base_names")
		}
		return err
	}

	return nil
}

func (m *MutableSubscriber) contextValidateActivePolicies(ctx context.Context, formats strfmt.Registry) error {

	if err := m.ActivePolicies.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("active_policies")
		}
		return err
	}

	return nil
}

func (m *MutableSubscriber) contextValidateActivePoliciesByAPN(ctx context.Context, formats strfmt.Registry) error {

	if err := m.ActivePoliciesByAPN.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("active_policies_by_apn")
		}
		return err
	}

	return nil
}

func (m *MutableSubscriber) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := m.ID.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("id")
		}
		return err
	}

	return nil
}

func (m *MutableSubscriber) contextValidateLTE(ctx context.Context, formats strfmt.Registry) error {

	if m.LTE != nil {
		if err := m.LTE.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lte")
			}
			return err
		}
	}

	return nil
}

func (m *MutableSubscriber) contextValidateStaticIps(ctx context.Context, formats strfmt.Registry) error {

	if err := m.StaticIps.ContextValidate(ctx, formats); err != nil {
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
