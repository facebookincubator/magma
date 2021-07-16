// Code generated by go-swagger; DO NOT EDIT.

package subscribers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"magma/orc8r/cloud/go/obsidian/swagger/v1/models"
)

// DeleteLTENetworkIDSubscribersSubscriberIDReader is a Reader for the DeleteLTENetworkIDSubscribersSubscriberID structure.
type DeleteLTENetworkIDSubscribersSubscriberIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteLTENetworkIDSubscribersSubscriberIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteLTENetworkIDSubscribersSubscriberIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteLTENetworkIDSubscribersSubscriberIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteLTENetworkIDSubscribersSubscriberIDNoContent creates a DeleteLTENetworkIDSubscribersSubscriberIDNoContent with default headers values
func NewDeleteLTENetworkIDSubscribersSubscriberIDNoContent() *DeleteLTENetworkIDSubscribersSubscriberIDNoContent {
	return &DeleteLTENetworkIDSubscribersSubscriberIDNoContent{}
}

/* DeleteLTENetworkIDSubscribersSubscriberIDNoContent describes a response with status code 204, with default header values.

Success
*/
type DeleteLTENetworkIDSubscribersSubscriberIDNoContent struct {
}

func (o *DeleteLTENetworkIDSubscribersSubscriberIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /lte/{network_id}/subscribers/{subscriber_id}][%d] deleteLteNetworkIdSubscribersSubscriberIdNoContent ", 204)
}

func (o *DeleteLTENetworkIDSubscribersSubscriberIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteLTENetworkIDSubscribersSubscriberIDDefault creates a DeleteLTENetworkIDSubscribersSubscriberIDDefault with default headers values
func NewDeleteLTENetworkIDSubscribersSubscriberIDDefault(code int) *DeleteLTENetworkIDSubscribersSubscriberIDDefault {
	return &DeleteLTENetworkIDSubscribersSubscriberIDDefault{
		_statusCode: code,
	}
}

/* DeleteLTENetworkIDSubscribersSubscriberIDDefault describes a response with status code -1, with default header values.

Unexpected Error
*/
type DeleteLTENetworkIDSubscribersSubscriberIDDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the delete LTE network ID subscribers subscriber ID default response
func (o *DeleteLTENetworkIDSubscribersSubscriberIDDefault) Code() int {
	return o._statusCode
}

func (o *DeleteLTENetworkIDSubscribersSubscriberIDDefault) Error() string {
	return fmt.Sprintf("[DELETE /lte/{network_id}/subscribers/{subscriber_id}][%d] DeleteLTENetworkIDSubscribersSubscriberID default  %+v", o._statusCode, o.Payload)
}
func (o *DeleteLTENetworkIDSubscribersSubscriberIDDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteLTENetworkIDSubscribersSubscriberIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
