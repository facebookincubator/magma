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

// NewPutNetworksNetworkIDDNSRecordsParams creates a new PutNetworksNetworkIDDNSRecordsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutNetworksNetworkIDDNSRecordsParams() *PutNetworksNetworkIDDNSRecordsParams {
	return &PutNetworksNetworkIDDNSRecordsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutNetworksNetworkIDDNSRecordsParamsWithTimeout creates a new PutNetworksNetworkIDDNSRecordsParams object
// with the ability to set a timeout on a request.
func NewPutNetworksNetworkIDDNSRecordsParamsWithTimeout(timeout time.Duration) *PutNetworksNetworkIDDNSRecordsParams {
	return &PutNetworksNetworkIDDNSRecordsParams{
		timeout: timeout,
	}
}

// NewPutNetworksNetworkIDDNSRecordsParamsWithContext creates a new PutNetworksNetworkIDDNSRecordsParams object
// with the ability to set a context for a request.
func NewPutNetworksNetworkIDDNSRecordsParamsWithContext(ctx context.Context) *PutNetworksNetworkIDDNSRecordsParams {
	return &PutNetworksNetworkIDDNSRecordsParams{
		Context: ctx,
	}
}

// NewPutNetworksNetworkIDDNSRecordsParamsWithHTTPClient creates a new PutNetworksNetworkIDDNSRecordsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutNetworksNetworkIDDNSRecordsParamsWithHTTPClient(client *http.Client) *PutNetworksNetworkIDDNSRecordsParams {
	return &PutNetworksNetworkIDDNSRecordsParams{
		HTTPClient: client,
	}
}

/* PutNetworksNetworkIDDNSRecordsParams contains all the parameters to send to the API endpoint
   for the put networks network ID DNS records operation.

   Typically these are written to a http.Request.
*/
type PutNetworksNetworkIDDNSRecordsParams struct {

	/* NetworkID.

	   Network ID
	*/
	NetworkID string

	/* Records.

	   Custom DNS records for the network
	*/
	Records models.NetworkDNSRecords

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put networks network ID DNS records params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutNetworksNetworkIDDNSRecordsParams) WithDefaults() *PutNetworksNetworkIDDNSRecordsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put networks network ID DNS records params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutNetworksNetworkIDDNSRecordsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put networks network ID DNS records params
func (o *PutNetworksNetworkIDDNSRecordsParams) WithTimeout(timeout time.Duration) *PutNetworksNetworkIDDNSRecordsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put networks network ID DNS records params
func (o *PutNetworksNetworkIDDNSRecordsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put networks network ID DNS records params
func (o *PutNetworksNetworkIDDNSRecordsParams) WithContext(ctx context.Context) *PutNetworksNetworkIDDNSRecordsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put networks network ID DNS records params
func (o *PutNetworksNetworkIDDNSRecordsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put networks network ID DNS records params
func (o *PutNetworksNetworkIDDNSRecordsParams) WithHTTPClient(client *http.Client) *PutNetworksNetworkIDDNSRecordsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put networks network ID DNS records params
func (o *PutNetworksNetworkIDDNSRecordsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the put networks network ID DNS records params
func (o *PutNetworksNetworkIDDNSRecordsParams) WithNetworkID(networkID string) *PutNetworksNetworkIDDNSRecordsParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the put networks network ID DNS records params
func (o *PutNetworksNetworkIDDNSRecordsParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithRecords adds the records to the put networks network ID DNS records params
func (o *PutNetworksNetworkIDDNSRecordsParams) WithRecords(records models.NetworkDNSRecords) *PutNetworksNetworkIDDNSRecordsParams {
	o.SetRecords(records)
	return o
}

// SetRecords adds the records to the put networks network ID DNS records params
func (o *PutNetworksNetworkIDDNSRecordsParams) SetRecords(records models.NetworkDNSRecords) {
	o.Records = records
}

// WriteToRequest writes these params to a swagger request
func (o *PutNetworksNetworkIDDNSRecordsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param network_id
	if err := r.SetPathParam("network_id", o.NetworkID); err != nil {
		return err
	}
	if o.Records != nil {
		if err := r.SetBodyParam(o.Records); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
