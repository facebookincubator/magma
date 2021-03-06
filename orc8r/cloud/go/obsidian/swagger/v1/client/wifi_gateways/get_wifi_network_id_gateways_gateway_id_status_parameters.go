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

// NewGetWifiNetworkIDGatewaysGatewayIDStatusParams creates a new GetWifiNetworkIDGatewaysGatewayIDStatusParams object
// with the default values initialized.
func NewGetWifiNetworkIDGatewaysGatewayIDStatusParams() *GetWifiNetworkIDGatewaysGatewayIDStatusParams {
	var ()
	return &GetWifiNetworkIDGatewaysGatewayIDStatusParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetWifiNetworkIDGatewaysGatewayIDStatusParamsWithTimeout creates a new GetWifiNetworkIDGatewaysGatewayIDStatusParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetWifiNetworkIDGatewaysGatewayIDStatusParamsWithTimeout(timeout time.Duration) *GetWifiNetworkIDGatewaysGatewayIDStatusParams {
	var ()
	return &GetWifiNetworkIDGatewaysGatewayIDStatusParams{

		timeout: timeout,
	}
}

// NewGetWifiNetworkIDGatewaysGatewayIDStatusParamsWithContext creates a new GetWifiNetworkIDGatewaysGatewayIDStatusParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetWifiNetworkIDGatewaysGatewayIDStatusParamsWithContext(ctx context.Context) *GetWifiNetworkIDGatewaysGatewayIDStatusParams {
	var ()
	return &GetWifiNetworkIDGatewaysGatewayIDStatusParams{

		Context: ctx,
	}
}

// NewGetWifiNetworkIDGatewaysGatewayIDStatusParamsWithHTTPClient creates a new GetWifiNetworkIDGatewaysGatewayIDStatusParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetWifiNetworkIDGatewaysGatewayIDStatusParamsWithHTTPClient(client *http.Client) *GetWifiNetworkIDGatewaysGatewayIDStatusParams {
	var ()
	return &GetWifiNetworkIDGatewaysGatewayIDStatusParams{
		HTTPClient: client,
	}
}

/*GetWifiNetworkIDGatewaysGatewayIDStatusParams contains all the parameters to send to the API endpoint
for the get wifi network ID gateways gateway ID status operation typically these are written to a http.Request
*/
type GetWifiNetworkIDGatewaysGatewayIDStatusParams struct {

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

// WithTimeout adds the timeout to the get wifi network ID gateways gateway ID status params
func (o *GetWifiNetworkIDGatewaysGatewayIDStatusParams) WithTimeout(timeout time.Duration) *GetWifiNetworkIDGatewaysGatewayIDStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get wifi network ID gateways gateway ID status params
func (o *GetWifiNetworkIDGatewaysGatewayIDStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get wifi network ID gateways gateway ID status params
func (o *GetWifiNetworkIDGatewaysGatewayIDStatusParams) WithContext(ctx context.Context) *GetWifiNetworkIDGatewaysGatewayIDStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get wifi network ID gateways gateway ID status params
func (o *GetWifiNetworkIDGatewaysGatewayIDStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get wifi network ID gateways gateway ID status params
func (o *GetWifiNetworkIDGatewaysGatewayIDStatusParams) WithHTTPClient(client *http.Client) *GetWifiNetworkIDGatewaysGatewayIDStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get wifi network ID gateways gateway ID status params
func (o *GetWifiNetworkIDGatewaysGatewayIDStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGatewayID adds the gatewayID to the get wifi network ID gateways gateway ID status params
func (o *GetWifiNetworkIDGatewaysGatewayIDStatusParams) WithGatewayID(gatewayID string) *GetWifiNetworkIDGatewaysGatewayIDStatusParams {
	o.SetGatewayID(gatewayID)
	return o
}

// SetGatewayID adds the gatewayId to the get wifi network ID gateways gateway ID status params
func (o *GetWifiNetworkIDGatewaysGatewayIDStatusParams) SetGatewayID(gatewayID string) {
	o.GatewayID = gatewayID
}

// WithNetworkID adds the networkID to the get wifi network ID gateways gateway ID status params
func (o *GetWifiNetworkIDGatewaysGatewayIDStatusParams) WithNetworkID(networkID string) *GetWifiNetworkIDGatewaysGatewayIDStatusParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get wifi network ID gateways gateway ID status params
func (o *GetWifiNetworkIDGatewaysGatewayIDStatusParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *GetWifiNetworkIDGatewaysGatewayIDStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
