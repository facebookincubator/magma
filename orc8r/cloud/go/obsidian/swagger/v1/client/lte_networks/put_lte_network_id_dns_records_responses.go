// Code generated by go-swagger; DO NOT EDIT.

package lte_networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"magma/orc8r/cloud/go/obsidian/swagger/v1/models"
)

// PutLTENetworkIDDNSRecordsReader is a Reader for the PutLTENetworkIDDNSRecords structure.
type PutLTENetworkIDDNSRecordsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutLTENetworkIDDNSRecordsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPutLTENetworkIDDNSRecordsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPutLTENetworkIDDNSRecordsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutLTENetworkIDDNSRecordsNoContent creates a PutLTENetworkIDDNSRecordsNoContent with default headers values
func NewPutLTENetworkIDDNSRecordsNoContent() *PutLTENetworkIDDNSRecordsNoContent {
	return &PutLTENetworkIDDNSRecordsNoContent{}
}

/* PutLTENetworkIDDNSRecordsNoContent describes a response with status code 204, with default header values.

Success
*/
type PutLTENetworkIDDNSRecordsNoContent struct {
}

func (o *PutLTENetworkIDDNSRecordsNoContent) Error() string {
	return fmt.Sprintf("[PUT /lte/{network_id}/dns/records][%d] putLteNetworkIdDnsRecordsNoContent ", 204)
}

func (o *PutLTENetworkIDDNSRecordsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutLTENetworkIDDNSRecordsDefault creates a PutLTENetworkIDDNSRecordsDefault with default headers values
func NewPutLTENetworkIDDNSRecordsDefault(code int) *PutLTENetworkIDDNSRecordsDefault {
	return &PutLTENetworkIDDNSRecordsDefault{
		_statusCode: code,
	}
}

/* PutLTENetworkIDDNSRecordsDefault describes a response with status code -1, with default header values.

Unexpected Error
*/
type PutLTENetworkIDDNSRecordsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the put LTE network ID DNS records default response
func (o *PutLTENetworkIDDNSRecordsDefault) Code() int {
	return o._statusCode
}

func (o *PutLTENetworkIDDNSRecordsDefault) Error() string {
	return fmt.Sprintf("[PUT /lte/{network_id}/dns/records][%d] PutLTENetworkIDDNSRecords default  %+v", o._statusCode, o.Payload)
}
func (o *PutLTENetworkIDDNSRecordsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutLTENetworkIDDNSRecordsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
