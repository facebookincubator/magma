// Code generated by go-swagger; DO NOT EDIT.

package e2e

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "magma/orc8r/cloud/go/obsidian/swagger/v1/models"
)

// GetTestsE2eEnodebdTestPkReader is a Reader for the GetTestsE2eEnodebdTestPk structure.
type GetTestsE2eEnodebdTestPkReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTestsE2eEnodebdTestPkReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTestsE2eEnodebdTestPkOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetTestsE2eEnodebdTestPkDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetTestsE2eEnodebdTestPkOK creates a GetTestsE2eEnodebdTestPkOK with default headers values
func NewGetTestsE2eEnodebdTestPkOK() *GetTestsE2eEnodebdTestPkOK {
	return &GetTestsE2eEnodebdTestPkOK{}
}

/*GetTestsE2eEnodebdTestPkOK handles this case with default header values.

Requested enodebd test case
*/
type GetTestsE2eEnodebdTestPkOK struct {
	Payload *models.EnodebdTestConfig
}

func (o *GetTestsE2eEnodebdTestPkOK) Error() string {
	return fmt.Sprintf("[GET /tests/e2e/enodebd/{test_pk}][%d] getTestsE2eEnodebdTestPkOK  %+v", 200, o.Payload)
}

func (o *GetTestsE2eEnodebdTestPkOK) GetPayload() *models.EnodebdTestConfig {
	return o.Payload
}

func (o *GetTestsE2eEnodebdTestPkOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EnodebdTestConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTestsE2eEnodebdTestPkDefault creates a GetTestsE2eEnodebdTestPkDefault with default headers values
func NewGetTestsE2eEnodebdTestPkDefault(code int) *GetTestsE2eEnodebdTestPkDefault {
	return &GetTestsE2eEnodebdTestPkDefault{
		_statusCode: code,
	}
}

/*GetTestsE2eEnodebdTestPkDefault handles this case with default header values.

Unexpected Error
*/
type GetTestsE2eEnodebdTestPkDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get tests e2e enodebd test pk default response
func (o *GetTestsE2eEnodebdTestPkDefault) Code() int {
	return o._statusCode
}

func (o *GetTestsE2eEnodebdTestPkDefault) Error() string {
	return fmt.Sprintf("[GET /tests/e2e/enodebd/{test_pk}][%d] GetTestsE2eEnodebdTestPk default  %+v", o._statusCode, o.Payload)
}

func (o *GetTestsE2eEnodebdTestPkDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetTestsE2eEnodebdTestPkDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
