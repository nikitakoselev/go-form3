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

// QueryResponseAdmission query response admission
// swagger:model QueryResponseAdmission
type QueryResponseAdmission struct {

	// type
	// Required: true
	Type *QueryResponseAdmissionResourceType `json:"type"`

	// id
	// Required: true
	// Format: uuid
	ID *strfmt.UUID `json:"id"`

	// version
	// Required: true
	// Minimum: 0
	Version *int64 `json:"version"`

	// organisation id
	// Required: true
	// Format: uuid
	OrganisationID *strfmt.UUID `json:"organisation_id"`

	// created on
	// Format: date-time
	CreatedOn strfmt.DateTime `json:"created_on,omitempty"`

	// modified on
	// Format: date-time
	ModifiedOn strfmt.DateTime `json:"modified_on,omitempty"`

	// attributes
	// Required: true
	Attributes *QueryResponseAdmissionAttributes `json:"attributes"`

	// relationships
	Relationships *QueryResponseAdmissionRelationships `json:"relationships,omitempty"`
}

func QueryResponseAdmissionWithDefaults(defaults client.Defaults) *QueryResponseAdmission {
	return &QueryResponseAdmission{

		// TODO Type: QueryResponseAdmissionResourceType,

		ID: defaults.GetStrfmtUUIDPtr("QueryResponseAdmission", "id"),

		Version: defaults.GetInt64Ptr("QueryResponseAdmission", "version"),

		OrganisationID: defaults.GetStrfmtUUIDPtr("QueryResponseAdmission", "organisation_id"),

		CreatedOn: defaults.GetStrfmtDateTime("QueryResponseAdmission", "created_on"),

		ModifiedOn: defaults.GetStrfmtDateTime("QueryResponseAdmission", "modified_on"),

		Attributes: QueryResponseAdmissionAttributesWithDefaults(defaults),

		Relationships: QueryResponseAdmissionRelationshipsWithDefaults(defaults),
	}
}

func (m *QueryResponseAdmission) WithType(typeVar QueryResponseAdmissionResourceType) *QueryResponseAdmission {

	m.Type = &typeVar

	return m
}

func (m *QueryResponseAdmission) WithoutType() *QueryResponseAdmission {
	m.Type = nil
	return m
}

func (m *QueryResponseAdmission) WithID(id strfmt.UUID) *QueryResponseAdmission {

	m.ID = &id

	return m
}

func (m *QueryResponseAdmission) WithoutID() *QueryResponseAdmission {
	m.ID = nil
	return m
}

func (m *QueryResponseAdmission) WithVersion(version int64) *QueryResponseAdmission {

	m.Version = &version

	return m
}

func (m *QueryResponseAdmission) WithoutVersion() *QueryResponseAdmission {
	m.Version = nil
	return m
}

func (m *QueryResponseAdmission) WithOrganisationID(organisationID strfmt.UUID) *QueryResponseAdmission {

	m.OrganisationID = &organisationID

	return m
}

func (m *QueryResponseAdmission) WithoutOrganisationID() *QueryResponseAdmission {
	m.OrganisationID = nil
	return m
}

func (m *QueryResponseAdmission) WithCreatedOn(createdOn strfmt.DateTime) *QueryResponseAdmission {

	m.CreatedOn = createdOn

	return m
}

func (m *QueryResponseAdmission) WithModifiedOn(modifiedOn strfmt.DateTime) *QueryResponseAdmission {

	m.ModifiedOn = modifiedOn

	return m
}

func (m *QueryResponseAdmission) WithAttributes(attributes QueryResponseAdmissionAttributes) *QueryResponseAdmission {

	m.Attributes = &attributes

	return m
}

func (m *QueryResponseAdmission) WithoutAttributes() *QueryResponseAdmission {
	m.Attributes = nil
	return m
}

func (m *QueryResponseAdmission) WithRelationships(relationships QueryResponseAdmissionRelationships) *QueryResponseAdmission {

	m.Relationships = &relationships

	return m
}

func (m *QueryResponseAdmission) WithoutRelationships() *QueryResponseAdmission {
	m.Relationships = nil
	return m
}

// Validate validates this query response admission
func (m *QueryResponseAdmission) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrganisationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedOn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModifiedOn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAttributes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRelationships(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *QueryResponseAdmission) validateType(formats strfmt.Registry) error {

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

func (m *QueryResponseAdmission) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *QueryResponseAdmission) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version); err != nil {
		return err
	}

	if err := validate.MinimumInt("version", "body", int64(*m.Version), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *QueryResponseAdmission) validateOrganisationID(formats strfmt.Registry) error {

	if err := validate.Required("organisation_id", "body", m.OrganisationID); err != nil {
		return err
	}

	if err := validate.FormatOf("organisation_id", "body", "uuid", m.OrganisationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *QueryResponseAdmission) validateCreatedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("created_on", "body", "date-time", m.CreatedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *QueryResponseAdmission) validateModifiedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.ModifiedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("modified_on", "body", "date-time", m.ModifiedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *QueryResponseAdmission) validateAttributes(formats strfmt.Registry) error {

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

func (m *QueryResponseAdmission) validateRelationships(formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *QueryResponseAdmission) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *QueryResponseAdmission) UnmarshalBinary(b []byte) error {
	var res QueryResponseAdmission
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *QueryResponseAdmission) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
