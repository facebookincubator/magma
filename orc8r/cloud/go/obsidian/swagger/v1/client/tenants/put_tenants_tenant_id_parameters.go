// Code generated by go-swagger; DO NOT EDIT.

package tenants

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
	"github.com/go-openapi/swag"

	"magma/orc8r/cloud/go/obsidian/swagger/v1/models"
)

// NewPutTenantsTenantIDParams creates a new PutTenantsTenantIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutTenantsTenantIDParams() *PutTenantsTenantIDParams {
	return &PutTenantsTenantIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutTenantsTenantIDParamsWithTimeout creates a new PutTenantsTenantIDParams object
// with the ability to set a timeout on a request.
func NewPutTenantsTenantIDParamsWithTimeout(timeout time.Duration) *PutTenantsTenantIDParams {
	return &PutTenantsTenantIDParams{
		timeout: timeout,
	}
}

// NewPutTenantsTenantIDParamsWithContext creates a new PutTenantsTenantIDParams object
// with the ability to set a context for a request.
func NewPutTenantsTenantIDParamsWithContext(ctx context.Context) *PutTenantsTenantIDParams {
	return &PutTenantsTenantIDParams{
		Context: ctx,
	}
}

// NewPutTenantsTenantIDParamsWithHTTPClient creates a new PutTenantsTenantIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutTenantsTenantIDParamsWithHTTPClient(client *http.Client) *PutTenantsTenantIDParams {
	return &PutTenantsTenantIDParams{
		HTTPClient: client,
	}
}

/* PutTenantsTenantIDParams contains all the parameters to send to the API endpoint
   for the put tenants tenant ID operation.

   Typically these are written to a http.Request.
*/
type PutTenantsTenantIDParams struct {

	/* Tenant.

	   Tenant to be updated
	*/
	Tenant *models.Tenant

	/* TenantID.

	   Tenant ID
	*/
	TenantID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put tenants tenant ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutTenantsTenantIDParams) WithDefaults() *PutTenantsTenantIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put tenants tenant ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutTenantsTenantIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put tenants tenant ID params
func (o *PutTenantsTenantIDParams) WithTimeout(timeout time.Duration) *PutTenantsTenantIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put tenants tenant ID params
func (o *PutTenantsTenantIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put tenants tenant ID params
func (o *PutTenantsTenantIDParams) WithContext(ctx context.Context) *PutTenantsTenantIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put tenants tenant ID params
func (o *PutTenantsTenantIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put tenants tenant ID params
func (o *PutTenantsTenantIDParams) WithHTTPClient(client *http.Client) *PutTenantsTenantIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put tenants tenant ID params
func (o *PutTenantsTenantIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTenant adds the tenant to the put tenants tenant ID params
func (o *PutTenantsTenantIDParams) WithTenant(tenant *models.Tenant) *PutTenantsTenantIDParams {
	o.SetTenant(tenant)
	return o
}

// SetTenant adds the tenant to the put tenants tenant ID params
func (o *PutTenantsTenantIDParams) SetTenant(tenant *models.Tenant) {
	o.Tenant = tenant
}

// WithTenantID adds the tenantID to the put tenants tenant ID params
func (o *PutTenantsTenantIDParams) WithTenantID(tenantID int64) *PutTenantsTenantIDParams {
	o.SetTenantID(tenantID)
	return o
}

// SetTenantID adds the tenantId to the put tenants tenant ID params
func (o *PutTenantsTenantIDParams) SetTenantID(tenantID int64) {
	o.TenantID = tenantID
}

// WriteToRequest writes these params to a swagger request
func (o *PutTenantsTenantIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Tenant != nil {
		if err := r.SetBodyParam(o.Tenant); err != nil {
			return err
		}
	}

	// path param tenant_id
	if err := r.SetPathParam("tenant_id", swag.FormatInt64(o.TenantID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
