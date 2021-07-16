// Code generated by go-swagger; DO NOT EDIT.

package networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"magma/orc8r/cloud/go/obsidian/swagger/v1/models"
)

// PutNetworksNetworkIDSentryReader is a Reader for the PutNetworksNetworkIDSentry structure.
type PutNetworksNetworkIDSentryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutNetworksNetworkIDSentryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPutNetworksNetworkIDSentryNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPutNetworksNetworkIDSentryDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutNetworksNetworkIDSentryNoContent creates a PutNetworksNetworkIDSentryNoContent with default headers values
func NewPutNetworksNetworkIDSentryNoContent() *PutNetworksNetworkIDSentryNoContent {
	return &PutNetworksNetworkIDSentryNoContent{}
}

/* PutNetworksNetworkIDSentryNoContent describes a response with status code 204, with default header values.

Success
*/
type PutNetworksNetworkIDSentryNoContent struct {
}

func (o *PutNetworksNetworkIDSentryNoContent) Error() string {
	return fmt.Sprintf("[PUT /networks/{network_id}/sentry][%d] putNetworksNetworkIdSentryNoContent ", 204)
}

func (o *PutNetworksNetworkIDSentryNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutNetworksNetworkIDSentryDefault creates a PutNetworksNetworkIDSentryDefault with default headers values
func NewPutNetworksNetworkIDSentryDefault(code int) *PutNetworksNetworkIDSentryDefault {
	return &PutNetworksNetworkIDSentryDefault{
		_statusCode: code,
	}
}

/* PutNetworksNetworkIDSentryDefault describes a response with status code -1, with default header values.

Unexpected Error
*/
type PutNetworksNetworkIDSentryDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the put networks network ID sentry default response
func (o *PutNetworksNetworkIDSentryDefault) Code() int {
	return o._statusCode
}

func (o *PutNetworksNetworkIDSentryDefault) Error() string {
	return fmt.Sprintf("[PUT /networks/{network_id}/sentry][%d] PutNetworksNetworkIDSentry default  %+v", o._statusCode, o.Payload)
}
func (o *PutNetworksNetworkIDSentryDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutNetworksNetworkIDSentryDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
