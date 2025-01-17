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

// NewReversalSubmission new reversal submission
// swagger:model NewReversalSubmission
type NewReversalSubmission struct {

	// Unique resource ID
	// Required: true
	// Format: uuid
	ID *strfmt.UUID `json:"id"`

	// Unique ID of the organisation this resource is created by
	// Required: true
	// Format: uuid
	OrganisationID *strfmt.UUID `json:"organisation_id"`

	// Name of the resource type
	// Pattern: ^[A-Za-z_]*$
	Type string `json:"type,omitempty"`

	// Version number
	// Minimum: 0
	Version *int64 `json:"version,omitempty"`
}

func NewReversalSubmissionWithDefaults(defaults client.Defaults) *NewReversalSubmission {
	return &NewReversalSubmission{

		ID: defaults.GetStrfmtUUIDPtr("NewReversalSubmission", "id"),

		OrganisationID: defaults.GetStrfmtUUIDPtr("NewReversalSubmission", "organisation_id"),

		Type: defaults.GetString("NewReversalSubmission", "type"),

		Version: defaults.GetInt64Ptr("NewReversalSubmission", "version"),
	}
}

func (m *NewReversalSubmission) WithID(id strfmt.UUID) *NewReversalSubmission {

	m.ID = &id

	return m
}

func (m *NewReversalSubmission) WithoutID() *NewReversalSubmission {
	m.ID = nil
	return m
}

func (m *NewReversalSubmission) WithOrganisationID(organisationID strfmt.UUID) *NewReversalSubmission {

	m.OrganisationID = &organisationID

	return m
}

func (m *NewReversalSubmission) WithoutOrganisationID() *NewReversalSubmission {
	m.OrganisationID = nil
	return m
}

func (m *NewReversalSubmission) WithType(typeVar string) *NewReversalSubmission {

	m.Type = typeVar

	return m
}

func (m *NewReversalSubmission) WithVersion(version int64) *NewReversalSubmission {

	m.Version = &version

	return m
}

func (m *NewReversalSubmission) WithoutVersion() *NewReversalSubmission {
	m.Version = nil
	return m
}

// Validate validates this new reversal submission
func (m *NewReversalSubmission) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrganisationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewReversalSubmission) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *NewReversalSubmission) validateOrganisationID(formats strfmt.Registry) error {

	if err := validate.Required("organisation_id", "body", m.OrganisationID); err != nil {
		return err
	}

	if err := validate.FormatOf("organisation_id", "body", "uuid", m.OrganisationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *NewReversalSubmission) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := validate.Pattern("type", "body", string(m.Type), `^[A-Za-z_]*$`); err != nil {
		return err
	}

	return nil
}

func (m *NewReversalSubmission) validateVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.Version) { // not required
		return nil
	}

	if err := validate.MinimumInt("version", "body", int64(*m.Version), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NewReversalSubmission) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewReversalSubmission) UnmarshalBinary(b []byte) error {
	var res NewReversalSubmission
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *NewReversalSubmission) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
