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

// ClaimReversalSubmissionDetailsResponse claim reversal submission details response
// swagger:model ClaimReversalSubmissionDetailsResponse
type ClaimReversalSubmissionDetailsResponse struct {

	// data
	Data *ClaimReversalSubmission `json:"data,omitempty"`

	// links
	Links *Links `json:"links,omitempty"`
}

func ClaimReversalSubmissionDetailsResponseWithDefaults(defaults client.Defaults) *ClaimReversalSubmissionDetailsResponse {
	return &ClaimReversalSubmissionDetailsResponse{

		Data: ClaimReversalSubmissionWithDefaults(defaults),

		Links: LinksWithDefaults(defaults),
	}
}

func (m *ClaimReversalSubmissionDetailsResponse) WithData(data ClaimReversalSubmission) *ClaimReversalSubmissionDetailsResponse {

	m.Data = &data

	return m
}

func (m *ClaimReversalSubmissionDetailsResponse) WithoutData() *ClaimReversalSubmissionDetailsResponse {
	m.Data = nil
	return m
}

func (m *ClaimReversalSubmissionDetailsResponse) WithLinks(links Links) *ClaimReversalSubmissionDetailsResponse {

	m.Links = &links

	return m
}

func (m *ClaimReversalSubmissionDetailsResponse) WithoutLinks() *ClaimReversalSubmissionDetailsResponse {
	m.Links = nil
	return m
}

// Validate validates this claim reversal submission details response
func (m *ClaimReversalSubmissionDetailsResponse) Validate(formats strfmt.Registry) error {
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

func (m *ClaimReversalSubmissionDetailsResponse) validateData(formats strfmt.Registry) error {

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

func (m *ClaimReversalSubmissionDetailsResponse) validateLinks(formats strfmt.Registry) error {

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
func (m *ClaimReversalSubmissionDetailsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClaimReversalSubmissionDetailsResponse) UnmarshalBinary(b []byte) error {
	var res ClaimReversalSubmissionDetailsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *ClaimReversalSubmissionDetailsResponse) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
