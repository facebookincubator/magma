// Code generated by go-swagger; DO NOT EDIT.

package federation_gateways

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

// NewPostFegNetworkIDGatewaysParams creates a new PostFegNetworkIDGatewaysParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostFegNetworkIDGatewaysParams() *PostFegNetworkIDGatewaysParams {
	return &PostFegNetworkIDGatewaysParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostFegNetworkIDGatewaysParamsWithTimeout creates a new PostFegNetworkIDGatewaysParams object
// with the ability to set a timeout on a request.
func NewPostFegNetworkIDGatewaysParamsWithTimeout(timeout time.Duration) *PostFegNetworkIDGatewaysParams {
	return &PostFegNetworkIDGatewaysParams{
		timeout: timeout,
	}
}

// NewPostFegNetworkIDGatewaysParamsWithContext creates a new PostFegNetworkIDGatewaysParams object
// with the ability to set a context for a request.
func NewPostFegNetworkIDGatewaysParamsWithContext(ctx context.Context) *PostFegNetworkIDGatewaysParams {
	return &PostFegNetworkIDGatewaysParams{
		Context: ctx,
	}
}

// NewPostFegNetworkIDGatewaysParamsWithHTTPClient creates a new PostFegNetworkIDGatewaysParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostFegNetworkIDGatewaysParamsWithHTTPClient(client *http.Client) *PostFegNetworkIDGatewaysParams {
	return &PostFegNetworkIDGatewaysParams{
		HTTPClient: client,
	}
}

/* PostFegNetworkIDGatewaysParams contains all the parameters to send to the API endpoint
   for the post feg network ID gateways operation.

   Typically these are written to a http.Request.
*/
type PostFegNetworkIDGatewaysParams struct {

	/* Gateway.

	   Full desired configuration of the gateway
	*/
	Gateway *models.MutableFederationGateway

	/* NetworkID.

	   Network ID
	*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post feg network ID gateways params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostFegNetworkIDGatewaysParams) WithDefaults() *PostFegNetworkIDGatewaysParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post feg network ID gateways params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostFegNetworkIDGatewaysParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post feg network ID gateways params
func (o *PostFegNetworkIDGatewaysParams) WithTimeout(timeout time.Duration) *PostFegNetworkIDGatewaysParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post feg network ID gateways params
func (o *PostFegNetworkIDGatewaysParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post feg network ID gateways params
func (o *PostFegNetworkIDGatewaysParams) WithContext(ctx context.Context) *PostFegNetworkIDGatewaysParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post feg network ID gateways params
func (o *PostFegNetworkIDGatewaysParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post feg network ID gateways params
func (o *PostFegNetworkIDGatewaysParams) WithHTTPClient(client *http.Client) *PostFegNetworkIDGatewaysParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post feg network ID gateways params
func (o *PostFegNetworkIDGatewaysParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGateway adds the gateway to the post feg network ID gateways params
func (o *PostFegNetworkIDGatewaysParams) WithGateway(gateway *models.MutableFederationGateway) *PostFegNetworkIDGatewaysParams {
	o.SetGateway(gateway)
	return o
}

// SetGateway adds the gateway to the post feg network ID gateways params
func (o *PostFegNetworkIDGatewaysParams) SetGateway(gateway *models.MutableFederationGateway) {
	o.Gateway = gateway
}

// WithNetworkID adds the networkID to the post feg network ID gateways params
func (o *PostFegNetworkIDGatewaysParams) WithNetworkID(networkID string) *PostFegNetworkIDGatewaysParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the post feg network ID gateways params
func (o *PostFegNetworkIDGatewaysParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *PostFegNetworkIDGatewaysParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
