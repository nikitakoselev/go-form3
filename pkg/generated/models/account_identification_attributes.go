// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"

	"github.com/form3tech-oss/go-form3/v4/pkg/client"
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AccountIdentificationAttributes account identification attributes
// swagger:model AccountIdentificationAttributes
type AccountIdentificationAttributes struct {

	// Used to validate where expecting an exact match against payment's reference information.
	// Required: true
	// Max Length: 35
	// Min Length: 1
	SecondaryIdentification *string `json:"secondary_identification"`
}

func AccountIdentificationAttributesWithDefaults(defaults client.Defaults) *AccountIdentificationAttributes {
	return &AccountIdentificationAttributes{

		SecondaryIdentification: defaults.GetStringPtr("AccountIdentificationAttributes", "secondary_identification"),
	}
}

func (m *AccountIdentificationAttributes) WithSecondaryIdentification(secondaryIdentification string) *AccountIdentificationAttributes {

	m.SecondaryIdentification = &secondaryIdentification

	return m
}

func (m *AccountIdentificationAttributes) WithoutSecondaryIdentification() *AccountIdentificationAttributes {
	m.SecondaryIdentification = nil
	return m
}

// Validate validates this account identification attributes
func (m *AccountIdentificationAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSecondaryIdentification(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountIdentificationAttributes) validateSecondaryIdentification(formats strfmt.Registry) error {

	if err := validate.Required("secondary_identification", "body", m.SecondaryIdentification); err != nil {
		return err
	}

	if err := validate.MinLength("secondary_identification", "body", string(*m.SecondaryIdentification), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("secondary_identification", "body", string(*m.SecondaryIdentification), 35); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AccountIdentificationAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountIdentificationAttributes) UnmarshalBinary(b []byte) error {
	var res AccountIdentificationAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *AccountIdentificationAttributes) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
