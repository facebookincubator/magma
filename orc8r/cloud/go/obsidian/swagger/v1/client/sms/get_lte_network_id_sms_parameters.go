// Code generated by go-swagger; DO NOT EDIT.

package sms

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

// NewGetLTENetworkIDSMSParams creates a new GetLTENetworkIDSMSParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetLTENetworkIDSMSParams() *GetLTENetworkIDSMSParams {
	return &GetLTENetworkIDSMSParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetLTENetworkIDSMSParamsWithTimeout creates a new GetLTENetworkIDSMSParams object
// with the ability to set a timeout on a request.
func NewGetLTENetworkIDSMSParamsWithTimeout(timeout time.Duration) *GetLTENetworkIDSMSParams {
	return &GetLTENetworkIDSMSParams{
		timeout: timeout,
	}
}

// NewGetLTENetworkIDSMSParamsWithContext creates a new GetLTENetworkIDSMSParams object
// with the ability to set a context for a request.
func NewGetLTENetworkIDSMSParamsWithContext(ctx context.Context) *GetLTENetworkIDSMSParams {
	return &GetLTENetworkIDSMSParams{
		Context: ctx,
	}
}

// NewGetLTENetworkIDSMSParamsWithHTTPClient creates a new GetLTENetworkIDSMSParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetLTENetworkIDSMSParamsWithHTTPClient(client *http.Client) *GetLTENetworkIDSMSParams {
	return &GetLTENetworkIDSMSParams{
		HTTPClient: client,
	}
}

/* GetLTENetworkIDSMSParams contains all the parameters to send to the API endpoint
   for the get LTE network ID SMS operation.

   Typically these are written to a http.Request.
*/
type GetLTENetworkIDSMSParams struct {

	/* NetworkID.

	   Network ID
	*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get LTE network ID SMS params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLTENetworkIDSMSParams) WithDefaults() *GetLTENetworkIDSMSParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get LTE network ID SMS params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLTENetworkIDSMSParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get LTE network ID SMS params
func (o *GetLTENetworkIDSMSParams) WithTimeout(timeout time.Duration) *GetLTENetworkIDSMSParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get LTE network ID SMS params
func (o *GetLTENetworkIDSMSParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get LTE network ID SMS params
func (o *GetLTENetworkIDSMSParams) WithContext(ctx context.Context) *GetLTENetworkIDSMSParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get LTE network ID SMS params
func (o *GetLTENetworkIDSMSParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get LTE network ID SMS params
func (o *GetLTENetworkIDSMSParams) WithHTTPClient(client *http.Client) *GetLTENetworkIDSMSParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get LTE network ID SMS params
func (o *GetLTENetworkIDSMSParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the get LTE network ID SMS params
func (o *GetLTENetworkIDSMSParams) WithNetworkID(networkID string) *GetLTENetworkIDSMSParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get LTE network ID SMS params
func (o *GetLTENetworkIDSMSParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *GetLTENetworkIDSMSParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
