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

// PutNetworksNetworkIDGatewaysGatewayIDNameReader is a Reader for the PutNetworksNetworkIDGatewaysGatewayIDName structure.
type PutNetworksNetworkIDGatewaysGatewayIDNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutNetworksNetworkIDGatewaysGatewayIDNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPutNetworksNetworkIDGatewaysGatewayIDNameNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPutNetworksNetworkIDGatewaysGatewayIDNameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutNetworksNetworkIDGatewaysGatewayIDNameNoContent creates a PutNetworksNetworkIDGatewaysGatewayIDNameNoContent with default headers values
func NewPutNetworksNetworkIDGatewaysGatewayIDNameNoContent() *PutNetworksNetworkIDGatewaysGatewayIDNameNoContent {
	return &PutNetworksNetworkIDGatewaysGatewayIDNameNoContent{}
}

/* PutNetworksNetworkIDGatewaysGatewayIDNameNoContent describes a response with status code 204, with default header values.

Success
*/
type PutNetworksNetworkIDGatewaysGatewayIDNameNoContent struct {
}

func (o *PutNetworksNetworkIDGatewaysGatewayIDNameNoContent) Error() string {
	return fmt.Sprintf("[PUT /networks/{network_id}/gateways/{gateway_id}/name][%d] putNetworksNetworkIdGatewaysGatewayIdNameNoContent ", 204)
}

func (o *PutNetworksNetworkIDGatewaysGatewayIDNameNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutNetworksNetworkIDGatewaysGatewayIDNameDefault creates a PutNetworksNetworkIDGatewaysGatewayIDNameDefault with default headers values
func NewPutNetworksNetworkIDGatewaysGatewayIDNameDefault(code int) *PutNetworksNetworkIDGatewaysGatewayIDNameDefault {
	return &PutNetworksNetworkIDGatewaysGatewayIDNameDefault{
		_statusCode: code,
	}
}

/* PutNetworksNetworkIDGatewaysGatewayIDNameDefault describes a response with status code -1, with default header values.

Unexpected Error
*/
type PutNetworksNetworkIDGatewaysGatewayIDNameDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the put networks network ID gateways gateway ID name default response
func (o *PutNetworksNetworkIDGatewaysGatewayIDNameDefault) Code() int {
	return o._statusCode
}

func (o *PutNetworksNetworkIDGatewaysGatewayIDNameDefault) Error() string {
	return fmt.Sprintf("[PUT /networks/{network_id}/gateways/{gateway_id}/name][%d] PutNetworksNetworkIDGatewaysGatewayIDName default  %+v", o._statusCode, o.Payload)
}
func (o *PutNetworksNetworkIDGatewaysGatewayIDNameDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutNetworksNetworkIDGatewaysGatewayIDNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
