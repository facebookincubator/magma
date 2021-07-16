// Code generated by go-swagger; DO NOT EDIT.

package wifi_gateways

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"magma/orc8r/cloud/go/obsidian/swagger/v1/models"
)

// PutWifiNetworkIDGatewaysGatewayIDDeviceReader is a Reader for the PutWifiNetworkIDGatewaysGatewayIDDevice structure.
type PutWifiNetworkIDGatewaysGatewayIDDeviceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutWifiNetworkIDGatewaysGatewayIDDeviceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPutWifiNetworkIDGatewaysGatewayIDDeviceNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPutWifiNetworkIDGatewaysGatewayIDDeviceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutWifiNetworkIDGatewaysGatewayIDDeviceNoContent creates a PutWifiNetworkIDGatewaysGatewayIDDeviceNoContent with default headers values
func NewPutWifiNetworkIDGatewaysGatewayIDDeviceNoContent() *PutWifiNetworkIDGatewaysGatewayIDDeviceNoContent {
	return &PutWifiNetworkIDGatewaysGatewayIDDeviceNoContent{}
}

/* PutWifiNetworkIDGatewaysGatewayIDDeviceNoContent describes a response with status code 204, with default header values.

Success
*/
type PutWifiNetworkIDGatewaysGatewayIDDeviceNoContent struct {
}

func (o *PutWifiNetworkIDGatewaysGatewayIDDeviceNoContent) Error() string {
	return fmt.Sprintf("[PUT /wifi/{network_id}/gateways/{gateway_id}/device][%d] putWifiNetworkIdGatewaysGatewayIdDeviceNoContent ", 204)
}

func (o *PutWifiNetworkIDGatewaysGatewayIDDeviceNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutWifiNetworkIDGatewaysGatewayIDDeviceDefault creates a PutWifiNetworkIDGatewaysGatewayIDDeviceDefault with default headers values
func NewPutWifiNetworkIDGatewaysGatewayIDDeviceDefault(code int) *PutWifiNetworkIDGatewaysGatewayIDDeviceDefault {
	return &PutWifiNetworkIDGatewaysGatewayIDDeviceDefault{
		_statusCode: code,
	}
}

/* PutWifiNetworkIDGatewaysGatewayIDDeviceDefault describes a response with status code -1, with default header values.

Unexpected Error
*/
type PutWifiNetworkIDGatewaysGatewayIDDeviceDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the put wifi network ID gateways gateway ID device default response
func (o *PutWifiNetworkIDGatewaysGatewayIDDeviceDefault) Code() int {
	return o._statusCode
}

func (o *PutWifiNetworkIDGatewaysGatewayIDDeviceDefault) Error() string {
	return fmt.Sprintf("[PUT /wifi/{network_id}/gateways/{gateway_id}/device][%d] PutWifiNetworkIDGatewaysGatewayIDDevice default  %+v", o._statusCode, o.Payload)
}
func (o *PutWifiNetworkIDGatewaysGatewayIDDeviceDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutWifiNetworkIDGatewaysGatewayIDDeviceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
