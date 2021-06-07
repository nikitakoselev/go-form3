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

// NewQueryResponse new query response
// swagger:model NewQueryResponse
type NewQueryResponse struct {

	// attributes
	// Required: true
	Attributes *QueryResponseAttributes `json:"attributes"`

	// id
	// Required: true
	// Format: uuid
	ID *strfmt.UUID `json:"id"`

	// organisation id
	// Required: true
	// Format: uuid
	OrganisationID *strfmt.UUID `json:"organisation_id"`

	// type
	// Required: true
	Type *QueryResponseResourceType `json:"type"`
}

func NewQueryResponseWithDefaults(defaults client.Defaults) *NewQueryResponse {
	return &NewQueryResponse{

		Attributes: QueryResponseAttributesWithDefaults(defaults),

		ID: defaults.GetStrfmtUUIDPtr("NewQueryResponse", "id"),

		OrganisationID: defaults.GetStrfmtUUIDPtr("NewQueryResponse", "organisation_id"),

		// TODO Type: QueryResponseResourceType,

	}
}

func (m *NewQueryResponse) WithAttributes(attributes QueryResponseAttributes) *NewQueryResponse {

	m.Attributes = &attributes

	return m
}

func (m *NewQueryResponse) WithoutAttributes() *NewQueryResponse {
	m.Attributes = nil
	return m
}

func (m *NewQueryResponse) WithID(id strfmt.UUID) *NewQueryResponse {

	m.ID = &id

	return m
}

func (m *NewQueryResponse) WithoutID() *NewQueryResponse {
	m.ID = nil
	return m
}

func (m *NewQueryResponse) WithOrganisationID(organisationID strfmt.UUID) *NewQueryResponse {

	m.OrganisationID = &organisationID

	return m
}

func (m *NewQueryResponse) WithoutOrganisationID() *NewQueryResponse {
	m.OrganisationID = nil
	return m
}

func (m *NewQueryResponse) WithType(typeVar QueryResponseResourceType) *NewQueryResponse {

	m.Type = &typeVar

	return m
}

func (m *NewQueryResponse) WithoutType() *NewQueryResponse {
	m.Type = nil
	return m
}

// Validate validates this new query response
func (m *NewQueryResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttributes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrganisationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewQueryResponse) validateAttributes(formats strfmt.Registry) error {

	if err := validate.Required("attributes", "body", m.Attributes); err != nil {
		return err
	}

	if m.Attributes != nil {
		if err := m.Attributes.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attributes")
			}
			return err
		}
	}

	return nil
}

func (m *NewQueryResponse) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *NewQueryResponse) validateOrganisationID(formats strfmt.Registry) error {

	if err := validate.Required("organisation_id", "body", m.OrganisationID); err != nil {
		return err
	}

	if err := validate.FormatOf("organisation_id", "body", "uuid", m.OrganisationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *NewQueryResponse) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	if m.Type != nil {
		if err := m.Type.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NewQueryResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewQueryResponse) UnmarshalBinary(b []byte) error {
	var res NewQueryResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *NewQueryResponse) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
