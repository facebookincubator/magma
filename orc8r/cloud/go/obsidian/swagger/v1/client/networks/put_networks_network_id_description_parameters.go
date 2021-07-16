// Code generated by go-swagger; DO NOT EDIT.

package networks

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

// NewPutNetworksNetworkIDDescriptionParams creates a new PutNetworksNetworkIDDescriptionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutNetworksNetworkIDDescriptionParams() *PutNetworksNetworkIDDescriptionParams {
	return &PutNetworksNetworkIDDescriptionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutNetworksNetworkIDDescriptionParamsWithTimeout creates a new PutNetworksNetworkIDDescriptionParams object
// with the ability to set a timeout on a request.
func NewPutNetworksNetworkIDDescriptionParamsWithTimeout(timeout time.Duration) *PutNetworksNetworkIDDescriptionParams {
	return &PutNetworksNetworkIDDescriptionParams{
		timeout: timeout,
	}
}

// NewPutNetworksNetworkIDDescriptionParamsWithContext creates a new PutNetworksNetworkIDDescriptionParams object
// with the ability to set a context for a request.
func NewPutNetworksNetworkIDDescriptionParamsWithContext(ctx context.Context) *PutNetworksNetworkIDDescriptionParams {
	return &PutNetworksNetworkIDDescriptionParams{
		Context: ctx,
	}
}

// NewPutNetworksNetworkIDDescriptionParamsWithHTTPClient creates a new PutNetworksNetworkIDDescriptionParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutNetworksNetworkIDDescriptionParamsWithHTTPClient(client *http.Client) *PutNetworksNetworkIDDescriptionParams {
	return &PutNetworksNetworkIDDescriptionParams{
		HTTPClient: client,
	}
}

/* PutNetworksNetworkIDDescriptionParams contains all the parameters to send to the API endpoint
   for the put networks network ID description operation.

   Typically these are written to a http.Request.
*/
type PutNetworksNetworkIDDescriptionParams struct {

	/* Description.

	   New name for the network
	*/
	Description models.NetworkDescription

	/* NetworkID.

	   Network ID
	*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put networks network ID description params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutNetworksNetworkIDDescriptionParams) WithDefaults() *PutNetworksNetworkIDDescriptionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put networks network ID description params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutNetworksNetworkIDDescriptionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put networks network ID description params
func (o *PutNetworksNetworkIDDescriptionParams) WithTimeout(timeout time.Duration) *PutNetworksNetworkIDDescriptionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put networks network ID description params
func (o *PutNetworksNetworkIDDescriptionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put networks network ID description params
func (o *PutNetworksNetworkIDDescriptionParams) WithContext(ctx context.Context) *PutNetworksNetworkIDDescriptionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put networks network ID description params
func (o *PutNetworksNetworkIDDescriptionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put networks network ID description params
func (o *PutNetworksNetworkIDDescriptionParams) WithHTTPClient(client *http.Client) *PutNetworksNetworkIDDescriptionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put networks network ID description params
func (o *PutNetworksNetworkIDDescriptionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDescription adds the description to the put networks network ID description params
func (o *PutNetworksNetworkIDDescriptionParams) WithDescription(description models.NetworkDescription) *PutNetworksNetworkIDDescriptionParams {
	o.SetDescription(description)
	return o
}

// SetDescription adds the description to the put networks network ID description params
func (o *PutNetworksNetworkIDDescriptionParams) SetDescription(description models.NetworkDescription) {
	o.Description = description
}

// WithNetworkID adds the networkID to the put networks network ID description params
func (o *PutNetworksNetworkIDDescriptionParams) WithNetworkID(networkID string) *PutNetworksNetworkIDDescriptionParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the put networks network ID description params
func (o *PutNetworksNetworkIDDescriptionParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *PutNetworksNetworkIDDescriptionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.Description); err != nil {
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
