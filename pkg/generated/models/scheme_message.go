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

// SchemeMessage scheme message
// swagger:model SchemeMessage
type SchemeMessage struct {

	// attributes
	// Required: true
	Attributes *SchemeMessageAttributes `json:"attributes"`

	// id
	// Required: true
	// Format: uuid
	ID strfmt.UUID `json:"id"`

	// organisation id
	// Required: true
	// Format: uuid
	OrganisationID strfmt.UUID `json:"organisation_id"`

	// relationships
	Relationships *SchemeMessageRelationships `json:"relationships,omitempty"`

	// type
	// Enum: [scheme_messages]
	Type string `json:"type,omitempty"`

	// version
	// Minimum: 0
	Version *int64 `json:"version,omitempty"`
}

func SchemeMessageWithDefaults(defaults client.Defaults) *SchemeMessage {
	return &SchemeMessage{

		Attributes: SchemeMessageAttributesWithDefaults(defaults),

		ID: defaults.GetStrfmtUUID("SchemeMessage", "id"),

		OrganisationID: defaults.GetStrfmtUUID("SchemeMessage", "organisation_id"),

		Relationships: SchemeMessageRelationshipsWithDefaults(defaults),

		Type: defaults.GetString("SchemeMessage", "type"),

		Version: defaults.GetInt64Ptr("SchemeMessage", "version"),
	}
}

func (m *SchemeMessage) WithAttributes(attributes SchemeMessageAttributes) *SchemeMessage {

	m.Attributes = &attributes

	return m
}

func (m *SchemeMessage) WithoutAttributes() *SchemeMessage {
	m.Attributes = nil
	return m
}

func (m *SchemeMessage) WithID(id strfmt.UUID) *SchemeMessage {

	m.ID = id

	return m
}

func (m *SchemeMessage) WithOrganisationID(organisationID strfmt.UUID) *SchemeMessage {

	m.OrganisationID = organisationID

	return m
}

func (m *SchemeMessage) WithRelationships(relationships SchemeMessageRelationships) *SchemeMessage {

	m.Relationships = &relationships

	return m
}

func (m *SchemeMessage) WithoutRelationships() *SchemeMessage {
	m.Relationships = nil
	return m
}

func (m *SchemeMessage) WithType(typeVar string) *SchemeMessage {

	m.Type = typeVar

	return m
}

func (m *SchemeMessage) WithVersion(version int64) *SchemeMessage {

	m.Version = &version

	return m
}

func (m *SchemeMessage) WithoutVersion() *SchemeMessage {
	m.Version = nil
	return m
}

// Validate validates this scheme message
func (m *SchemeMessage) Validate(formats strfmt.Registry) error {
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

	if err := m.validateRelationships(formats); err != nil {
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

func (m *SchemeMessage) validateAttributes(formats strfmt.Registry) error {

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

func (m *SchemeMessage) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", strfmt.UUID(m.ID)); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SchemeMessage) validateOrganisationID(formats strfmt.Registry) error {

	if err := validate.Required("organisation_id", "body", strfmt.UUID(m.OrganisationID)); err != nil {
		return err
	}

	if err := validate.FormatOf("organisation_id", "body", "uuid", m.OrganisationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SchemeMessage) validateRelationships(formats strfmt.Registry) error {

	if swag.IsZero(m.Relationships) { // not required
		return nil
	}

	if m.Relationships != nil {
		if err := m.Relationships.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("relationships")
			}
			return err
		}
	}

	return nil
}

var schemeMessageTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["scheme_messages"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		schemeMessageTypeTypePropEnum = append(schemeMessageTypeTypePropEnum, v)
	}
}

const (

	// SchemeMessageTypeSchemeMessages captures enum value "scheme_messages"
	SchemeMessageTypeSchemeMessages string = "scheme_messages"
)

// prop value enum
func (m *SchemeMessage) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, schemeMessageTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SchemeMessage) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *SchemeMessage) validateVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.Version) { // not required
		return nil
	}

	if err := validate.MinimumInt("version", "body", int64(*m.Version), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SchemeMessage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SchemeMessage) UnmarshalBinary(b []byte) error {
	var res SchemeMessage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *SchemeMessage) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
