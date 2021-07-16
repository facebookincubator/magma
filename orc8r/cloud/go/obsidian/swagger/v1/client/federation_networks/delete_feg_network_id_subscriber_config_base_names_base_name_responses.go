// Code generated by go-swagger; DO NOT EDIT.

package federation_networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"magma/orc8r/cloud/go/obsidian/swagger/v1/models"
)

// DeleteFegNetworkIDSubscriberConfigBaseNamesBaseNameReader is a Reader for the DeleteFegNetworkIDSubscriberConfigBaseNamesBaseName structure.
type DeleteFegNetworkIDSubscriberConfigBaseNamesBaseNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteFegNetworkIDSubscriberConfigBaseNamesBaseNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteFegNetworkIDSubscriberConfigBaseNamesBaseNameNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteFegNetworkIDSubscriberConfigBaseNamesBaseNameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteFegNetworkIDSubscriberConfigBaseNamesBaseNameNoContent creates a DeleteFegNetworkIDSubscriberConfigBaseNamesBaseNameNoContent with default headers values
func NewDeleteFegNetworkIDSubscriberConfigBaseNamesBaseNameNoContent() *DeleteFegNetworkIDSubscriberConfigBaseNamesBaseNameNoContent {
	return &DeleteFegNetworkIDSubscriberConfigBaseNamesBaseNameNoContent{}
}

/* DeleteFegNetworkIDSubscriberConfigBaseNamesBaseNameNoContent describes a response with status code 204, with default header values.

Success
*/
type DeleteFegNetworkIDSubscriberConfigBaseNamesBaseNameNoContent struct {
}

func (o *DeleteFegNetworkIDSubscriberConfigBaseNamesBaseNameNoContent) Error() string {
	return fmt.Sprintf("[DELETE /feg/{network_id}/subscriber_config/base_names/{base_name}][%d] deleteFegNetworkIdSubscriberConfigBaseNamesBaseNameNoContent ", 204)
}

func (o *DeleteFegNetworkIDSubscriberConfigBaseNamesBaseNameNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteFegNetworkIDSubscriberConfigBaseNamesBaseNameDefault creates a DeleteFegNetworkIDSubscriberConfigBaseNamesBaseNameDefault with default headers values
func NewDeleteFegNetworkIDSubscriberConfigBaseNamesBaseNameDefault(code int) *DeleteFegNetworkIDSubscriberConfigBaseNamesBaseNameDefault {
	return &DeleteFegNetworkIDSubscriberConfigBaseNamesBaseNameDefault{
		_statusCode: code,
	}
}

/* DeleteFegNetworkIDSubscriberConfigBaseNamesBaseNameDefault describes a response with status code -1, with default header values.

Unexpected Error
*/
type DeleteFegNetworkIDSubscriberConfigBaseNamesBaseNameDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the delete feg network ID subscriber config base names base name default response
func (o *DeleteFegNetworkIDSubscriberConfigBaseNamesBaseNameDefault) Code() int {
	return o._statusCode
}

func (o *DeleteFegNetworkIDSubscriberConfigBaseNamesBaseNameDefault) Error() string {
	return fmt.Sprintf("[DELETE /feg/{network_id}/subscriber_config/base_names/{base_name}][%d] DeleteFegNetworkIDSubscriberConfigBaseNamesBaseName default  %+v", o._statusCode, o.Payload)
}
func (o *DeleteFegNetworkIDSubscriberConfigBaseNamesBaseNameDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteFegNetworkIDSubscriberConfigBaseNamesBaseNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
