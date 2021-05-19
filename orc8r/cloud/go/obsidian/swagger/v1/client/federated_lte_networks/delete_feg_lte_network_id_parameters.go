// Code generated by go-swagger; DO NOT EDIT.

package federated_lte_networks

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

// NewDeleteFegLTENetworkIDParams creates a new DeleteFegLTENetworkIDParams object
// with the default values initialized.
func NewDeleteFegLTENetworkIDParams() *DeleteFegLTENetworkIDParams {
	var ()
	return &DeleteFegLTENetworkIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteFegLTENetworkIDParamsWithTimeout creates a new DeleteFegLTENetworkIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteFegLTENetworkIDParamsWithTimeout(timeout time.Duration) *DeleteFegLTENetworkIDParams {
	var ()
	return &DeleteFegLTENetworkIDParams{

		timeout: timeout,
	}
}

// NewDeleteFegLTENetworkIDParamsWithContext creates a new DeleteFegLTENetworkIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteFegLTENetworkIDParamsWithContext(ctx context.Context) *DeleteFegLTENetworkIDParams {
	var ()
	return &DeleteFegLTENetworkIDParams{

		Context: ctx,
	}
}

// NewDeleteFegLTENetworkIDParamsWithHTTPClient creates a new DeleteFegLTENetworkIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteFegLTENetworkIDParamsWithHTTPClient(client *http.Client) *DeleteFegLTENetworkIDParams {
	var ()
	return &DeleteFegLTENetworkIDParams{
		HTTPClient: client,
	}
}

/*DeleteFegLTENetworkIDParams contains all the parameters to send to the API endpoint
for the delete feg LTE network ID operation typically these are written to a http.Request
*/
type DeleteFegLTENetworkIDParams struct {

	/*NetworkID
	  Network ID

	*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete feg LTE network ID params
func (o *DeleteFegLTENetworkIDParams) WithTimeout(timeout time.Duration) *DeleteFegLTENetworkIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete feg LTE network ID params
func (o *DeleteFegLTENetworkIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete feg LTE network ID params
func (o *DeleteFegLTENetworkIDParams) WithContext(ctx context.Context) *DeleteFegLTENetworkIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete feg LTE network ID params
func (o *DeleteFegLTENetworkIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete feg LTE network ID params
func (o *DeleteFegLTENetworkIDParams) WithHTTPClient(client *http.Client) *DeleteFegLTENetworkIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete feg LTE network ID params
func (o *DeleteFegLTENetworkIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the delete feg LTE network ID params
func (o *DeleteFegLTENetworkIDParams) WithNetworkID(networkID string) *DeleteFegLTENetworkIDParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the delete feg LTE network ID params
func (o *DeleteFegLTENetworkIDParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteFegLTENetworkIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
