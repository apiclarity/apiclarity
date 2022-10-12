// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/openclarity/apiclarity/api/server/models"
)

// GetControlGatewaysHandlerFunc turns a function with the right signature into a get control gateways handler
type GetControlGatewaysHandlerFunc func(GetControlGatewaysParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetControlGatewaysHandlerFunc) Handle(params GetControlGatewaysParams) middleware.Responder {
	return fn(params)
}

// GetControlGatewaysHandler interface for that can handle valid get control gateways params
type GetControlGatewaysHandler interface {
	Handle(GetControlGatewaysParams) middleware.Responder
}

// NewGetControlGateways creates a new http.Handler for the get control gateways operation
func NewGetControlGateways(ctx *middleware.Context, handler GetControlGatewaysHandler) *GetControlGateways {
	return &GetControlGateways{Context: ctx, Handler: handler}
}

/* GetControlGateways swagger:route GET /control/gateways getControlGateways

List of configured gateways

*/
type GetControlGateways struct {
	Context *middleware.Context
	Handler GetControlGatewaysHandler
}

func (o *GetControlGateways) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetControlGatewaysParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// GetControlGatewaysOKBody get control gateways o k body
//
// swagger:model GetControlGatewaysOKBody
type GetControlGatewaysOKBody struct {

	// List of gateways
	// Required: true
	Gateways []*models.APIGateway `json:"gateways"`
}

// Validate validates this get control gateways o k body
func (o *GetControlGatewaysOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateGateways(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetControlGatewaysOKBody) validateGateways(formats strfmt.Registry) error {

	if err := validate.Required("getControlGatewaysOK"+"."+"gateways", "body", o.Gateways); err != nil {
		return err
	}

	for i := 0; i < len(o.Gateways); i++ {
		if swag.IsZero(o.Gateways[i]) { // not required
			continue
		}

		if o.Gateways[i] != nil {
			if err := o.Gateways[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getControlGatewaysOK" + "." + "gateways" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get control gateways o k body based on the context it is used
func (o *GetControlGatewaysOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateGateways(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetControlGatewaysOKBody) contextValidateGateways(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Gateways); i++ {

		if o.Gateways[i] != nil {
			if err := o.Gateways[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getControlGatewaysOK" + "." + "gateways" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetControlGatewaysOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetControlGatewaysOKBody) UnmarshalBinary(b []byte) error {
	var res GetControlGatewaysOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
