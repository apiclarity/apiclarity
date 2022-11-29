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
)

// NewGetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDParams creates a new GetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDParams() *GetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDParams {
	return &GetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDParamsWithTimeout creates a new GetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDParams object
// with the ability to set a timeout on a request.
func NewGetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDParamsWithTimeout(timeout time.Duration) *GetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDParams {
	return &GetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDParams{
		timeout: timeout,
	}
}

// NewGetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDParamsWithContext creates a new GetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDParams object
// with the ability to set a context for a request.
func NewGetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDParamsWithContext(ctx context.Context) *GetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDParams {
	return &GetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDParams{
		Context: ctx,
	}
}

// NewGetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDParamsWithHTTPClient creates a new GetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDParamsWithHTTPClient(client *http.Client) *GetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDParams {
	return &GetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDParams{
		HTTPClient: client,
	}
}

/* GetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDParams contains all the parameters to send to the API endpoint
   for the get API inventory API ID from host and port and trace source ID operation.

   Typically these are written to a http.Request.
*/
type GetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDParams struct {

	/* Host.

	   api host name
	*/
	Host string

	/* Port.

	   api port
	*/
	Port string

	/* TraceSourceID.

	   Trace Source ID

	   Format: uuid
	*/
	TraceSourceID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get API inventory API ID from host and port and trace source ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDParams) WithDefaults() *GetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get API inventory API ID from host and port and trace source ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get API inventory API ID from host and port and trace source ID params
func (o *GetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDParams) WithTimeout(timeout time.Duration) *GetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get API inventory API ID from host and port and trace source ID params
func (o *GetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get API inventory API ID from host and port and trace source ID params
func (o *GetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDParams) WithContext(ctx context.Context) *GetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get API inventory API ID from host and port and trace source ID params
func (o *GetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get API inventory API ID from host and port and trace source ID params
func (o *GetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDParams) WithHTTPClient(client *http.Client) *GetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get API inventory API ID from host and port and trace source ID params
func (o *GetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithHost adds the host to the get API inventory API ID from host and port and trace source ID params
func (o *GetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDParams) WithHost(host string) *GetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDParams {
	o.SetHost(host)
	return o
}

// SetHost adds the host to the get API inventory API ID from host and port and trace source ID params
func (o *GetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDParams) SetHost(host string) {
	o.Host = host
}

// WithPort adds the port to the get API inventory API ID from host and port and trace source ID params
func (o *GetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDParams) WithPort(port string) *GetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDParams {
	o.SetPort(port)
	return o
}

// SetPort adds the port to the get API inventory API ID from host and port and trace source ID params
func (o *GetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDParams) SetPort(port string) {
	o.Port = port
}

// WithTraceSourceID adds the traceSourceID to the get API inventory API ID from host and port and trace source ID params
func (o *GetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDParams) WithTraceSourceID(traceSourceID strfmt.UUID) *GetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDParams {
	o.SetTraceSourceID(traceSourceID)
	return o
}

// SetTraceSourceID adds the traceSourceId to the get API inventory API ID from host and port and trace source ID params
func (o *GetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDParams) SetTraceSourceID(traceSourceID strfmt.UUID) {
	o.TraceSourceID = traceSourceID
}

// WriteToRequest writes these params to a swagger request
func (o *GetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param host
	qrHost := o.Host
	qHost := qrHost
	if qHost != "" {

		if err := r.SetQueryParam("host", qHost); err != nil {
			return err
		}
	}

	// query param port
	qrPort := o.Port
	qPort := qrPort
	if qPort != "" {

		if err := r.SetQueryParam("port", qPort); err != nil {
			return err
		}
	}

	// query param traceSourceId
	qrTraceSourceID := o.TraceSourceID
	qTraceSourceID := qrTraceSourceID.String()
	if qTraceSourceID != "" {

		if err := r.SetQueryParam("traceSourceId", qTraceSourceID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
