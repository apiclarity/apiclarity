// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// APIInfo Api info
//
// swagger:model ApiInfo
type APIInfo struct {

	// destination namespace
	DestinationNamespace string `json:"destinationNamespace,omitempty"`

	// has provided spec
	HasProvidedSpec *bool `json:"hasProvidedSpec,omitempty"`

	// has reconstructed spec
	HasReconstructedSpec *bool `json:"hasReconstructedSpec,omitempty"`

	// id
	ID uint32 `json:"id,omitempty"`

	// API name
	Name string `json:"name,omitempty"`

	// port
	Port int64 `json:"port,omitempty"`

	// Trace Source ID which created this API. Null UUID 0 means it has been created by APIClarity (from the UI for example)
	// Format: uuid
	TraceSourceID strfmt.UUID `json:"traceSourceId,omitempty"`

	// traceSourceName
	TraceSourceName string `json:"traceSourceName,omitempty"`
}

// Validate validates this Api info
func (m *APIInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTraceSourceID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *APIInfo) validateTraceSourceID(formats strfmt.Registry) error {
	if swag.IsZero(m.TraceSourceID) { // not required
		return nil
	}

	if err := validate.FormatOf("traceSourceId", "body", "uuid", m.TraceSourceID.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this Api info based on context it is used
func (m *APIInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *APIInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIInfo) UnmarshalBinary(b []byte) error {
	var res APIInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
