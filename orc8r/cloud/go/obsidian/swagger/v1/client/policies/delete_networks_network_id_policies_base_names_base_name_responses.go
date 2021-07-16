// Code generated by go-swagger; DO NOT EDIT.

package policies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"magma/orc8r/cloud/go/obsidian/swagger/v1/models"
)

// DeleteNetworksNetworkIDPoliciesBaseNamesBaseNameReader is a Reader for the DeleteNetworksNetworkIDPoliciesBaseNamesBaseName structure.
type DeleteNetworksNetworkIDPoliciesBaseNamesBaseNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteNetworksNetworkIDPoliciesBaseNamesBaseNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteNetworksNetworkIDPoliciesBaseNamesBaseNameNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteNetworksNetworkIDPoliciesBaseNamesBaseNameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteNetworksNetworkIDPoliciesBaseNamesBaseNameNoContent creates a DeleteNetworksNetworkIDPoliciesBaseNamesBaseNameNoContent with default headers values
func NewDeleteNetworksNetworkIDPoliciesBaseNamesBaseNameNoContent() *DeleteNetworksNetworkIDPoliciesBaseNamesBaseNameNoContent {
	return &DeleteNetworksNetworkIDPoliciesBaseNamesBaseNameNoContent{}
}

/* DeleteNetworksNetworkIDPoliciesBaseNamesBaseNameNoContent describes a response with status code 204, with default header values.

Success
*/
type DeleteNetworksNetworkIDPoliciesBaseNamesBaseNameNoContent struct {
}

func (o *DeleteNetworksNetworkIDPoliciesBaseNamesBaseNameNoContent) Error() string {
	return fmt.Sprintf("[DELETE /networks/{network_id}/policies/base_names/{base_name}][%d] deleteNetworksNetworkIdPoliciesBaseNamesBaseNameNoContent ", 204)
}

func (o *DeleteNetworksNetworkIDPoliciesBaseNamesBaseNameNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteNetworksNetworkIDPoliciesBaseNamesBaseNameDefault creates a DeleteNetworksNetworkIDPoliciesBaseNamesBaseNameDefault with default headers values
func NewDeleteNetworksNetworkIDPoliciesBaseNamesBaseNameDefault(code int) *DeleteNetworksNetworkIDPoliciesBaseNamesBaseNameDefault {
	return &DeleteNetworksNetworkIDPoliciesBaseNamesBaseNameDefault{
		_statusCode: code,
	}
}

/* DeleteNetworksNetworkIDPoliciesBaseNamesBaseNameDefault describes a response with status code -1, with default header values.

Unexpected Error
*/
type DeleteNetworksNetworkIDPoliciesBaseNamesBaseNameDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the delete networks network ID policies base names base name default response
func (o *DeleteNetworksNetworkIDPoliciesBaseNamesBaseNameDefault) Code() int {
	return o._statusCode
}

func (o *DeleteNetworksNetworkIDPoliciesBaseNamesBaseNameDefault) Error() string {
	return fmt.Sprintf("[DELETE /networks/{network_id}/policies/base_names/{base_name}][%d] DeleteNetworksNetworkIDPoliciesBaseNamesBaseName default  %+v", o._statusCode, o.Payload)
}
func (o *DeleteNetworksNetworkIDPoliciesBaseNamesBaseNameDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteNetworksNetworkIDPoliciesBaseNamesBaseNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
