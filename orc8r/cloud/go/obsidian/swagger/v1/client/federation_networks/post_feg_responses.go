// Code generated by go-swagger; DO NOT EDIT.

package federation_networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "magma/orc8r/cloud/go/obsidian/swagger/v1/models"
)

// PostFegReader is a Reader for the PostFeg structure.
type PostFegReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostFegReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostFegCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostFegDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostFegCreated creates a PostFegCreated with default headers values
func NewPostFegCreated() *PostFegCreated {
	return &PostFegCreated{}
}

/*PostFegCreated handles this case with default header values.

Success
*/
type PostFegCreated struct {
}

func (o *PostFegCreated) Error() string {
	return fmt.Sprintf("[POST /feg][%d] postFegCreated ", 201)
}

func (o *PostFegCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostFegDefault creates a PostFegDefault with default headers values
func NewPostFegDefault(code int) *PostFegDefault {
	return &PostFegDefault{
		_statusCode: code,
	}
}

/*PostFegDefault handles this case with default header values.

Unexpected Error
*/
type PostFegDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the post feg default response
func (o *PostFegDefault) Code() int {
	return o._statusCode
}

func (o *PostFegDefault) Error() string {
	return fmt.Sprintf("[POST /feg][%d] PostFeg default  %+v", o._statusCode, o.Payload)
}

func (o *PostFegDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostFegDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
