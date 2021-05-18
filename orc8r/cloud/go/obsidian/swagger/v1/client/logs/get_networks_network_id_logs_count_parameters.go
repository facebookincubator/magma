// Code generated by go-swagger; DO NOT EDIT.

package logs

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

// NewGetNetworksNetworkIDLogsCountParams creates a new GetNetworksNetworkIDLogsCountParams object
// with the default values initialized.
func NewGetNetworksNetworkIDLogsCountParams() *GetNetworksNetworkIDLogsCountParams {
	var ()
	return &GetNetworksNetworkIDLogsCountParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetworksNetworkIDLogsCountParamsWithTimeout creates a new GetNetworksNetworkIDLogsCountParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetNetworksNetworkIDLogsCountParamsWithTimeout(timeout time.Duration) *GetNetworksNetworkIDLogsCountParams {
	var ()
	return &GetNetworksNetworkIDLogsCountParams{

		timeout: timeout,
	}
}

// NewGetNetworksNetworkIDLogsCountParamsWithContext creates a new GetNetworksNetworkIDLogsCountParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetNetworksNetworkIDLogsCountParamsWithContext(ctx context.Context) *GetNetworksNetworkIDLogsCountParams {
	var ()
	return &GetNetworksNetworkIDLogsCountParams{

		Context: ctx,
	}
}

// NewGetNetworksNetworkIDLogsCountParamsWithHTTPClient creates a new GetNetworksNetworkIDLogsCountParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetNetworksNetworkIDLogsCountParamsWithHTTPClient(client *http.Client) *GetNetworksNetworkIDLogsCountParams {
	var ()
	return &GetNetworksNetworkIDLogsCountParams{
		HTTPClient: client,
	}
}

/*GetNetworksNetworkIDLogsCountParams contains all the parameters to send to the API endpoint
for the get networks network ID logs count operation typically these are written to a http.Request
*/
type GetNetworksNetworkIDLogsCountParams struct {

	/*End
	  Time to end searching

	*/
	End *string
	/*Fields
	  Comma-separated list of fields to search with the simple query. Defaults to the log field.

	*/
	Fields *string
	/*Filters
	  Comma-separated list of key:value pairs to filter the query with.

	*/
	Filters *string
	/*NetworkID
	  Network ID

	*/
	NetworkID string
	/*SimpleQuery
	  Simple Query String to execute

	*/
	SimpleQuery *string
	/*Start
	  Time to start searching

	*/
	Start *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get networks network ID logs count params
func (o *GetNetworksNetworkIDLogsCountParams) WithTimeout(timeout time.Duration) *GetNetworksNetworkIDLogsCountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get networks network ID logs count params
func (o *GetNetworksNetworkIDLogsCountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get networks network ID logs count params
func (o *GetNetworksNetworkIDLogsCountParams) WithContext(ctx context.Context) *GetNetworksNetworkIDLogsCountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get networks network ID logs count params
func (o *GetNetworksNetworkIDLogsCountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get networks network ID logs count params
func (o *GetNetworksNetworkIDLogsCountParams) WithHTTPClient(client *http.Client) *GetNetworksNetworkIDLogsCountParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get networks network ID logs count params
func (o *GetNetworksNetworkIDLogsCountParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEnd adds the end to the get networks network ID logs count params
func (o *GetNetworksNetworkIDLogsCountParams) WithEnd(end *string) *GetNetworksNetworkIDLogsCountParams {
	o.SetEnd(end)
	return o
}

// SetEnd adds the end to the get networks network ID logs count params
func (o *GetNetworksNetworkIDLogsCountParams) SetEnd(end *string) {
	o.End = end
}

// WithFields adds the fields to the get networks network ID logs count params
func (o *GetNetworksNetworkIDLogsCountParams) WithFields(fields *string) *GetNetworksNetworkIDLogsCountParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get networks network ID logs count params
func (o *GetNetworksNetworkIDLogsCountParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithFilters adds the filters to the get networks network ID logs count params
func (o *GetNetworksNetworkIDLogsCountParams) WithFilters(filters *string) *GetNetworksNetworkIDLogsCountParams {
	o.SetFilters(filters)
	return o
}

// SetFilters adds the filters to the get networks network ID logs count params
func (o *GetNetworksNetworkIDLogsCountParams) SetFilters(filters *string) {
	o.Filters = filters
}

// WithNetworkID adds the networkID to the get networks network ID logs count params
func (o *GetNetworksNetworkIDLogsCountParams) WithNetworkID(networkID string) *GetNetworksNetworkIDLogsCountParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get networks network ID logs count params
func (o *GetNetworksNetworkIDLogsCountParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithSimpleQuery adds the simpleQuery to the get networks network ID logs count params
func (o *GetNetworksNetworkIDLogsCountParams) WithSimpleQuery(simpleQuery *string) *GetNetworksNetworkIDLogsCountParams {
	o.SetSimpleQuery(simpleQuery)
	return o
}

// SetSimpleQuery adds the simpleQuery to the get networks network ID logs count params
func (o *GetNetworksNetworkIDLogsCountParams) SetSimpleQuery(simpleQuery *string) {
	o.SimpleQuery = simpleQuery
}

// WithStart adds the start to the get networks network ID logs count params
func (o *GetNetworksNetworkIDLogsCountParams) WithStart(start *string) *GetNetworksNetworkIDLogsCountParams {
	o.SetStart(start)
	return o
}

// SetStart adds the start to the get networks network ID logs count params
func (o *GetNetworksNetworkIDLogsCountParams) SetStart(start *string) {
	o.Start = start
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetworksNetworkIDLogsCountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Fields != nil {

		// query param fields
		var qrFields string
		if o.Fields != nil {
			qrFields = *o.Fields
		}
		qFields := qrFields
		if qFields != "" {
			if err := r.SetQueryParam("fields", qFields); err != nil {
				return err
			}
		}

	}

	if o.Filters != nil {

		// query param filters
		var qrFilters string
		if o.Filters != nil {
			qrFilters = *o.Filters
		}
		qFilters := qrFilters
		if qFilters != "" {
			if err := r.SetQueryParam("filters", qFilters); err != nil {
				return err
			}
		}

	}

	// path param network_id
	if err := r.SetPathParam("network_id", o.NetworkID); err != nil {
		return err
	}

	if o.SimpleQuery != nil {

		// query param simple_query
		var qrSimpleQuery string
		if o.SimpleQuery != nil {
			qrSimpleQuery = *o.SimpleQuery
		}
		qSimpleQuery := qrSimpleQuery
		if qSimpleQuery != "" {
			if err := r.SetQueryParam("simple_query", qSimpleQuery); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
