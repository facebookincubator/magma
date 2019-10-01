// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NetworkCarrierWifiConfigs Carrier WiFi configuration for a network
// swagger:model network_carrier_wifi_configs
type NetworkCarrierWifiConfigs struct {

	// aaa server
	// Required: true
	AaaServer *AaaServer `json:"aaa_server"`

	// default rule id
	// Required: true
	DefaultRuleID *string `json:"default_rule_id"`

	// eap aka
	// Required: true
	EapAka *EapAka `json:"eap_aka"`

	// Configuration for network services. Services will be instantiated in the listed order.
	// Required: true
	NetworkServices []string `json:"network_services"`
}

// Validate validates this network carrier wifi configs
func (m *NetworkCarrierWifiConfigs) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAaaServer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDefaultRuleID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEapAka(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetworkServices(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NetworkCarrierWifiConfigs) validateAaaServer(formats strfmt.Registry) error {

	if err := validate.Required("aaa_server", "body", m.AaaServer); err != nil {
		return err
	}

	if m.AaaServer != nil {
		if err := m.AaaServer.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("aaa_server")
			}
			return err
		}
	}

	return nil
}

func (m *NetworkCarrierWifiConfigs) validateDefaultRuleID(formats strfmt.Registry) error {

	if err := validate.Required("default_rule_id", "body", m.DefaultRuleID); err != nil {
		return err
	}

	return nil
}

func (m *NetworkCarrierWifiConfigs) validateEapAka(formats strfmt.Registry) error {

	if err := validate.Required("eap_aka", "body", m.EapAka); err != nil {
		return err
	}

	if m.EapAka != nil {
		if err := m.EapAka.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("eap_aka")
			}
			return err
		}
	}

	return nil
}

var networkCarrierWifiConfigsNetworkServicesItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["metering","dpi","policy_enforcement"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		networkCarrierWifiConfigsNetworkServicesItemsEnum = append(networkCarrierWifiConfigsNetworkServicesItemsEnum, v)
	}
}

func (m *NetworkCarrierWifiConfigs) validateNetworkServicesItemsEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, networkCarrierWifiConfigsNetworkServicesItemsEnum); err != nil {
		return err
	}
	return nil
}

func (m *NetworkCarrierWifiConfigs) validateNetworkServices(formats strfmt.Registry) error {

	if err := validate.Required("network_services", "body", m.NetworkServices); err != nil {
		return err
	}

	for i := 0; i < len(m.NetworkServices); i++ {

		// value enum
		if err := m.validateNetworkServicesItemsEnum("network_services"+"."+strconv.Itoa(i), "body", m.NetworkServices[i]); err != nil {
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *NetworkCarrierWifiConfigs) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetworkCarrierWifiConfigs) UnmarshalBinary(b []byte) error {
	var res NetworkCarrierWifiConfigs
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
