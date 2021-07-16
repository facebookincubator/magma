// Code generated by go-swagger; DO NOT EDIT.

package policies

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

// NewPutNetworksNetworkIDPoliciesBaseNamesBaseNameParams creates a new PutNetworksNetworkIDPoliciesBaseNamesBaseNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutNetworksNetworkIDPoliciesBaseNamesBaseNameParams() *PutNetworksNetworkIDPoliciesBaseNamesBaseNameParams {
	return &PutNetworksNetworkIDPoliciesBaseNamesBaseNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutNetworksNetworkIDPoliciesBaseNamesBaseNameParamsWithTimeout creates a new PutNetworksNetworkIDPoliciesBaseNamesBaseNameParams object
// with the ability to set a timeout on a request.
func NewPutNetworksNetworkIDPoliciesBaseNamesBaseNameParamsWithTimeout(timeout time.Duration) *PutNetworksNetworkIDPoliciesBaseNamesBaseNameParams {
	return &PutNetworksNetworkIDPoliciesBaseNamesBaseNameParams{
		timeout: timeout,
	}
}

// NewPutNetworksNetworkIDPoliciesBaseNamesBaseNameParamsWithContext creates a new PutNetworksNetworkIDPoliciesBaseNamesBaseNameParams object
// with the ability to set a context for a request.
func NewPutNetworksNetworkIDPoliciesBaseNamesBaseNameParamsWithContext(ctx context.Context) *PutNetworksNetworkIDPoliciesBaseNamesBaseNameParams {
	return &PutNetworksNetworkIDPoliciesBaseNamesBaseNameParams{
		Context: ctx,
	}
}

// NewPutNetworksNetworkIDPoliciesBaseNamesBaseNameParamsWithHTTPClient creates a new PutNetworksNetworkIDPoliciesBaseNamesBaseNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutNetworksNetworkIDPoliciesBaseNamesBaseNameParamsWithHTTPClient(client *http.Client) *PutNetworksNetworkIDPoliciesBaseNamesBaseNameParams {
	return &PutNetworksNetworkIDPoliciesBaseNamesBaseNameParams{
		HTTPClient: client,
	}
}

/* PutNetworksNetworkIDPoliciesBaseNamesBaseNameParams contains all the parameters to send to the API endpoint
   for the put networks network ID policies base names base name operation.

   Typically these are written to a http.Request.
*/
type PutNetworksNetworkIDPoliciesBaseNamesBaseNameParams struct {

	/* BaseName.

	   Charging Rule Base Name
	*/
	BaseName string

	/* BaseNameRecord.

	   Charging Rule Base Name
	*/
	BaseNameRecord *models.BaseNameRecord

	/* NetworkID.

	   Network ID
	*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put networks network ID policies base names base name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutNetworksNetworkIDPoliciesBaseNamesBaseNameParams) WithDefaults() *PutNetworksNetworkIDPoliciesBaseNamesBaseNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put networks network ID policies base names base name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutNetworksNetworkIDPoliciesBaseNamesBaseNameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put networks network ID policies base names base name params
func (o *PutNetworksNetworkIDPoliciesBaseNamesBaseNameParams) WithTimeout(timeout time.Duration) *PutNetworksNetworkIDPoliciesBaseNamesBaseNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put networks network ID policies base names base name params
func (o *PutNetworksNetworkIDPoliciesBaseNamesBaseNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put networks network ID policies base names base name params
func (o *PutNetworksNetworkIDPoliciesBaseNamesBaseNameParams) WithContext(ctx context.Context) *PutNetworksNetworkIDPoliciesBaseNamesBaseNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put networks network ID policies base names base name params
func (o *PutNetworksNetworkIDPoliciesBaseNamesBaseNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put networks network ID policies base names base name params
func (o *PutNetworksNetworkIDPoliciesBaseNamesBaseNameParams) WithHTTPClient(client *http.Client) *PutNetworksNetworkIDPoliciesBaseNamesBaseNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put networks network ID policies base names base name params
func (o *PutNetworksNetworkIDPoliciesBaseNamesBaseNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBaseName adds the baseName to the put networks network ID policies base names base name params
func (o *PutNetworksNetworkIDPoliciesBaseNamesBaseNameParams) WithBaseName(baseName string) *PutNetworksNetworkIDPoliciesBaseNamesBaseNameParams {
	o.SetBaseName(baseName)
	return o
}

// SetBaseName adds the baseName to the put networks network ID policies base names base name params
func (o *PutNetworksNetworkIDPoliciesBaseNamesBaseNameParams) SetBaseName(baseName string) {
	o.BaseName = baseName
}

// WithBaseNameRecord adds the baseNameRecord to the put networks network ID policies base names base name params
func (o *PutNetworksNetworkIDPoliciesBaseNamesBaseNameParams) WithBaseNameRecord(baseNameRecord *models.BaseNameRecord) *PutNetworksNetworkIDPoliciesBaseNamesBaseNameParams {
	o.SetBaseNameRecord(baseNameRecord)
	return o
}

// SetBaseNameRecord adds the baseNameRecord to the put networks network ID policies base names base name params
func (o *PutNetworksNetworkIDPoliciesBaseNamesBaseNameParams) SetBaseNameRecord(baseNameRecord *models.BaseNameRecord) {
	o.BaseNameRecord = baseNameRecord
}

// WithNetworkID adds the networkID to the put networks network ID policies base names base name params
func (o *PutNetworksNetworkIDPoliciesBaseNamesBaseNameParams) WithNetworkID(networkID string) *PutNetworksNetworkIDPoliciesBaseNamesBaseNameParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the put networks network ID policies base names base name params
func (o *PutNetworksNetworkIDPoliciesBaseNamesBaseNameParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *PutNetworksNetworkIDPoliciesBaseNamesBaseNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param base_name
	if err := r.SetPathParam("base_name", o.BaseName); err != nil {
		return err
	}
	if o.BaseNameRecord != nil {
		if err := r.SetBodyParam(o.BaseNameRecord); err != nil {
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
