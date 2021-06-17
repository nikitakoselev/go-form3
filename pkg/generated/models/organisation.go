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
	"github.com/go-openapi/validate"
)

// Organisation organisation
// swagger:model Organisation
type Organisation struct {

	// attributes
	Attributes *OrganisationAttributes `json:"attributes,omitempty"`

	// Unique resource ID
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// Unique ID of the organisation this resource is created by
	// Format: uuid
	OrganisationID strfmt.UUID `json:"organisation_id,omitempty"`

	// Name of the resource type
	Type string `json:"type,omitempty"`

	// Version number
	// Minimum: 0
	Version *int64 `json:"version,omitempty"`
}

func OrganisationWithDefaults(defaults client.Defaults) *Organisation {
	return &Organisation{

		Attributes: OrganisationAttributesWithDefaults(defaults),

		ID: defaults.GetStrfmtUUID("Organisation", "id"),

		OrganisationID: defaults.GetStrfmtUUID("Organisation", "organisation_id"),

		Type: defaults.GetString("Organisation", "type"),

		Version: defaults.GetInt64Ptr("Organisation", "version"),
	}
}

func (m *Organisation) WithAttributes(attributes OrganisationAttributes) *Organisation {

	m.Attributes = &attributes

	return m
}

func (m *Organisation) WithoutAttributes() *Organisation {
	m.Attributes = nil
	return m
}

func (m *Organisation) WithID(id strfmt.UUID) *Organisation {

	m.ID = id

	return m
}

func (m *Organisation) WithOrganisationID(organisationID strfmt.UUID) *Organisation {

	m.OrganisationID = organisationID

	return m
}

func (m *Organisation) WithType(typeVar string) *Organisation {

	m.Type = typeVar

	return m
}

func (m *Organisation) WithVersion(version int64) *Organisation {

	m.Version = &version

	return m
}

func (m *Organisation) WithoutVersion() *Organisation {
	m.Version = nil
	return m
}

// Validate validates this organisation
func (m *Organisation) Validate(formats strfmt.Registry) error {
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

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Organisation) validateAttributes(formats strfmt.Registry) error {

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

func (m *Organisation) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Organisation) validateOrganisationID(formats strfmt.Registry) error {

	if swag.IsZero(m.OrganisationID) { // not required
		return nil
	}

	if err := validate.FormatOf("organisation_id", "body", "uuid", m.OrganisationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Organisation) validateVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.Version) { // not required
		return nil
	}

	if err := validate.MinimumInt("version", "body", int64(*m.Version), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Organisation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Organisation) UnmarshalBinary(b []byte) error {
	var res Organisation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *Organisation) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
