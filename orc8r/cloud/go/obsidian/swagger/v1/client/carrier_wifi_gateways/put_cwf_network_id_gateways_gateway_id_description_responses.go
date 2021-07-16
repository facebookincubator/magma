// Code generated by go-swagger; DO NOT EDIT.

package carrier_wifi_gateways

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"magma/orc8r/cloud/go/obsidian/swagger/v1/models"
)

// PutCwfNetworkIDGatewaysGatewayIDDescriptionReader is a Reader for the PutCwfNetworkIDGatewaysGatewayIDDescription structure.
type PutCwfNetworkIDGatewaysGatewayIDDescriptionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutCwfNetworkIDGatewaysGatewayIDDescriptionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPutCwfNetworkIDGatewaysGatewayIDDescriptionNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPutCwfNetworkIDGatewaysGatewayIDDescriptionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutCwfNetworkIDGatewaysGatewayIDDescriptionNoContent creates a PutCwfNetworkIDGatewaysGatewayIDDescriptionNoContent with default headers values
func NewPutCwfNetworkIDGatewaysGatewayIDDescriptionNoContent() *PutCwfNetworkIDGatewaysGatewayIDDescriptionNoContent {
	return &PutCwfNetworkIDGatewaysGatewayIDDescriptionNoContent{}
}

/* PutCwfNetworkIDGatewaysGatewayIDDescriptionNoContent describes a response with status code 204, with default header values.

Success
*/
type PutCwfNetworkIDGatewaysGatewayIDDescriptionNoContent struct {
}

func (o *PutCwfNetworkIDGatewaysGatewayIDDescriptionNoContent) Error() string {
	return fmt.Sprintf("[PUT /cwf/{network_id}/gateways/{gateway_id}/description][%d] putCwfNetworkIdGatewaysGatewayIdDescriptionNoContent ", 204)
}

func (o *PutCwfNetworkIDGatewaysGatewayIDDescriptionNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutCwfNetworkIDGatewaysGatewayIDDescriptionDefault creates a PutCwfNetworkIDGatewaysGatewayIDDescriptionDefault with default headers values
func NewPutCwfNetworkIDGatewaysGatewayIDDescriptionDefault(code int) *PutCwfNetworkIDGatewaysGatewayIDDescriptionDefault {
	return &PutCwfNetworkIDGatewaysGatewayIDDescriptionDefault{
		_statusCode: code,
	}
}

/* PutCwfNetworkIDGatewaysGatewayIDDescriptionDefault describes a response with status code -1, with default header values.

Unexpected Error
*/
type PutCwfNetworkIDGatewaysGatewayIDDescriptionDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the put cwf network ID gateways gateway ID description default response
func (o *PutCwfNetworkIDGatewaysGatewayIDDescriptionDefault) Code() int {
	return o._statusCode
}

func (o *PutCwfNetworkIDGatewaysGatewayIDDescriptionDefault) Error() string {
	return fmt.Sprintf("[PUT /cwf/{network_id}/gateways/{gateway_id}/description][%d] PutCwfNetworkIDGatewaysGatewayIDDescription default  %+v", o._statusCode, o.Payload)
}
func (o *PutCwfNetworkIDGatewaysGatewayIDDescriptionDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutCwfNetworkIDGatewaysGatewayIDDescriptionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
