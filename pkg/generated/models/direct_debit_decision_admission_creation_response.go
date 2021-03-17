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
)

// DirectDebitDecisionAdmissionCreationResponse direct debit decision admission creation response
// swagger:model DirectDebitDecisionAdmissionCreationResponse
type DirectDebitDecisionAdmissionCreationResponse struct {

	// data
	Data *DirectDebitDecisionAdmission `json:"data,omitempty"`

	// links
	Links *Links `json:"links,omitempty"`
}

func DirectDebitDecisionAdmissionCreationResponseWithDefaults(defaults client.Defaults) *DirectDebitDecisionAdmissionCreationResponse {
	return &DirectDebitDecisionAdmissionCreationResponse{

		Data: DirectDebitDecisionAdmissionWithDefaults(defaults),

		Links: LinksWithDefaults(defaults),
	}
}

func (m *DirectDebitDecisionAdmissionCreationResponse) WithData(data DirectDebitDecisionAdmission) *DirectDebitDecisionAdmissionCreationResponse {

	m.Data = &data

	return m
}

func (m *DirectDebitDecisionAdmissionCreationResponse) WithoutData() *DirectDebitDecisionAdmissionCreationResponse {
	m.Data = nil
	return m
}

func (m *DirectDebitDecisionAdmissionCreationResponse) WithLinks(links Links) *DirectDebitDecisionAdmissionCreationResponse {

	m.Links = &links

	return m
}

func (m *DirectDebitDecisionAdmissionCreationResponse) WithoutLinks() *DirectDebitDecisionAdmissionCreationResponse {
	m.Links = nil
	return m
}

// Validate validates this direct debit decision admission creation response
func (m *DirectDebitDecisionAdmissionCreationResponse) Validate(formats strfmt.Registry) error {
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

func (m *DirectDebitDecisionAdmissionCreationResponse) validateData(formats strfmt.Registry) error {

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

func (m *DirectDebitDecisionAdmissionCreationResponse) validateLinks(formats strfmt.Registry) error {

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
func (m *DirectDebitDecisionAdmissionCreationResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DirectDebitDecisionAdmissionCreationResponse) UnmarshalBinary(b []byte) error {
	var res DirectDebitDecisionAdmissionCreationResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *DirectDebitDecisionAdmissionCreationResponse) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}