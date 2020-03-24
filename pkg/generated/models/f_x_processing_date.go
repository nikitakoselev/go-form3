// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// FXProcessingDate The date for which the rate would be valid
// swagger:model FXProcessingDate
type FXProcessingDate strfmt.Date

// UnmarshalJSON sets a FXProcessingDate value from JSON input
func (m *FXProcessingDate) UnmarshalJSON(b []byte) error {
	return ((*strfmt.Date)(m)).UnmarshalJSON(b)
}

// MarshalJSON retrieves a FXProcessingDate value as JSON output
func (m FXProcessingDate) MarshalJSON() ([]byte, error) {
	return (strfmt.Date(m)).MarshalJSON()
}

// Validate validates this f x processing date
func (m FXProcessingDate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := validate.FormatOf("", "body", "date", strfmt.Date(m).String(), formats); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *FXProcessingDate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FXProcessingDate) UnmarshalBinary(b []byte) error {
	var res FXProcessingDate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *FXProcessingDate) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}