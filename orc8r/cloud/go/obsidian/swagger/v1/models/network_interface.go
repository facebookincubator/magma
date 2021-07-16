// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NetworkInterface network interface
//
// swagger:model network_interface
type NetworkInterface struct {

	// ip addresses
	// Example: ["10.10.10.1","10.0.0.1"]
	IPAddresses []string `json:"ip_addresses,omitempty" magma_alt_name:"IpAddresses"`

	// ipv6 addresses
	// Example: ["fe80::a00:27ff:fe1e:8332","fe80::a00:27ff:fe1e:8432"]
	IPV6Addresses []string `json:"ipv6_addresses,omitempty" magma_alt_name:"Ipv6Addresses"`

	// mac address
	// Example: 08:00:27:1e:8a:32
	MacAddress string `json:"mac_address,omitempty"`

	// network interface id
	// Example: gtp_br0
	NetworkInterfaceID string `json:"network_interface_id,omitempty" magma_alt_name:"NetworkInterfaceId"`

	// status
	// Example: UP
	// Enum: [UP DOWN UNKNOWN]
	Status string `json:"status,omitempty"`
}

// Validate validates this network interface
func (m *NetworkInterface) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var networkInterfaceTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["UP","DOWN","UNKNOWN"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		networkInterfaceTypeStatusPropEnum = append(networkInterfaceTypeStatusPropEnum, v)
	}
}

const (

	// NetworkInterfaceStatusUP captures enum value "UP"
	NetworkInterfaceStatusUP string = "UP"

	// NetworkInterfaceStatusDOWN captures enum value "DOWN"
	NetworkInterfaceStatusDOWN string = "DOWN"

	// NetworkInterfaceStatusUNKNOWN captures enum value "UNKNOWN"
	NetworkInterfaceStatusUNKNOWN string = "UNKNOWN"
)

// prop value enum
func (m *NetworkInterface) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, networkInterfaceTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *NetworkInterface) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this network interface based on context it is used
func (m *NetworkInterface) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NetworkInterface) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetworkInterface) UnmarshalBinary(b []byte) error {
	var res NetworkInterface
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
