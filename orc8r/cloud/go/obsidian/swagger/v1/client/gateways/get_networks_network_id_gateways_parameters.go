// Code generated by go-swagger; DO NOT EDIT.

package gateways

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

// NewGetNetworksNetworkIDGatewaysParams creates a new GetNetworksNetworkIDGatewaysParams object
// with the default values initialized.
func NewGetNetworksNetworkIDGatewaysParams() *GetNetworksNetworkIDGatewaysParams {
	var ()
	return &GetNetworksNetworkIDGatewaysParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetworksNetworkIDGatewaysParamsWithTimeout creates a new GetNetworksNetworkIDGatewaysParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetNetworksNetworkIDGatewaysParamsWithTimeout(timeout time.Duration) *GetNetworksNetworkIDGatewaysParams {
	var ()
	return &GetNetworksNetworkIDGatewaysParams{

		timeout: timeout,
	}
}

// NewGetNetworksNetworkIDGatewaysParamsWithContext creates a new GetNetworksNetworkIDGatewaysParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetNetworksNetworkIDGatewaysParamsWithContext(ctx context.Context) *GetNetworksNetworkIDGatewaysParams {
	var ()
	return &GetNetworksNetworkIDGatewaysParams{

		Context: ctx,
	}
}

// NewGetNetworksNetworkIDGatewaysParamsWithHTTPClient creates a new GetNetworksNetworkIDGatewaysParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetNetworksNetworkIDGatewaysParamsWithHTTPClient(client *http.Client) *GetNetworksNetworkIDGatewaysParams {
	var ()
	return &GetNetworksNetworkIDGatewaysParams{
		HTTPClient: client,
	}
}

/*GetNetworksNetworkIDGatewaysParams contains all the parameters to send to the API endpoint
for the get networks network ID gateways operation typically these are written to a http.Request
*/
type GetNetworksNetworkIDGatewaysParams struct {

	/*NetworkID
	  Network ID

	*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get networks network ID gateways params
func (o *GetNetworksNetworkIDGatewaysParams) WithTimeout(timeout time.Duration) *GetNetworksNetworkIDGatewaysParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get networks network ID gateways params
func (o *GetNetworksNetworkIDGatewaysParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get networks network ID gateways params
func (o *GetNetworksNetworkIDGatewaysParams) WithContext(ctx context.Context) *GetNetworksNetworkIDGatewaysParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get networks network ID gateways params
func (o *GetNetworksNetworkIDGatewaysParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get networks network ID gateways params
func (o *GetNetworksNetworkIDGatewaysParams) WithHTTPClient(client *http.Client) *GetNetworksNetworkIDGatewaysParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get networks network ID gateways params
func (o *GetNetworksNetworkIDGatewaysParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the get networks network ID gateways params
func (o *GetNetworksNetworkIDGatewaysParams) WithNetworkID(networkID string) *GetNetworksNetworkIDGatewaysParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get networks network ID gateways params
func (o *GetNetworksNetworkIDGatewaysParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetworksNetworkIDGatewaysParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
