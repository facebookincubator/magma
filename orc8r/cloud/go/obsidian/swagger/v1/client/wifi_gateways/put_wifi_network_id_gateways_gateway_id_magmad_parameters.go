// Code generated by go-swagger; DO NOT EDIT.

package wifi_gateways

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

// NewPutWifiNetworkIDGatewaysGatewayIDMagmadParams creates a new PutWifiNetworkIDGatewaysGatewayIDMagmadParams object
// with the default values initialized.
func NewPutWifiNetworkIDGatewaysGatewayIDMagmadParams() *PutWifiNetworkIDGatewaysGatewayIDMagmadParams {
	var ()
	return &PutWifiNetworkIDGatewaysGatewayIDMagmadParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutWifiNetworkIDGatewaysGatewayIDMagmadParamsWithTimeout creates a new PutWifiNetworkIDGatewaysGatewayIDMagmadParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutWifiNetworkIDGatewaysGatewayIDMagmadParamsWithTimeout(timeout time.Duration) *PutWifiNetworkIDGatewaysGatewayIDMagmadParams {
	var ()
	return &PutWifiNetworkIDGatewaysGatewayIDMagmadParams{

		timeout: timeout,
	}
}

// NewPutWifiNetworkIDGatewaysGatewayIDMagmadParamsWithContext creates a new PutWifiNetworkIDGatewaysGatewayIDMagmadParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutWifiNetworkIDGatewaysGatewayIDMagmadParamsWithContext(ctx context.Context) *PutWifiNetworkIDGatewaysGatewayIDMagmadParams {
	var ()
	return &PutWifiNetworkIDGatewaysGatewayIDMagmadParams{

		Context: ctx,
	}
}

// NewPutWifiNetworkIDGatewaysGatewayIDMagmadParamsWithHTTPClient creates a new PutWifiNetworkIDGatewaysGatewayIDMagmadParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutWifiNetworkIDGatewaysGatewayIDMagmadParamsWithHTTPClient(client *http.Client) *PutWifiNetworkIDGatewaysGatewayIDMagmadParams {
	var ()
	return &PutWifiNetworkIDGatewaysGatewayIDMagmadParams{
		HTTPClient: client,
	}
}

/*PutWifiNetworkIDGatewaysGatewayIDMagmadParams contains all the parameters to send to the API endpoint
for the put wifi network ID gateways gateway ID magmad operation typically these are written to a http.Request
*/
type PutWifiNetworkIDGatewaysGatewayIDMagmadParams struct {

	/*GatewayID
	  Gateway ID

	*/
	GatewayID string
	/*Magmad
	  New magmad configuration

	*/
	Magmad *models.MagmadGatewayConfigs
	/*NetworkID
	  Network ID

	*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put wifi network ID gateways gateway ID magmad params
func (o *PutWifiNetworkIDGatewaysGatewayIDMagmadParams) WithTimeout(timeout time.Duration) *PutWifiNetworkIDGatewaysGatewayIDMagmadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put wifi network ID gateways gateway ID magmad params
func (o *PutWifiNetworkIDGatewaysGatewayIDMagmadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put wifi network ID gateways gateway ID magmad params
func (o *PutWifiNetworkIDGatewaysGatewayIDMagmadParams) WithContext(ctx context.Context) *PutWifiNetworkIDGatewaysGatewayIDMagmadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put wifi network ID gateways gateway ID magmad params
func (o *PutWifiNetworkIDGatewaysGatewayIDMagmadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put wifi network ID gateways gateway ID magmad params
func (o *PutWifiNetworkIDGatewaysGatewayIDMagmadParams) WithHTTPClient(client *http.Client) *PutWifiNetworkIDGatewaysGatewayIDMagmadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put wifi network ID gateways gateway ID magmad params
func (o *PutWifiNetworkIDGatewaysGatewayIDMagmadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGatewayID adds the gatewayID to the put wifi network ID gateways gateway ID magmad params
func (o *PutWifiNetworkIDGatewaysGatewayIDMagmadParams) WithGatewayID(gatewayID string) *PutWifiNetworkIDGatewaysGatewayIDMagmadParams {
	o.SetGatewayID(gatewayID)
	return o
}

// SetGatewayID adds the gatewayId to the put wifi network ID gateways gateway ID magmad params
func (o *PutWifiNetworkIDGatewaysGatewayIDMagmadParams) SetGatewayID(gatewayID string) {
	o.GatewayID = gatewayID
}

// WithMagmad adds the magmad to the put wifi network ID gateways gateway ID magmad params
func (o *PutWifiNetworkIDGatewaysGatewayIDMagmadParams) WithMagmad(magmad *models.MagmadGatewayConfigs) *PutWifiNetworkIDGatewaysGatewayIDMagmadParams {
	o.SetMagmad(magmad)
	return o
}

// SetMagmad adds the magmad to the put wifi network ID gateways gateway ID magmad params
func (o *PutWifiNetworkIDGatewaysGatewayIDMagmadParams) SetMagmad(magmad *models.MagmadGatewayConfigs) {
	o.Magmad = magmad
}

// WithNetworkID adds the networkID to the put wifi network ID gateways gateway ID magmad params
func (o *PutWifiNetworkIDGatewaysGatewayIDMagmadParams) WithNetworkID(networkID string) *PutWifiNetworkIDGatewaysGatewayIDMagmadParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the put wifi network ID gateways gateway ID magmad params
func (o *PutWifiNetworkIDGatewaysGatewayIDMagmadParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *PutWifiNetworkIDGatewaysGatewayIDMagmadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param gateway_id
	if err := r.SetPathParam("gateway_id", o.GatewayID); err != nil {
		return err
	}

	if o.Magmad != nil {
		if err := r.SetBodyParam(o.Magmad); err != nil {
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
