// Code generated by go-swagger; DO NOT EDIT.

package network_probes

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

// NewGetLTENetworkIDNetworkProbeDestinationsParams creates a new GetLTENetworkIDNetworkProbeDestinationsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetLTENetworkIDNetworkProbeDestinationsParams() *GetLTENetworkIDNetworkProbeDestinationsParams {
	return &GetLTENetworkIDNetworkProbeDestinationsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetLTENetworkIDNetworkProbeDestinationsParamsWithTimeout creates a new GetLTENetworkIDNetworkProbeDestinationsParams object
// with the ability to set a timeout on a request.
func NewGetLTENetworkIDNetworkProbeDestinationsParamsWithTimeout(timeout time.Duration) *GetLTENetworkIDNetworkProbeDestinationsParams {
	return &GetLTENetworkIDNetworkProbeDestinationsParams{
		timeout: timeout,
	}
}

// NewGetLTENetworkIDNetworkProbeDestinationsParamsWithContext creates a new GetLTENetworkIDNetworkProbeDestinationsParams object
// with the ability to set a context for a request.
func NewGetLTENetworkIDNetworkProbeDestinationsParamsWithContext(ctx context.Context) *GetLTENetworkIDNetworkProbeDestinationsParams {
	return &GetLTENetworkIDNetworkProbeDestinationsParams{
		Context: ctx,
	}
}

// NewGetLTENetworkIDNetworkProbeDestinationsParamsWithHTTPClient creates a new GetLTENetworkIDNetworkProbeDestinationsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetLTENetworkIDNetworkProbeDestinationsParamsWithHTTPClient(client *http.Client) *GetLTENetworkIDNetworkProbeDestinationsParams {
	return &GetLTENetworkIDNetworkProbeDestinationsParams{
		HTTPClient: client,
	}
}

/* GetLTENetworkIDNetworkProbeDestinationsParams contains all the parameters to send to the API endpoint
   for the get LTE network ID network probe destinations operation.

   Typically these are written to a http.Request.
*/
type GetLTENetworkIDNetworkProbeDestinationsParams struct {

	/* NetworkID.

	   Network ID
	*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get LTE network ID network probe destinations params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLTENetworkIDNetworkProbeDestinationsParams) WithDefaults() *GetLTENetworkIDNetworkProbeDestinationsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get LTE network ID network probe destinations params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLTENetworkIDNetworkProbeDestinationsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get LTE network ID network probe destinations params
func (o *GetLTENetworkIDNetworkProbeDestinationsParams) WithTimeout(timeout time.Duration) *GetLTENetworkIDNetworkProbeDestinationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get LTE network ID network probe destinations params
func (o *GetLTENetworkIDNetworkProbeDestinationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get LTE network ID network probe destinations params
func (o *GetLTENetworkIDNetworkProbeDestinationsParams) WithContext(ctx context.Context) *GetLTENetworkIDNetworkProbeDestinationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get LTE network ID network probe destinations params
func (o *GetLTENetworkIDNetworkProbeDestinationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get LTE network ID network probe destinations params
func (o *GetLTENetworkIDNetworkProbeDestinationsParams) WithHTTPClient(client *http.Client) *GetLTENetworkIDNetworkProbeDestinationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get LTE network ID network probe destinations params
func (o *GetLTENetworkIDNetworkProbeDestinationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the get LTE network ID network probe destinations params
func (o *GetLTENetworkIDNetworkProbeDestinationsParams) WithNetworkID(networkID string) *GetLTENetworkIDNetworkProbeDestinationsParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get LTE network ID network probe destinations params
func (o *GetLTENetworkIDNetworkProbeDestinationsParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *GetLTENetworkIDNetworkProbeDestinationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
