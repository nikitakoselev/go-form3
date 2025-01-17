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
)

// PaymentAdmissionTaskDetailsResponse payment admission task details response
// swagger:model PaymentAdmissionTaskDetailsResponse
type PaymentAdmissionTaskDetailsResponse struct {

	// data
	Data *PaymentAdmissionTask `json:"data,omitempty"`

	// links
	Links *Links `json:"links,omitempty"`
}

func PaymentAdmissionTaskDetailsResponseWithDefaults(defaults client.Defaults) *PaymentAdmissionTaskDetailsResponse {
	return &PaymentAdmissionTaskDetailsResponse{

		Data: PaymentAdmissionTaskWithDefaults(defaults),

		Links: LinksWithDefaults(defaults),
	}
}

func (m *PaymentAdmissionTaskDetailsResponse) WithData(data PaymentAdmissionTask) *PaymentAdmissionTaskDetailsResponse {

	m.Data = &data

	return m
}

func (m *PaymentAdmissionTaskDetailsResponse) WithoutData() *PaymentAdmissionTaskDetailsResponse {
	m.Data = nil
	return m
}

func (m *PaymentAdmissionTaskDetailsResponse) WithLinks(links Links) *PaymentAdmissionTaskDetailsResponse {

	m.Links = &links

	return m
}

func (m *PaymentAdmissionTaskDetailsResponse) WithoutLinks() *PaymentAdmissionTaskDetailsResponse {
	m.Links = nil
	return m
}

// Validate validates this payment admission task details response
func (m *PaymentAdmissionTaskDetailsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PaymentAdmissionTaskDetailsResponse) validateData(formats strfmt.Registry) error {

	if swag.IsZero(m.Data) { // not required
		return nil
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

func (m *PaymentAdmissionTaskDetailsResponse) validateLinks(formats strfmt.Registry) error {

	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PaymentAdmissionTaskDetailsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentAdmissionTaskDetailsResponse) UnmarshalBinary(b []byte) error {
	var res PaymentAdmissionTaskDetailsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *PaymentAdmissionTaskDetailsResponse) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
