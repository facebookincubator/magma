// Code generated by go-swagger; DO NOT EDIT.

package gateways

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"magma/orc8r/cloud/go/obsidian/swagger/v1/models"
)

// PutNetworksNetworkIDGatewaysGatewayIDReader is a Reader for the PutNetworksNetworkIDGatewaysGatewayID structure.
type PutNetworksNetworkIDGatewaysGatewayIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutNetworksNetworkIDGatewaysGatewayIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPutNetworksNetworkIDGatewaysGatewayIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPutNetworksNetworkIDGatewaysGatewayIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutNetworksNetworkIDGatewaysGatewayIDNoContent creates a PutNetworksNetworkIDGatewaysGatewayIDNoContent with default headers values
func NewPutNetworksNetworkIDGatewaysGatewayIDNoContent() *PutNetworksNetworkIDGatewaysGatewayIDNoContent {
	return &PutNetworksNetworkIDGatewaysGatewayIDNoContent{}
}

/* PutNetworksNetworkIDGatewaysGatewayIDNoContent describes a response with status code 204, with default header values.

Success
*/
type PutNetworksNetworkIDGatewaysGatewayIDNoContent struct {
}

func (o *PutNetworksNetworkIDGatewaysGatewayIDNoContent) Error() string {
	return fmt.Sprintf("[PUT /networks/{network_id}/gateways/{gateway_id}][%d] putNetworksNetworkIdGatewaysGatewayIdNoContent ", 204)
}

func (o *PutNetworksNetworkIDGatewaysGatewayIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutNetworksNetworkIDGatewaysGatewayIDDefault creates a PutNetworksNetworkIDGatewaysGatewayIDDefault with default headers values
func NewPutNetworksNetworkIDGatewaysGatewayIDDefault(code int) *PutNetworksNetworkIDGatewaysGatewayIDDefault {
	return &PutNetworksNetworkIDGatewaysGatewayIDDefault{
		_statusCode: code,
	}
}

/* PutNetworksNetworkIDGatewaysGatewayIDDefault describes a response with status code -1, with default header values.

Unexpected Error
*/
type PutNetworksNetworkIDGatewaysGatewayIDDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the put networks network ID gateways gateway ID default response
func (o *PutNetworksNetworkIDGatewaysGatewayIDDefault) Code() int {
	return o._statusCode
}

func (o *PutNetworksNetworkIDGatewaysGatewayIDDefault) Error() string {
	return fmt.Sprintf("[PUT /networks/{network_id}/gateways/{gateway_id}][%d] PutNetworksNetworkIDGatewaysGatewayID default  %+v", o._statusCode, o.Payload)
}
func (o *PutNetworksNetworkIDGatewaysGatewayIDDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutNetworksNetworkIDGatewaysGatewayIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
