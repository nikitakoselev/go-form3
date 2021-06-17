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

// RecallDecisionAdmission recall decision admission
// swagger:model RecallDecisionAdmission
type RecallDecisionAdmission struct {

	// attributes
	Attributes *RecallDecisionAdmissionAttributes `json:"attributes,omitempty"`

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
	Relationships *RecallDecisionAdmissionRelationships `json:"relationships,omitempty"`

	// Name of the resource type
	// Pattern: ^[A-Za-z_]*$
	Type string `json:"type,omitempty"`

	// Version number
	// Minimum: 0
	Version *int64 `json:"version,omitempty"`
}

func RecallDecisionAdmissionWithDefaults(defaults client.Defaults) *RecallDecisionAdmission {
	return &RecallDecisionAdmission{

		Attributes: RecallDecisionAdmissionAttributesWithDefaults(defaults),

		CreatedOn: defaults.GetStrfmtDateTimePtr("RecallDecisionAdmission", "created_on"),

		ID: defaults.GetStrfmtUUIDPtr("RecallDecisionAdmission", "id"),

		ModifiedOn: defaults.GetStrfmtDateTimePtr("RecallDecisionAdmission", "modified_on"),

		OrganisationID: defaults.GetStrfmtUUIDPtr("RecallDecisionAdmission", "organisation_id"),

		Relationships: RecallDecisionAdmissionRelationshipsWithDefaults(defaults),

		Type: defaults.GetString("RecallDecisionAdmission", "type"),

		Version: defaults.GetInt64Ptr("RecallDecisionAdmission", "version"),
	}
}

func (m *RecallDecisionAdmission) WithAttributes(attributes RecallDecisionAdmissionAttributes) *RecallDecisionAdmission {

	m.Attributes = &attributes

	return m
}

func (m *RecallDecisionAdmission) WithoutAttributes() *RecallDecisionAdmission {
	m.Attributes = nil
	return m
}

func (m *RecallDecisionAdmission) WithCreatedOn(createdOn strfmt.DateTime) *RecallDecisionAdmission {

	m.CreatedOn = &createdOn

	return m
}

func (m *RecallDecisionAdmission) WithoutCreatedOn() *RecallDecisionAdmission {
	m.CreatedOn = nil
	return m
}

func (m *RecallDecisionAdmission) WithID(id strfmt.UUID) *RecallDecisionAdmission {

	m.ID = &id

	return m
}

func (m *RecallDecisionAdmission) WithoutID() *RecallDecisionAdmission {
	m.ID = nil
	return m
}

func (m *RecallDecisionAdmission) WithModifiedOn(modifiedOn strfmt.DateTime) *RecallDecisionAdmission {

	m.ModifiedOn = &modifiedOn

	return m
}

func (m *RecallDecisionAdmission) WithoutModifiedOn() *RecallDecisionAdmission {
	m.ModifiedOn = nil
	return m
}

func (m *RecallDecisionAdmission) WithOrganisationID(organisationID strfmt.UUID) *RecallDecisionAdmission {

	m.OrganisationID = &organisationID

	return m
}

func (m *RecallDecisionAdmission) WithoutOrganisationID() *RecallDecisionAdmission {
	m.OrganisationID = nil
	return m
}

func (m *RecallDecisionAdmission) WithRelationships(relationships RecallDecisionAdmissionRelationships) *RecallDecisionAdmission {

	m.Relationships = &relationships

	return m
}

func (m *RecallDecisionAdmission) WithoutRelationships() *RecallDecisionAdmission {
	m.Relationships = nil
	return m
}

func (m *RecallDecisionAdmission) WithType(typeVar string) *RecallDecisionAdmission {

	m.Type = typeVar

	return m
}

func (m *RecallDecisionAdmission) WithVersion(version int64) *RecallDecisionAdmission {

	m.Version = &version

	return m
}

func (m *RecallDecisionAdmission) WithoutVersion() *RecallDecisionAdmission {
	m.Version = nil
	return m
}

