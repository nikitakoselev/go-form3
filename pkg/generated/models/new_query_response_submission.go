// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"

	"github.com/form3tech-oss/go-form3/pkg/client"
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NewQueryResponseSubmission new query response submission
// swagger:model NewQueryResponseSubmission
type NewQueryResponseSubmission struct {

	// attributes
	Attributes *QueryResponseSubmissionAttributes `json:"attributes,omitempty"`

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
	Type QueryResponseSubmissionResourceType `json:"type"`
}

func NewQueryResponseSubmissionWithDefaults(defaults client.Defaults) *NewQueryResponseSubmission {
	return &NewQueryResponseSubmission{

		Attributes: QueryResponseSubmissionAttributesWithDefaults(defaults),

		ID: defaults.GetStrfmtUUIDPtr("NewQueryResponseSubmission", "id"),

		OrganisationID: defaults.GetStrfmtUUIDPtr("NewQueryResponseSubmission", "organisation_id"),

		// TODO Type: QueryResponseSubmissionResourceType,

	}
}

func (m *NewQueryResponseSubmission) WithAttributes(attributes QueryResponseSubmissionAttributes) *NewQueryResponseSubmission {

	m.Attributes = &attributes

	return m
}

func (m *NewQueryResponseSubmission) WithoutAttributes() *NewQueryResponseSubmission {
	m.Attributes = nil
	return m
}

func (m *NewQueryResponseSubmission) WithID(id strfmt.UUID) *NewQueryResponseSubmission {

	m.ID = &id

	return m
}

func (m *NewQueryResponseSubmission) WithoutID() *NewQueryResponseSubmission {
	m.ID = nil
	return m
}

func (m *NewQueryResponseSubmission) WithOrganisationID(organisationID strfmt.UUID) *NewQueryResponseSubmission {

	m.OrganisationID = &organisationID

	return m
}

func (m *NewQueryResponseSubmission) WithoutOrganisationID() *NewQueryResponseSubmission {
	m.OrganisationID = nil
	return m
}

func (m *NewQueryResponseSubmission) WithType(typeVar QueryResponseSubmissionResourceType) *NewQueryResponseSubmission {

	m.Type = typeVar

	return m
}

// Validate validates this new query response submission
func (m *NewQueryResponseSubmission) Validate(formats strfmt.Registry) error {
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

func (m *NewQueryResponseSubmission) validateAttributes(formats strfmt.Registry) error {

	if swag.IsZero(m.Attributes) { // not required
		return nil
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

func (m *NewQueryResponseSubmission) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *NewQueryResponseSubmission) validateOrganisationID(formats strfmt.Registry) error {

	if err := validate.Required("organisation_id", "body", m.OrganisationID); err != nil {
		return err
	}

	if err := validate.FormatOf("organisation_id", "body", "uuid", m.OrganisationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *NewQueryResponseSubmission) validateType(formats strfmt.Registry) error {

	if err := m.Type.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("type")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NewQueryResponseSubmission) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewQueryResponseSubmission) UnmarshalBinary(b []byte) error {
	var res NewQueryResponseSubmission
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *NewQueryResponseSubmission) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}