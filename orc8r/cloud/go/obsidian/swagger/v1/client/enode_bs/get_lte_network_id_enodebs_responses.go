// Code generated by go-swagger; DO NOT EDIT.

package enode_bs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "magma/orc8r/cloud/go/obsidian/swagger/v1/models"
)

// GetLTENetworkIDEnodebsReader is a Reader for the GetLTENetworkIDEnodebs structure.
type GetLTENetworkIDEnodebsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLTENetworkIDEnodebsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLTENetworkIDEnodebsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetLTENetworkIDEnodebsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetLTENetworkIDEnodebsOK creates a GetLTENetworkIDEnodebsOK with default headers values
func NewGetLTENetworkIDEnodebsOK() *GetLTENetworkIDEnodebsOK {
	return &GetLTENetworkIDEnodebsOK{}
}

/*GetLTENetworkIDEnodebsOK handles this case with default header values.

All enodeBs registered in the network
*/
type GetLTENetworkIDEnodebsOK struct {
	Payload map[string]models.ENODEB
}

func (o *GetLTENetworkIDEnodebsOK) Error() string {
	return fmt.Sprintf("[GET /lte/{network_id}/enodebs][%d] getLteNetworkIdEnodebsOK  %+v", 200, o.Payload)
}

func (o *GetLTENetworkIDEnodebsOK) GetPayload() map[string]models.ENODEB {
	return o.Payload
}

func (o *GetLTENetworkIDEnodebsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLTENetworkIDEnodebsDefault creates a GetLTENetworkIDEnodebsDefault with default headers values
func NewGetLTENetworkIDEnodebsDefault(code int) *GetLTENetworkIDEnodebsDefault {
	return &GetLTENetworkIDEnodebsDefault{
		_statusCode: code,
	}
}

/*GetLTENetworkIDEnodebsDefault handles this case with default header values.

Unexpected Error
*/
type GetLTENetworkIDEnodebsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get LTE network ID enodebs default response
func (o *GetLTENetworkIDEnodebsDefault) Code() int {
	return o._statusCode
}

func (o *GetLTENetworkIDEnodebsDefault) Error() string {
	return fmt.Sprintf("[GET /lte/{network_id}/enodebs][%d] GetLTENetworkIDEnodebs default  %+v", o._statusCode, o.Payload)
}

func (o *GetLTENetworkIDEnodebsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetLTENetworkIDEnodebsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
