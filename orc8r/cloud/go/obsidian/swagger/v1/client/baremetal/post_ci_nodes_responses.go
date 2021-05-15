// Code generated by go-swagger; DO NOT EDIT.

package baremetal

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "magma/orc8r/cloud/go/obsidian/swagger/v1/models"
)

// PostCiNodesReader is a Reader for the PostCiNodes structure.
type PostCiNodesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostCiNodesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostCiNodesCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostCiNodesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostCiNodesCreated creates a PostCiNodesCreated with default headers values
func NewPostCiNodesCreated() *PostCiNodesCreated {
	return &PostCiNodesCreated{}
}

/*PostCiNodesCreated handles this case with default header values.

Created
*/
type PostCiNodesCreated struct {
}

func (o *PostCiNodesCreated) Error() string {
	return fmt.Sprintf("[POST /ci/nodes][%d] postCiNodesCreated ", 201)
}

func (o *PostCiNodesCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostCiNodesDefault creates a PostCiNodesDefault with default headers values
func NewPostCiNodesDefault(code int) *PostCiNodesDefault {
	return &PostCiNodesDefault{
		_statusCode: code,
	}
}

/*PostCiNodesDefault handles this case with default header values.

Unexpected Error
*/
type PostCiNodesDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the post ci nodes default response
func (o *PostCiNodesDefault) Code() int {
	return o._statusCode
}

func (o *PostCiNodesDefault) Error() string {
	return fmt.Sprintf("[POST /ci/nodes][%d] PostCiNodes default  %+v", o._statusCode, o.Payload)
}

func (o *PostCiNodesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostCiNodesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
