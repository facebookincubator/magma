// Code generated by go-swagger; DO NOT EDIT.

package events

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

// NewGetEventsNetworkIDParams creates a new GetEventsNetworkIDParams object
// with the default values initialized.
func NewGetEventsNetworkIDParams() *GetEventsNetworkIDParams {
	var ()
	return &GetEventsNetworkIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetEventsNetworkIDParamsWithTimeout creates a new GetEventsNetworkIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetEventsNetworkIDParamsWithTimeout(timeout time.Duration) *GetEventsNetworkIDParams {
	var ()
	return &GetEventsNetworkIDParams{

		timeout: timeout,
	}
}

// NewGetEventsNetworkIDParamsWithContext creates a new GetEventsNetworkIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetEventsNetworkIDParamsWithContext(ctx context.Context) *GetEventsNetworkIDParams {
	var ()
	return &GetEventsNetworkIDParams{

		Context: ctx,
	}
}

// NewGetEventsNetworkIDParamsWithHTTPClient creates a new GetEventsNetworkIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetEventsNetworkIDParamsWithHTTPClient(client *http.Client) *GetEventsNetworkIDParams {
	var ()
	return &GetEventsNetworkIDParams{
		HTTPClient: client,
	}
}

/*GetEventsNetworkIDParams contains all the parameters to send to the API endpoint
for the get events network ID operation typically these are written to a http.Request
*/
type GetEventsNetworkIDParams struct {

	/*End
	  End time for the query, in RFC3339 or ISO8601 format

	*/
	End *string
	/*Events
	  Comma-separated list of event types to query

	*/
	Events *string
	/*From
	  Index to start the query from

	*/
	From *string
	/*HwIds
	  Comma-separated list of hardware IDs to query

	*/
	HwIds *string
	/*NetworkID
	  Network ID

	*/
	NetworkID string
	/*Size
	  Maximum size to limit the query result to. This defaults to 50

	*/
	Size *string
	/*Start
	  Start time for the query, in RFC3339 or ISO8601 format

	*/
	Start *string
	/*Streams
	  Comma-separated list of streams to query

	*/
	Streams *string
	/*Tags
	  Comma-separated list of tags to query

	*/
	Tags *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get events network ID params
func (o *GetEventsNetworkIDParams) WithTimeout(timeout time.Duration) *GetEventsNetworkIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get events network ID params
func (o *GetEventsNetworkIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get events network ID params
func (o *GetEventsNetworkIDParams) WithContext(ctx context.Context) *GetEventsNetworkIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get events network ID params
func (o *GetEventsNetworkIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get events network ID params
func (o *GetEventsNetworkIDParams) WithHTTPClient(client *http.Client) *GetEventsNetworkIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get events network ID params
func (o *GetEventsNetworkIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEnd adds the end to the get events network ID params
func (o *GetEventsNetworkIDParams) WithEnd(end *string) *GetEventsNetworkIDParams {
	o.SetEnd(end)
	return o
}

// SetEnd adds the end to the get events network ID params
func (o *GetEventsNetworkIDParams) SetEnd(end *string) {
	o.End = end
}

// WithEvents adds the events to the get events network ID params
func (o *GetEventsNetworkIDParams) WithEvents(events *string) *GetEventsNetworkIDParams {
	o.SetEvents(events)
	return o
}

// SetEvents adds the events to the get events network ID params
func (o *GetEventsNetworkIDParams) SetEvents(events *string) {
	o.Events = events
}

// WithFrom adds the from to the get events network ID params
func (o *GetEventsNetworkIDParams) WithFrom(from *string) *GetEventsNetworkIDParams {
	o.SetFrom(from)
	return o
}

// SetFrom adds the from to the get events network ID params
func (o *GetEventsNetworkIDParams) SetFrom(from *string) {
	o.From = from
}

// WithHwIds adds the hwIds to the get events network ID params
func (o *GetEventsNetworkIDParams) WithHwIds(hwIds *string) *GetEventsNetworkIDParams {
	o.SetHwIds(hwIds)
	return o
}

// SetHwIds adds the hwIds to the get events network ID params
func (o *GetEventsNetworkIDParams) SetHwIds(hwIds *string) {
	o.HwIds = hwIds
}

// WithNetworkID adds the networkID to the get events network ID params
func (o *GetEventsNetworkIDParams) WithNetworkID(networkID string) *GetEventsNetworkIDParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get events network ID params
func (o *GetEventsNetworkIDParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithSize adds the size to the get events network ID params
func (o *GetEventsNetworkIDParams) WithSize(size *string) *GetEventsNetworkIDParams {
	o.SetSize(size)
	return o
}

// SetSize adds the size to the get events network ID params
func (o *GetEventsNetworkIDParams) SetSize(size *string) {
	o.Size = size
}

// WithStart adds the start to the get events network ID params
func (o *GetEventsNetworkIDParams) WithStart(start *string) *GetEventsNetworkIDParams {
	o.SetStart(start)
	return o
}

// SetStart adds the start to the get events network ID params
func (o *GetEventsNetworkIDParams) SetStart(start *string) {
	o.Start = start
}

// WithStreams adds the streams to the get events network ID params
func (o *GetEventsNetworkIDParams) WithStreams(streams *string) *GetEventsNetworkIDParams {
	o.SetStreams(streams)
	return o
}

// SetStreams adds the streams to the get events network ID params
func (o *GetEventsNetworkIDParams) SetStreams(streams *string) {
	o.Streams = streams
}

// WithTags adds the tags to the get events network ID params
func (o *GetEventsNetworkIDParams) WithTags(tags *string) *GetEventsNetworkIDParams {
	o.SetTags(tags)
	return o
}

// SetTags adds the tags to the get events network ID params
func (o *GetEventsNetworkIDParams) SetTags(tags *string) {
	o.Tags = tags
}

// WriteToRequest writes these params to a swagger request
func (o *GetEventsNetworkIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.End != nil {

		// query param end
		var qrEnd string
		if o.End != nil {
			qrEnd = *o.End
		}
		qEnd := qrEnd
		if qEnd != "" {
			if err := r.SetQueryParam("end", qEnd); err != nil {
				return err
			}
		}

	}

	if o.Events != nil {

		// query param events
		var qrEvents string
		if o.Events != nil {
			qrEvents = *o.Events
		}
		qEvents := qrEvents
		if qEvents != "" {
			if err := r.SetQueryParam("events", qEvents); err != nil {
				return err
			}
		}

	}

	if o.From != nil {

		// query param from
		var qrFrom string
		if o.From != nil {
			qrFrom = *o.From
		}
		qFrom := qrFrom
		if qFrom != "" {
			if err := r.SetQueryParam("from", qFrom); err != nil {
				return err
			}
		}

	}

	if o.HwIds != nil {

		// query param hw_ids
		var qrHwIds string
		if o.HwIds != nil {
			qrHwIds = *o.HwIds
		}
		qHwIds := qrHwIds
		if qHwIds != "" {
			if err := r.SetQueryParam("hw_ids", qHwIds); err != nil {
				return err
			}
		}

	}

	// path param network_id
	if err := r.SetPathParam("network_id", o.NetworkID); err != nil {
		return err
	}

	if o.Size != nil {

		// query param size
		var qrSize string
		if o.Size != nil {
			qrSize = *o.Size
		}
		qSize := qrSize
		if qSize != "" {
			if err := r.SetQueryParam("size", qSize); err != nil {
				return err
			}
		}

	}

	if o.Start != nil {

		// query param start
		var qrStart string
		if o.Start != nil {
			qrStart = *o.Start
		}
		qStart := qrStart
		if qStart != "" {
			if err := r.SetQueryParam("start", qStart); err != nil {
				return err
			}
		}

	}

	if o.Streams != nil {

		// query param streams
		var qrStreams string
		if o.Streams != nil {
			qrStreams = *o.Streams
		}
		qStreams := qrStreams
		if qStreams != "" {
			if err := r.SetQueryParam("streams", qStreams); err != nil {
				return err
			}
		}

	}

	if o.Tags != nil {

		// query param tags
		var qrTags string
		if o.Tags != nil {
			qrTags = *o.Tags
		}
		qTags := qrTags
		if qTags != "" {
			if err := r.SetQueryParam("tags", qTags); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
