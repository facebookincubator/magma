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

// PaginatedSubscribers Page of subscribers
//
// swagger:model paginated_subscribers
type PaginatedSubscribers struct {

	// next page token
	// Required: true
	NextPageToken *NextPageToken `json:"next_page_token"`

	// subscribers
	// Required: true
	Subscribers map[string]*Subscriber `json:"subscribers"`

	// estimated total number of subscriber entries
	// Example: 10
	// Required: true
	TotalCount int64 `json:"total_count"`
}

// Validate validates this paginated subscribers
func (m *PaginatedSubscribers) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNextPageToken(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubscribers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTotalCount(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PaginatedSubscribers) validateNextPageToken(formats strfmt.Registry) error {

	if err := validate.Required("next_page_token", "body", m.NextPageToken); err != nil {
		return err
	}

	if err := validate.Required("next_page_token", "body", m.NextPageToken); err != nil {
		return err
	}

	if m.NextPageToken != nil {
		if err := m.NextPageToken.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("next_page_token")
			}
			return err
		}
	}

	return nil
}

func (m *PaginatedSubscribers) validateSubscribers(formats strfmt.Registry) error {

	if err := validate.Required("subscribers", "body", m.Subscribers); err != nil {
		return err
	}

	for k := range m.Subscribers {

		if err := validate.Required("subscribers"+"."+k, "body", m.Subscribers[k]); err != nil {
			return err
		}
		if val, ok := m.Subscribers[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *PaginatedSubscribers) validateTotalCount(formats strfmt.Registry) error {

	if err := validate.Required("total_count", "body", int64(m.TotalCount)); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this paginated subscribers based on the context it is used
func (m *PaginatedSubscribers) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNextPageToken(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSubscribers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PaginatedSubscribers) contextValidateNextPageToken(ctx context.Context, formats strfmt.Registry) error {

	if m.NextPageToken != nil {
		if err := m.NextPageToken.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("next_page_token")
			}
			return err
		}
	}

	return nil
}

func (m *PaginatedSubscribers) contextValidateSubscribers(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.Required("subscribers", "body", m.Subscribers); err != nil {
		return err
	}

	for k := range m.Subscribers {

		if val, ok := m.Subscribers[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PaginatedSubscribers) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaginatedSubscribers) UnmarshalBinary(b []byte) error {
	var res PaginatedSubscribers
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
