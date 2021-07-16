// Code generated by go-swagger; DO NOT EDIT.

package lte_gateways

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"magma/orc8r/cloud/go/obsidian/swagger/v1/models"
)

// PutLTENetworkIDGatewaysGatewayIDReader is a Reader for the PutLTENetworkIDGatewaysGatewayID structure.
type PutLTENetworkIDGatewaysGatewayIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutLTENetworkIDGatewaysGatewayIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPutLTENetworkIDGatewaysGatewayIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPutLTENetworkIDGatewaysGatewayIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutLTENetworkIDGatewaysGatewayIDNoContent creates a PutLTENetworkIDGatewaysGatewayIDNoContent with default headers values
func NewPutLTENetworkIDGatewaysGatewayIDNoContent() *PutLTENetworkIDGatewaysGatewayIDNoContent {
	return &PutLTENetworkIDGatewaysGatewayIDNoContent{}
}

/* PutLTENetworkIDGatewaysGatewayIDNoContent describes a response with status code 204, with default header values.

Success
*/
type PutLTENetworkIDGatewaysGatewayIDNoContent struct {
}

func (o *PutLTENetworkIDGatewaysGatewayIDNoContent) Error() string {
	return fmt.Sprintf("[PUT /lte/{network_id}/gateways/{gateway_id}][%d] putLteNetworkIdGatewaysGatewayIdNoContent ", 204)
}

func (o *PutLTENetworkIDGatewaysGatewayIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutLTENetworkIDGatewaysGatewayIDDefault creates a PutLTENetworkIDGatewaysGatewayIDDefault with default headers values
func NewPutLTENetworkIDGatewaysGatewayIDDefault(code int) *PutLTENetworkIDGatewaysGatewayIDDefault {
	return &PutLTENetworkIDGatewaysGatewayIDDefault{
		_statusCode: code,
	}
}

/* PutLTENetworkIDGatewaysGatewayIDDefault describes a response with status code -1, with default header values.

Unexpected Error
*/
type PutLTENetworkIDGatewaysGatewayIDDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the put LTE network ID gateways gateway ID default response
func (o *PutLTENetworkIDGatewaysGatewayIDDefault) Code() int {
	return o._statusCode
}

func (o *PutLTENetworkIDGatewaysGatewayIDDefault) Error() string {
	return fmt.Sprintf("[PUT /lte/{network_id}/gateways/{gateway_id}][%d] PutLTENetworkIDGatewaysGatewayID default  %+v", o._statusCode, o.Payload)
}
func (o *PutLTENetworkIDGatewaysGatewayIDDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutLTENetworkIDGatewaysGatewayIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
