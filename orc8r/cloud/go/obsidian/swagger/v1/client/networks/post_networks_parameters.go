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

// NewPostNetworksParams creates a new PostNetworksParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostNetworksParams() *PostNetworksParams {
	return &PostNetworksParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostNetworksParamsWithTimeout creates a new PostNetworksParams object
// with the ability to set a timeout on a request.
func NewPostNetworksParamsWithTimeout(timeout time.Duration) *PostNetworksParams {
	return &PostNetworksParams{
		timeout: timeout,
	}
}

// NewPostNetworksParamsWithContext creates a new PostNetworksParams object
// with the ability to set a context for a request.
func NewPostNetworksParamsWithContext(ctx context.Context) *PostNetworksParams {
	return &PostNetworksParams{
		Context: ctx,
	}
}

// NewPostNetworksParamsWithHTTPClient creates a new PostNetworksParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostNetworksParamsWithHTTPClient(client *http.Client) *PostNetworksParams {
	return &PostNetworksParams{
		HTTPClient: client,
	}
}

/* PostNetworksParams contains all the parameters to send to the API endpoint
   for the post networks operation.

   Typically these are written to a http.Request.
*/
type PostNetworksParams struct {

	/* Network.

	   Configuration of the network to create
	*/
	Network *models.Network

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post networks params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostNetworksParams) WithDefaults() *PostNetworksParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post networks params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostNetworksParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post networks params
func (o *PostNetworksParams) WithTimeout(timeout time.Duration) *PostNetworksParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post networks params
func (o *PostNetworksParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post networks params
func (o *PostNetworksParams) WithContext(ctx context.Context) *PostNetworksParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post networks params
func (o *PostNetworksParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post networks params
func (o *PostNetworksParams) WithHTTPClient(client *http.Client) *PostNetworksParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post networks params
func (o *PostNetworksParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetwork adds the network to the post networks params
func (o *PostNetworksParams) WithNetwork(network *models.Network) *PostNetworksParams {
	o.SetNetwork(network)
	return o
}

// SetNetwork adds the network to the post networks params
func (o *PostNetworksParams) SetNetwork(network *models.Network) {
	o.Network = network
}

// WriteToRequest writes these params to a swagger request
func (o *PostNetworksParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Network != nil {
		if err := r.SetBodyParam(o.Network); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
