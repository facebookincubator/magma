// Code generated by go-swagger; DO NOT EDIT.

package alerts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "magma/orc8r/cloud/go/obsidian/swagger/v1/models"
)

// GetNetworksNetworkIDPrometheusAlertConfigReader is a Reader for the GetNetworksNetworkIDPrometheusAlertConfig structure.
type GetNetworksNetworkIDPrometheusAlertConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworksNetworkIDPrometheusAlertConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworksNetworkIDPrometheusAlertConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetNetworksNetworkIDPrometheusAlertConfigDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetNetworksNetworkIDPrometheusAlertConfigOK creates a GetNetworksNetworkIDPrometheusAlertConfigOK with default headers values
func NewGetNetworksNetworkIDPrometheusAlertConfigOK() *GetNetworksNetworkIDPrometheusAlertConfigOK {
	return &GetNetworksNetworkIDPrometheusAlertConfigOK{}
}

/*GetNetworksNetworkIDPrometheusAlertConfigOK handles this case with default header values.

List of alert configurations
*/
type GetNetworksNetworkIDPrometheusAlertConfigOK struct {
	Payload models.PromAlertConfigList
}

func (o *GetNetworksNetworkIDPrometheusAlertConfigOK) Error() string {
	return fmt.Sprintf("[GET /networks/{network_id}/prometheus/alert_config][%d] getNetworksNetworkIdPrometheusAlertConfigOK  %+v", 200, o.Payload)
}

func (o *GetNetworksNetworkIDPrometheusAlertConfigOK) GetPayload() models.PromAlertConfigList {
	return o.Payload
}

func (o *GetNetworksNetworkIDPrometheusAlertConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNetworksNetworkIDPrometheusAlertConfigDefault creates a GetNetworksNetworkIDPrometheusAlertConfigDefault with default headers values
func NewGetNetworksNetworkIDPrometheusAlertConfigDefault(code int) *GetNetworksNetworkIDPrometheusAlertConfigDefault {
	return &GetNetworksNetworkIDPrometheusAlertConfigDefault{
		_statusCode: code,
	}
}

/*GetNetworksNetworkIDPrometheusAlertConfigDefault handles this case with default header values.

Unexpected Error
*/
type GetNetworksNetworkIDPrometheusAlertConfigDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get networks network ID prometheus alert config default response
func (o *GetNetworksNetworkIDPrometheusAlertConfigDefault) Code() int {
	return o._statusCode
}

func (o *GetNetworksNetworkIDPrometheusAlertConfigDefault) Error() string {
	return fmt.Sprintf("[GET /networks/{network_id}/prometheus/alert_config][%d] GetNetworksNetworkIDPrometheusAlertConfig default  %+v", o._statusCode, o.Payload)
}

func (o *GetNetworksNetworkIDPrometheusAlertConfigDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetNetworksNetworkIDPrometheusAlertConfigDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
