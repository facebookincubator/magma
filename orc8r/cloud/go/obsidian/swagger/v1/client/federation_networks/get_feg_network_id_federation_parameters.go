// Code generated by go-swagger; DO NOT EDIT.

package federation_networks

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

// NewGetFegNetworkIDFederationParams creates a new GetFegNetworkIDFederationParams object
// with the default values initialized.
func NewGetFegNetworkIDFederationParams() *GetFegNetworkIDFederationParams {
	var ()
	return &GetFegNetworkIDFederationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetFegNetworkIDFederationParamsWithTimeout creates a new GetFegNetworkIDFederationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetFegNetworkIDFederationParamsWithTimeout(timeout time.Duration) *GetFegNetworkIDFederationParams {
	var ()
	return &GetFegNetworkIDFederationParams{

		timeout: timeout,
	}
}

// NewGetFegNetworkIDFederationParamsWithContext creates a new GetFegNetworkIDFederationParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetFegNetworkIDFederationParamsWithContext(ctx context.Context) *GetFegNetworkIDFederationParams {
	var ()
	return &GetFegNetworkIDFederationParams{

		Context: ctx,
	}
}

// NewGetFegNetworkIDFederationParamsWithHTTPClient creates a new GetFegNetworkIDFederationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetFegNetworkIDFederationParamsWithHTTPClient(client *http.Client) *GetFegNetworkIDFederationParams {
	var ()
	return &GetFegNetworkIDFederationParams{
		HTTPClient: client,
	}
}

/*GetFegNetworkIDFederationParams contains all the parameters to send to the API endpoint
for the get feg network ID federation operation typically these are written to a http.Request
*/
type GetFegNetworkIDFederationParams struct {

	/*NetworkID
	  Network ID

	*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get feg network ID federation params
func (o *GetFegNetworkIDFederationParams) WithTimeout(timeout time.Duration) *GetFegNetworkIDFederationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get feg network ID federation params
func (o *GetFegNetworkIDFederationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get feg network ID federation params
func (o *GetFegNetworkIDFederationParams) WithContext(ctx context.Context) *GetFegNetworkIDFederationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get feg network ID federation params
func (o *GetFegNetworkIDFederationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get feg network ID federation params
func (o *GetFegNetworkIDFederationParams) WithHTTPClient(client *http.Client) *GetFegNetworkIDFederationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get feg network ID federation params
func (o *GetFegNetworkIDFederationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the get feg network ID federation params
func (o *GetFegNetworkIDFederationParams) WithNetworkID(networkID string) *GetFegNetworkIDFederationParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get feg network ID federation params
func (o *GetFegNetworkIDFederationParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *GetFegNetworkIDFederationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
