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

// BranchDetailsResponse branch details response
// swagger:model BranchDetailsResponse
type BranchDetailsResponse struct {

	// data
	Data *Branch `json:"data,omitempty"`

	// links
	Links *Links `json:"links,omitempty"`
}

func BranchDetailsResponseWithDefaults(defaults client.Defaults) *BranchDetailsResponse {
	return &BranchDetailsResponse{

		Data: BranchWithDefaults(defaults),

		Links: LinksWithDefaults(defaults),
	}
}

func (m *BranchDetailsResponse) WithData(data Branch) *BranchDetailsResponse {

	m.Data = &data

	return m
}

func (m *BranchDetailsResponse) WithoutData() *BranchDetailsResponse {
	m.Data = nil
	return m
}

func (m *BranchDetailsResponse) WithLinks(links Links) *BranchDetailsResponse {

	m.Links = &links

	return m
}

func (m *BranchDetailsResponse) WithoutLinks() *BranchDetailsResponse {
	m.Links = nil
	return m
}

// Validate validates this branch details response
func (m *BranchDetailsResponse) Validate(formats strfmt.Registry) error {
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

func (m *BranchDetailsResponse) validateData(formats strfmt.Registry) error {

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

func (m *BranchDetailsResponse) validateLinks(formats strfmt.Registry) error {

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
func (m *BranchDetailsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BranchDetailsResponse) UnmarshalBinary(b []byte) error {
	var res BranchDetailsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *BranchDetailsResponse) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
