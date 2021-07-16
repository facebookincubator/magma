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
)

// NewGetNetworksNetworkIDDescriptionParams creates a new GetNetworksNetworkIDDescriptionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetNetworksNetworkIDDescriptionParams() *GetNetworksNetworkIDDescriptionParams {
	return &GetNetworksNetworkIDDescriptionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetworksNetworkIDDescriptionParamsWithTimeout creates a new GetNetworksNetworkIDDescriptionParams object
// with the ability to set a timeout on a request.
func NewGetNetworksNetworkIDDescriptionParamsWithTimeout(timeout time.Duration) *GetNetworksNetworkIDDescriptionParams {
	return &GetNetworksNetworkIDDescriptionParams{
		timeout: timeout,
	}
}

// NewGetNetworksNetworkIDDescriptionParamsWithContext creates a new GetNetworksNetworkIDDescriptionParams object
// with the ability to set a context for a request.
func NewGetNetworksNetworkIDDescriptionParamsWithContext(ctx context.Context) *GetNetworksNetworkIDDescriptionParams {
	return &GetNetworksNetworkIDDescriptionParams{
		Context: ctx,
	}
}

// NewGetNetworksNetworkIDDescriptionParamsWithHTTPClient creates a new GetNetworksNetworkIDDescriptionParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetNetworksNetworkIDDescriptionParamsWithHTTPClient(client *http.Client) *GetNetworksNetworkIDDescriptionParams {
	return &GetNetworksNetworkIDDescriptionParams{
		HTTPClient: client,
	}
}

/* GetNetworksNetworkIDDescriptionParams contains all the parameters to send to the API endpoint
   for the get networks network ID description operation.

   Typically these are written to a http.Request.
*/
type GetNetworksNetworkIDDescriptionParams struct {

	/* NetworkID.

	   Network ID
	*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get networks network ID description params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworksNetworkIDDescriptionParams) WithDefaults() *GetNetworksNetworkIDDescriptionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get networks network ID description params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworksNetworkIDDescriptionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get networks network ID description params
func (o *GetNetworksNetworkIDDescriptionParams) WithTimeout(timeout time.Duration) *GetNetworksNetworkIDDescriptionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get networks network ID description params
func (o *GetNetworksNetworkIDDescriptionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get networks network ID description params
func (o *GetNetworksNetworkIDDescriptionParams) WithContext(ctx context.Context) *GetNetworksNetworkIDDescriptionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get networks network ID description params
func (o *GetNetworksNetworkIDDescriptionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get networks network ID description params
func (o *GetNetworksNetworkIDDescriptionParams) WithHTTPClient(client *http.Client) *GetNetworksNetworkIDDescriptionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get networks network ID description params
func (o *GetNetworksNetworkIDDescriptionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the get networks network ID description params
func (o *GetNetworksNetworkIDDescriptionParams) WithNetworkID(networkID string) *GetNetworksNetworkIDDescriptionParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get networks network ID description params
func (o *GetNetworksNetworkIDDescriptionParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetworksNetworkIDDescriptionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param network_id
	if err := r.SetPathParam("network_id", o.NetworkID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
