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

// GetFegLTENetworkIDSubscriberConfigBaseNamesReader is a Reader for the GetFegLTENetworkIDSubscriberConfigBaseNames structure.
type GetFegLTENetworkIDSubscriberConfigBaseNamesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFegLTENetworkIDSubscriberConfigBaseNamesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFegLTENetworkIDSubscriberConfigBaseNamesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetFegLTENetworkIDSubscriberConfigBaseNamesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetFegLTENetworkIDSubscriberConfigBaseNamesOK creates a GetFegLTENetworkIDSubscriberConfigBaseNamesOK with default headers values
func NewGetFegLTENetworkIDSubscriberConfigBaseNamesOK() *GetFegLTENetworkIDSubscriberConfigBaseNamesOK {
	return &GetFegLTENetworkIDSubscriberConfigBaseNamesOK{}
}

/* GetFegLTENetworkIDSubscriberConfigBaseNamesOK describes a response with status code 200, with default header values.

Subscriber Config
*/
type GetFegLTENetworkIDSubscriberConfigBaseNamesOK struct {
	Payload models.BaseNames
}

func (o *GetFegLTENetworkIDSubscriberConfigBaseNamesOK) Error() string {
	return fmt.Sprintf("[GET /feg_lte/{network_id}/subscriber_config/base_names][%d] getFegLteNetworkIdSubscriberConfigBaseNamesOK  %+v", 200, o.Payload)
}
func (o *GetFegLTENetworkIDSubscriberConfigBaseNamesOK) GetPayload() models.BaseNames {
	return o.Payload
}

func (o *GetFegLTENetworkIDSubscriberConfigBaseNamesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFegLTENetworkIDSubscriberConfigBaseNamesDefault creates a GetFegLTENetworkIDSubscriberConfigBaseNamesDefault with default headers values
func NewGetFegLTENetworkIDSubscriberConfigBaseNamesDefault(code int) *GetFegLTENetworkIDSubscriberConfigBaseNamesDefault {
	return &GetFegLTENetworkIDSubscriberConfigBaseNamesDefault{
		_statusCode: code,
	}
}

/* GetFegLTENetworkIDSubscriberConfigBaseNamesDefault describes a response with status code -1, with default header values.

Unexpected Error
*/
type GetFegLTENetworkIDSubscriberConfigBaseNamesDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get feg LTE network ID subscriber config base names default response
func (o *GetFegLTENetworkIDSubscriberConfigBaseNamesDefault) Code() int {
	return o._statusCode
}

func (o *GetFegLTENetworkIDSubscriberConfigBaseNamesDefault) Error() string {
	return fmt.Sprintf("[GET /feg_lte/{network_id}/subscriber_config/base_names][%d] GetFegLTENetworkIDSubscriberConfigBaseNames default  %+v", o._statusCode, o.Payload)
}
func (o *GetFegLTENetworkIDSubscriberConfigBaseNamesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetFegLTENetworkIDSubscriberConfigBaseNamesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
