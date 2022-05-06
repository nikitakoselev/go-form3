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

// ReturnReversalAdmissionDetailsResponse return reversal admission details response
// swagger:model ReturnReversalAdmissionDetailsResponse
type ReturnReversalAdmissionDetailsResponse struct {

	// data
	Data *ReturnReversalAdmission `json:"data,omitempty"`

	// links
	Links *Links `json:"links,omitempty"`
}

func ReturnReversalAdmissionDetailsResponseWithDefaults(defaults client.Defaults) *ReturnReversalAdmissionDetailsResponse {
	return &ReturnReversalAdmissionDetailsResponse{

		Data: ReturnReversalAdmissionWithDefaults(defaults),

		Links: LinksWithDefaults(defaults),
	}
}

func (m *ReturnReversalAdmissionDetailsResponse) WithData(data ReturnReversalAdmission) *ReturnReversalAdmissionDetailsResponse {

	m.Data = &data

	return m
}

func (m *ReturnReversalAdmissionDetailsResponse) WithoutData() *ReturnReversalAdmissionDetailsResponse {
	m.Data = nil
	return m
}

func (m *ReturnReversalAdmissionDetailsResponse) WithLinks(links Links) *ReturnReversalAdmissionDetailsResponse {

	m.Links = &links

	return m
}

func (m *ReturnReversalAdmissionDetailsResponse) WithoutLinks() *ReturnReversalAdmissionDetailsResponse {
	m.Links = nil
	return m
}

// Validate validates this return reversal admission details response
func (m *ReturnReversalAdmissionDetailsResponse) Validate(formats strfmt.Registry) error {
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

func (m *ReturnReversalAdmissionDetailsResponse) validateData(formats strfmt.Registry) error {

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

func (m *ReturnReversalAdmissionDetailsResponse) validateLinks(formats strfmt.Registry) error {

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
func (m *ReturnReversalAdmissionDetailsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReturnReversalAdmissionDetailsResponse) UnmarshalBinary(b []byte) error {
	var res ReturnReversalAdmissionDetailsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *ReturnReversalAdmissionDetailsResponse) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
