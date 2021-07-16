// Code generated by go-swagger; DO NOT EDIT.

package subscribers

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

// NewPostLTENetworkIDSubscribersV2Params creates a new PostLTENetworkIDSubscribersV2Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostLTENetworkIDSubscribersV2Params() *PostLTENetworkIDSubscribersV2Params {
	return &PostLTENetworkIDSubscribersV2Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostLTENetworkIDSubscribersV2ParamsWithTimeout creates a new PostLTENetworkIDSubscribersV2Params object
// with the ability to set a timeout on a request.
func NewPostLTENetworkIDSubscribersV2ParamsWithTimeout(timeout time.Duration) *PostLTENetworkIDSubscribersV2Params {
	return &PostLTENetworkIDSubscribersV2Params{
		timeout: timeout,
	}
}

// NewPostLTENetworkIDSubscribersV2ParamsWithContext creates a new PostLTENetworkIDSubscribersV2Params object
// with the ability to set a context for a request.
func NewPostLTENetworkIDSubscribersV2ParamsWithContext(ctx context.Context) *PostLTENetworkIDSubscribersV2Params {
	return &PostLTENetworkIDSubscribersV2Params{
		Context: ctx,
	}
}

// NewPostLTENetworkIDSubscribersV2ParamsWithHTTPClient creates a new PostLTENetworkIDSubscribersV2Params object
// with the ability to set a custom HTTPClient for a request.
func NewPostLTENetworkIDSubscribersV2ParamsWithHTTPClient(client *http.Client) *PostLTENetworkIDSubscribersV2Params {
	return &PostLTENetworkIDSubscribersV2Params{
		HTTPClient: client,
	}
}

/* PostLTENetworkIDSubscribersV2Params contains all the parameters to send to the API endpoint
   for the post LTE network ID subscribers v2 operation.

   Typically these are written to a http.Request.
*/
type PostLTENetworkIDSubscribersV2Params struct {

	/* NetworkID.

	   Network ID
	*/
	NetworkID string

	/* Subscribers.

	   Subscribers to add
	*/
	Subscribers models.MutableSubscribers

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post LTE network ID subscribers v2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostLTENetworkIDSubscribersV2Params) WithDefaults() *PostLTENetworkIDSubscribersV2Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post LTE network ID subscribers v2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostLTENetworkIDSubscribersV2Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post LTE network ID subscribers v2 params
func (o *PostLTENetworkIDSubscribersV2Params) WithTimeout(timeout time.Duration) *PostLTENetworkIDSubscribersV2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post LTE network ID subscribers v2 params
func (o *PostLTENetworkIDSubscribersV2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post LTE network ID subscribers v2 params
func (o *PostLTENetworkIDSubscribersV2Params) WithContext(ctx context.Context) *PostLTENetworkIDSubscribersV2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post LTE network ID subscribers v2 params
func (o *PostLTENetworkIDSubscribersV2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post LTE network ID subscribers v2 params
func (o *PostLTENetworkIDSubscribersV2Params) WithHTTPClient(client *http.Client) *PostLTENetworkIDSubscribersV2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post LTE network ID subscribers v2 params
func (o *PostLTENetworkIDSubscribersV2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the post LTE network ID subscribers v2 params
func (o *PostLTENetworkIDSubscribersV2Params) WithNetworkID(networkID string) *PostLTENetworkIDSubscribersV2Params {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the post LTE network ID subscribers v2 params
func (o *PostLTENetworkIDSubscribersV2Params) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithSubscribers adds the subscribers to the post LTE network ID subscribers v2 params
func (o *PostLTENetworkIDSubscribersV2Params) WithSubscribers(subscribers models.MutableSubscribers) *PostLTENetworkIDSubscribersV2Params {
	o.SetSubscribers(subscribers)
	return o
}

// SetSubscribers adds the subscribers to the post LTE network ID subscribers v2 params
func (o *PostLTENetworkIDSubscribersV2Params) SetSubscribers(subscribers models.MutableSubscribers) {
	o.Subscribers = subscribers
}

// WriteToRequest writes these params to a swagger request
func (o *PostLTENetworkIDSubscribersV2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param network_id
	if err := r.SetPathParam("network_id", o.NetworkID); err != nil {
		return err
	}
	if o.Subscribers != nil {
		if err := r.SetBodyParam(o.Subscribers); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
