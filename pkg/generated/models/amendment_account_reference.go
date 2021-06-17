// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"

	"github.com/form3tech-oss/go-form3/v3/pkg/client"
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AmendmentAccountReference Account Reference
// swagger:model AmendmentAccountReference
type AmendmentAccountReference struct {

	// Account ID
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// type
	Type ResourceType `json:"type,omitempty"`

	// Version Number
	// Minimum: 0
	Version int64 `json:"version"`
}

func AmendmentAccountReferenceWithDefaults(defaults client.Defaults) *AmendmentAccountReference {
	return &AmendmentAccountReference{

		ID: defaults.GetStrfmtUUID("AmendmentAccountReference", "id"),

		// TODO Type: ResourceType,

		Version: defaults.GetInt64("AmendmentAccountReference", "version"),
	}
}

func (m *AmendmentAccountReference) WithID(id strfmt.UUID) *AmendmentAccountReference {

	m.ID = id

	return m
}

func (m *AmendmentAccountReference) WithType(typeVar ResourceType) *AmendmentAccountReference {

	m.Type = typeVar

	return m
}

func (m *AmendmentAccountReference) WithVersion(version int64) *AmendmentAccountReference {

	m.Version = version

	return m
}

// Validate validates this amendment account reference
func (m *AmendmentAccountReference) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AmendmentAccountReference) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *AmendmentAccountReference) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := m.Type.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("type")
		}
		return err
	}

	return nil
}

func (m *AmendmentAccountReference) validateVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.Version) { // not required
		return nil
	}

	if err := validate.MinimumInt("version", "body", int64(m.Version), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AmendmentAccountReference) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AmendmentAccountReference) UnmarshalBinary(b []byte) error {
	var res AmendmentAccountReference
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *AmendmentAccountReference) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
