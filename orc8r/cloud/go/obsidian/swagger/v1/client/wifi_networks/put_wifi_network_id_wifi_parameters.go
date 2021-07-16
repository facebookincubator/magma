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

// NewPutWifiNetworkIDWifiParams creates a new PutWifiNetworkIDWifiParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutWifiNetworkIDWifiParams() *PutWifiNetworkIDWifiParams {
	return &PutWifiNetworkIDWifiParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutWifiNetworkIDWifiParamsWithTimeout creates a new PutWifiNetworkIDWifiParams object
// with the ability to set a timeout on a request.
func NewPutWifiNetworkIDWifiParamsWithTimeout(timeout time.Duration) *PutWifiNetworkIDWifiParams {
	return &PutWifiNetworkIDWifiParams{
		timeout: timeout,
	}
}

// NewPutWifiNetworkIDWifiParamsWithContext creates a new PutWifiNetworkIDWifiParams object
// with the ability to set a context for a request.
func NewPutWifiNetworkIDWifiParamsWithContext(ctx context.Context) *PutWifiNetworkIDWifiParams {
	return &PutWifiNetworkIDWifiParams{
		Context: ctx,
	}
}

// NewPutWifiNetworkIDWifiParamsWithHTTPClient creates a new PutWifiNetworkIDWifiParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutWifiNetworkIDWifiParamsWithHTTPClient(client *http.Client) *PutWifiNetworkIDWifiParams {
	return &PutWifiNetworkIDWifiParams{
		HTTPClient: client,
	}
}

/* PutWifiNetworkIDWifiParams contains all the parameters to send to the API endpoint
   for the put wifi network ID wifi operation.

   Typically these are written to a http.Request.
*/
type PutWifiNetworkIDWifiParams struct {

	/* Config.

	   New wifi configuration for the network
	*/
	Config *models.NetworkWifiConfigs

	/* NetworkID.

	   Network ID
	*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put wifi network ID wifi params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutWifiNetworkIDWifiParams) WithDefaults() *PutWifiNetworkIDWifiParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put wifi network ID wifi params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutWifiNetworkIDWifiParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put wifi network ID wifi params
func (o *PutWifiNetworkIDWifiParams) WithTimeout(timeout time.Duration) *PutWifiNetworkIDWifiParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put wifi network ID wifi params
func (o *PutWifiNetworkIDWifiParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put wifi network ID wifi params
func (o *PutWifiNetworkIDWifiParams) WithContext(ctx context.Context) *PutWifiNetworkIDWifiParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put wifi network ID wifi params
func (o *PutWifiNetworkIDWifiParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put wifi network ID wifi params
func (o *PutWifiNetworkIDWifiParams) WithHTTPClient(client *http.Client) *PutWifiNetworkIDWifiParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put wifi network ID wifi params
func (o *PutWifiNetworkIDWifiParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConfig adds the config to the put wifi network ID wifi params
func (o *PutWifiNetworkIDWifiParams) WithConfig(config *models.NetworkWifiConfigs) *PutWifiNetworkIDWifiParams {
	o.SetConfig(config)
	return o
}

// SetConfig adds the config to the put wifi network ID wifi params
func (o *PutWifiNetworkIDWifiParams) SetConfig(config *models.NetworkWifiConfigs) {
	o.Config = config
}

// WithNetworkID adds the networkID to the put wifi network ID wifi params
func (o *PutWifiNetworkIDWifiParams) WithNetworkID(networkID string) *PutWifiNetworkIDWifiParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the put wifi network ID wifi params
func (o *PutWifiNetworkIDWifiParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *PutWifiNetworkIDWifiParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Config != nil {
		if err := r.SetBodyParam(o.Config); err != nil {
			return err
		}
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
