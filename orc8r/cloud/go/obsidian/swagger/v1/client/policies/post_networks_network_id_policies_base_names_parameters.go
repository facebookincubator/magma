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

// NewPostNetworksNetworkIDPoliciesBaseNamesParams creates a new PostNetworksNetworkIDPoliciesBaseNamesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostNetworksNetworkIDPoliciesBaseNamesParams() *PostNetworksNetworkIDPoliciesBaseNamesParams {
	return &PostNetworksNetworkIDPoliciesBaseNamesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostNetworksNetworkIDPoliciesBaseNamesParamsWithTimeout creates a new PostNetworksNetworkIDPoliciesBaseNamesParams object
// with the ability to set a timeout on a request.
func NewPostNetworksNetworkIDPoliciesBaseNamesParamsWithTimeout(timeout time.Duration) *PostNetworksNetworkIDPoliciesBaseNamesParams {
	return &PostNetworksNetworkIDPoliciesBaseNamesParams{
		timeout: timeout,
	}
}

// NewPostNetworksNetworkIDPoliciesBaseNamesParamsWithContext creates a new PostNetworksNetworkIDPoliciesBaseNamesParams object
// with the ability to set a context for a request.
func NewPostNetworksNetworkIDPoliciesBaseNamesParamsWithContext(ctx context.Context) *PostNetworksNetworkIDPoliciesBaseNamesParams {
	return &PostNetworksNetworkIDPoliciesBaseNamesParams{
		Context: ctx,
	}
}

// NewPostNetworksNetworkIDPoliciesBaseNamesParamsWithHTTPClient creates a new PostNetworksNetworkIDPoliciesBaseNamesParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostNetworksNetworkIDPoliciesBaseNamesParamsWithHTTPClient(client *http.Client) *PostNetworksNetworkIDPoliciesBaseNamesParams {
	return &PostNetworksNetworkIDPoliciesBaseNamesParams{
		HTTPClient: client,
	}
}

/* PostNetworksNetworkIDPoliciesBaseNamesParams contains all the parameters to send to the API endpoint
   for the post networks network ID policies base names operation.

   Typically these are written to a http.Request.
*/
type PostNetworksNetworkIDPoliciesBaseNamesParams struct {

	/* BaseNameRecord.

	   Charging Rule Base Name to add
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

// WithDefaults hydrates default values in the post networks network ID policies base names params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostNetworksNetworkIDPoliciesBaseNamesParams) WithDefaults() *PostNetworksNetworkIDPoliciesBaseNamesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post networks network ID policies base names params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostNetworksNetworkIDPoliciesBaseNamesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post networks network ID policies base names params
func (o *PostNetworksNetworkIDPoliciesBaseNamesParams) WithTimeout(timeout time.Duration) *PostNetworksNetworkIDPoliciesBaseNamesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post networks network ID policies base names params
func (o *PostNetworksNetworkIDPoliciesBaseNamesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post networks network ID policies base names params
func (o *PostNetworksNetworkIDPoliciesBaseNamesParams) WithContext(ctx context.Context) *PostNetworksNetworkIDPoliciesBaseNamesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post networks network ID policies base names params
func (o *PostNetworksNetworkIDPoliciesBaseNamesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post networks network ID policies base names params
func (o *PostNetworksNetworkIDPoliciesBaseNamesParams) WithHTTPClient(client *http.Client) *PostNetworksNetworkIDPoliciesBaseNamesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post networks network ID policies base names params
func (o *PostNetworksNetworkIDPoliciesBaseNamesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBaseNameRecord adds the baseNameRecord to the post networks network ID policies base names params
func (o *PostNetworksNetworkIDPoliciesBaseNamesParams) WithBaseNameRecord(baseNameRecord *models.BaseNameRecord) *PostNetworksNetworkIDPoliciesBaseNamesParams {
	o.SetBaseNameRecord(baseNameRecord)
	return o
}

// SetBaseNameRecord adds the baseNameRecord to the post networks network ID policies base names params
func (o *PostNetworksNetworkIDPoliciesBaseNamesParams) SetBaseNameRecord(baseNameRecord *models.BaseNameRecord) {
	o.BaseNameRecord = baseNameRecord
}

// WithNetworkID adds the networkID to the post networks network ID policies base names params
func (o *PostNetworksNetworkIDPoliciesBaseNamesParams) WithNetworkID(networkID string) *PostNetworksNetworkIDPoliciesBaseNamesParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the post networks network ID policies base names params
func (o *PostNetworksNetworkIDPoliciesBaseNamesParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *PostNetworksNetworkIDPoliciesBaseNamesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
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
