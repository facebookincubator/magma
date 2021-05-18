// Code generated by go-swagger; DO NOT EDIT.

package baremetal

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "magma/orc8r/cloud/go/obsidian/swagger/v1/models"
)

// GetCiNodesReader is a Reader for the GetCiNodes structure.
type GetCiNodesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCiNodesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCiNodesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetCiNodesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetCiNodesOK creates a GetCiNodesOK with default headers values
func NewGetCiNodesOK() *GetCiNodesOK {
	return &GetCiNodesOK{}
}

/*GetCiNodesOK handles this case with default header values.

All CI worker nodes
*/
type GetCiNodesOK struct {
	Payload []*models.CiNode
}

func (o *GetCiNodesOK) Error() string {
	return fmt.Sprintf("[GET /ci/nodes][%d] getCiNodesOK  %+v", 200, o.Payload)
}

func (o *GetCiNodesOK) GetPayload() []*models.CiNode {
	return o.Payload
}

func (o *GetCiNodesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCiNodesDefault creates a GetCiNodesDefault with default headers values
func NewGetCiNodesDefault(code int) *GetCiNodesDefault {
	return &GetCiNodesDefault{
		_statusCode: code,
	}
}

/*GetCiNodesDefault handles this case with default header values.

Unexpected Error
*/
type GetCiNodesDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get ci nodes default response
func (o *GetCiNodesDefault) Code() int {
	return o._statusCode
}

func (o *GetCiNodesDefault) Error() string {
	return fmt.Sprintf("[GET /ci/nodes][%d] GetCiNodes default  %+v", o._statusCode, o.Payload)
}

func (o *GetCiNodesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetCiNodesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
