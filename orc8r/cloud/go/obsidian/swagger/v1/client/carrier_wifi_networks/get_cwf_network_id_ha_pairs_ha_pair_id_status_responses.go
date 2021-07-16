// Code generated by go-swagger; DO NOT EDIT.

package carrier_wifi_networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"magma/orc8r/cloud/go/obsidian/swagger/v1/models"
)

// GetCwfNetworkIDHaPairsHaPairIDStatusReader is a Reader for the GetCwfNetworkIDHaPairsHaPairIDStatus structure.
type GetCwfNetworkIDHaPairsHaPairIDStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCwfNetworkIDHaPairsHaPairIDStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCwfNetworkIDHaPairsHaPairIDStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetCwfNetworkIDHaPairsHaPairIDStatusDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetCwfNetworkIDHaPairsHaPairIDStatusOK creates a GetCwfNetworkIDHaPairsHaPairIDStatusOK with default headers values
func NewGetCwfNetworkIDHaPairsHaPairIDStatusOK() *GetCwfNetworkIDHaPairsHaPairIDStatusOK {
	return &GetCwfNetworkIDHaPairsHaPairIDStatusOK{}
}

/* GetCwfNetworkIDHaPairsHaPairIDStatusOK describes a response with status code 200, with default header values.

Status of a HA pair
*/
type GetCwfNetworkIDHaPairsHaPairIDStatusOK struct {
	Payload *models.CarrierWifiHaPairStatus
}

func (o *GetCwfNetworkIDHaPairsHaPairIDStatusOK) Error() string {
	return fmt.Sprintf("[GET /cwf/{network_id}/ha_pairs/{ha_pair_id}/status][%d] getCwfNetworkIdHaPairsHaPairIdStatusOK  %+v", 200, o.Payload)
}
func (o *GetCwfNetworkIDHaPairsHaPairIDStatusOK) GetPayload() *models.CarrierWifiHaPairStatus {
	return o.Payload
}

func (o *GetCwfNetworkIDHaPairsHaPairIDStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CarrierWifiHaPairStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCwfNetworkIDHaPairsHaPairIDStatusDefault creates a GetCwfNetworkIDHaPairsHaPairIDStatusDefault with default headers values
func NewGetCwfNetworkIDHaPairsHaPairIDStatusDefault(code int) *GetCwfNetworkIDHaPairsHaPairIDStatusDefault {
	return &GetCwfNetworkIDHaPairsHaPairIDStatusDefault{
		_statusCode: code,
	}
}

/* GetCwfNetworkIDHaPairsHaPairIDStatusDefault describes a response with status code -1, with default header values.

Unexpected Error
*/
type GetCwfNetworkIDHaPairsHaPairIDStatusDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get cwf network ID ha pairs ha pair ID status default response
func (o *GetCwfNetworkIDHaPairsHaPairIDStatusDefault) Code() int {
	return o._statusCode
}

func (o *GetCwfNetworkIDHaPairsHaPairIDStatusDefault) Error() string {
	return fmt.Sprintf("[GET /cwf/{network_id}/ha_pairs/{ha_pair_id}/status][%d] GetCwfNetworkIDHaPairsHaPairIDStatus default  %+v", o._statusCode, o.Payload)
}
func (o *GetCwfNetworkIDHaPairsHaPairIDStatusDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetCwfNetworkIDHaPairsHaPairIDStatusDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
