// Code generated by go-swagger; DO NOT EDIT.

package enode_bs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"magma/orc8r/cloud/go/obsidian/swagger/v1/models"
)

// PostLTENetworkIDEnodebsReader is a Reader for the PostLTENetworkIDEnodebs structure.
type PostLTENetworkIDEnodebsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostLTENetworkIDEnodebsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostLTENetworkIDEnodebsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostLTENetworkIDEnodebsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostLTENetworkIDEnodebsCreated creates a PostLTENetworkIDEnodebsCreated with default headers values
func NewPostLTENetworkIDEnodebsCreated() *PostLTENetworkIDEnodebsCreated {
	return &PostLTENetworkIDEnodebsCreated{}
}

/* PostLTENetworkIDEnodebsCreated describes a response with status code 201, with default header values.

Success
*/
type PostLTENetworkIDEnodebsCreated struct {
}

func (o *PostLTENetworkIDEnodebsCreated) Error() string {
	return fmt.Sprintf("[POST /lte/{network_id}/enodebs][%d] postLteNetworkIdEnodebsCreated ", 201)
}

func (o *PostLTENetworkIDEnodebsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostLTENetworkIDEnodebsDefault creates a PostLTENetworkIDEnodebsDefault with default headers values
func NewPostLTENetworkIDEnodebsDefault(code int) *PostLTENetworkIDEnodebsDefault {
	return &PostLTENetworkIDEnodebsDefault{
		_statusCode: code,
	}
}

/* PostLTENetworkIDEnodebsDefault describes a response with status code -1, with default header values.

Unexpected Error
*/
type PostLTENetworkIDEnodebsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the post LTE network ID enodebs default response
func (o *PostLTENetworkIDEnodebsDefault) Code() int {
	return o._statusCode
}

func (o *PostLTENetworkIDEnodebsDefault) Error() string {
	return fmt.Sprintf("[POST /lte/{network_id}/enodebs][%d] PostLTENetworkIDEnodebs default  %+v", o._statusCode, o.Payload)
}
func (o *PostLTENetworkIDEnodebsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostLTENetworkIDEnodebsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
