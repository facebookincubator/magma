// Code generated by go-swagger; DO NOT EDIT.

package wifi_meshes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"magma/orc8r/cloud/go/obsidian/swagger/v1/models"
)

// GetWifiNetworkIDMeshesMeshIDNameReader is a Reader for the GetWifiNetworkIDMeshesMeshIDName structure.
type GetWifiNetworkIDMeshesMeshIDNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWifiNetworkIDMeshesMeshIDNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWifiNetworkIDMeshesMeshIDNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetWifiNetworkIDMeshesMeshIDNameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetWifiNetworkIDMeshesMeshIDNameOK creates a GetWifiNetworkIDMeshesMeshIDNameOK with default headers values
func NewGetWifiNetworkIDMeshesMeshIDNameOK() *GetWifiNetworkIDMeshesMeshIDNameOK {
	return &GetWifiNetworkIDMeshesMeshIDNameOK{}
}

/* GetWifiNetworkIDMeshesMeshIDNameOK describes a response with status code 200, with default header values.

The name of the mesh
*/
type GetWifiNetworkIDMeshesMeshIDNameOK struct {
	Payload models.MeshName
}

func (o *GetWifiNetworkIDMeshesMeshIDNameOK) Error() string {
	return fmt.Sprintf("[GET /wifi/{network_id}/meshes/{mesh_id}/name][%d] getWifiNetworkIdMeshesMeshIdNameOK  %+v", 200, o.Payload)
}
func (o *GetWifiNetworkIDMeshesMeshIDNameOK) GetPayload() models.MeshName {
	return o.Payload
}

func (o *GetWifiNetworkIDMeshesMeshIDNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWifiNetworkIDMeshesMeshIDNameDefault creates a GetWifiNetworkIDMeshesMeshIDNameDefault with default headers values
func NewGetWifiNetworkIDMeshesMeshIDNameDefault(code int) *GetWifiNetworkIDMeshesMeshIDNameDefault {
	return &GetWifiNetworkIDMeshesMeshIDNameDefault{
		_statusCode: code,
	}
}

/* GetWifiNetworkIDMeshesMeshIDNameDefault describes a response with status code -1, with default header values.

Unexpected Error
*/
type GetWifiNetworkIDMeshesMeshIDNameDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get wifi network ID meshes mesh ID name default response
func (o *GetWifiNetworkIDMeshesMeshIDNameDefault) Code() int {
	return o._statusCode
}

func (o *GetWifiNetworkIDMeshesMeshIDNameDefault) Error() string {
	return fmt.Sprintf("[GET /wifi/{network_id}/meshes/{mesh_id}/name][%d] GetWifiNetworkIDMeshesMeshIDName default  %+v", o._statusCode, o.Payload)
}
func (o *GetWifiNetworkIDMeshesMeshIDNameDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetWifiNetworkIDMeshesMeshIDNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
