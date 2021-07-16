// Code generated by go-swagger; DO NOT EDIT.

package network_probes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"magma/orc8r/cloud/go/obsidian/swagger/v1/models"
)

// GetLTENetworkIDNetworkProbeTasksReader is a Reader for the GetLTENetworkIDNetworkProbeTasks structure.
type GetLTENetworkIDNetworkProbeTasksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLTENetworkIDNetworkProbeTasksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLTENetworkIDNetworkProbeTasksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetLTENetworkIDNetworkProbeTasksDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetLTENetworkIDNetworkProbeTasksOK creates a GetLTENetworkIDNetworkProbeTasksOK with default headers values
func NewGetLTENetworkIDNetworkProbeTasksOK() *GetLTENetworkIDNetworkProbeTasksOK {
	return &GetLTENetworkIDNetworkProbeTasksOK{}
}

/* GetLTENetworkIDNetworkProbeTasksOK describes a response with status code 200, with default header values.

Provisioned NetworkProbeTasks
*/
type GetLTENetworkIDNetworkProbeTasksOK struct {
	Payload *models.NetworkProbeTask
}

func (o *GetLTENetworkIDNetworkProbeTasksOK) Error() string {
	return fmt.Sprintf("[GET /lte/{network_id}/network_probe/tasks][%d] getLteNetworkIdNetworkProbeTasksOK  %+v", 200, o.Payload)
}
func (o *GetLTENetworkIDNetworkProbeTasksOK) GetPayload() *models.NetworkProbeTask {
	return o.Payload
}

func (o *GetLTENetworkIDNetworkProbeTasksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NetworkProbeTask)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLTENetworkIDNetworkProbeTasksDefault creates a GetLTENetworkIDNetworkProbeTasksDefault with default headers values
func NewGetLTENetworkIDNetworkProbeTasksDefault(code int) *GetLTENetworkIDNetworkProbeTasksDefault {
	return &GetLTENetworkIDNetworkProbeTasksDefault{
		_statusCode: code,
	}
}

/* GetLTENetworkIDNetworkProbeTasksDefault describes a response with status code -1, with default header values.

Unexpected Error
*/
type GetLTENetworkIDNetworkProbeTasksDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get LTE network ID network probe tasks default response
func (o *GetLTENetworkIDNetworkProbeTasksDefault) Code() int {
	return o._statusCode
}

func (o *GetLTENetworkIDNetworkProbeTasksDefault) Error() string {
	return fmt.Sprintf("[GET /lte/{network_id}/network_probe/tasks][%d] GetLTENetworkIDNetworkProbeTasks default  %+v", o._statusCode, o.Payload)
}
func (o *GetLTENetworkIDNetworkProbeTasksDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetLTENetworkIDNetworkProbeTasksDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
