// Code generated by go-swagger; DO NOT EDIT.

package carrier_wifi_networks

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

// NewPutCwfNetworkIDSubscriberConfigBaseNamesParams creates a new PutCwfNetworkIDSubscriberConfigBaseNamesParams object
// with the default values initialized.
func NewPutCwfNetworkIDSubscriberConfigBaseNamesParams() *PutCwfNetworkIDSubscriberConfigBaseNamesParams {
	var ()
	return &PutCwfNetworkIDSubscriberConfigBaseNamesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutCwfNetworkIDSubscriberConfigBaseNamesParamsWithTimeout creates a new PutCwfNetworkIDSubscriberConfigBaseNamesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutCwfNetworkIDSubscriberConfigBaseNamesParamsWithTimeout(timeout time.Duration) *PutCwfNetworkIDSubscriberConfigBaseNamesParams {
	var ()
	return &PutCwfNetworkIDSubscriberConfigBaseNamesParams{

		timeout: timeout,
	}
}

// NewPutCwfNetworkIDSubscriberConfigBaseNamesParamsWithContext creates a new PutCwfNetworkIDSubscriberConfigBaseNamesParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutCwfNetworkIDSubscriberConfigBaseNamesParamsWithContext(ctx context.Context) *PutCwfNetworkIDSubscriberConfigBaseNamesParams {
	var ()
	return &PutCwfNetworkIDSubscriberConfigBaseNamesParams{

		Context: ctx,
	}
}

// NewPutCwfNetworkIDSubscriberConfigBaseNamesParamsWithHTTPClient creates a new PutCwfNetworkIDSubscriberConfigBaseNamesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutCwfNetworkIDSubscriberConfigBaseNamesParamsWithHTTPClient(client *http.Client) *PutCwfNetworkIDSubscriberConfigBaseNamesParams {
	var ()
	return &PutCwfNetworkIDSubscriberConfigBaseNamesParams{
		HTTPClient: client,
	}
}

/*PutCwfNetworkIDSubscriberConfigBaseNamesParams contains all the parameters to send to the API endpoint
for the put cwf network ID subscriber config base names operation typically these are written to a http.Request
*/
type PutCwfNetworkIDSubscriberConfigBaseNamesParams struct {

	/*NetworkID
	  Network ID

	*/
	NetworkID string
	/*Record
	  Subscriber Config for the Network

	*/
	Record models.BaseNames

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put cwf network ID subscriber config base names params
func (o *PutCwfNetworkIDSubscriberConfigBaseNamesParams) WithTimeout(timeout time.Duration) *PutCwfNetworkIDSubscriberConfigBaseNamesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put cwf network ID subscriber config base names params
func (o *PutCwfNetworkIDSubscriberConfigBaseNamesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put cwf network ID subscriber config base names params
func (o *PutCwfNetworkIDSubscriberConfigBaseNamesParams) WithContext(ctx context.Context) *PutCwfNetworkIDSubscriberConfigBaseNamesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put cwf network ID subscriber config base names params
func (o *PutCwfNetworkIDSubscriberConfigBaseNamesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put cwf network ID subscriber config base names params
func (o *PutCwfNetworkIDSubscriberConfigBaseNamesParams) WithHTTPClient(client *http.Client) *PutCwfNetworkIDSubscriberConfigBaseNamesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put cwf network ID subscriber config base names params
func (o *PutCwfNetworkIDSubscriberConfigBaseNamesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the put cwf network ID subscriber config base names params
func (o *PutCwfNetworkIDSubscriberConfigBaseNamesParams) WithNetworkID(networkID string) *PutCwfNetworkIDSubscriberConfigBaseNamesParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the put cwf network ID subscriber config base names params
func (o *PutCwfNetworkIDSubscriberConfigBaseNamesParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithRecord adds the record to the put cwf network ID subscriber config base names params
func (o *PutCwfNetworkIDSubscriberConfigBaseNamesParams) WithRecord(record models.BaseNames) *PutCwfNetworkIDSubscriberConfigBaseNamesParams {
	o.SetRecord(record)
	return o
}

// SetRecord adds the record to the put cwf network ID subscriber config base names params
func (o *PutCwfNetworkIDSubscriberConfigBaseNamesParams) SetRecord(record models.BaseNames) {
	o.Record = record
}

// WriteToRequest writes these params to a swagger request
func (o *PutCwfNetworkIDSubscriberConfigBaseNamesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param network_id
	if err := r.SetPathParam("network_id", o.NetworkID); err != nil {
		return err
	}

	if o.Record != nil {
		if err := r.SetBodyParam(o.Record); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
