// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// OASVersion OpenAPI version
//
// swagger:model OASVersion
type OASVersion string

func NewOASVersion(value OASVersion) *OASVersion {
	v := value
	return &v
}

const (

	// OASVersionOASv2Dot0 captures enum value "OASv2.0"
	OASVersionOASv2Dot0 OASVersion = "OASv2.0"

	// OASVersionOASv3Dot0 captures enum value "OASv3.0"
	OASVersionOASv3Dot0 OASVersion = "OASv3.0"
)

// for schema
var oASVersionEnum []interface{}

func init() {
	var res []OASVersion
	if err := json.Unmarshal([]byte(`["OASv2.0","OASv3.0"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		oASVersionEnum = append(oASVersionEnum, v)
	}
}

func (m OASVersion) validateOASVersionEnum(path, location string, value OASVersion) error {
	if err := validate.EnumCase(path, location, value, oASVersionEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this o a s version
func (m OASVersion) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateOASVersionEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this o a s version based on context it is used
func (m OASVersion) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
