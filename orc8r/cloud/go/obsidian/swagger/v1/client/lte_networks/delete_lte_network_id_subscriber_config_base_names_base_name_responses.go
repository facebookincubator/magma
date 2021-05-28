// Code generated by go-swagger; DO NOT EDIT.

package lte_networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "magma/orc8r/cloud/go/obsidian/swagger/v1/models"
)

// DeleteLTENetworkIDSubscriberConfigBaseNamesBaseNameReader is a Reader for the DeleteLTENetworkIDSubscriberConfigBaseNamesBaseName structure.
type DeleteLTENetworkIDSubscriberConfigBaseNamesBaseNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteLTENetworkIDSubscriberConfigBaseNamesBaseNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteLTENetworkIDSubscriberConfigBaseNamesBaseNameNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteLTENetworkIDSubscriberConfigBaseNamesBaseNameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteLTENetworkIDSubscriberConfigBaseNamesBaseNameNoContent creates a DeleteLTENetworkIDSubscriberConfigBaseNamesBaseNameNoContent with default headers values
func NewDeleteLTENetworkIDSubscriberConfigBaseNamesBaseNameNoContent() *DeleteLTENetworkIDSubscriberConfigBaseNamesBaseNameNoContent {
	return &DeleteLTENetworkIDSubscriberConfigBaseNamesBaseNameNoContent{}
}

/*DeleteLTENetworkIDSubscriberConfigBaseNamesBaseNameNoContent handles this case with default header values.

Success
*/
type DeleteLTENetworkIDSubscriberConfigBaseNamesBaseNameNoContent struct {
}

func (o *DeleteLTENetworkIDSubscriberConfigBaseNamesBaseNameNoContent) Error() string {
	return fmt.Sprintf("[DELETE /lte/{network_id}/subscriber_config/base_names/{base_name}][%d] deleteLteNetworkIdSubscriberConfigBaseNamesBaseNameNoContent ", 204)
}

func (o *DeleteLTENetworkIDSubscriberConfigBaseNamesBaseNameNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteLTENetworkIDSubscriberConfigBaseNamesBaseNameDefault creates a DeleteLTENetworkIDSubscriberConfigBaseNamesBaseNameDefault with default headers values
func NewDeleteLTENetworkIDSubscriberConfigBaseNamesBaseNameDefault(code int) *DeleteLTENetworkIDSubscriberConfigBaseNamesBaseNameDefault {
	return &DeleteLTENetworkIDSubscriberConfigBaseNamesBaseNameDefault{
		_statusCode: code,
	}
}

/*DeleteLTENetworkIDSubscriberConfigBaseNamesBaseNameDefault handles this case with default header values.

Unexpected Error
*/
type DeleteLTENetworkIDSubscriberConfigBaseNamesBaseNameDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the delete LTE network ID subscriber config base names base name default response
func (o *DeleteLTENetworkIDSubscriberConfigBaseNamesBaseNameDefault) Code() int {
	return o._statusCode
}

func (o *DeleteLTENetworkIDSubscriberConfigBaseNamesBaseNameDefault) Error() string {
	return fmt.Sprintf("[DELETE /lte/{network_id}/subscriber_config/base_names/{base_name}][%d] DeleteLTENetworkIDSubscriberConfigBaseNamesBaseName default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteLTENetworkIDSubscriberConfigBaseNamesBaseNameDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteLTENetworkIDSubscriberConfigBaseNamesBaseNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
