// Code generated by go-swagger; DO NOT EDIT.

package lte_networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "magma/orc8r/cloud/go/obsidian/swagger/v1/models"
)

// PutLTENetworkIDDescriptionReader is a Reader for the PutLTENetworkIDDescription structure.
type PutLTENetworkIDDescriptionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutLTENetworkIDDescriptionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPutLTENetworkIDDescriptionNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPutLTENetworkIDDescriptionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutLTENetworkIDDescriptionNoContent creates a PutLTENetworkIDDescriptionNoContent with default headers values
func NewPutLTENetworkIDDescriptionNoContent() *PutLTENetworkIDDescriptionNoContent {
	return &PutLTENetworkIDDescriptionNoContent{}
}

/*PutLTENetworkIDDescriptionNoContent handles this case with default header values.

Success
*/
type PutLTENetworkIDDescriptionNoContent struct {
}

func (o *PutLTENetworkIDDescriptionNoContent) Error() string {
	return fmt.Sprintf("[PUT /lte/{network_id}/description][%d] putLteNetworkIdDescriptionNoContent ", 204)
}

func (o *PutLTENetworkIDDescriptionNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutLTENetworkIDDescriptionDefault creates a PutLTENetworkIDDescriptionDefault with default headers values
func NewPutLTENetworkIDDescriptionDefault(code int) *PutLTENetworkIDDescriptionDefault {
	return &PutLTENetworkIDDescriptionDefault{
		_statusCode: code,
	}
}

/*PutLTENetworkIDDescriptionDefault handles this case with default header values.

Unexpected Error
*/
type PutLTENetworkIDDescriptionDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the put LTE network ID description default response
func (o *PutLTENetworkIDDescriptionDefault) Code() int {
	return o._statusCode
}

func (o *PutLTENetworkIDDescriptionDefault) Error() string {
	return fmt.Sprintf("[PUT /lte/{network_id}/description][%d] PutLTENetworkIDDescription default  %+v", o._statusCode, o.Payload)
}

func (o *PutLTENetworkIDDescriptionDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutLTENetworkIDDescriptionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
