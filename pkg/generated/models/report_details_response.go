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

// ReportDetailsResponse report details response
// swagger:model ReportDetailsResponse
type ReportDetailsResponse struct {

	// data
	// Required: true
	Data *Report `json:"data"`
}

func ReportDetailsResponseWithDefaults(defaults client.Defaults) *ReportDetailsResponse {
	return &ReportDetailsResponse{

		Data: ReportWithDefaults(defaults),
	}
}

func (m *ReportDetailsResponse) WithData(data Report) *ReportDetailsResponse {

	m.Data = &data

	return m
}

func (m *ReportDetailsResponse) WithoutData() *ReportDetailsResponse {
	m.Data = nil
	return m
}

// Validate validates this report details response
func (m *ReportDetailsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReportDetailsResponse) validateData(formats strfmt.Registry) error {

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
func (m *ReportDetailsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReportDetailsResponse) UnmarshalBinary(b []byte) error {
	var res ReportDetailsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *ReportDetailsResponse) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
