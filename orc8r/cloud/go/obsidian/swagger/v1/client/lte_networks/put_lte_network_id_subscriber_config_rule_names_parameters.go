// Code generated by go-swagger; DO NOT EDIT.

package lte_networks

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

// NewPutLTENetworkIDSubscriberConfigRuleNamesParams creates a new PutLTENetworkIDSubscriberConfigRuleNamesParams object
// with the default values initialized.
func NewPutLTENetworkIDSubscriberConfigRuleNamesParams() *PutLTENetworkIDSubscriberConfigRuleNamesParams {
	var ()
	return &PutLTENetworkIDSubscriberConfigRuleNamesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutLTENetworkIDSubscriberConfigRuleNamesParamsWithTimeout creates a new PutLTENetworkIDSubscriberConfigRuleNamesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutLTENetworkIDSubscriberConfigRuleNamesParamsWithTimeout(timeout time.Duration) *PutLTENetworkIDSubscriberConfigRuleNamesParams {
	var ()
	return &PutLTENetworkIDSubscriberConfigRuleNamesParams{

		timeout: timeout,
	}
}

// NewPutLTENetworkIDSubscriberConfigRuleNamesParamsWithContext creates a new PutLTENetworkIDSubscriberConfigRuleNamesParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutLTENetworkIDSubscriberConfigRuleNamesParamsWithContext(ctx context.Context) *PutLTENetworkIDSubscriberConfigRuleNamesParams {
	var ()
	return &PutLTENetworkIDSubscriberConfigRuleNamesParams{

		Context: ctx,
	}
}

// NewPutLTENetworkIDSubscriberConfigRuleNamesParamsWithHTTPClient creates a new PutLTENetworkIDSubscriberConfigRuleNamesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutLTENetworkIDSubscriberConfigRuleNamesParamsWithHTTPClient(client *http.Client) *PutLTENetworkIDSubscriberConfigRuleNamesParams {
	var ()
	return &PutLTENetworkIDSubscriberConfigRuleNamesParams{
		HTTPClient: client,
	}
}

/*PutLTENetworkIDSubscriberConfigRuleNamesParams contains all the parameters to send to the API endpoint
for the put LTE network ID subscriber config rule names operation typically these are written to a http.Request
*/
type PutLTENetworkIDSubscriberConfigRuleNamesParams struct {

	/*NetworkID
	  Network ID

	*/
	NetworkID string
	/*Record
	  Subscriber Config for the Network

	*/
	Record models.RuleNames

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put LTE network ID subscriber config rule names params
func (o *PutLTENetworkIDSubscriberConfigRuleNamesParams) WithTimeout(timeout time.Duration) *PutLTENetworkIDSubscriberConfigRuleNamesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put LTE network ID subscriber config rule names params
func (o *PutLTENetworkIDSubscriberConfigRuleNamesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put LTE network ID subscriber config rule names params
func (o *PutLTENetworkIDSubscriberConfigRuleNamesParams) WithContext(ctx context.Context) *PutLTENetworkIDSubscriberConfigRuleNamesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put LTE network ID subscriber config rule names params
func (o *PutLTENetworkIDSubscriberConfigRuleNamesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put LTE network ID subscriber config rule names params
func (o *PutLTENetworkIDSubscriberConfigRuleNamesParams) WithHTTPClient(client *http.Client) *PutLTENetworkIDSubscriberConfigRuleNamesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put LTE network ID subscriber config rule names params
func (o *PutLTENetworkIDSubscriberConfigRuleNamesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the put LTE network ID subscriber config rule names params
func (o *PutLTENetworkIDSubscriberConfigRuleNamesParams) WithNetworkID(networkID string) *PutLTENetworkIDSubscriberConfigRuleNamesParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the put LTE network ID subscriber config rule names params
func (o *PutLTENetworkIDSubscriberConfigRuleNamesParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithRecord adds the record to the put LTE network ID subscriber config rule names params
func (o *PutLTENetworkIDSubscriberConfigRuleNamesParams) WithRecord(record models.RuleNames) *PutLTENetworkIDSubscriberConfigRuleNamesParams {
	o.SetRecord(record)
	return o
}

// SetRecord adds the record to the put LTE network ID subscriber config rule names params
func (o *PutLTENetworkIDSubscriberConfigRuleNamesParams) SetRecord(record models.RuleNames) {
	o.Record = record
}

// WriteToRequest writes these params to a swagger request
func (o *PutLTENetworkIDSubscriberConfigRuleNamesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