// Validate validates this recall decision admission
func (m *RecallDecisionAdmission) Validate(formats strfmt.Registry) error {
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

func (m *RecallDecisionAdmission) validateAttributes(formats strfmt.Registry) error {

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

func (m *RecallDecisionAdmission) validateCreatedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("created_on", "body", "date-time", m.CreatedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *RecallDecisionAdmission) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *RecallDecisionAdmission) validateModifiedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.ModifiedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("modified_on", "body", "date-time", m.ModifiedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *RecallDecisionAdmission) validateOrganisationID(formats strfmt.Registry) error {

	if err := validate.Required("organisation_id", "body", m.OrganisationID); err != nil {
		return err
	}

	if err := validate.FormatOf("organisation_id", "body", "uuid", m.OrganisationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *RecallDecisionAdmission) validateRelationships(formats strfmt.Registry) error {

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

func (m *RecallDecisionAdmission) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := validate.Pattern("type", "body", string(m.Type), `^[A-Za-z_]*$`); err != nil {
		return err
	}

	return nil
}

func (m *RecallDecisionAdmission) validateVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.Version) { // not required
		return nil
	}

	if err := validate.MinimumInt("version", "body", int64(*m.Version), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RecallDecisionAdmission) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RecallDecisionAdmission) UnmarshalBinary(b []byte) error {
	var res RecallDecisionAdmission
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *RecallDecisionAdmission) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// RecallDecisionAdmissionAttributes recall decision admission attributes
// swagger:model RecallDecisionAdmissionAttributes
type RecallDecisionAdmissionAttributes struct {

	// Date and time the recall decision admission was created
	// Read Only: true
	// Format: date-time
	AdmissionDatetime *strfmt.DateTime `json:"admission_datetime,omitempty"`

	// source gateway
	SourceGateway string `json:"source_gateway,omitempty"`

	// status
	Status RecallDecisionAdmissionStatus `json:"status,omitempty"`

	// Human-readable reason for failure if status is failed.
	StatusReason string `json:"status_reason,omitempty"`
}

func RecallDecisionAdmissionAttributesWithDefaults(defaults client.Defaults) *RecallDecisionAdmissionAttributes {
	return &RecallDecisionAdmissionAttributes{

		AdmissionDatetime: defaults.GetStrfmtDateTimePtr("RecallDecisionAdmissionAttributes", "admission_datetime"),

		SourceGateway: defaults.GetString("RecallDecisionAdmissionAttributes", "source_gateway"),

		// TODO Status: RecallDecisionAdmissionStatus,

		StatusReason: defaults.GetString("RecallDecisionAdmissionAttributes", "status_reason"),
	}
}

func (m *RecallDecisionAdmissionAttributes) WithAdmissionDatetime(admissionDatetime strfmt.DateTime) *RecallDecisionAdmissionAttributes {

	m.AdmissionDatetime = &admissionDatetime

	return m
}

func (m *RecallDecisionAdmissionAttributes) WithoutAdmissionDatetime() *RecallDecisionAdmissionAttributes {
	m.AdmissionDatetime = nil
	return m
}

func (m *RecallDecisionAdmissionAttributes) WithSourceGateway(sourceGateway string) *RecallDecisionAdmissionAttributes {

	m.SourceGateway = sourceGateway

	return m
}

func (m *RecallDecisionAdmissionAttributes) WithStatus(status RecallDecisionAdmissionStatus) *RecallDecisionAdmissionAttributes {

	m.Status = status

	return m
}

func (m *RecallDecisionAdmissionAttributes) WithStatusReason(statusReason string) *RecallDecisionAdmissionAttributes {

	m.StatusReason = statusReason

	return m
}

// Validate validates this recall decision admission attributes
func (m *RecallDecisionAdmissionAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdmissionDatetime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecallDecisionAdmissionAttributes) validateAdmissionDatetime(formats strfmt.Registry) error {

	if swag.IsZero(m.AdmissionDatetime) { // not required
		return nil
	}

	if err := validate.FormatOf("attributes"+"."+"admission_datetime", "body", "date-time", m.AdmissionDatetime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *RecallDecisionAdmissionAttributes) validateStatus(formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *RecallDecisionAdmissionAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RecallDecisionAdmissionAttributes) UnmarshalBinary(b []byte) error {
	var res RecallDecisionAdmissionAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *RecallDecisionAdmissionAttributes) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
