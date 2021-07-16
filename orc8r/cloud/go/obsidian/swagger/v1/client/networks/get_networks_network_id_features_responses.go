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

// GetNetworksNetworkIDFeaturesReader is a Reader for the GetNetworksNetworkIDFeatures structure.
type GetNetworksNetworkIDFeaturesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworksNetworkIDFeaturesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworksNetworkIDFeaturesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetNetworksNetworkIDFeaturesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetNetworksNetworkIDFeaturesOK creates a GetNetworksNetworkIDFeaturesOK with default headers values
func NewGetNetworksNetworkIDFeaturesOK() *GetNetworksNetworkIDFeaturesOK {
	return &GetNetworksNetworkIDFeaturesOK{}
}

/* GetNetworksNetworkIDFeaturesOK describes a response with status code 200, with default header values.

Feature flags of the network
*/
type GetNetworksNetworkIDFeaturesOK struct {
	Payload *models.NetworkFeatures
}

func (o *GetNetworksNetworkIDFeaturesOK) Error() string {
	return fmt.Sprintf("[GET /networks/{network_id}/features][%d] getNetworksNetworkIdFeaturesOK  %+v", 200, o.Payload)
}
func (o *GetNetworksNetworkIDFeaturesOK) GetPayload() *models.NetworkFeatures {
	return o.Payload
}

func (o *GetNetworksNetworkIDFeaturesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NetworkFeatures)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNetworksNetworkIDFeaturesDefault creates a GetNetworksNetworkIDFeaturesDefault with default headers values
func NewGetNetworksNetworkIDFeaturesDefault(code int) *GetNetworksNetworkIDFeaturesDefault {
	return &GetNetworksNetworkIDFeaturesDefault{
		_statusCode: code,
	}
}

/* GetNetworksNetworkIDFeaturesDefault describes a response with status code -1, with default header values.

Unexpected Error
*/
type GetNetworksNetworkIDFeaturesDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get networks network ID features default response
func (o *GetNetworksNetworkIDFeaturesDefault) Code() int {
	return o._statusCode
}

func (o *GetNetworksNetworkIDFeaturesDefault) Error() string {
	return fmt.Sprintf("[GET /networks/{network_id}/features][%d] GetNetworksNetworkIDFeatures default  %+v", o._statusCode, o.Payload)
}
func (o *GetNetworksNetworkIDFeaturesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetNetworksNetworkIDFeaturesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
