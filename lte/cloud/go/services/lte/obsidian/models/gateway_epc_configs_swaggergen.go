// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GatewayEpcConfigs EPC configuration for an LTE gateway
// swagger:model gateway_epc_configs
type GatewayEpcConfigs struct {

	// dns primary
	// Max Length: 45
	// Min Length: 5
	DNSPrimary string `json:"dns_primary,omitempty"`

	// dns secondary
	// Max Length: 45
	// Min Length: 5
	DNSSecondary string `json:"dns_secondary,omitempty"`

	// ip block
	// Required: true
	// Max Length: 49
	// Min Length: 5
	IPBlock string `json:"ip_block"`

	// nat enabled
	// Required: true
	NatEnabled *bool `json:"nat_enabled"`

	// IP address for management interface on the AGW, If not specified AGW uses DHCP to configure it.
	// Max Length: 49
	// Min Length: 5
	SgiManagementIfaceStaticIP string `json:"sgi_management_iface_static_ip,omitempty"`

	// VLAN ID for management interface traffic on the AGW
	// Max Length: 4
	// Min Length: 1
	SgiManagementIfaceVlan string `json:"sgi_management_iface_vlan,omitempty"`
}

// Validate validates this gateway epc configs
func (m *GatewayEpcConfigs) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDNSPrimary(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDNSSecondary(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIPBlock(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNatEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSgiManagementIfaceStaticIP(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSgiManagementIfaceVlan(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GatewayEpcConfigs) validateDNSPrimary(formats strfmt.Registry) error {

	if swag.IsZero(m.DNSPrimary) { // not required
		return nil
	}

	if err := validate.MinLength("dns_primary", "body", string(m.DNSPrimary), 5); err != nil {
		return err
	}

	if err := validate.MaxLength("dns_primary", "body", string(m.DNSPrimary), 45); err != nil {
		return err
	}

	return nil
}

func (m *GatewayEpcConfigs) validateDNSSecondary(formats strfmt.Registry) error {

	if swag.IsZero(m.DNSSecondary) { // not required
		return nil
	}

	if err := validate.MinLength("dns_secondary", "body", string(m.DNSSecondary), 5); err != nil {
		return err
	}

	if err := validate.MaxLength("dns_secondary", "body", string(m.DNSSecondary), 45); err != nil {
		return err
	}

	return nil
}

func (m *GatewayEpcConfigs) validateIPBlock(formats strfmt.Registry) error {

	if err := validate.RequiredString("ip_block", "body", string(m.IPBlock)); err != nil {
		return err
	}

	if err := validate.MinLength("ip_block", "body", string(m.IPBlock), 5); err != nil {
		return err
	}

	if err := validate.MaxLength("ip_block", "body", string(m.IPBlock), 49); err != nil {
		return err
	}

	return nil
}

func (m *GatewayEpcConfigs) validateNatEnabled(formats strfmt.Registry) error {

	if err := validate.Required("nat_enabled", "body", m.NatEnabled); err != nil {
		return err
	}

	return nil
}

func (m *GatewayEpcConfigs) validateSgiManagementIfaceStaticIP(formats strfmt.Registry) error {

	if swag.IsZero(m.SgiManagementIfaceStaticIP) { // not required
		return nil
	}

	if err := validate.MinLength("sgi_management_iface_static_ip", "body", string(m.SgiManagementIfaceStaticIP), 5); err != nil {
		return err
	}

	if err := validate.MaxLength("sgi_management_iface_static_ip", "body", string(m.SgiManagementIfaceStaticIP), 49); err != nil {
		return err
	}

	return nil
}

func (m *GatewayEpcConfigs) validateSgiManagementIfaceVlan(formats strfmt.Registry) error {

	if swag.IsZero(m.SgiManagementIfaceVlan) { // not required
		return nil
	}

	if err := validate.MinLength("sgi_management_iface_vlan", "body", string(m.SgiManagementIfaceVlan), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("sgi_management_iface_vlan", "body", string(m.SgiManagementIfaceVlan), 4); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GatewayEpcConfigs) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GatewayEpcConfigs) UnmarshalBinary(b []byte) error {
	var res GatewayEpcConfigs
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
