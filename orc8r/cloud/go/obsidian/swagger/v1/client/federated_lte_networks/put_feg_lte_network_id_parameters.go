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

	models "magma/orc8r/cloud/go/obsidian/swagger/v1/models"
)

// NewPutFegLTENetworkIDParams creates a new PutFegLTENetworkIDParams object
// with the default values initialized.
func NewPutFegLTENetworkIDParams() *PutFegLTENetworkIDParams {
	var ()
	return &PutFegLTENetworkIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutFegLTENetworkIDParamsWithTimeout creates a new PutFegLTENetworkIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutFegLTENetworkIDParamsWithTimeout(timeout time.Duration) *PutFegLTENetworkIDParams {
	var ()
	return &PutFegLTENetworkIDParams{

		timeout: timeout,
	}
}

// NewPutFegLTENetworkIDParamsWithContext creates a new PutFegLTENetworkIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutFegLTENetworkIDParamsWithContext(ctx context.Context) *PutFegLTENetworkIDParams {
	var ()
	return &PutFegLTENetworkIDParams{

		Context: ctx,
	}
}

// NewPutFegLTENetworkIDParamsWithHTTPClient creates a new PutFegLTENetworkIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutFegLTENetworkIDParamsWithHTTPClient(client *http.Client) *PutFegLTENetworkIDParams {
	var ()
	return &PutFegLTENetworkIDParams{
		HTTPClient: client,
	}
}

/*PutFegLTENetworkIDParams contains all the parameters to send to the API endpoint
for the put feg LTE network ID operation typically these are written to a http.Request
*/
type PutFegLTENetworkIDParams struct {

	/*LTENetwork
	  Full desired configuration of the network

	*/
	LTENetwork *models.FegLTENetwork
	/*NetworkID
	  Network ID

	*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put feg LTE network ID params
func (o *PutFegLTENetworkIDParams) WithTimeout(timeout time.Duration) *PutFegLTENetworkIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put feg LTE network ID params
func (o *PutFegLTENetworkIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put feg LTE network ID params
func (o *PutFegLTENetworkIDParams) WithContext(ctx context.Context) *PutFegLTENetworkIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put feg LTE network ID params
func (o *PutFegLTENetworkIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put feg LTE network ID params
func (o *PutFegLTENetworkIDParams) WithHTTPClient(client *http.Client) *PutFegLTENetworkIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put feg LTE network ID params
func (o *PutFegLTENetworkIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLTENetwork adds the lTENetwork to the put feg LTE network ID params
func (o *PutFegLTENetworkIDParams) WithLTENetwork(lTENetwork *models.FegLTENetwork) *PutFegLTENetworkIDParams {
	o.SetLTENetwork(lTENetwork)
	return o
}

// SetLTENetwork adds the lteNetwork to the put feg LTE network ID params
func (o *PutFegLTENetworkIDParams) SetLTENetwork(lTENetwork *models.FegLTENetwork) {
	o.LTENetwork = lTENetwork
}

// WithNetworkID adds the networkID to the put feg LTE network ID params
func (o *PutFegLTENetworkIDParams) WithNetworkID(networkID string) *PutFegLTENetworkIDParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the put feg LTE network ID params
func (o *PutFegLTENetworkIDParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *PutFegLTENetworkIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.LTENetwork != nil {
		if err := r.SetBodyParam(o.LTENetwork); err != nil {
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
