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
	"github.com/go-openapi/strfmt"
)

// NewGetLTENetworkIDGatewaysGatewayIDCellularDNSRecordsParams creates a new GetLTENetworkIDGatewaysGatewayIDCellularDNSRecordsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetLTENetworkIDGatewaysGatewayIDCellularDNSRecordsParams() *GetLTENetworkIDGatewaysGatewayIDCellularDNSRecordsParams {
	return &GetLTENetworkIDGatewaysGatewayIDCellularDNSRecordsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetLTENetworkIDGatewaysGatewayIDCellularDNSRecordsParamsWithTimeout creates a new GetLTENetworkIDGatewaysGatewayIDCellularDNSRecordsParams object
// with the ability to set a timeout on a request.
func NewGetLTENetworkIDGatewaysGatewayIDCellularDNSRecordsParamsWithTimeout(timeout time.Duration) *GetLTENetworkIDGatewaysGatewayIDCellularDNSRecordsParams {
	return &GetLTENetworkIDGatewaysGatewayIDCellularDNSRecordsParams{
		timeout: timeout,
	}
}

// NewGetLTENetworkIDGatewaysGatewayIDCellularDNSRecordsParamsWithContext creates a new GetLTENetworkIDGatewaysGatewayIDCellularDNSRecordsParams object
// with the ability to set a context for a request.
func NewGetLTENetworkIDGatewaysGatewayIDCellularDNSRecordsParamsWithContext(ctx context.Context) *GetLTENetworkIDGatewaysGatewayIDCellularDNSRecordsParams {
	return &GetLTENetworkIDGatewaysGatewayIDCellularDNSRecordsParams{
		Context: ctx,
	}
}

// NewGetLTENetworkIDGatewaysGatewayIDCellularDNSRecordsParamsWithHTTPClient creates a new GetLTENetworkIDGatewaysGatewayIDCellularDNSRecordsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetLTENetworkIDGatewaysGatewayIDCellularDNSRecordsParamsWithHTTPClient(client *http.Client) *GetLTENetworkIDGatewaysGatewayIDCellularDNSRecordsParams {
	return &GetLTENetworkIDGatewaysGatewayIDCellularDNSRecordsParams{
		HTTPClient: client,
	}
}

/* GetLTENetworkIDGatewaysGatewayIDCellularDNSRecordsParams contains all the parameters to send to the API endpoint
   for the get LTE network ID gateways gateway ID cellular DNS records operation.

   Typically these are written to a http.Request.
*/
type GetLTENetworkIDGatewaysGatewayIDCellularDNSRecordsParams struct {

	/* GatewayID.

	   Gateway ID
	*/
	GatewayID string

	/* NetworkID.

	   Network ID
	*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get LTE network ID gateways gateway ID cellular DNS records params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLTENetworkIDGatewaysGatewayIDCellularDNSRecordsParams) WithDefaults() *GetLTENetworkIDGatewaysGatewayIDCellularDNSRecordsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get LTE network ID gateways gateway ID cellular DNS records params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLTENetworkIDGatewaysGatewayIDCellularDNSRecordsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get LTE network ID gateways gateway ID cellular DNS records params
func (o *GetLTENetworkIDGatewaysGatewayIDCellularDNSRecordsParams) WithTimeout(timeout time.Duration) *GetLTENetworkIDGatewaysGatewayIDCellularDNSRecordsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get LTE network ID gateways gateway ID cellular DNS records params
func (o *GetLTENetworkIDGatewaysGatewayIDCellularDNSRecordsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get LTE network ID gateways gateway ID cellular DNS records params
func (o *GetLTENetworkIDGatewaysGatewayIDCellularDNSRecordsParams) WithContext(ctx context.Context) *GetLTENetworkIDGatewaysGatewayIDCellularDNSRecordsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get LTE network ID gateways gateway ID cellular DNS records params
func (o *GetLTENetworkIDGatewaysGatewayIDCellularDNSRecordsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get LTE network ID gateways gateway ID cellular DNS records params
func (o *GetLTENetworkIDGatewaysGatewayIDCellularDNSRecordsParams) WithHTTPClient(client *http.Client) *GetLTENetworkIDGatewaysGatewayIDCellularDNSRecordsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get LTE network ID gateways gateway ID cellular DNS records params
func (o *GetLTENetworkIDGatewaysGatewayIDCellularDNSRecordsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGatewayID adds the gatewayID to the get LTE network ID gateways gateway ID cellular DNS records params
func (o *GetLTENetworkIDGatewaysGatewayIDCellularDNSRecordsParams) WithGatewayID(gatewayID string) *GetLTENetworkIDGatewaysGatewayIDCellularDNSRecordsParams {
	o.SetGatewayID(gatewayID)
	return o
}

// SetGatewayID adds the gatewayId to the get LTE network ID gateways gateway ID cellular DNS records params
func (o *GetLTENetworkIDGatewaysGatewayIDCellularDNSRecordsParams) SetGatewayID(gatewayID string) {
	o.GatewayID = gatewayID
}

// WithNetworkID adds the networkID to the get LTE network ID gateways gateway ID cellular DNS records params
func (o *GetLTENetworkIDGatewaysGatewayIDCellularDNSRecordsParams) WithNetworkID(networkID string) *GetLTENetworkIDGatewaysGatewayIDCellularDNSRecordsParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get LTE network ID gateways gateway ID cellular DNS records params
func (o *GetLTENetworkIDGatewaysGatewayIDCellularDNSRecordsParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *GetLTENetworkIDGatewaysGatewayIDCellularDNSRecordsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
