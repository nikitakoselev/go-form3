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

// AccountCreation account creation
// swagger:model AccountCreation
type AccountCreation struct {

	// data
	// Required: true
	Data *Account `json:"data"`
}

func AccountCreationWithDefaults(defaults client.Defaults) *AccountCreation {
	return &AccountCreation{

		Data: AccountWithDefaults(defaults),
	}
}

func (m *AccountCreation) WithData(data Account) *AccountCreation {

	m.Data = &data

	return m
}

func (m *AccountCreation) WithoutData() *AccountCreation {
	m.Data = nil
	return m
}

// Validate validates this account creation
func (m *AccountCreation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountCreation) validateData(formats strfmt.Registry) error {

	if err := validate.Required("data", "body", m.Data); err != nil {
		return err
	}

	if m.Data != nil {
		if err := m.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AccountCreation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountCreation) UnmarshalBinary(b []byte) error {
	var res AccountCreation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *AccountCreation) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
