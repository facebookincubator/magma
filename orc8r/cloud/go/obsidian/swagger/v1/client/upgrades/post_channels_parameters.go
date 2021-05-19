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

// NewPostChannelsParams creates a new PostChannelsParams object
// with the default values initialized.
func NewPostChannelsParams() *PostChannelsParams {
	var ()
	return &PostChannelsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostChannelsParamsWithTimeout creates a new PostChannelsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostChannelsParamsWithTimeout(timeout time.Duration) *PostChannelsParams {
	var ()
	return &PostChannelsParams{

		timeout: timeout,
	}
}

// NewPostChannelsParamsWithContext creates a new PostChannelsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostChannelsParamsWithContext(ctx context.Context) *PostChannelsParams {
	var ()
	return &PostChannelsParams{

		Context: ctx,
	}
}

// NewPostChannelsParamsWithHTTPClient creates a new PostChannelsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostChannelsParamsWithHTTPClient(client *http.Client) *PostChannelsParams {
	var ()
	return &PostChannelsParams{
		HTTPClient: client,
	}
}

/*PostChannelsParams contains all the parameters to send to the API endpoint
for the post channels operation typically these are written to a http.Request
*/
type PostChannelsParams struct {

	/*Channel
	  The release channel to create

	*/
	Channel models.ReleaseChannel

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post channels params
func (o *PostChannelsParams) WithTimeout(timeout time.Duration) *PostChannelsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post channels params
func (o *PostChannelsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post channels params
func (o *PostChannelsParams) WithContext(ctx context.Context) *PostChannelsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post channels params
func (o *PostChannelsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post channels params
func (o *PostChannelsParams) WithHTTPClient(client *http.Client) *PostChannelsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post channels params
func (o *PostChannelsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithChannel adds the channel to the post channels params
func (o *PostChannelsParams) WithChannel(channel models.ReleaseChannel) *PostChannelsParams {
	o.SetChannel(channel)
	return o
}

// SetChannel adds the channel to the post channels params
func (o *PostChannelsParams) SetChannel(channel models.ReleaseChannel) {
	o.Channel = channel
}

// WriteToRequest writes these params to a swagger request
func (o *PostChannelsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.Channel); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
