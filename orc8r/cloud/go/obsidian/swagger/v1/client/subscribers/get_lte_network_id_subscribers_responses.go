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

// GetLTENetworkIDSubscribersReader is a Reader for the GetLTENetworkIDSubscribers structure.
type GetLTENetworkIDSubscribersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLTENetworkIDSubscribersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLTENetworkIDSubscribersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetLTENetworkIDSubscribersDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetLTENetworkIDSubscribersOK creates a GetLTENetworkIDSubscribersOK with default headers values
func NewGetLTENetworkIDSubscribersOK() *GetLTENetworkIDSubscribersOK {
	return &GetLTENetworkIDSubscribersOK{}
}

/* GetLTENetworkIDSubscribersOK describes a response with status code 200, with default header values.

List of all the subscribers in the network
*/
type GetLTENetworkIDSubscribersOK struct {
	Payload map[string]models.Subscriber
}

func (o *GetLTENetworkIDSubscribersOK) Error() string {
	return fmt.Sprintf("[GET /lte/{network_id}/subscribers][%d] getLteNetworkIdSubscribersOK  %+v", 200, o.Payload)
}
func (o *GetLTENetworkIDSubscribersOK) GetPayload() map[string]models.Subscriber {
	return o.Payload
}

func (o *GetLTENetworkIDSubscribersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLTENetworkIDSubscribersDefault creates a GetLTENetworkIDSubscribersDefault with default headers values
func NewGetLTENetworkIDSubscribersDefault(code int) *GetLTENetworkIDSubscribersDefault {
	return &GetLTENetworkIDSubscribersDefault{
		_statusCode: code,
	}
}

/* GetLTENetworkIDSubscribersDefault describes a response with status code -1, with default header values.

Unexpected Error
*/
type GetLTENetworkIDSubscribersDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get LTE network ID subscribers default response
func (o *GetLTENetworkIDSubscribersDefault) Code() int {
	return o._statusCode
}

func (o *GetLTENetworkIDSubscribersDefault) Error() string {
	return fmt.Sprintf("[GET /lte/{network_id}/subscribers][%d] GetLTENetworkIDSubscribers default  %+v", o._statusCode, o.Payload)
}
func (o *GetLTENetworkIDSubscribersDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetLTENetworkIDSubscribersDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
