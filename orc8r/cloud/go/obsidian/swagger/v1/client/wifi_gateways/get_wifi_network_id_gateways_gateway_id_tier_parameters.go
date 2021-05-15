// Code generated by go-swagger; DO NOT EDIT.

package wifi_gateways

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetWifiNetworkIDGatewaysGatewayIDTierParams creates a new GetWifiNetworkIDGatewaysGatewayIDTierParams object
// with the default values initialized.
func NewGetWifiNetworkIDGatewaysGatewayIDTierParams() *GetWifiNetworkIDGatewaysGatewayIDTierParams {
	var ()
	return &GetWifiNetworkIDGatewaysGatewayIDTierParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetWifiNetworkIDGatewaysGatewayIDTierParamsWithTimeout creates a new GetWifiNetworkIDGatewaysGatewayIDTierParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetWifiNetworkIDGatewaysGatewayIDTierParamsWithTimeout(timeout time.Duration) *GetWifiNetworkIDGatewaysGatewayIDTierParams {
	var ()
	return &GetWifiNetworkIDGatewaysGatewayIDTierParams{

		timeout: timeout,
	}
}

// NewGetWifiNetworkIDGatewaysGatewayIDTierParamsWithContext creates a new GetWifiNetworkIDGatewaysGatewayIDTierParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetWifiNetworkIDGatewaysGatewayIDTierParamsWithContext(ctx context.Context) *GetWifiNetworkIDGatewaysGatewayIDTierParams {
	var ()
	return &GetWifiNetworkIDGatewaysGatewayIDTierParams{

		Context: ctx,
	}
}

// NewGetWifiNetworkIDGatewaysGatewayIDTierParamsWithHTTPClient creates a new GetWifiNetworkIDGatewaysGatewayIDTierParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetWifiNetworkIDGatewaysGatewayIDTierParamsWithHTTPClient(client *http.Client) *GetWifiNetworkIDGatewaysGatewayIDTierParams {
	var ()
	return &GetWifiNetworkIDGatewaysGatewayIDTierParams{
		HTTPClient: client,
	}
}

/*GetWifiNetworkIDGatewaysGatewayIDTierParams contains all the parameters to send to the API endpoint
for the get wifi network ID gateways gateway ID tier operation typically these are written to a http.Request
*/
type GetWifiNetworkIDGatewaysGatewayIDTierParams struct {

	/*GatewayID
	  Gateway ID

	*/
	GatewayID string
	/*NetworkID
	  Network ID

	*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get wifi network ID gateways gateway ID tier params
func (o *GetWifiNetworkIDGatewaysGatewayIDTierParams) WithTimeout(timeout time.Duration) *GetWifiNetworkIDGatewaysGatewayIDTierParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get wifi network ID gateways gateway ID tier params
func (o *GetWifiNetworkIDGatewaysGatewayIDTierParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get wifi network ID gateways gateway ID tier params
func (o *GetWifiNetworkIDGatewaysGatewayIDTierParams) WithContext(ctx context.Context) *GetWifiNetworkIDGatewaysGatewayIDTierParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get wifi network ID gateways gateway ID tier params
func (o *GetWifiNetworkIDGatewaysGatewayIDTierParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get wifi network ID gateways gateway ID tier params
func (o *GetWifiNetworkIDGatewaysGatewayIDTierParams) WithHTTPClient(client *http.Client) *GetWifiNetworkIDGatewaysGatewayIDTierParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get wifi network ID gateways gateway ID tier params
func (o *GetWifiNetworkIDGatewaysGatewayIDTierParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGatewayID adds the gatewayID to the get wifi network ID gateways gateway ID tier params
func (o *GetWifiNetworkIDGatewaysGatewayIDTierParams) WithGatewayID(gatewayID string) *GetWifiNetworkIDGatewaysGatewayIDTierParams {
	o.SetGatewayID(gatewayID)
	return o
}

// SetGatewayID adds the gatewayId to the get wifi network ID gateways gateway ID tier params
func (o *GetWifiNetworkIDGatewaysGatewayIDTierParams) SetGatewayID(gatewayID string) {
	o.GatewayID = gatewayID
}

// WithNetworkID adds the networkID to the get wifi network ID gateways gateway ID tier params
func (o *GetWifiNetworkIDGatewaysGatewayIDTierParams) WithNetworkID(networkID string) *GetWifiNetworkIDGatewaysGatewayIDTierParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get wifi network ID gateways gateway ID tier params
func (o *GetWifiNetworkIDGatewaysGatewayIDTierParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *GetWifiNetworkIDGatewaysGatewayIDTierParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param gateway_id
	if err := r.SetPathParam("gateway_id", o.GatewayID); err != nil {
		return err
	}

	// path param network_id
	if err := r.SetPathParam("network_id", o.NetworkID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
