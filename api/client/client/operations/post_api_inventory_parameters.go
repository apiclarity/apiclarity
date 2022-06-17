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

	"github.com/openclarity/apiclarity/api/client/models"
)

// NewPostAPIInventoryParams creates a new PostAPIInventoryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostAPIInventoryParams() *PostAPIInventoryParams {
	return &PostAPIInventoryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostAPIInventoryParamsWithTimeout creates a new PostAPIInventoryParams object
// with the ability to set a timeout on a request.
func NewPostAPIInventoryParamsWithTimeout(timeout time.Duration) *PostAPIInventoryParams {
	return &PostAPIInventoryParams{
		timeout: timeout,
	}
}

// NewPostAPIInventoryParamsWithContext creates a new PostAPIInventoryParams object
// with the ability to set a context for a request.
func NewPostAPIInventoryParamsWithContext(ctx context.Context) *PostAPIInventoryParams {
	return &PostAPIInventoryParams{
		Context: ctx,
	}
}

// NewPostAPIInventoryParamsWithHTTPClient creates a new PostAPIInventoryParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostAPIInventoryParamsWithHTTPClient(client *http.Client) *PostAPIInventoryParams {
	return &PostAPIInventoryParams{
		HTTPClient: client,
	}
}

/* PostAPIInventoryParams contains all the parameters to send to the API endpoint
   for the post API inventory operation.

   Typically these are written to a http.Request.
*/
type PostAPIInventoryParams struct {

	// Body.
	Body *models.APIInfoWithType

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post API inventory params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostAPIInventoryParams) WithDefaults() *PostAPIInventoryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post API inventory params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostAPIInventoryParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post API inventory params
func (o *PostAPIInventoryParams) WithTimeout(timeout time.Duration) *PostAPIInventoryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post API inventory params
func (o *PostAPIInventoryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post API inventory params
func (o *PostAPIInventoryParams) WithContext(ctx context.Context) *PostAPIInventoryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post API inventory params
func (o *PostAPIInventoryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post API inventory params
func (o *PostAPIInventoryParams) WithHTTPClient(client *http.Client) *PostAPIInventoryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post API inventory params
func (o *PostAPIInventoryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post API inventory params
func (o *PostAPIInventoryParams) WithBody(body *models.APIInfoWithType) *PostAPIInventoryParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post API inventory params
func (o *PostAPIInventoryParams) SetBody(body *models.APIInfoWithType) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostAPIInventoryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
