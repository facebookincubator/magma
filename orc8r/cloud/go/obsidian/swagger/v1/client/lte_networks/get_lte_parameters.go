// Code generated by go-swagger; DO NOT EDIT.

package lte_networks

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

// NewGetLTEParams creates a new GetLTEParams object
// with the default values initialized.
func NewGetLTEParams() *GetLTEParams {

	return &GetLTEParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLTEParamsWithTimeout creates a new GetLTEParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLTEParamsWithTimeout(timeout time.Duration) *GetLTEParams {

	return &GetLTEParams{

		timeout: timeout,
	}
}

// NewGetLTEParamsWithContext creates a new GetLTEParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLTEParamsWithContext(ctx context.Context) *GetLTEParams {

	return &GetLTEParams{

		Context: ctx,
	}
}

// NewGetLTEParamsWithHTTPClient creates a new GetLTEParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLTEParamsWithHTTPClient(client *http.Client) *GetLTEParams {

	return &GetLTEParams{
		HTTPClient: client,
	}
}

/*GetLTEParams contains all the parameters to send to the API endpoint
for the get LTE operation typically these are written to a http.Request
*/
type GetLTEParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get LTE params
func (o *GetLTEParams) WithTimeout(timeout time.Duration) *GetLTEParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get LTE params
func (o *GetLTEParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get LTE params
func (o *GetLTEParams) WithContext(ctx context.Context) *GetLTEParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get LTE params
func (o *GetLTEParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get LTE params
func (o *GetLTEParams) WithHTTPClient(client *http.Client) *GetLTEParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get LTE params
func (o *GetLTEParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetLTEParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
