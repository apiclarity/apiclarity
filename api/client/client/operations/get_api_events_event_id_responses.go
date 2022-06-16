// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openclarity/apiclarity/api/client/models"
)

// GetAPIEventsEventIDReader is a Reader for the GetAPIEventsEventID structure.
type GetAPIEventsEventIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPIEventsEventIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAPIEventsEventIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetAPIEventsEventIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetAPIEventsEventIDOK creates a GetAPIEventsEventIDOK with default headers values
func NewGetAPIEventsEventIDOK() *GetAPIEventsEventIDOK {
	return &GetAPIEventsEventIDOK{}
}

/* GetAPIEventsEventIDOK describes a response with status code 200, with default header values.

Success
*/
type GetAPIEventsEventIDOK struct {
	Payload *models.APIEvent
}

func (o *GetAPIEventsEventIDOK) Error() string {
	return fmt.Sprintf("[GET /apiEvents/{eventId}][%d] getApiEventsEventIdOK  %+v", 200, o.Payload)
}
func (o *GetAPIEventsEventIDOK) GetPayload() *models.APIEvent {
	return o.Payload
}

func (o *GetAPIEventsEventIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIEvent)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIEventsEventIDDefault creates a GetAPIEventsEventIDDefault with default headers values
func NewGetAPIEventsEventIDDefault(code int) *GetAPIEventsEventIDDefault {
	return &GetAPIEventsEventIDDefault{
		_statusCode: code,
	}
}

/* GetAPIEventsEventIDDefault describes a response with status code -1, with default header values.

unknown error
*/
type GetAPIEventsEventIDDefault struct {
	_statusCode int

	Payload *models.APIResponse
}

// Code gets the status code for the get API events event ID default response
func (o *GetAPIEventsEventIDDefault) Code() int {
	return o._statusCode
}

func (o *GetAPIEventsEventIDDefault) Error() string {
	return fmt.Sprintf("[GET /apiEvents/{eventId}][%d] GetAPIEventsEventID default  %+v", o._statusCode, o.Payload)
}
func (o *GetAPIEventsEventIDDefault) GetPayload() *models.APIResponse {
	return o.Payload
}

func (o *GetAPIEventsEventIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
