// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openclarity/apiclarity/api/server/models"
)

// GetControlGatewaysOKCode is the HTTP code returned for type GetControlGatewaysOK
const GetControlGatewaysOKCode int = 200

/*GetControlGatewaysOK Success

swagger:response getControlGatewaysOK
*/
type GetControlGatewaysOK struct {

	/*
	  In: Body
	*/
	Payload *GetControlGatewaysOKBody `json:"body,omitempty"`
}

// NewGetControlGatewaysOK creates GetControlGatewaysOK with default headers values
func NewGetControlGatewaysOK() *GetControlGatewaysOK {

	return &GetControlGatewaysOK{}
}

// WithPayload adds the payload to the get control gateways o k response
func (o *GetControlGatewaysOK) WithPayload(payload *GetControlGatewaysOKBody) *GetControlGatewaysOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get control gateways o k response
func (o *GetControlGatewaysOK) SetPayload(payload *GetControlGatewaysOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetControlGatewaysOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetControlGatewaysDefault unknown error

swagger:response getControlGatewaysDefault
*/
type GetControlGatewaysDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.APIResponse `json:"body,omitempty"`
}

// NewGetControlGatewaysDefault creates GetControlGatewaysDefault with default headers values
func NewGetControlGatewaysDefault(code int) *GetControlGatewaysDefault {
	if code <= 0 {
		code = 500
	}

	return &GetControlGatewaysDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get control gateways default response
func (o *GetControlGatewaysDefault) WithStatusCode(code int) *GetControlGatewaysDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get control gateways default response
func (o *GetControlGatewaysDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get control gateways default response
func (o *GetControlGatewaysDefault) WithPayload(payload *models.APIResponse) *GetControlGatewaysDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get control gateways default response
func (o *GetControlGatewaysDefault) SetPayload(payload *models.APIResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetControlGatewaysDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
