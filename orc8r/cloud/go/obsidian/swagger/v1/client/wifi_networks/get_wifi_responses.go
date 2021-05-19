// Code generated by go-swagger; DO NOT EDIT.

package wifi_networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "magma/orc8r/cloud/go/obsidian/swagger/v1/models"
)

// GetWifiReader is a Reader for the GetWifi structure.
type GetWifiReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWifiReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWifiOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetWifiDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetWifiOK creates a GetWifiOK with default headers values
func NewGetWifiOK() *GetWifiOK {
	return &GetWifiOK{}
}

/*GetWifiOK handles this case with default header values.

List of Wifi network IDs
*/
type GetWifiOK struct {
	Payload []string
}

func (o *GetWifiOK) Error() string {
	return fmt.Sprintf("[GET /wifi][%d] getWifiOK  %+v", 200, o.Payload)
}

func (o *GetWifiOK) GetPayload() []string {
	return o.Payload
}

func (o *GetWifiOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWifiDefault creates a GetWifiDefault with default headers values
func NewGetWifiDefault(code int) *GetWifiDefault {
	return &GetWifiDefault{
		_statusCode: code,
	}
}

/*GetWifiDefault handles this case with default header values.

Unexpected Error
*/
type GetWifiDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get wifi default response
func (o *GetWifiDefault) Code() int {
	return o._statusCode
}

func (o *GetWifiDefault) Error() string {
	return fmt.Sprintf("[GET /wifi][%d] GetWifi default  %+v", o._statusCode, o.Payload)
}

func (o *GetWifiDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetWifiDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
