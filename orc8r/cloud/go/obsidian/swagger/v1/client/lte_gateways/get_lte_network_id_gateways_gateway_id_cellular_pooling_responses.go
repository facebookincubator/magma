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

// GetLTENetworkIDGatewaysGatewayIDCellularPoolingReader is a Reader for the GetLTENetworkIDGatewaysGatewayIDCellularPooling structure.
type GetLTENetworkIDGatewaysGatewayIDCellularPoolingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLTENetworkIDGatewaysGatewayIDCellularPoolingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLTENetworkIDGatewaysGatewayIDCellularPoolingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetLTENetworkIDGatewaysGatewayIDCellularPoolingDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetLTENetworkIDGatewaysGatewayIDCellularPoolingOK creates a GetLTENetworkIDGatewaysGatewayIDCellularPoolingOK with default headers values
func NewGetLTENetworkIDGatewaysGatewayIDCellularPoolingOK() *GetLTENetworkIDGatewaysGatewayIDCellularPoolingOK {
	return &GetLTENetworkIDGatewaysGatewayIDCellularPoolingOK{}
}

/*GetLTENetworkIDGatewaysGatewayIDCellularPoolingOK handles this case with default header values.

Gateway pool records
*/
type GetLTENetworkIDGatewaysGatewayIDCellularPoolingOK struct {
	Payload []*models.CellularGatewayPoolRecord
}

func (o *GetLTENetworkIDGatewaysGatewayIDCellularPoolingOK) Error() string {
	return fmt.Sprintf("[GET /lte/{network_id}/gateways/{gateway_id}/cellular/pooling][%d] getLteNetworkIdGatewaysGatewayIdCellularPoolingOK  %+v", 200, o.Payload)
}

func (o *GetLTENetworkIDGatewaysGatewayIDCellularPoolingOK) GetPayload() []*models.CellularGatewayPoolRecord {
	return o.Payload
}

func (o *GetLTENetworkIDGatewaysGatewayIDCellularPoolingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLTENetworkIDGatewaysGatewayIDCellularPoolingDefault creates a GetLTENetworkIDGatewaysGatewayIDCellularPoolingDefault with default headers values
func NewGetLTENetworkIDGatewaysGatewayIDCellularPoolingDefault(code int) *GetLTENetworkIDGatewaysGatewayIDCellularPoolingDefault {
	return &GetLTENetworkIDGatewaysGatewayIDCellularPoolingDefault{
		_statusCode: code,
	}
}

/*GetLTENetworkIDGatewaysGatewayIDCellularPoolingDefault handles this case with default header values.

Unexpected Error
*/
type GetLTENetworkIDGatewaysGatewayIDCellularPoolingDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get LTE network ID gateways gateway ID cellular pooling default response
func (o *GetLTENetworkIDGatewaysGatewayIDCellularPoolingDefault) Code() int {
	return o._statusCode
}

func (o *GetLTENetworkIDGatewaysGatewayIDCellularPoolingDefault) Error() string {
	return fmt.Sprintf("[GET /lte/{network_id}/gateways/{gateway_id}/cellular/pooling][%d] GetLTENetworkIDGatewaysGatewayIDCellularPooling default  %+v", o._statusCode, o.Payload)
}

func (o *GetLTENetworkIDGatewaysGatewayIDCellularPoolingDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetLTENetworkIDGatewaysGatewayIDCellularPoolingDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
