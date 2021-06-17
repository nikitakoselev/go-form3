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
)

// DirectDebitDecisionSubmissionDetailsResponse direct debit decision submission details response
// swagger:model DirectDebitDecisionSubmissionDetailsResponse
type DirectDebitDecisionSubmissionDetailsResponse struct {

	// data
	Data *DirectDebitDecisionSubmission `json:"data,omitempty"`

	// links
	Links *Links `json:"links,omitempty"`
}

func DirectDebitDecisionSubmissionDetailsResponseWithDefaults(defaults client.Defaults) *DirectDebitDecisionSubmissionDetailsResponse {
	return &DirectDebitDecisionSubmissionDetailsResponse{

		Data: DirectDebitDecisionSubmissionWithDefaults(defaults),

		Links: LinksWithDefaults(defaults),
	}
}

func (m *DirectDebitDecisionSubmissionDetailsResponse) WithData(data DirectDebitDecisionSubmission) *DirectDebitDecisionSubmissionDetailsResponse {

	m.Data = &data

	return m
}

func (m *DirectDebitDecisionSubmissionDetailsResponse) WithoutData() *DirectDebitDecisionSubmissionDetailsResponse {
	m.Data = nil
	return m
}

func (m *DirectDebitDecisionSubmissionDetailsResponse) WithLinks(links Links) *DirectDebitDecisionSubmissionDetailsResponse {

	m.Links = &links

	return m
}

func (m *DirectDebitDecisionSubmissionDetailsResponse) WithoutLinks() *DirectDebitDecisionSubmissionDetailsResponse {
	m.Links = nil
	return m
}

// Validate validates this direct debit decision submission details response
func (m *DirectDebitDecisionSubmissionDetailsResponse) Validate(formats strfmt.Registry) error {
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

func (m *DirectDebitDecisionSubmissionDetailsResponse) validateData(formats strfmt.Registry) error {

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

func (m *DirectDebitDecisionSubmissionDetailsResponse) validateLinks(formats strfmt.Registry) error {

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
func (m *DirectDebitDecisionSubmissionDetailsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DirectDebitDecisionSubmissionDetailsResponse) UnmarshalBinary(b []byte) error {
	var res DirectDebitDecisionSubmissionDetailsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *DirectDebitDecisionSubmissionDetailsResponse) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
