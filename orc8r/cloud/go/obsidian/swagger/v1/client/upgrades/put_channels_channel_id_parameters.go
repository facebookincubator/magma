// Code generated by go-swagger; DO NOT EDIT.

package upgrades

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

// NewPutChannelsChannelIDParams creates a new PutChannelsChannelIDParams object
// with the default values initialized.
func NewPutChannelsChannelIDParams() *PutChannelsChannelIDParams {
	var ()
	return &PutChannelsChannelIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutChannelsChannelIDParamsWithTimeout creates a new PutChannelsChannelIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutChannelsChannelIDParamsWithTimeout(timeout time.Duration) *PutChannelsChannelIDParams {
	var ()
	return &PutChannelsChannelIDParams{

		timeout: timeout,
	}
}

// NewPutChannelsChannelIDParamsWithContext creates a new PutChannelsChannelIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutChannelsChannelIDParamsWithContext(ctx context.Context) *PutChannelsChannelIDParams {
	var ()
	return &PutChannelsChannelIDParams{

		Context: ctx,
	}
}

// NewPutChannelsChannelIDParamsWithHTTPClient creates a new PutChannelsChannelIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutChannelsChannelIDParamsWithHTTPClient(client *http.Client) *PutChannelsChannelIDParams {
	var ()
	return &PutChannelsChannelIDParams{
		HTTPClient: client,
	}
}

/*PutChannelsChannelIDParams contains all the parameters to send to the API endpoint
for the put channels channel ID operation typically these are written to a http.Request
*/
type PutChannelsChannelIDParams struct {

	/*ChannelID
	  Release Channel ID

	*/
	ChannelID string
	/*ReleaseChannel
	  New channel configuration

	*/
	ReleaseChannel models.ReleaseChannel

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put channels channel ID params
func (o *PutChannelsChannelIDParams) WithTimeout(timeout time.Duration) *PutChannelsChannelIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put channels channel ID params
func (o *PutChannelsChannelIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put channels channel ID params
func (o *PutChannelsChannelIDParams) WithContext(ctx context.Context) *PutChannelsChannelIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put channels channel ID params
func (o *PutChannelsChannelIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put channels channel ID params
func (o *PutChannelsChannelIDParams) WithHTTPClient(client *http.Client) *PutChannelsChannelIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put channels channel ID params
func (o *PutChannelsChannelIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithChannelID adds the channelID to the put channels channel ID params
func (o *PutChannelsChannelIDParams) WithChannelID(channelID string) *PutChannelsChannelIDParams {
	o.SetChannelID(channelID)
	return o
}

// SetChannelID adds the channelId to the put channels channel ID params
func (o *PutChannelsChannelIDParams) SetChannelID(channelID string) {
	o.ChannelID = channelID
}

// WithReleaseChannel adds the releaseChannel to the put channels channel ID params
func (o *PutChannelsChannelIDParams) WithReleaseChannel(releaseChannel models.ReleaseChannel) *PutChannelsChannelIDParams {
	o.SetReleaseChannel(releaseChannel)
	return o
}

// SetReleaseChannel adds the releaseChannel to the put channels channel ID params
func (o *PutChannelsChannelIDParams) SetReleaseChannel(releaseChannel models.ReleaseChannel) {
	o.ReleaseChannel = releaseChannel
}

// WriteToRequest writes these params to a swagger request
func (o *PutChannelsChannelIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param channel_id
	if err := r.SetPathParam("channel_id", o.ChannelID); err != nil {
		return err
	}

	if err := r.SetBodyParam(o.ReleaseChannel); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
