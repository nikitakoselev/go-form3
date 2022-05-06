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

// RecallSubmission recall submission
// swagger:model RecallSubmission
type RecallSubmission struct {

	// attributes
	Attributes *RecallSubmissionAttributes `json:"attributes,omitempty"`

	// created on
	// Format: date-time
	CreatedOn *strfmt.DateTime `json:"created_on,omitempty"`

	// Unique resource ID
	// Required: true
	// Format: uuid
	ID *strfmt.UUID `json:"id"`

	// modified on
	// Format: date-time
	ModifiedOn *strfmt.DateTime `json:"modified_on,omitempty"`

	// Unique ID of the organisation this resource is created by
	// Required: true
	// Format: uuid
	OrganisationID *strfmt.UUID `json:"organisation_id"`

	// relationships
	Relationships *RecallSubmissionRelationships `json:"relationships,omitempty"`

	// Name of the resource type
	// Pattern: ^[A-Za-z_]*$
	Type string `json:"type,omitempty"`

	// Version number
	// Minimum: 0
	Version *int64 `json:"version,omitempty"`
}

func RecallSubmissionWithDefaults(defaults client.Defaults) *RecallSubmission {
	return &RecallSubmission{

		Attributes: RecallSubmissionAttributesWithDefaults(defaults),

		CreatedOn: defaults.GetStrfmtDateTimePtr("RecallSubmission", "created_on"),

		ID: defaults.GetStrfmtUUIDPtr("RecallSubmission", "id"),

		ModifiedOn: defaults.GetStrfmtDateTimePtr("RecallSubmission", "modified_on"),

		OrganisationID: defaults.GetStrfmtUUIDPtr("RecallSubmission", "organisation_id"),

		Relationships: RecallSubmissionRelationshipsWithDefaults(defaults),

		Type: defaults.GetString("RecallSubmission", "type"),

		Version: defaults.GetInt64Ptr("RecallSubmission", "version"),
	}
}

func (m *RecallSubmission) WithAttributes(attributes RecallSubmissionAttributes) *RecallSubmission {

	m.Attributes = &attributes

	return m
}

func (m *RecallSubmission) WithoutAttributes() *RecallSubmission {
	m.Attributes = nil
	return m
}

func (m *RecallSubmission) WithCreatedOn(createdOn strfmt.DateTime) *RecallSubmission {

	m.CreatedOn = &createdOn

	return m
}

func (m *RecallSubmission) WithoutCreatedOn() *RecallSubmission {
	m.CreatedOn = nil
	return m
}

func (m *RecallSubmission) WithID(id strfmt.UUID) *RecallSubmission {

	m.ID = &id

	return m
}

func (m *RecallSubmission) WithoutID() *RecallSubmission {
	m.ID = nil
	return m
}

func (m *RecallSubmission) WithModifiedOn(modifiedOn strfmt.DateTime) *RecallSubmission {

	m.ModifiedOn = &modifiedOn

	return m
}

func (m *RecallSubmission) WithoutModifiedOn() *RecallSubmission {
	m.ModifiedOn = nil
	return m
}

func (m *RecallSubmission) WithOrganisationID(organisationID strfmt.UUID) *RecallSubmission {

	m.OrganisationID = &organisationID

	return m
}

func (m *RecallSubmission) WithoutOrganisationID() *RecallSubmission {
	m.OrganisationID = nil
	return m
}

func (m *RecallSubmission) WithRelationships(relationships RecallSubmissionRelationships) *RecallSubmission {

	m.Relationships = &relationships

	return m
}

func (m *RecallSubmission) WithoutRelationships() *RecallSubmission {
	m.Relationships = nil
	return m
}

func (m *RecallSubmission) WithType(typeVar string) *RecallSubmission {

	m.Type = typeVar

	return m
}

func (m *RecallSubmission) WithVersion(version int64) *RecallSubmission {

	m.Version = &version

	return m
}

func (m *RecallSubmission) WithoutVersion() *RecallSubmission {
	m.Version = nil
	return m
}

// Validate validates this recall submission
func (m *RecallSubmission) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttributes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedOn(formats); err != nil {
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

func (m *RecallSubmission) validateAttributes(formats strfmt.Registry) error {

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

func (m *RecallSubmission) validateCreatedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("created_on", "body", "date-time", m.CreatedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *RecallSubmission) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *RecallSubmission) validateModifiedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.ModifiedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("modified_on", "body", "date-time", m.ModifiedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *RecallSubmission) validateOrganisationID(formats strfmt.Registry) error {

	if err := validate.Required("organisation_id", "body", m.OrganisationID); err != nil {
		return err
	}

	if err := validate.FormatOf("organisation_id", "body", "uuid", m.OrganisationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *RecallSubmission) validateRelationships(formats strfmt.Registry) error {

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

func (m *RecallSubmission) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := validate.Pattern("type", "body", string(m.Type), `^[A-Za-z_]*$`); err != nil {
		return err
	}

	return nil
}

func (m *RecallSubmission) validateVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.Version) { // not required
		return nil
	}

	if err := validate.MinimumInt("version", "body", int64(*m.Version), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RecallSubmission) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RecallSubmission) UnmarshalBinary(b []byte) error {
	var res RecallSubmission
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *RecallSubmission) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// RecallSubmissionAttributes recall submission attributes
// swagger:model RecallSubmissionAttributes
type RecallSubmissionAttributes struct {

	// status
	Status RecallSubmissionStatus `json:"status,omitempty"`

	// Reason for submission failure if submission status is `delivery_failed`
	StatusReason string `json:"status_reason,omitempty"`

	// Date and time of the submission
	// Read Only: true
	// Format: date-time
	SubmissionDatetime *strfmt.DateTime `json:"submission_datetime,omitempty"`
}

func RecallSubmissionAttributesWithDefaults(defaults client.Defaults) *RecallSubmissionAttributes {
	return &RecallSubmissionAttributes{

		// TODO Status: RecallSubmissionStatus,

		StatusReason: defaults.GetString("RecallSubmissionAttributes", "status_reason"),

		SubmissionDatetime: defaults.GetStrfmtDateTimePtr("RecallSubmissionAttributes", "submission_datetime"),
	}
}

func (m *RecallSubmissionAttributes) WithStatus(status RecallSubmissionStatus) *RecallSubmissionAttributes {

	m.Status = status

	return m
}

func (m *RecallSubmissionAttributes) WithStatusReason(statusReason string) *RecallSubmissionAttributes {

	m.StatusReason = statusReason

	return m
}

func (m *RecallSubmissionAttributes) WithSubmissionDatetime(submissionDatetime strfmt.DateTime) *RecallSubmissionAttributes {

	m.SubmissionDatetime = &submissionDatetime

	return m
}

func (m *RecallSubmissionAttributes) WithoutSubmissionDatetime() *RecallSubmissionAttributes {
	m.SubmissionDatetime = nil
	return m
}

// Validate validates this recall submission attributes
func (m *RecallSubmissionAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubmissionDatetime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecallSubmissionAttributes) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if err := m.Status.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("attributes" + "." + "status")
		}
		return err
	}

	return nil
}

func (m *RecallSubmissionAttributes) validateSubmissionDatetime(formats strfmt.Registry) error {

	if swag.IsZero(m.SubmissionDatetime) { // not required
		return nil
	}

	if err := validate.FormatOf("attributes"+"."+"submission_datetime", "body", "date-time", m.SubmissionDatetime.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RecallSubmissionAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RecallSubmissionAttributes) UnmarshalBinary(b []byte) error {
	var res RecallSubmissionAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *RecallSubmissionAttributes) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
