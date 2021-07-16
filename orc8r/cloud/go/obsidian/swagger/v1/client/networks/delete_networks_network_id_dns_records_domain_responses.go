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

// DeleteNetworksNetworkIDDNSRecordsDomainReader is a Reader for the DeleteNetworksNetworkIDDNSRecordsDomain structure.
type DeleteNetworksNetworkIDDNSRecordsDomainReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteNetworksNetworkIDDNSRecordsDomainReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteNetworksNetworkIDDNSRecordsDomainNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteNetworksNetworkIDDNSRecordsDomainDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteNetworksNetworkIDDNSRecordsDomainNoContent creates a DeleteNetworksNetworkIDDNSRecordsDomainNoContent with default headers values
func NewDeleteNetworksNetworkIDDNSRecordsDomainNoContent() *DeleteNetworksNetworkIDDNSRecordsDomainNoContent {
	return &DeleteNetworksNetworkIDDNSRecordsDomainNoContent{}
}

/* DeleteNetworksNetworkIDDNSRecordsDomainNoContent describes a response with status code 204, with default header values.

Success
*/
type DeleteNetworksNetworkIDDNSRecordsDomainNoContent struct {
}

func (o *DeleteNetworksNetworkIDDNSRecordsDomainNoContent) Error() string {
	return fmt.Sprintf("[DELETE /networks/{network_id}/dns/records/{domain}][%d] deleteNetworksNetworkIdDnsRecordsDomainNoContent ", 204)
}

func (o *DeleteNetworksNetworkIDDNSRecordsDomainNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteNetworksNetworkIDDNSRecordsDomainDefault creates a DeleteNetworksNetworkIDDNSRecordsDomainDefault with default headers values
func NewDeleteNetworksNetworkIDDNSRecordsDomainDefault(code int) *DeleteNetworksNetworkIDDNSRecordsDomainDefault {
	return &DeleteNetworksNetworkIDDNSRecordsDomainDefault{
		_statusCode: code,
	}
}

/* DeleteNetworksNetworkIDDNSRecordsDomainDefault describes a response with status code -1, with default header values.

Unexpected Error
*/
type DeleteNetworksNetworkIDDNSRecordsDomainDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the delete networks network ID DNS records domain default response
func (o *DeleteNetworksNetworkIDDNSRecordsDomainDefault) Code() int {
	return o._statusCode
}

func (o *DeleteNetworksNetworkIDDNSRecordsDomainDefault) Error() string {
	return fmt.Sprintf("[DELETE /networks/{network_id}/dns/records/{domain}][%d] DeleteNetworksNetworkIDDNSRecordsDomain default  %+v", o._statusCode, o.Payload)
}
func (o *DeleteNetworksNetworkIDDNSRecordsDomainDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteNetworksNetworkIDDNSRecordsDomainDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
