// Code generated by go-swagger; DO NOT EDIT.

package tenants

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"magma/orc8r/cloud/go/obsidian/swagger/v1/models"
)

// GetTenantsTenantIDReader is a Reader for the GetTenantsTenantID structure.
type GetTenantsTenantIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTenantsTenantIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTenantsTenantIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetTenantsTenantIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetTenantsTenantIDOK creates a GetTenantsTenantIDOK with default headers values
func NewGetTenantsTenantIDOK() *GetTenantsTenantIDOK {
	return &GetTenantsTenantIDOK{}
}

/* GetTenantsTenantIDOK describes a response with status code 200, with default header values.

Requested Tenant Information
*/
type GetTenantsTenantIDOK struct {
	Payload *models.Tenant
}

func (o *GetTenantsTenantIDOK) Error() string {
	return fmt.Sprintf("[GET /tenants/{tenant_id}][%d] getTenantsTenantIdOK  %+v", 200, o.Payload)
}
func (o *GetTenantsTenantIDOK) GetPayload() *models.Tenant {
	return o.Payload
}

func (o *GetTenantsTenantIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Tenant)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTenantsTenantIDDefault creates a GetTenantsTenantIDDefault with default headers values
func NewGetTenantsTenantIDDefault(code int) *GetTenantsTenantIDDefault {
	return &GetTenantsTenantIDDefault{
		_statusCode: code,
	}
}

/* GetTenantsTenantIDDefault describes a response with status code -1, with default header values.

Unexpected Error
*/
type GetTenantsTenantIDDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get tenants tenant ID default response
func (o *GetTenantsTenantIDDefault) Code() int {
	return o._statusCode
}

func (o *GetTenantsTenantIDDefault) Error() string {
	return fmt.Sprintf("[GET /tenants/{tenant_id}][%d] GetTenantsTenantID default  %+v", o._statusCode, o.Payload)
}
func (o *GetTenantsTenantIDDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetTenantsTenantIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
