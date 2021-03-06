// Code generated by go-swagger; DO NOT EDIT.

package lte_gateways

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

// NewGetLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsParams creates a new GetLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsParams object
// with the default values initialized.
func NewGetLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsParams() *GetLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsParams {
	var ()
	return &GetLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsParamsWithTimeout creates a new GetLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsParamsWithTimeout(timeout time.Duration) *GetLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsParams {
	var ()
	return &GetLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsParams{

		timeout: timeout,
	}
}

// NewGetLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsParamsWithContext creates a new GetLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsParamsWithContext(ctx context.Context) *GetLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsParams {
	var ()
	return &GetLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsParams{

		Context: ctx,
	}
}

// NewGetLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsParamsWithHTTPClient creates a new GetLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsParamsWithHTTPClient(client *http.Client) *GetLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsParams {
	var ()
	return &GetLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsParams{
		HTTPClient: client,
	}
}

/*GetLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsParams contains all the parameters to send to the API endpoint
for the get LTE network ID gateways gateway ID connected ENODEB serials operation typically these are written to a http.Request
*/
type GetLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsParams struct {

	/*GatewayID
	  Gateway ID

	*/
	GatewayID string
	/*NetworkID
	  Network ID

	*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get LTE network ID gateways gateway ID connected ENODEB serials params
func (o *GetLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsParams) WithTimeout(timeout time.Duration) *GetLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get LTE network ID gateways gateway ID connected ENODEB serials params
func (o *GetLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get LTE network ID gateways gateway ID connected ENODEB serials params
func (o *GetLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsParams) WithContext(ctx context.Context) *GetLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get LTE network ID gateways gateway ID connected ENODEB serials params
func (o *GetLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get LTE network ID gateways gateway ID connected ENODEB serials params
func (o *GetLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsParams) WithHTTPClient(client *http.Client) *GetLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get LTE network ID gateways gateway ID connected ENODEB serials params
func (o *GetLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGatewayID adds the gatewayID to the get LTE network ID gateways gateway ID connected ENODEB serials params
func (o *GetLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsParams) WithGatewayID(gatewayID string) *GetLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsParams {
	o.SetGatewayID(gatewayID)
	return o
}

// SetGatewayID adds the gatewayId to the get LTE network ID gateways gateway ID connected ENODEB serials params
func (o *GetLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsParams) SetGatewayID(gatewayID string) {
	o.GatewayID = gatewayID
}

// WithNetworkID adds the networkID to the get LTE network ID gateways gateway ID connected ENODEB serials params
func (o *GetLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsParams) WithNetworkID(networkID string) *GetLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get LTE network ID gateways gateway ID connected ENODEB serials params
func (o *GetLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *GetLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param gateway_id
	if err := r.SetPathParam("gateway_id", o.GatewayID); err != nil {
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
