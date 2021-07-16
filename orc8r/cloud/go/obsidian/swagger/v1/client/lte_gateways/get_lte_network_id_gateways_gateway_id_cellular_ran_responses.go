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

// GetLTENetworkIDGatewaysGatewayIDCellularRanReader is a Reader for the GetLTENetworkIDGatewaysGatewayIDCellularRan structure.
type GetLTENetworkIDGatewaysGatewayIDCellularRanReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLTENetworkIDGatewaysGatewayIDCellularRanReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewGetLTENetworkIDGatewaysGatewayIDCellularRanNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetLTENetworkIDGatewaysGatewayIDCellularRanDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetLTENetworkIDGatewaysGatewayIDCellularRanNoContent creates a GetLTENetworkIDGatewaysGatewayIDCellularRanNoContent with default headers values
func NewGetLTENetworkIDGatewaysGatewayIDCellularRanNoContent() *GetLTENetworkIDGatewaysGatewayIDCellularRanNoContent {
	return &GetLTENetworkIDGatewaysGatewayIDCellularRanNoContent{}
}

/* GetLTENetworkIDGatewaysGatewayIDCellularRanNoContent describes a response with status code 204, with default header values.

RAN configuration
*/
type GetLTENetworkIDGatewaysGatewayIDCellularRanNoContent struct {
	Payload *models.GatewayRanConfigs
}

func (o *GetLTENetworkIDGatewaysGatewayIDCellularRanNoContent) Error() string {
	return fmt.Sprintf("[GET /lte/{network_id}/gateways/{gateway_id}/cellular/ran][%d] getLteNetworkIdGatewaysGatewayIdCellularRanNoContent  %+v", 204, o.Payload)
}
func (o *GetLTENetworkIDGatewaysGatewayIDCellularRanNoContent) GetPayload() *models.GatewayRanConfigs {
	return o.Payload
}

func (o *GetLTENetworkIDGatewaysGatewayIDCellularRanNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayRanConfigs)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLTENetworkIDGatewaysGatewayIDCellularRanDefault creates a GetLTENetworkIDGatewaysGatewayIDCellularRanDefault with default headers values
func NewGetLTENetworkIDGatewaysGatewayIDCellularRanDefault(code int) *GetLTENetworkIDGatewaysGatewayIDCellularRanDefault {
	return &GetLTENetworkIDGatewaysGatewayIDCellularRanDefault{
		_statusCode: code,
	}
}

/* GetLTENetworkIDGatewaysGatewayIDCellularRanDefault describes a response with status code -1, with default header values.

Unexpected Error
*/
type GetLTENetworkIDGatewaysGatewayIDCellularRanDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get LTE network ID gateways gateway ID cellular ran default response
func (o *GetLTENetworkIDGatewaysGatewayIDCellularRanDefault) Code() int {
	return o._statusCode
}

func (o *GetLTENetworkIDGatewaysGatewayIDCellularRanDefault) Error() string {
	return fmt.Sprintf("[GET /lte/{network_id}/gateways/{gateway_id}/cellular/ran][%d] GetLTENetworkIDGatewaysGatewayIDCellularRan default  %+v", o._statusCode, o.Payload)
}
func (o *GetLTENetworkIDGatewaysGatewayIDCellularRanDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetLTENetworkIDGatewaysGatewayIDCellularRanDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
