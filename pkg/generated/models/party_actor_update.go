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

// PartyActorUpdate party actor update
// swagger:model PartyActorUpdate
type PartyActorUpdate struct {

	// attributes
	Attributes *PartyActorPatchAttributes `json:"attributes,omitempty"`

	// id
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// organisation id
	// Format: uuid
	OrganisationID strfmt.UUID `json:"organisation_id,omitempty"`

	// type
	// Enum: [party_actors]
	Type string `json:"type,omitempty"`

	// version
	// Minimum: 0
	Version *int64 `json:"version,omitempty"`
}

func PartyActorUpdateWithDefaults(defaults client.Defaults) *PartyActorUpdate {
	return &PartyActorUpdate{

		Attributes: PartyActorPatchAttributesWithDefaults(defaults),

		ID: defaults.GetStrfmtUUID("PartyActorUpdate", "id"),

		OrganisationID: defaults.GetStrfmtUUID("PartyActorUpdate", "organisation_id"),

		Type: defaults.GetString("PartyActorUpdate", "type"),

		Version: defaults.GetInt64Ptr("PartyActorUpdate", "version"),
	}
}

func (m *PartyActorUpdate) WithAttributes(attributes PartyActorPatchAttributes) *PartyActorUpdate {

	m.Attributes = &attributes

	return m
}

func (m *PartyActorUpdate) WithoutAttributes() *PartyActorUpdate {
	m.Attributes = nil
	return m
}

func (m *PartyActorUpdate) WithID(id strfmt.UUID) *PartyActorUpdate {

	m.ID = id

	return m
}

func (m *PartyActorUpdate) WithOrganisationID(organisationID strfmt.UUID) *PartyActorUpdate {

	m.OrganisationID = organisationID

	return m
}

func (m *PartyActorUpdate) WithType(typeVar string) *PartyActorUpdate {

	m.Type = typeVar

	return m
}

func (m *PartyActorUpdate) WithVersion(version int64) *PartyActorUpdate {

	m.Version = &version

	return m
}

func (m *PartyActorUpdate) WithoutVersion() *PartyActorUpdate {
	m.Version = nil
	return m
}

// Validate validates this party actor update
func (m *PartyActorUpdate) Validate(formats strfmt.Registry) error {
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

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PartyActorUpdate) validateAttributes(formats strfmt.Registry) error {

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

func (m *PartyActorUpdate) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PartyActorUpdate) validateOrganisationID(formats strfmt.Registry) error {

	if swag.IsZero(m.OrganisationID) { // not required
		return nil
	}

	if err := validate.FormatOf("organisation_id", "body", "uuid", m.OrganisationID.String(), formats); err != nil {
		return err
	}

	return nil
}

var partyActorUpdateTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["party_actors"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		partyActorUpdateTypeTypePropEnum = append(partyActorUpdateTypeTypePropEnum, v)
	}
}

const (

	// PartyActorUpdateTypePartyActors captures enum value "party_actors"
	PartyActorUpdateTypePartyActors string = "party_actors"
)

// prop value enum
func (m *PartyActorUpdate) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, partyActorUpdateTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PartyActorUpdate) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *PartyActorUpdate) validateVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.Version) { // not required
		return nil
	}

	if err := validate.MinimumInt("version", "body", int64(*m.Version), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PartyActorUpdate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PartyActorUpdate) UnmarshalBinary(b []byte) error {
	var res PartyActorUpdate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *PartyActorUpdate) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
