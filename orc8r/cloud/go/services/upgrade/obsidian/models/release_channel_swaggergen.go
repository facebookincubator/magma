// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// ReleaseChannel release channel
// swagger:model release_channel
type ReleaseChannel struct {

	// name
	Name string `json:"name,omitempty"`

	// supported versions
	SupportedVersions []string `json:"supported_versions"`
}

// Validate validates this release channel
func (m *ReleaseChannel) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ReleaseChannel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReleaseChannel) UnmarshalBinary(b []byte) error {
	var res ReleaseChannel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
