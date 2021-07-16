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

// PutCwfNetworkIDGatewaysGatewayIDMagmadReader is a Reader for the PutCwfNetworkIDGatewaysGatewayIDMagmad structure.
type PutCwfNetworkIDGatewaysGatewayIDMagmadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutCwfNetworkIDGatewaysGatewayIDMagmadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPutCwfNetworkIDGatewaysGatewayIDMagmadNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPutCwfNetworkIDGatewaysGatewayIDMagmadDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutCwfNetworkIDGatewaysGatewayIDMagmadNoContent creates a PutCwfNetworkIDGatewaysGatewayIDMagmadNoContent with default headers values
func NewPutCwfNetworkIDGatewaysGatewayIDMagmadNoContent() *PutCwfNetworkIDGatewaysGatewayIDMagmadNoContent {
	return &PutCwfNetworkIDGatewaysGatewayIDMagmadNoContent{}
}

/* PutCwfNetworkIDGatewaysGatewayIDMagmadNoContent describes a response with status code 204, with default header values.

Success
*/
type PutCwfNetworkIDGatewaysGatewayIDMagmadNoContent struct {
}

func (o *PutCwfNetworkIDGatewaysGatewayIDMagmadNoContent) Error() string {
	return fmt.Sprintf("[PUT /cwf/{network_id}/gateways/{gateway_id}/magmad][%d] putCwfNetworkIdGatewaysGatewayIdMagmadNoContent ", 204)
}

func (o *PutCwfNetworkIDGatewaysGatewayIDMagmadNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutCwfNetworkIDGatewaysGatewayIDMagmadDefault creates a PutCwfNetworkIDGatewaysGatewayIDMagmadDefault with default headers values
func NewPutCwfNetworkIDGatewaysGatewayIDMagmadDefault(code int) *PutCwfNetworkIDGatewaysGatewayIDMagmadDefault {
	return &PutCwfNetworkIDGatewaysGatewayIDMagmadDefault{
		_statusCode: code,
	}
}

/* PutCwfNetworkIDGatewaysGatewayIDMagmadDefault describes a response with status code -1, with default header values.

Unexpected Error
*/
type PutCwfNetworkIDGatewaysGatewayIDMagmadDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the put cwf network ID gateways gateway ID magmad default response
func (o *PutCwfNetworkIDGatewaysGatewayIDMagmadDefault) Code() int {
	return o._statusCode
}

func (o *PutCwfNetworkIDGatewaysGatewayIDMagmadDefault) Error() string {
	return fmt.Sprintf("[PUT /cwf/{network_id}/gateways/{gateway_id}/magmad][%d] PutCwfNetworkIDGatewaysGatewayIDMagmad default  %+v", o._statusCode, o.Payload)
}
func (o *PutCwfNetworkIDGatewaysGatewayIDMagmadDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutCwfNetworkIDGatewaysGatewayIDMagmadDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
