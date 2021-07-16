// Code generated by go-swagger; DO NOT EDIT.

package federated_lte_networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"magma/orc8r/cloud/go/obsidian/swagger/v1/models"
)

// PutFegLTENetworkIDFederationReader is a Reader for the PutFegLTENetworkIDFederation structure.
type PutFegLTENetworkIDFederationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutFegLTENetworkIDFederationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutFegLTENetworkIDFederationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPutFegLTENetworkIDFederationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutFegLTENetworkIDFederationOK creates a PutFegLTENetworkIDFederationOK with default headers values
func NewPutFegLTENetworkIDFederationOK() *PutFegLTENetworkIDFederationOK {
	return &PutFegLTENetworkIDFederationOK{}
}

/* PutFegLTENetworkIDFederationOK describes a response with status code 200, with default header values.

Success
*/
type PutFegLTENetworkIDFederationOK struct {
}

func (o *PutFegLTENetworkIDFederationOK) Error() string {
	return fmt.Sprintf("[PUT /feg_lte/{network_id}/federation][%d] putFegLteNetworkIdFederationOK ", 200)
}

func (o *PutFegLTENetworkIDFederationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutFegLTENetworkIDFederationDefault creates a PutFegLTENetworkIDFederationDefault with default headers values
func NewPutFegLTENetworkIDFederationDefault(code int) *PutFegLTENetworkIDFederationDefault {
	return &PutFegLTENetworkIDFederationDefault{
		_statusCode: code,
	}
}

/* PutFegLTENetworkIDFederationDefault describes a response with status code -1, with default header values.

Unexpected Error
*/
type PutFegLTENetworkIDFederationDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the put feg LTE network ID federation default response
func (o *PutFegLTENetworkIDFederationDefault) Code() int {
	return o._statusCode
}

func (o *PutFegLTENetworkIDFederationDefault) Error() string {
	return fmt.Sprintf("[PUT /feg_lte/{network_id}/federation][%d] PutFegLTENetworkIDFederation default  %+v", o._statusCode, o.Payload)
}
func (o *PutFegLTENetworkIDFederationDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutFegLTENetworkIDFederationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
