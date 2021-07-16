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

// GetWifiNetworkIDGatewaysGatewayIDReader is a Reader for the GetWifiNetworkIDGatewaysGatewayID structure.
type GetWifiNetworkIDGatewaysGatewayIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWifiNetworkIDGatewaysGatewayIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWifiNetworkIDGatewaysGatewayIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetWifiNetworkIDGatewaysGatewayIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetWifiNetworkIDGatewaysGatewayIDOK creates a GetWifiNetworkIDGatewaysGatewayIDOK with default headers values
func NewGetWifiNetworkIDGatewaysGatewayIDOK() *GetWifiNetworkIDGatewaysGatewayIDOK {
	return &GetWifiNetworkIDGatewaysGatewayIDOK{}
}

/* GetWifiNetworkIDGatewaysGatewayIDOK describes a response with status code 200, with default header values.

The requested Wifi gateway
*/
type GetWifiNetworkIDGatewaysGatewayIDOK struct {
	Payload *models.WifiGateway
}

func (o *GetWifiNetworkIDGatewaysGatewayIDOK) Error() string {
	return fmt.Sprintf("[GET /wifi/{network_id}/gateways/{gateway_id}][%d] getWifiNetworkIdGatewaysGatewayIdOK  %+v", 200, o.Payload)
}
func (o *GetWifiNetworkIDGatewaysGatewayIDOK) GetPayload() *models.WifiGateway {
	return o.Payload
}

func (o *GetWifiNetworkIDGatewaysGatewayIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WifiGateway)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWifiNetworkIDGatewaysGatewayIDDefault creates a GetWifiNetworkIDGatewaysGatewayIDDefault with default headers values
func NewGetWifiNetworkIDGatewaysGatewayIDDefault(code int) *GetWifiNetworkIDGatewaysGatewayIDDefault {
	return &GetWifiNetworkIDGatewaysGatewayIDDefault{
		_statusCode: code,
	}
}

/* GetWifiNetworkIDGatewaysGatewayIDDefault describes a response with status code -1, with default header values.

Unexpected Error
*/
type GetWifiNetworkIDGatewaysGatewayIDDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get wifi network ID gateways gateway ID default response
func (o *GetWifiNetworkIDGatewaysGatewayIDDefault) Code() int {
	return o._statusCode
}

func (o *GetWifiNetworkIDGatewaysGatewayIDDefault) Error() string {
	return fmt.Sprintf("[GET /wifi/{network_id}/gateways/{gateway_id}][%d] GetWifiNetworkIDGatewaysGatewayID default  %+v", o._statusCode, o.Payload)
}
func (o *GetWifiNetworkIDGatewaysGatewayIDDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetWifiNetworkIDGatewaysGatewayIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
