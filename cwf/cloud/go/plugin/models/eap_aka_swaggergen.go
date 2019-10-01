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

// EapAka eap_aka configuration
// swagger:model eapAka
type EapAka struct {

	// plmn ids
	PlmnIds []string `json:"plmn_ids"`

	// timeout
	Timeout *EapAkaTimeout `json:"timeout,omitempty"`
}

// Validate validates this eap aka
func (m *EapAka) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePlmnIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimeout(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EapAka) validatePlmnIds(formats strfmt.Registry) error {

	if swag.IsZero(m.PlmnIds) { // not required
		return nil
	}

	for i := 0; i < len(m.PlmnIds); i++ {

		if err := validate.MinLength("plmn_ids"+"."+strconv.Itoa(i), "body", string(m.PlmnIds[i]), 5); err != nil {
			return err
		}

		if err := validate.MaxLength("plmn_ids"+"."+strconv.Itoa(i), "body", string(m.PlmnIds[i]), 6); err != nil {
			return err
		}

		if err := validate.Pattern("plmn_ids"+"."+strconv.Itoa(i), "body", string(m.PlmnIds[i]), `^(\d{5,6})$`); err != nil {
			return err
		}

	}

	return nil
}

func (m *EapAka) validateTimeout(formats strfmt.Registry) error {

	if swag.IsZero(m.Timeout) { // not required
		return nil
	}

	if m.Timeout != nil {
		if err := m.Timeout.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("timeout")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EapAka) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EapAka) UnmarshalBinary(b []byte) error {
	var res EapAka
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// EapAkaTimeout eap aka timeout
// swagger:model EapAkaTimeout
type EapAkaTimeout struct {

	// challenge ms
	ChallengeMs uint32 `json:"challenge_ms,omitempty"`

	// error notification ms
	ErrorNotificationMs uint32 `json:"error_notification_ms,omitempty"`

	// session authenticated ms
	SessionAuthenticatedMs uint32 `json:"session_authenticated_ms,omitempty"`

	// session ms
	SessionMs uint32 `json:"session_ms,omitempty"`
}

// Validate validates this eap aka timeout
func (m *EapAkaTimeout) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EapAkaTimeout) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EapAkaTimeout) UnmarshalBinary(b []byte) error {
	var res EapAkaTimeout
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
