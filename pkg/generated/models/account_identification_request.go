// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"

	"github.com/form3tech-oss/go-form3/v2/pkg/client"
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AccountIdentificationRequest account identification request
// swagger:model AccountIdentificationRequest
type AccountIdentificationRequest struct {

	// data
	// Required: true
	Data *AccountIdentification `json:"data"`
}

func AccountIdentificationRequestWithDefaults(defaults client.Defaults) *AccountIdentificationRequest {
	return &AccountIdentificationRequest{

		Data: AccountIdentificationWithDefaults(defaults),
	}
}

func (m *AccountIdentificationRequest) WithData(data AccountIdentification) *AccountIdentificationRequest {

	m.Data = &data

	return m
}

func (m *AccountIdentificationRequest) WithoutData() *AccountIdentificationRequest {
	m.Data = nil
	return m
}

// Validate validates this account identification request
func (m *AccountIdentificationRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountIdentificationRequest) validateData(formats strfmt.Registry) error {

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
func (m *AccountIdentificationRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountIdentificationRequest) UnmarshalBinary(b []byte) error {
	var res AccountIdentificationRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *AccountIdentificationRequest) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
