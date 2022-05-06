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

// AccountAmendment account amendment
// swagger:model AccountAmendment
type AccountAmendment struct {

	// attributes
	// Required: true
	Attributes *AccountAmendmentAttributes `json:"attributes"`

	// Date when the resource was created
	// Read Only: true
	// Format: date-time
	CreatedOn strfmt.DateTime `json:"created_on,omitempty"`

	// data
	Data *AccountUpdate `json:"data,omitempty"`

	// Account Request ID
	// Required: true
	// Format: uuid
	ID strfmt.UUID `json:"id"`

	// Date when the resource was last modified
	// Read Only: true
	// Format: date-time
	ModifiedOn strfmt.DateTime `json:"modified_on,omitempty"`

	// Unique ID of the organisation this resource is created by
	// Required: true
	// Format: uuid
	OrganisationID strfmt.UUID `json:"organisation_id"`

	// relationships
	Relationships *AccountAmendmentRelationships `json:"relationships,omitempty"`

	// type
	// Required: true
	Type ResourceType `json:"type"`

	// Version Number
	// Minimum: 0
	Version int64 `json:"version"`
}

func AccountAmendmentWithDefaults(defaults client.Defaults) *AccountAmendment {
	return &AccountAmendment{

		Attributes: AccountAmendmentAttributesWithDefaults(defaults),

		CreatedOn: defaults.GetStrfmtDateTime("AccountAmendment", "created_on"),

		Data: AccountUpdateWithDefaults(defaults),

		ID: defaults.GetStrfmtUUID("AccountAmendment", "id"),

		ModifiedOn: defaults.GetStrfmtDateTime("AccountAmendment", "modified_on"),

		OrganisationID: defaults.GetStrfmtUUID("AccountAmendment", "organisation_id"),

		Relationships: AccountAmendmentRelationshipsWithDefaults(defaults),

		// TODO Type: ResourceType,

		Version: defaults.GetInt64("AccountAmendment", "version"),
	}
}

func (m *AccountAmendment) WithAttributes(attributes AccountAmendmentAttributes) *AccountAmendment {

	m.Attributes = &attributes

	return m
}

func (m *AccountAmendment) WithoutAttributes() *AccountAmendment {
	m.Attributes = nil
	return m
}

func (m *AccountAmendment) WithCreatedOn(createdOn strfmt.DateTime) *AccountAmendment {

	m.CreatedOn = createdOn

	return m
}

func (m *AccountAmendment) WithData(data AccountUpdate) *AccountAmendment {

	m.Data = &data

	return m
}

func (m *AccountAmendment) WithoutData() *AccountAmendment {
	m.Data = nil
	return m
}

func (m *AccountAmendment) WithID(id strfmt.UUID) *AccountAmendment {

	m.ID = id

	return m
}

func (m *AccountAmendment) WithModifiedOn(modifiedOn strfmt.DateTime) *AccountAmendment {

	m.ModifiedOn = modifiedOn

	return m
}

func (m *AccountAmendment) WithOrganisationID(organisationID strfmt.UUID) *AccountAmendment {

	m.OrganisationID = organisationID

	return m
}

func (m *AccountAmendment) WithRelationships(relationships AccountAmendmentRelationships) *AccountAmendment {

	m.Relationships = &relationships

	return m
}

func (m *AccountAmendment) WithoutRelationships() *AccountAmendment {
	m.Relationships = nil
	return m
}

func (m *AccountAmendment) WithType(typeVar ResourceType) *AccountAmendment {

	m.Type = typeVar

	return m
}

func (m *AccountAmendment) WithVersion(version int64) *AccountAmendment {

	m.Version = version

	return m
}

// Validate validates this account amendment
func (m *AccountAmendment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttributes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedOn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModifiedOn(formats); err != nil {
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

func (m *AccountAmendment) validateAttributes(formats strfmt.Registry) error {

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

func (m *AccountAmendment) validateCreatedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("created_on", "body", "date-time", m.CreatedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *AccountAmendment) validateData(formats strfmt.Registry) error {

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

func (m *AccountAmendment) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", strfmt.UUID(m.ID)); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *AccountAmendment) validateModifiedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.ModifiedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("modified_on", "body", "date-time", m.ModifiedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *AccountAmendment) validateOrganisationID(formats strfmt.Registry) error {

	if err := validate.Required("organisation_id", "body", strfmt.UUID(m.OrganisationID)); err != nil {
		return err
	}

	if err := validate.FormatOf("organisation_id", "body", "uuid", m.OrganisationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *AccountAmendment) validateRelationships(formats strfmt.Registry) error {

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

func (m *AccountAmendment) validateType(formats strfmt.Registry) error {

	if err := m.Type.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("type")
		}
		return err
	}

	return nil
}

func (m *AccountAmendment) validateVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.Version) { // not required
		return nil
	}

	if err := validate.MinimumInt("version", "body", int64(m.Version), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AccountAmendment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountAmendment) UnmarshalBinary(b []byte) error {
	var res AccountAmendment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *AccountAmendment) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
