// Code generated by go-swagger; DO NOT EDIT.

package metrics

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"magma/orc8r/cloud/go/obsidian/swagger/v1/models"
)

// GetTenantsTenantIDMetricsQueryReader is a Reader for the GetTenantsTenantIDMetricsQuery structure.
type GetTenantsTenantIDMetricsQueryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTenantsTenantIDMetricsQueryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTenantsTenantIDMetricsQueryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetTenantsTenantIDMetricsQueryDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetTenantsTenantIDMetricsQueryOK creates a GetTenantsTenantIDMetricsQueryOK with default headers values
func NewGetTenantsTenantIDMetricsQueryOK() *GetTenantsTenantIDMetricsQueryOK {
	return &GetTenantsTenantIDMetricsQueryOK{}
}

/* GetTenantsTenantIDMetricsQueryOK describes a response with status code 200, with default header values.

List of PromQL metrics results
*/
type GetTenantsTenantIDMetricsQueryOK struct {
	Payload *models.PromqlReturnObject
}

func (o *GetTenantsTenantIDMetricsQueryOK) Error() string {
	return fmt.Sprintf("[GET /tenants/{tenant_id}/metrics/query][%d] getTenantsTenantIdMetricsQueryOK  %+v", 200, o.Payload)
}
func (o *GetTenantsTenantIDMetricsQueryOK) GetPayload() *models.PromqlReturnObject {
	return o.Payload
}

func (o *GetTenantsTenantIDMetricsQueryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PromqlReturnObject)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTenantsTenantIDMetricsQueryDefault creates a GetTenantsTenantIDMetricsQueryDefault with default headers values
func NewGetTenantsTenantIDMetricsQueryDefault(code int) *GetTenantsTenantIDMetricsQueryDefault {
	return &GetTenantsTenantIDMetricsQueryDefault{
		_statusCode: code,
	}
}

/* GetTenantsTenantIDMetricsQueryDefault describes a response with status code -1, with default header values.

Unexpected Error
*/
type GetTenantsTenantIDMetricsQueryDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get tenants tenant ID metrics query default response
func (o *GetTenantsTenantIDMetricsQueryDefault) Code() int {
	return o._statusCode
}

func (o *GetTenantsTenantIDMetricsQueryDefault) Error() string {
	return fmt.Sprintf("[GET /tenants/{tenant_id}/metrics/query][%d] GetTenantsTenantIDMetricsQuery default  %+v", o._statusCode, o.Payload)
}
func (o *GetTenantsTenantIDMetricsQueryDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetTenantsTenantIDMetricsQueryDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
