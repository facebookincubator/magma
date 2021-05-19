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

// PostNetworksNetworkIDPrometheusAlertReceiverRouteReader is a Reader for the PostNetworksNetworkIDPrometheusAlertReceiverRoute structure.
type PostNetworksNetworkIDPrometheusAlertReceiverRouteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostNetworksNetworkIDPrometheusAlertReceiverRouteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostNetworksNetworkIDPrometheusAlertReceiverRouteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostNetworksNetworkIDPrometheusAlertReceiverRouteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostNetworksNetworkIDPrometheusAlertReceiverRouteOK creates a PostNetworksNetworkIDPrometheusAlertReceiverRouteOK with default headers values
func NewPostNetworksNetworkIDPrometheusAlertReceiverRouteOK() *PostNetworksNetworkIDPrometheusAlertReceiverRouteOK {
	return &PostNetworksNetworkIDPrometheusAlertReceiverRouteOK{}
}

/*PostNetworksNetworkIDPrometheusAlertReceiverRouteOK handles this case with default header values.

OK
*/
type PostNetworksNetworkIDPrometheusAlertReceiverRouteOK struct {
}

func (o *PostNetworksNetworkIDPrometheusAlertReceiverRouteOK) Error() string {
	return fmt.Sprintf("[POST /networks/{network_id}/prometheus/alert_receiver/route][%d] postNetworksNetworkIdPrometheusAlertReceiverRouteOK ", 200)
}

func (o *PostNetworksNetworkIDPrometheusAlertReceiverRouteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostNetworksNetworkIDPrometheusAlertReceiverRouteDefault creates a PostNetworksNetworkIDPrometheusAlertReceiverRouteDefault with default headers values
func NewPostNetworksNetworkIDPrometheusAlertReceiverRouteDefault(code int) *PostNetworksNetworkIDPrometheusAlertReceiverRouteDefault {
	return &PostNetworksNetworkIDPrometheusAlertReceiverRouteDefault{
		_statusCode: code,
	}
}

/*PostNetworksNetworkIDPrometheusAlertReceiverRouteDefault handles this case with default header values.

Unexpected Error
*/
type PostNetworksNetworkIDPrometheusAlertReceiverRouteDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the post networks network ID prometheus alert receiver route default response
func (o *PostNetworksNetworkIDPrometheusAlertReceiverRouteDefault) Code() int {
	return o._statusCode
}

func (o *PostNetworksNetworkIDPrometheusAlertReceiverRouteDefault) Error() string {
	return fmt.Sprintf("[POST /networks/{network_id}/prometheus/alert_receiver/route][%d] PostNetworksNetworkIDPrometheusAlertReceiverRoute default  %+v", o._statusCode, o.Payload)
}

func (o *PostNetworksNetworkIDPrometheusAlertReceiverRouteDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostNetworksNetworkIDPrometheusAlertReceiverRouteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
