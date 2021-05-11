// Code generated by go-swagger; DO NOT EDIT.

package carrier_wifi_networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "magma/orc8r/cloud/go/obsidian/swagger/v1/models"
)

// PostCwfNetworkIDHaPairsReader is a Reader for the PostCwfNetworkIDHaPairs structure.
type PostCwfNetworkIDHaPairsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostCwfNetworkIDHaPairsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostCwfNetworkIDHaPairsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostCwfNetworkIDHaPairsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostCwfNetworkIDHaPairsCreated creates a PostCwfNetworkIDHaPairsCreated with default headers values
func NewPostCwfNetworkIDHaPairsCreated() *PostCwfNetworkIDHaPairsCreated {
	return &PostCwfNetworkIDHaPairsCreated{}
}

/*PostCwfNetworkIDHaPairsCreated handles this case with default header values.

Success
*/
type PostCwfNetworkIDHaPairsCreated struct {
}

func (o *PostCwfNetworkIDHaPairsCreated) Error() string {
	return fmt.Sprintf("[POST /cwf/{network_id}/ha_pairs][%d] postCwfNetworkIdHaPairsCreated ", 201)
}

func (o *PostCwfNetworkIDHaPairsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostCwfNetworkIDHaPairsDefault creates a PostCwfNetworkIDHaPairsDefault with default headers values
func NewPostCwfNetworkIDHaPairsDefault(code int) *PostCwfNetworkIDHaPairsDefault {
	return &PostCwfNetworkIDHaPairsDefault{
		_statusCode: code,
	}
}

/*PostCwfNetworkIDHaPairsDefault handles this case with default header values.

Unexpected Error
*/
type PostCwfNetworkIDHaPairsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the post cwf network ID ha pairs default response
func (o *PostCwfNetworkIDHaPairsDefault) Code() int {
	return o._statusCode
}

func (o *PostCwfNetworkIDHaPairsDefault) Error() string {
	return fmt.Sprintf("[POST /cwf/{network_id}/ha_pairs][%d] PostCwfNetworkIDHaPairs default  %+v", o._statusCode, o.Payload)
}

func (o *PostCwfNetworkIDHaPairsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostCwfNetworkIDHaPairsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
