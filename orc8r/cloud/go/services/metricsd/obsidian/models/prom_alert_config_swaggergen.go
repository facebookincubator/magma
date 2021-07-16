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

// PromAlertConfig prom alert config
//
// swagger:model prom_alert_config
type PromAlertConfig struct {

	// alert
	// Required: true
	Alert *string `json:"alert"`

	// annotations
	Annotations PromAlertLabels `json:"annotations,omitempty"`

	// expr
	// Required: true
	Expr *string `json:"expr"`

	// for
	For string `json:"for,omitempty"`

	// labels
	Labels PromAlertLabels `json:"labels,omitempty"`
}

// Validate validates this prom alert config
func (m *PromAlertConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAlert(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAnnotations(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExpr(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabels(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PromAlertConfig) validateAlert(formats strfmt.Registry) error {

	if err := validate.Required("alert", "body", m.Alert); err != nil {
		return err
	}

	return nil
}

func (m *PromAlertConfig) validateAnnotations(formats strfmt.Registry) error {
	if swag.IsZero(m.Annotations) { // not required
		return nil
	}

	if m.Annotations != nil {
		if err := m.Annotations.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("annotations")
			}
			return err
		}
	}

	return nil
}

func (m *PromAlertConfig) validateExpr(formats strfmt.Registry) error {

	if err := validate.Required("expr", "body", m.Expr); err != nil {
		return err
	}

	return nil
}

func (m *PromAlertConfig) validateLabels(formats strfmt.Registry) error {
	if swag.IsZero(m.Labels) { // not required
		return nil
	}

	if m.Labels != nil {
		if err := m.Labels.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("labels")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this prom alert config based on the context it is used
func (m *PromAlertConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAnnotations(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLabels(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PromAlertConfig) contextValidateAnnotations(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Annotations.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("annotations")
		}
		return err
	}

	return nil
}

func (m *PromAlertConfig) contextValidateLabels(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Labels.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("labels")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PromAlertConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PromAlertConfig) UnmarshalBinary(b []byte) error {
	var res PromAlertConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
