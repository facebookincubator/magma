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

// PutNetworksNetworkIDGatewaysGatewayIDMagmadReader is a Reader for the PutNetworksNetworkIDGatewaysGatewayIDMagmad structure.
type PutNetworksNetworkIDGatewaysGatewayIDMagmadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutNetworksNetworkIDGatewaysGatewayIDMagmadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPutNetworksNetworkIDGatewaysGatewayIDMagmadNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPutNetworksNetworkIDGatewaysGatewayIDMagmadDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutNetworksNetworkIDGatewaysGatewayIDMagmadNoContent creates a PutNetworksNetworkIDGatewaysGatewayIDMagmadNoContent with default headers values
func NewPutNetworksNetworkIDGatewaysGatewayIDMagmadNoContent() *PutNetworksNetworkIDGatewaysGatewayIDMagmadNoContent {
	return &PutNetworksNetworkIDGatewaysGatewayIDMagmadNoContent{}
}

/* PutNetworksNetworkIDGatewaysGatewayIDMagmadNoContent describes a response with status code 204, with default header values.

Success
*/
type PutNetworksNetworkIDGatewaysGatewayIDMagmadNoContent struct {
}

func (o *PutNetworksNetworkIDGatewaysGatewayIDMagmadNoContent) Error() string {
	return fmt.Sprintf("[PUT /networks/{network_id}/gateways/{gateway_id}/magmad][%d] putNetworksNetworkIdGatewaysGatewayIdMagmadNoContent ", 204)
}

func (o *PutNetworksNetworkIDGatewaysGatewayIDMagmadNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutNetworksNetworkIDGatewaysGatewayIDMagmadDefault creates a PutNetworksNetworkIDGatewaysGatewayIDMagmadDefault with default headers values
func NewPutNetworksNetworkIDGatewaysGatewayIDMagmadDefault(code int) *PutNetworksNetworkIDGatewaysGatewayIDMagmadDefault {
	return &PutNetworksNetworkIDGatewaysGatewayIDMagmadDefault{
		_statusCode: code,
	}
}

/* PutNetworksNetworkIDGatewaysGatewayIDMagmadDefault describes a response with status code -1, with default header values.

Unexpected Error
*/
type PutNetworksNetworkIDGatewaysGatewayIDMagmadDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the put networks network ID gateways gateway ID magmad default response
func (o *PutNetworksNetworkIDGatewaysGatewayIDMagmadDefault) Code() int {
	return o._statusCode
}

func (o *PutNetworksNetworkIDGatewaysGatewayIDMagmadDefault) Error() string {
	return fmt.Sprintf("[PUT /networks/{network_id}/gateways/{gateway_id}/magmad][%d] PutNetworksNetworkIDGatewaysGatewayIDMagmad default  %+v", o._statusCode, o.Payload)
}
func (o *PutNetworksNetworkIDGatewaysGatewayIDMagmadDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutNetworksNetworkIDGatewaysGatewayIDMagmadDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
