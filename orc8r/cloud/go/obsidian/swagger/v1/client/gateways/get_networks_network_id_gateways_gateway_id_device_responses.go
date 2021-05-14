// Code generated by go-swagger; DO NOT EDIT.

package gateways

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "magma/orc8r/cloud/go/obsidian/swagger/v1/models"
)

// GetNetworksNetworkIDGatewaysGatewayIDDeviceReader is a Reader for the GetNetworksNetworkIDGatewaysGatewayIDDevice structure.
type GetNetworksNetworkIDGatewaysGatewayIDDeviceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworksNetworkIDGatewaysGatewayIDDeviceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworksNetworkIDGatewaysGatewayIDDeviceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetNetworksNetworkIDGatewaysGatewayIDDeviceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetNetworksNetworkIDGatewaysGatewayIDDeviceOK creates a GetNetworksNetworkIDGatewaysGatewayIDDeviceOK with default headers values
func NewGetNetworksNetworkIDGatewaysGatewayIDDeviceOK() *GetNetworksNetworkIDGatewaysGatewayIDDeviceOK {
	return &GetNetworksNetworkIDGatewaysGatewayIDDeviceOK{}
}

/*GetNetworksNetworkIDGatewaysGatewayIDDeviceOK handles this case with default header values.

The physical device for the gateway
*/
type GetNetworksNetworkIDGatewaysGatewayIDDeviceOK struct {
	Payload *models.GatewayDevice
}

func (o *GetNetworksNetworkIDGatewaysGatewayIDDeviceOK) Error() string {
	return fmt.Sprintf("[GET /networks/{network_id}/gateways/{gateway_id}/device][%d] getNetworksNetworkIdGatewaysGatewayIdDeviceOK  %+v", 200, o.Payload)
}

func (o *GetNetworksNetworkIDGatewaysGatewayIDDeviceOK) GetPayload() *models.GatewayDevice {
	return o.Payload
}

func (o *GetNetworksNetworkIDGatewaysGatewayIDDeviceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayDevice)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNetworksNetworkIDGatewaysGatewayIDDeviceDefault creates a GetNetworksNetworkIDGatewaysGatewayIDDeviceDefault with default headers values
func NewGetNetworksNetworkIDGatewaysGatewayIDDeviceDefault(code int) *GetNetworksNetworkIDGatewaysGatewayIDDeviceDefault {
	return &GetNetworksNetworkIDGatewaysGatewayIDDeviceDefault{
		_statusCode: code,
	}
}

/*GetNetworksNetworkIDGatewaysGatewayIDDeviceDefault handles this case with default header values.

Unexpected Error
*/
type GetNetworksNetworkIDGatewaysGatewayIDDeviceDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get networks network ID gateways gateway ID device default response
func (o *GetNetworksNetworkIDGatewaysGatewayIDDeviceDefault) Code() int {
	return o._statusCode
}

func (o *GetNetworksNetworkIDGatewaysGatewayIDDeviceDefault) Error() string {
	return fmt.Sprintf("[GET /networks/{network_id}/gateways/{gateway_id}/device][%d] GetNetworksNetworkIDGatewaysGatewayIDDevice default  %+v", o._statusCode, o.Payload)
}

func (o *GetNetworksNetworkIDGatewaysGatewayIDDeviceDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetNetworksNetworkIDGatewaysGatewayIDDeviceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
