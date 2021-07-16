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

	"magma/orc8r/cloud/go/obsidian/swagger/v1/models"
)

// NewPutLTENetworkIDGatewaysGatewayIDNameParams creates a new PutLTENetworkIDGatewaysGatewayIDNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutLTENetworkIDGatewaysGatewayIDNameParams() *PutLTENetworkIDGatewaysGatewayIDNameParams {
	return &PutLTENetworkIDGatewaysGatewayIDNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutLTENetworkIDGatewaysGatewayIDNameParamsWithTimeout creates a new PutLTENetworkIDGatewaysGatewayIDNameParams object
// with the ability to set a timeout on a request.
func NewPutLTENetworkIDGatewaysGatewayIDNameParamsWithTimeout(timeout time.Duration) *PutLTENetworkIDGatewaysGatewayIDNameParams {
	return &PutLTENetworkIDGatewaysGatewayIDNameParams{
		timeout: timeout,
	}
}

// NewPutLTENetworkIDGatewaysGatewayIDNameParamsWithContext creates a new PutLTENetworkIDGatewaysGatewayIDNameParams object
// with the ability to set a context for a request.
func NewPutLTENetworkIDGatewaysGatewayIDNameParamsWithContext(ctx context.Context) *PutLTENetworkIDGatewaysGatewayIDNameParams {
	return &PutLTENetworkIDGatewaysGatewayIDNameParams{
		Context: ctx,
	}
}

// NewPutLTENetworkIDGatewaysGatewayIDNameParamsWithHTTPClient creates a new PutLTENetworkIDGatewaysGatewayIDNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutLTENetworkIDGatewaysGatewayIDNameParamsWithHTTPClient(client *http.Client) *PutLTENetworkIDGatewaysGatewayIDNameParams {
	return &PutLTENetworkIDGatewaysGatewayIDNameParams{
		HTTPClient: client,
	}
}

/* PutLTENetworkIDGatewaysGatewayIDNameParams contains all the parameters to send to the API endpoint
   for the put LTE network ID gateways gateway ID name operation.

   Typically these are written to a http.Request.
*/
type PutLTENetworkIDGatewaysGatewayIDNameParams struct {

	/* GatewayID.

	   Gateway ID
	*/
	GatewayID string

	// Name.
	Name models.GatewayName

	/* NetworkID.

	   Network ID
	*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put LTE network ID gateways gateway ID name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutLTENetworkIDGatewaysGatewayIDNameParams) WithDefaults() *PutLTENetworkIDGatewaysGatewayIDNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put LTE network ID gateways gateway ID name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutLTENetworkIDGatewaysGatewayIDNameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put LTE network ID gateways gateway ID name params
func (o *PutLTENetworkIDGatewaysGatewayIDNameParams) WithTimeout(timeout time.Duration) *PutLTENetworkIDGatewaysGatewayIDNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put LTE network ID gateways gateway ID name params
func (o *PutLTENetworkIDGatewaysGatewayIDNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put LTE network ID gateways gateway ID name params
func (o *PutLTENetworkIDGatewaysGatewayIDNameParams) WithContext(ctx context.Context) *PutLTENetworkIDGatewaysGatewayIDNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put LTE network ID gateways gateway ID name params
func (o *PutLTENetworkIDGatewaysGatewayIDNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put LTE network ID gateways gateway ID name params
func (o *PutLTENetworkIDGatewaysGatewayIDNameParams) WithHTTPClient(client *http.Client) *PutLTENetworkIDGatewaysGatewayIDNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put LTE network ID gateways gateway ID name params
func (o *PutLTENetworkIDGatewaysGatewayIDNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGatewayID adds the gatewayID to the put LTE network ID gateways gateway ID name params
func (o *PutLTENetworkIDGatewaysGatewayIDNameParams) WithGatewayID(gatewayID string) *PutLTENetworkIDGatewaysGatewayIDNameParams {
	o.SetGatewayID(gatewayID)
	return o
}

// SetGatewayID adds the gatewayId to the put LTE network ID gateways gateway ID name params
func (o *PutLTENetworkIDGatewaysGatewayIDNameParams) SetGatewayID(gatewayID string) {
	o.GatewayID = gatewayID
}

// WithName adds the name to the put LTE network ID gateways gateway ID name params
func (o *PutLTENetworkIDGatewaysGatewayIDNameParams) WithName(name models.GatewayName) *PutLTENetworkIDGatewaysGatewayIDNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the put LTE network ID gateways gateway ID name params
func (o *PutLTENetworkIDGatewaysGatewayIDNameParams) SetName(name models.GatewayName) {
	o.Name = name
}

// WithNetworkID adds the networkID to the put LTE network ID gateways gateway ID name params
func (o *PutLTENetworkIDGatewaysGatewayIDNameParams) WithNetworkID(networkID string) *PutLTENetworkIDGatewaysGatewayIDNameParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the put LTE network ID gateways gateway ID name params
func (o *PutLTENetworkIDGatewaysGatewayIDNameParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *PutLTENetworkIDGatewaysGatewayIDNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param gateway_id
	if err := r.SetPathParam("gateway_id", o.GatewayID); err != nil {
		return err
	}
	if err := r.SetBodyParam(o.Name); err != nil {
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
