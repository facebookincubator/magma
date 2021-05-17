// Code generated by go-swagger; DO NOT EDIT.

package lte_gateways

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

// NewGetLTENetworkIDGatewaysGatewayIDTierParams creates a new GetLTENetworkIDGatewaysGatewayIDTierParams object
// with the default values initialized.
func NewGetLTENetworkIDGatewaysGatewayIDTierParams() *GetLTENetworkIDGatewaysGatewayIDTierParams {
	var ()
	return &GetLTENetworkIDGatewaysGatewayIDTierParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLTENetworkIDGatewaysGatewayIDTierParamsWithTimeout creates a new GetLTENetworkIDGatewaysGatewayIDTierParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLTENetworkIDGatewaysGatewayIDTierParamsWithTimeout(timeout time.Duration) *GetLTENetworkIDGatewaysGatewayIDTierParams {
	var ()
	return &GetLTENetworkIDGatewaysGatewayIDTierParams{

		timeout: timeout,
	}
}

// NewGetLTENetworkIDGatewaysGatewayIDTierParamsWithContext creates a new GetLTENetworkIDGatewaysGatewayIDTierParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLTENetworkIDGatewaysGatewayIDTierParamsWithContext(ctx context.Context) *GetLTENetworkIDGatewaysGatewayIDTierParams {
	var ()
	return &GetLTENetworkIDGatewaysGatewayIDTierParams{

		Context: ctx,
	}
}

// NewGetLTENetworkIDGatewaysGatewayIDTierParamsWithHTTPClient creates a new GetLTENetworkIDGatewaysGatewayIDTierParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLTENetworkIDGatewaysGatewayIDTierParamsWithHTTPClient(client *http.Client) *GetLTENetworkIDGatewaysGatewayIDTierParams {
	var ()
	return &GetLTENetworkIDGatewaysGatewayIDTierParams{
		HTTPClient: client,
	}
}

/*GetLTENetworkIDGatewaysGatewayIDTierParams contains all the parameters to send to the API endpoint
for the get LTE network ID gateways gateway ID tier operation typically these are written to a http.Request
*/
type GetLTENetworkIDGatewaysGatewayIDTierParams struct {

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

// WithTimeout adds the timeout to the get LTE network ID gateways gateway ID tier params
func (o *GetLTENetworkIDGatewaysGatewayIDTierParams) WithTimeout(timeout time.Duration) *GetLTENetworkIDGatewaysGatewayIDTierParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get LTE network ID gateways gateway ID tier params
func (o *GetLTENetworkIDGatewaysGatewayIDTierParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get LTE network ID gateways gateway ID tier params
func (o *GetLTENetworkIDGatewaysGatewayIDTierParams) WithContext(ctx context.Context) *GetLTENetworkIDGatewaysGatewayIDTierParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get LTE network ID gateways gateway ID tier params
func (o *GetLTENetworkIDGatewaysGatewayIDTierParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get LTE network ID gateways gateway ID tier params
func (o *GetLTENetworkIDGatewaysGatewayIDTierParams) WithHTTPClient(client *http.Client) *GetLTENetworkIDGatewaysGatewayIDTierParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get LTE network ID gateways gateway ID tier params
func (o *GetLTENetworkIDGatewaysGatewayIDTierParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGatewayID adds the gatewayID to the get LTE network ID gateways gateway ID tier params
func (o *GetLTENetworkIDGatewaysGatewayIDTierParams) WithGatewayID(gatewayID string) *GetLTENetworkIDGatewaysGatewayIDTierParams {
	o.SetGatewayID(gatewayID)
	return o
}

// SetGatewayID adds the gatewayId to the get LTE network ID gateways gateway ID tier params
func (o *GetLTENetworkIDGatewaysGatewayIDTierParams) SetGatewayID(gatewayID string) {
	o.GatewayID = gatewayID
}

// WithNetworkID adds the networkID to the get LTE network ID gateways gateway ID tier params
func (o *GetLTENetworkIDGatewaysGatewayIDTierParams) WithNetworkID(networkID string) *GetLTENetworkIDGatewaysGatewayIDTierParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get LTE network ID gateways gateway ID tier params
func (o *GetLTENetworkIDGatewaysGatewayIDTierParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *GetLTENetworkIDGatewaysGatewayIDTierParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
