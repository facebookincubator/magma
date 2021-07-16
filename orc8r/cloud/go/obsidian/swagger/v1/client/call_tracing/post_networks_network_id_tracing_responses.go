// Code generated by go-swagger; DO NOT EDIT.

package call_tracing

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"magma/orc8r/cloud/go/obsidian/swagger/v1/models"
)

// PostNetworksNetworkIDTracingReader is a Reader for the PostNetworksNetworkIDTracing structure.
type PostNetworksNetworkIDTracingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostNetworksNetworkIDTracingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostNetworksNetworkIDTracingCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostNetworksNetworkIDTracingDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostNetworksNetworkIDTracingCreated creates a PostNetworksNetworkIDTracingCreated with default headers values
func NewPostNetworksNetworkIDTracingCreated() *PostNetworksNetworkIDTracingCreated {
	return &PostNetworksNetworkIDTracingCreated{}
}

/* PostNetworksNetworkIDTracingCreated describes a response with status code 201, with default header values.

ID of created call trace
*/
type PostNetworksNetworkIDTracingCreated struct {
	Payload string
}

func (o *PostNetworksNetworkIDTracingCreated) Error() string {
	return fmt.Sprintf("[POST /networks/{network_id}/tracing][%d] postNetworksNetworkIdTracingCreated  %+v", 201, o.Payload)
}
func (o *PostNetworksNetworkIDTracingCreated) GetPayload() string {
	return o.Payload
}

func (o *PostNetworksNetworkIDTracingCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostNetworksNetworkIDTracingDefault creates a PostNetworksNetworkIDTracingDefault with default headers values
func NewPostNetworksNetworkIDTracingDefault(code int) *PostNetworksNetworkIDTracingDefault {
	return &PostNetworksNetworkIDTracingDefault{
		_statusCode: code,
	}
}

/* PostNetworksNetworkIDTracingDefault describes a response with status code -1, with default header values.

Unexpected Error
*/
type PostNetworksNetworkIDTracingDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the post networks network ID tracing default response
func (o *PostNetworksNetworkIDTracingDefault) Code() int {
	return o._statusCode
}

func (o *PostNetworksNetworkIDTracingDefault) Error() string {
	return fmt.Sprintf("[POST /networks/{network_id}/tracing][%d] PostNetworksNetworkIDTracing default  %+v", o._statusCode, o.Payload)
}
func (o *PostNetworksNetworkIDTracingDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostNetworksNetworkIDTracingDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
