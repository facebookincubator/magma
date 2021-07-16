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

// SubscriberConfig subscriber config
//
// swagger:model subscriber_config
type SubscriberConfig struct {

	// lte
	// Required: true
	LTE *LTESubscription `json:"lte"`

	// static ips
	StaticIps SubscriberStaticIps `json:"static_ips,omitempty"`
}

// Validate validates this subscriber config
func (m *SubscriberConfig) Validate(formats strfmt.Registry) error {
	var res []error

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

func (m *SubscriberConfig) validateLTE(formats strfmt.Registry) error {

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

func (m *SubscriberConfig) validateStaticIps(formats strfmt.Registry) error {
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

// ContextValidate validate this subscriber config based on the context it is used
func (m *SubscriberConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

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

func (m *SubscriberConfig) contextValidateLTE(ctx context.Context, formats strfmt.Registry) error {

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

func (m *SubscriberConfig) contextValidateStaticIps(ctx context.Context, formats strfmt.Registry) error {

	if err := m.StaticIps.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("static_ips")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SubscriberConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SubscriberConfig) UnmarshalBinary(b []byte) error {
	var res SubscriberConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
