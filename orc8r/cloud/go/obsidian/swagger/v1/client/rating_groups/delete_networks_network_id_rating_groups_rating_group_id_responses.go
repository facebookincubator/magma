// Code generated by go-swagger; DO NOT EDIT.

package rating_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"magma/orc8r/cloud/go/obsidian/swagger/v1/models"
)

// DeleteNetworksNetworkIDRatingGroupsRatingGroupIDReader is a Reader for the DeleteNetworksNetworkIDRatingGroupsRatingGroupID structure.
type DeleteNetworksNetworkIDRatingGroupsRatingGroupIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteNetworksNetworkIDRatingGroupsRatingGroupIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteNetworksNetworkIDRatingGroupsRatingGroupIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteNetworksNetworkIDRatingGroupsRatingGroupIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteNetworksNetworkIDRatingGroupsRatingGroupIDNoContent creates a DeleteNetworksNetworkIDRatingGroupsRatingGroupIDNoContent with default headers values
func NewDeleteNetworksNetworkIDRatingGroupsRatingGroupIDNoContent() *DeleteNetworksNetworkIDRatingGroupsRatingGroupIDNoContent {
	return &DeleteNetworksNetworkIDRatingGroupsRatingGroupIDNoContent{}
}

/* DeleteNetworksNetworkIDRatingGroupsRatingGroupIDNoContent describes a response with status code 204, with default header values.

Success
*/
type DeleteNetworksNetworkIDRatingGroupsRatingGroupIDNoContent struct {
}

func (o *DeleteNetworksNetworkIDRatingGroupsRatingGroupIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /networks/{network_id}/rating_groups/{rating_group_id}][%d] deleteNetworksNetworkIdRatingGroupsRatingGroupIdNoContent ", 204)
}

func (o *DeleteNetworksNetworkIDRatingGroupsRatingGroupIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteNetworksNetworkIDRatingGroupsRatingGroupIDDefault creates a DeleteNetworksNetworkIDRatingGroupsRatingGroupIDDefault with default headers values
func NewDeleteNetworksNetworkIDRatingGroupsRatingGroupIDDefault(code int) *DeleteNetworksNetworkIDRatingGroupsRatingGroupIDDefault {
	return &DeleteNetworksNetworkIDRatingGroupsRatingGroupIDDefault{
		_statusCode: code,
	}
}

/* DeleteNetworksNetworkIDRatingGroupsRatingGroupIDDefault describes a response with status code -1, with default header values.

Unexpected Error
*/
type DeleteNetworksNetworkIDRatingGroupsRatingGroupIDDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the delete networks network ID rating groups rating group ID default response
func (o *DeleteNetworksNetworkIDRatingGroupsRatingGroupIDDefault) Code() int {
	return o._statusCode
}

func (o *DeleteNetworksNetworkIDRatingGroupsRatingGroupIDDefault) Error() string {
	return fmt.Sprintf("[DELETE /networks/{network_id}/rating_groups/{rating_group_id}][%d] DeleteNetworksNetworkIDRatingGroupsRatingGroupID default  %+v", o._statusCode, o.Payload)
}
func (o *DeleteNetworksNetworkIDRatingGroupsRatingGroupIDDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteNetworksNetworkIDRatingGroupsRatingGroupIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
