// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PostModulesSpecReconstructionAPIIDStartHandlerFunc turns a function with the right signature into a post modules spec reconstruction API ID start handler
type PostModulesSpecReconstructionAPIIDStartHandlerFunc func(PostModulesSpecReconstructionAPIIDStartParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostModulesSpecReconstructionAPIIDStartHandlerFunc) Handle(params PostModulesSpecReconstructionAPIIDStartParams) middleware.Responder {
	return fn(params)
}

// PostModulesSpecReconstructionAPIIDStartHandler interface for that can handle valid post modules spec reconstruction API ID start params
type PostModulesSpecReconstructionAPIIDStartHandler interface {
	Handle(PostModulesSpecReconstructionAPIIDStartParams) middleware.Responder
}

// NewPostModulesSpecReconstructionAPIIDStart creates a new http.Handler for the post modules spec reconstruction API ID start operation
func NewPostModulesSpecReconstructionAPIIDStart(ctx *middleware.Context, handler PostModulesSpecReconstructionAPIIDStartHandler) *PostModulesSpecReconstructionAPIIDStart {
	return &PostModulesSpecReconstructionAPIIDStart{Context: ctx, Handler: handler}
}

/* PostModulesSpecReconstructionAPIIDStart swagger:route POST /modules/spec_reconstruction/{apiId}/start postModulesSpecReconstructionApiIdStart

Start the spec reconstruction for this API.

Start the spec reconstruction for this API.

*/
type PostModulesSpecReconstructionAPIIDStart struct {
	Context *middleware.Context
	Handler PostModulesSpecReconstructionAPIIDStartHandler
}

func (o *PostModulesSpecReconstructionAPIIDStart) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPostModulesSpecReconstructionAPIIDStartParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
