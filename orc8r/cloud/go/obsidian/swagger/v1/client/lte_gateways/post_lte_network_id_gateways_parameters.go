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

// NewPostLTENetworkIDGatewaysParams creates a new PostLTENetworkIDGatewaysParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostLTENetworkIDGatewaysParams() *PostLTENetworkIDGatewaysParams {
	return &PostLTENetworkIDGatewaysParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostLTENetworkIDGatewaysParamsWithTimeout creates a new PostLTENetworkIDGatewaysParams object
// with the ability to set a timeout on a request.
func NewPostLTENetworkIDGatewaysParamsWithTimeout(timeout time.Duration) *PostLTENetworkIDGatewaysParams {
	return &PostLTENetworkIDGatewaysParams{
		timeout: timeout,
	}
}

// NewPostLTENetworkIDGatewaysParamsWithContext creates a new PostLTENetworkIDGatewaysParams object
// with the ability to set a context for a request.
func NewPostLTENetworkIDGatewaysParamsWithContext(ctx context.Context) *PostLTENetworkIDGatewaysParams {
	return &PostLTENetworkIDGatewaysParams{
		Context: ctx,
	}
}

// NewPostLTENetworkIDGatewaysParamsWithHTTPClient creates a new PostLTENetworkIDGatewaysParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostLTENetworkIDGatewaysParamsWithHTTPClient(client *http.Client) *PostLTENetworkIDGatewaysParams {
	return &PostLTENetworkIDGatewaysParams{
		HTTPClient: client,
	}
}

/* PostLTENetworkIDGatewaysParams contains all the parameters to send to the API endpoint
   for the post LTE network ID gateways operation.

   Typically these are written to a http.Request.
*/
type PostLTENetworkIDGatewaysParams struct {

	/* Gateway.

	   Full desired configuration of the gateway
	*/
	Gateway *models.MutableLTEGateway

	/* NetworkID.

	   Network ID
	*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post LTE network ID gateways params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostLTENetworkIDGatewaysParams) WithDefaults() *PostLTENetworkIDGatewaysParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post LTE network ID gateways params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostLTENetworkIDGatewaysParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post LTE network ID gateways params
func (o *PostLTENetworkIDGatewaysParams) WithTimeout(timeout time.Duration) *PostLTENetworkIDGatewaysParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post LTE network ID gateways params
func (o *PostLTENetworkIDGatewaysParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post LTE network ID gateways params
func (o *PostLTENetworkIDGatewaysParams) WithContext(ctx context.Context) *PostLTENetworkIDGatewaysParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post LTE network ID gateways params
func (o *PostLTENetworkIDGatewaysParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post LTE network ID gateways params
func (o *PostLTENetworkIDGatewaysParams) WithHTTPClient(client *http.Client) *PostLTENetworkIDGatewaysParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post LTE network ID gateways params
func (o *PostLTENetworkIDGatewaysParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGateway adds the gateway to the post LTE network ID gateways params
func (o *PostLTENetworkIDGatewaysParams) WithGateway(gateway *models.MutableLTEGateway) *PostLTENetworkIDGatewaysParams {
	o.SetGateway(gateway)
	return o
}

// SetGateway adds the gateway to the post LTE network ID gateways params
func (o *PostLTENetworkIDGatewaysParams) SetGateway(gateway *models.MutableLTEGateway) {
	o.Gateway = gateway
}

// WithNetworkID adds the networkID to the post LTE network ID gateways params
func (o *PostLTENetworkIDGatewaysParams) WithNetworkID(networkID string) *PostLTENetworkIDGatewaysParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the post LTE network ID gateways params
func (o *PostLTENetworkIDGatewaysParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *PostLTENetworkIDGatewaysParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Gateway != nil {
		if err := r.SetBodyParam(o.Gateway); err != nil {
			return err
		}
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
