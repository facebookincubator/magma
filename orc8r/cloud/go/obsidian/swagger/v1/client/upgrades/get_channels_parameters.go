// Code generated by go-swagger; DO NOT EDIT.

package upgrades

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

// NewGetChannelsParams creates a new GetChannelsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetChannelsParams() *GetChannelsParams {
	return &GetChannelsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetChannelsParamsWithTimeout creates a new GetChannelsParams object
// with the ability to set a timeout on a request.
func NewGetChannelsParamsWithTimeout(timeout time.Duration) *GetChannelsParams {
	return &GetChannelsParams{
		timeout: timeout,
	}
}

// NewGetChannelsParamsWithContext creates a new GetChannelsParams object
// with the ability to set a context for a request.
func NewGetChannelsParamsWithContext(ctx context.Context) *GetChannelsParams {
	return &GetChannelsParams{
		Context: ctx,
	}
}

// NewGetChannelsParamsWithHTTPClient creates a new GetChannelsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetChannelsParamsWithHTTPClient(client *http.Client) *GetChannelsParams {
	return &GetChannelsParams{
		HTTPClient: client,
	}
}

/* GetChannelsParams contains all the parameters to send to the API endpoint
   for the get channels operation.

   Typically these are written to a http.Request.
*/
type GetChannelsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get channels params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetChannelsParams) WithDefaults() *GetChannelsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get channels params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetChannelsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get channels params
func (o *GetChannelsParams) WithTimeout(timeout time.Duration) *GetChannelsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get channels params
func (o *GetChannelsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get channels params
func (o *GetChannelsParams) WithContext(ctx context.Context) *GetChannelsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get channels params
func (o *GetChannelsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get channels params
func (o *GetChannelsParams) WithHTTPClient(client *http.Client) *GetChannelsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get channels params
func (o *GetChannelsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetChannelsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
