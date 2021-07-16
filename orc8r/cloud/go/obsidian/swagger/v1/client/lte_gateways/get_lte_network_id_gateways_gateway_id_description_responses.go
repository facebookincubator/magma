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

// GetLTENetworkIDGatewaysGatewayIDDescriptionReader is a Reader for the GetLTENetworkIDGatewaysGatewayIDDescription structure.
type GetLTENetworkIDGatewaysGatewayIDDescriptionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLTENetworkIDGatewaysGatewayIDDescriptionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLTENetworkIDGatewaysGatewayIDDescriptionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetLTENetworkIDGatewaysGatewayIDDescriptionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetLTENetworkIDGatewaysGatewayIDDescriptionOK creates a GetLTENetworkIDGatewaysGatewayIDDescriptionOK with default headers values
func NewGetLTENetworkIDGatewaysGatewayIDDescriptionOK() *GetLTENetworkIDGatewaysGatewayIDDescriptionOK {
	return &GetLTENetworkIDGatewaysGatewayIDDescriptionOK{}
}

/* GetLTENetworkIDGatewaysGatewayIDDescriptionOK describes a response with status code 200, with default header values.

The description of the gateway
*/
type GetLTENetworkIDGatewaysGatewayIDDescriptionOK struct {
	Payload models.GatewayDescription
}

func (o *GetLTENetworkIDGatewaysGatewayIDDescriptionOK) Error() string {
	return fmt.Sprintf("[GET /lte/{network_id}/gateways/{gateway_id}/description][%d] getLteNetworkIdGatewaysGatewayIdDescriptionOK  %+v", 200, o.Payload)
}
func (o *GetLTENetworkIDGatewaysGatewayIDDescriptionOK) GetPayload() models.GatewayDescription {
	return o.Payload
}

func (o *GetLTENetworkIDGatewaysGatewayIDDescriptionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLTENetworkIDGatewaysGatewayIDDescriptionDefault creates a GetLTENetworkIDGatewaysGatewayIDDescriptionDefault with default headers values
func NewGetLTENetworkIDGatewaysGatewayIDDescriptionDefault(code int) *GetLTENetworkIDGatewaysGatewayIDDescriptionDefault {
	return &GetLTENetworkIDGatewaysGatewayIDDescriptionDefault{
		_statusCode: code,
	}
}

/* GetLTENetworkIDGatewaysGatewayIDDescriptionDefault describes a response with status code -1, with default header values.

Unexpected Error
*/
type GetLTENetworkIDGatewaysGatewayIDDescriptionDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get LTE network ID gateways gateway ID description default response
func (o *GetLTENetworkIDGatewaysGatewayIDDescriptionDefault) Code() int {
	return o._statusCode
}

func (o *GetLTENetworkIDGatewaysGatewayIDDescriptionDefault) Error() string {
	return fmt.Sprintf("[GET /lte/{network_id}/gateways/{gateway_id}/description][%d] GetLTENetworkIDGatewaysGatewayIDDescription default  %+v", o._statusCode, o.Payload)
}
func (o *GetLTENetworkIDGatewaysGatewayIDDescriptionDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetLTENetworkIDGatewaysGatewayIDDescriptionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
