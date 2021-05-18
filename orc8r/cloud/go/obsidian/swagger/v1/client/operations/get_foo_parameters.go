// Code generated by go-swagger; DO NOT EDIT.

package operations

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

// NewGetFooParams creates a new GetFooParams object
// with the default values initialized.
func NewGetFooParams() *GetFooParams {

	return &GetFooParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetFooParamsWithTimeout creates a new GetFooParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetFooParamsWithTimeout(timeout time.Duration) *GetFooParams {

	return &GetFooParams{

		timeout: timeout,
	}
}

// NewGetFooParamsWithContext creates a new GetFooParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetFooParamsWithContext(ctx context.Context) *GetFooParams {

	return &GetFooParams{

		Context: ctx,
	}
}

// NewGetFooParamsWithHTTPClient creates a new GetFooParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetFooParamsWithHTTPClient(client *http.Client) *GetFooParams {

	return &GetFooParams{
		HTTPClient: client,
	}
}

/*GetFooParams contains all the parameters to send to the API endpoint
for the get foo operation typically these are written to a http.Request
*/
type GetFooParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get foo params
func (o *GetFooParams) WithTimeout(timeout time.Duration) *GetFooParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get foo params
func (o *GetFooParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get foo params
func (o *GetFooParams) WithContext(ctx context.Context) *GetFooParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get foo params
func (o *GetFooParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get foo params
func (o *GetFooParams) WithHTTPClient(client *http.Client) *GetFooParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get foo params
func (o *GetFooParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetFooParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
