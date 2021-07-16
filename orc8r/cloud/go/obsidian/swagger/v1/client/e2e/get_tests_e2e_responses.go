// Code generated by go-swagger; DO NOT EDIT.

package e2e

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"magma/orc8r/cloud/go/obsidian/swagger/v1/models"
)

// GetTestsE2eReader is a Reader for the GetTestsE2e structure.
type GetTestsE2eReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTestsE2eReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTestsE2eOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetTestsE2eDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetTestsE2eOK creates a GetTestsE2eOK with default headers values
func NewGetTestsE2eOK() *GetTestsE2eOK {
	return &GetTestsE2eOK{}
}

/* GetTestsE2eOK describes a response with status code 200, with default header values.

List of test configurations
*/
type GetTestsE2eOK struct {
	Payload []*models.E2eTestCase
}

func (o *GetTestsE2eOK) Error() string {
	return fmt.Sprintf("[GET /tests/e2e][%d] getTestsE2eOK  %+v", 200, o.Payload)
}
func (o *GetTestsE2eOK) GetPayload() []*models.E2eTestCase {
	return o.Payload
}

func (o *GetTestsE2eOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTestsE2eDefault creates a GetTestsE2eDefault with default headers values
func NewGetTestsE2eDefault(code int) *GetTestsE2eDefault {
	return &GetTestsE2eDefault{
		_statusCode: code,
	}
}

/* GetTestsE2eDefault describes a response with status code -1, with default header values.

Unexpected Error
*/
type GetTestsE2eDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get tests e2e default response
func (o *GetTestsE2eDefault) Code() int {
	return o._statusCode
}

func (o *GetTestsE2eDefault) Error() string {
	return fmt.Sprintf("[GET /tests/e2e][%d] GetTestsE2e default  %+v", o._statusCode, o.Payload)
}
func (o *GetTestsE2eDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetTestsE2eDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
