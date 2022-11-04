// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewDeleteControlTraceSourcesTraceSourceIDParams creates a new DeleteControlTraceSourcesTraceSourceIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteControlTraceSourcesTraceSourceIDParams() *DeleteControlTraceSourcesTraceSourceIDParams {
	return &DeleteControlTraceSourcesTraceSourceIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteControlTraceSourcesTraceSourceIDParamsWithTimeout creates a new DeleteControlTraceSourcesTraceSourceIDParams object
// with the ability to set a timeout on a request.
func NewDeleteControlTraceSourcesTraceSourceIDParamsWithTimeout(timeout time.Duration) *DeleteControlTraceSourcesTraceSourceIDParams {
	return &DeleteControlTraceSourcesTraceSourceIDParams{
		timeout: timeout,
	}
}

// NewDeleteControlTraceSourcesTraceSourceIDParamsWithContext creates a new DeleteControlTraceSourcesTraceSourceIDParams object
// with the ability to set a context for a request.
func NewDeleteControlTraceSourcesTraceSourceIDParamsWithContext(ctx context.Context) *DeleteControlTraceSourcesTraceSourceIDParams {
	return &DeleteControlTraceSourcesTraceSourceIDParams{
		Context: ctx,
	}
}

// NewDeleteControlTraceSourcesTraceSourceIDParamsWithHTTPClient creates a new DeleteControlTraceSourcesTraceSourceIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteControlTraceSourcesTraceSourceIDParamsWithHTTPClient(client *http.Client) *DeleteControlTraceSourcesTraceSourceIDParams {
	return &DeleteControlTraceSourcesTraceSourceIDParams{
		HTTPClient: client,
	}
}

/* DeleteControlTraceSourcesTraceSourceIDParams contains all the parameters to send to the API endpoint
   for the delete control trace sources trace source ID operation.

   Typically these are written to a http.Request.
*/
type DeleteControlTraceSourcesTraceSourceIDParams struct {

	/* TraceSourceID.

	   Trace Source ID
	*/
	TraceSourceID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete control trace sources trace source ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteControlTraceSourcesTraceSourceIDParams) WithDefaults() *DeleteControlTraceSourcesTraceSourceIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete control trace sources trace source ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteControlTraceSourcesTraceSourceIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete control trace sources trace source ID params
func (o *DeleteControlTraceSourcesTraceSourceIDParams) WithTimeout(timeout time.Duration) *DeleteControlTraceSourcesTraceSourceIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete control trace sources trace source ID params
func (o *DeleteControlTraceSourcesTraceSourceIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete control trace sources trace source ID params
func (o *DeleteControlTraceSourcesTraceSourceIDParams) WithContext(ctx context.Context) *DeleteControlTraceSourcesTraceSourceIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete control trace sources trace source ID params
func (o *DeleteControlTraceSourcesTraceSourceIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete control trace sources trace source ID params
func (o *DeleteControlTraceSourcesTraceSourceIDParams) WithHTTPClient(client *http.Client) *DeleteControlTraceSourcesTraceSourceIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete control trace sources trace source ID params
func (o *DeleteControlTraceSourcesTraceSourceIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTraceSourceID adds the traceSourceID to the delete control trace sources trace source ID params
func (o *DeleteControlTraceSourcesTraceSourceIDParams) WithTraceSourceID(traceSourceID int64) *DeleteControlTraceSourcesTraceSourceIDParams {
	o.SetTraceSourceID(traceSourceID)
	return o
}

// SetTraceSourceID adds the traceSourceId to the delete control trace sources trace source ID params
func (o *DeleteControlTraceSourcesTraceSourceIDParams) SetTraceSourceID(traceSourceID int64) {
	o.TraceSourceID = traceSourceID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteControlTraceSourcesTraceSourceIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param traceSourceId
	if err := r.SetPathParam("traceSourceId", swag.FormatInt64(o.TraceSourceID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
