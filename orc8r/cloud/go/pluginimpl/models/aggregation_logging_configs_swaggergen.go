// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// AggregationLoggingConfigs Configuration for log aggregation
// swagger:model aggregation_logging_configs
type AggregationLoggingConfigs struct {

	// target files by tag
	TargetFilesByTag map[string]string `json:"target_files_by_tag,omitempty"`
}

// Validate validates this aggregation logging configs
func (m *AggregationLoggingConfigs) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AggregationLoggingConfigs) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AggregationLoggingConfigs) UnmarshalBinary(b []byte) error {
	var res AggregationLoggingConfigs
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
