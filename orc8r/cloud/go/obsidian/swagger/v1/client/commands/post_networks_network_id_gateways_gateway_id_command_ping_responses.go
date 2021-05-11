// Code generated by go-swagger; DO NOT EDIT.

package commands

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "magma/orc8r/cloud/go/obsidian/swagger/v1/models"
)

// PostNetworksNetworkIDGatewaysGatewayIDCommandPingReader is a Reader for the PostNetworksNetworkIDGatewaysGatewayIDCommandPing structure.
type PostNetworksNetworkIDGatewaysGatewayIDCommandPingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostNetworksNetworkIDGatewaysGatewayIDCommandPingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostNetworksNetworkIDGatewaysGatewayIDCommandPingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostNetworksNetworkIDGatewaysGatewayIDCommandPingDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostNetworksNetworkIDGatewaysGatewayIDCommandPingOK creates a PostNetworksNetworkIDGatewaysGatewayIDCommandPingOK with default headers values
func NewPostNetworksNetworkIDGatewaysGatewayIDCommandPingOK() *PostNetworksNetworkIDGatewaysGatewayIDCommandPingOK {
	return &PostNetworksNetworkIDGatewaysGatewayIDCommandPingOK{}
}

/*PostNetworksNetworkIDGatewaysGatewayIDCommandPingOK handles this case with default header values.

Success
*/
type PostNetworksNetworkIDGatewaysGatewayIDCommandPingOK struct {
	Payload *models.PingResponse
}

func (o *PostNetworksNetworkIDGatewaysGatewayIDCommandPingOK) Error() string {
	return fmt.Sprintf("[POST /networks/{network_id}/gateways/{gateway_id}/command/ping][%d] postNetworksNetworkIdGatewaysGatewayIdCommandPingOK  %+v", 200, o.Payload)
}

func (o *PostNetworksNetworkIDGatewaysGatewayIDCommandPingOK) GetPayload() *models.PingResponse {
	return o.Payload
}

func (o *PostNetworksNetworkIDGatewaysGatewayIDCommandPingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PingResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostNetworksNetworkIDGatewaysGatewayIDCommandPingDefault creates a PostNetworksNetworkIDGatewaysGatewayIDCommandPingDefault with default headers values
func NewPostNetworksNetworkIDGatewaysGatewayIDCommandPingDefault(code int) *PostNetworksNetworkIDGatewaysGatewayIDCommandPingDefault {
	return &PostNetworksNetworkIDGatewaysGatewayIDCommandPingDefault{
		_statusCode: code,
	}
}

/*PostNetworksNetworkIDGatewaysGatewayIDCommandPingDefault handles this case with default header values.

Unexpected Error
*/
type PostNetworksNetworkIDGatewaysGatewayIDCommandPingDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the post networks network ID gateways gateway ID command ping default response
func (o *PostNetworksNetworkIDGatewaysGatewayIDCommandPingDefault) Code() int {
	return o._statusCode
}

func (o *PostNetworksNetworkIDGatewaysGatewayIDCommandPingDefault) Error() string {
	return fmt.Sprintf("[POST /networks/{network_id}/gateways/{gateway_id}/command/ping][%d] PostNetworksNetworkIDGatewaysGatewayIDCommandPing default  %+v", o._statusCode, o.Payload)
}

func (o *PostNetworksNetworkIDGatewaysGatewayIDCommandPingDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostNetworksNetworkIDGatewaysGatewayIDCommandPingDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
