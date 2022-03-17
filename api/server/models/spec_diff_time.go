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

// SpecDiffTime spec diff time
//
// swagger:model SpecDiffTime
type SpecDiffTime struct {

	// api event Id
	APIEventID uint32 `json:"apiEventId,omitempty"`

	// api host name
	APIHostName string `json:"apiHostName,omitempty"`

	// diff type
	DiffType *DiffType `json:"diffType,omitempty"`

	// time
	// Format: date-time
	Time strfmt.DateTime `json:"time,omitempty"`
}

// Validate validates this spec diff time
func (m *SpecDiffTime) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDiffType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SpecDiffTime) validateDiffType(formats strfmt.Registry) error {
	if swag.IsZero(m.DiffType) { // not required
		return nil
	}

	if m.DiffType != nil {
		if err := m.DiffType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("diffType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("diffType")
			}
			return err
		}
	}

	return nil
}

func (m *SpecDiffTime) validateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.Time) { // not required
		return nil
	}

	if err := validate.FormatOf("time", "body", "date-time", m.Time.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this spec diff time based on the context it is used
func (m *SpecDiffTime) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDiffType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SpecDiffTime) contextValidateDiffType(ctx context.Context, formats strfmt.Registry) error {

	if m.DiffType != nil {
		if err := m.DiffType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("diffType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("diffType")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SpecDiffTime) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SpecDiffTime) UnmarshalBinary(b []byte) error {
	var res SpecDiffTime
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
