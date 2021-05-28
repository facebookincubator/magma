// Code generated by go-swagger; DO NOT EDIT.

package lte_gateways

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "magma/orc8r/cloud/go/obsidian/swagger/v1/models"
)

// PutLTENetworkIDGatewaysGatewayIDCellularNonEpsReader is a Reader for the PutLTENetworkIDGatewaysGatewayIDCellularNonEps structure.
type PutLTENetworkIDGatewaysGatewayIDCellularNonEpsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutLTENetworkIDGatewaysGatewayIDCellularNonEpsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPutLTENetworkIDGatewaysGatewayIDCellularNonEpsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPutLTENetworkIDGatewaysGatewayIDCellularNonEpsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutLTENetworkIDGatewaysGatewayIDCellularNonEpsNoContent creates a PutLTENetworkIDGatewaysGatewayIDCellularNonEpsNoContent with default headers values
func NewPutLTENetworkIDGatewaysGatewayIDCellularNonEpsNoContent() *PutLTENetworkIDGatewaysGatewayIDCellularNonEpsNoContent {
	return &PutLTENetworkIDGatewaysGatewayIDCellularNonEpsNoContent{}
}

/*PutLTENetworkIDGatewaysGatewayIDCellularNonEpsNoContent handles this case with default header values.

Success
*/
type PutLTENetworkIDGatewaysGatewayIDCellularNonEpsNoContent struct {
}

func (o *PutLTENetworkIDGatewaysGatewayIDCellularNonEpsNoContent) Error() string {
	return fmt.Sprintf("[PUT /lte/{network_id}/gateways/{gateway_id}/cellular/non_eps][%d] putLteNetworkIdGatewaysGatewayIdCellularNonEpsNoContent ", 204)
}

func (o *PutLTENetworkIDGatewaysGatewayIDCellularNonEpsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutLTENetworkIDGatewaysGatewayIDCellularNonEpsDefault creates a PutLTENetworkIDGatewaysGatewayIDCellularNonEpsDefault with default headers values
func NewPutLTENetworkIDGatewaysGatewayIDCellularNonEpsDefault(code int) *PutLTENetworkIDGatewaysGatewayIDCellularNonEpsDefault {
	return &PutLTENetworkIDGatewaysGatewayIDCellularNonEpsDefault{
		_statusCode: code,
	}
}

/*PutLTENetworkIDGatewaysGatewayIDCellularNonEpsDefault handles this case with default header values.

Unexpected Error
*/
type PutLTENetworkIDGatewaysGatewayIDCellularNonEpsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the put LTE network ID gateways gateway ID cellular non eps default response
func (o *PutLTENetworkIDGatewaysGatewayIDCellularNonEpsDefault) Code() int {
	return o._statusCode
}

func (o *PutLTENetworkIDGatewaysGatewayIDCellularNonEpsDefault) Error() string {
	return fmt.Sprintf("[PUT /lte/{network_id}/gateways/{gateway_id}/cellular/non_eps][%d] PutLTENetworkIDGatewaysGatewayIDCellularNonEps default  %+v", o._statusCode, o.Payload)
}

func (o *PutLTENetworkIDGatewaysGatewayIDCellularNonEpsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutLTENetworkIDGatewaysGatewayIDCellularNonEpsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
