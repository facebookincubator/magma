// Code generated by go-swagger; DO NOT EDIT.

package wifi_networks

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

// NewPutWifiNetworkIDNameParams creates a new PutWifiNetworkIDNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutWifiNetworkIDNameParams() *PutWifiNetworkIDNameParams {
	return &PutWifiNetworkIDNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutWifiNetworkIDNameParamsWithTimeout creates a new PutWifiNetworkIDNameParams object
// with the ability to set a timeout on a request.
func NewPutWifiNetworkIDNameParamsWithTimeout(timeout time.Duration) *PutWifiNetworkIDNameParams {
	return &PutWifiNetworkIDNameParams{
		timeout: timeout,
	}
}

// NewPutWifiNetworkIDNameParamsWithContext creates a new PutWifiNetworkIDNameParams object
// with the ability to set a context for a request.
func NewPutWifiNetworkIDNameParamsWithContext(ctx context.Context) *PutWifiNetworkIDNameParams {
	return &PutWifiNetworkIDNameParams{
		Context: ctx,
	}
}

// NewPutWifiNetworkIDNameParamsWithHTTPClient creates a new PutWifiNetworkIDNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutWifiNetworkIDNameParamsWithHTTPClient(client *http.Client) *PutWifiNetworkIDNameParams {
	return &PutWifiNetworkIDNameParams{
		HTTPClient: client,
	}
}

/* PutWifiNetworkIDNameParams contains all the parameters to send to the API endpoint
   for the put wifi network ID name operation.

   Typically these are written to a http.Request.
*/
type PutWifiNetworkIDNameParams struct {

	/* Name.

	   New name for the network
	*/
	Name models.NetworkName

	/* NetworkID.

	   Network ID
	*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put wifi network ID name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutWifiNetworkIDNameParams) WithDefaults() *PutWifiNetworkIDNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put wifi network ID name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutWifiNetworkIDNameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put wifi network ID name params
func (o *PutWifiNetworkIDNameParams) WithTimeout(timeout time.Duration) *PutWifiNetworkIDNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put wifi network ID name params
func (o *PutWifiNetworkIDNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put wifi network ID name params
func (o *PutWifiNetworkIDNameParams) WithContext(ctx context.Context) *PutWifiNetworkIDNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put wifi network ID name params
func (o *PutWifiNetworkIDNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put wifi network ID name params
func (o *PutWifiNetworkIDNameParams) WithHTTPClient(client *http.Client) *PutWifiNetworkIDNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put wifi network ID name params
func (o *PutWifiNetworkIDNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the put wifi network ID name params
func (o *PutWifiNetworkIDNameParams) WithName(name models.NetworkName) *PutWifiNetworkIDNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the put wifi network ID name params
func (o *PutWifiNetworkIDNameParams) SetName(name models.NetworkName) {
	o.Name = name
}

// WithNetworkID adds the networkID to the put wifi network ID name params
func (o *PutWifiNetworkIDNameParams) WithNetworkID(networkID string) *PutWifiNetworkIDNameParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the put wifi network ID name params
func (o *PutWifiNetworkIDNameParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *PutWifiNetworkIDNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.Name); err != nil {
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
