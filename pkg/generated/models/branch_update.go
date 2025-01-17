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

// BranchUpdate branch update
// swagger:model BranchUpdate
type BranchUpdate struct {

	// attributes
	// Required: true
	Attributes *BranchUpdateAttributes `json:"attributes"`

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
	// Required: true
	// Minimum: 0
	Version *int64 `json:"version"`
}

func BranchUpdateWithDefaults(defaults client.Defaults) *BranchUpdate {
	return &BranchUpdate{

		Attributes: BranchUpdateAttributesWithDefaults(defaults),

		ID: defaults.GetStrfmtUUIDPtr("BranchUpdate", "id"),

		OrganisationID: defaults.GetStrfmtUUIDPtr("BranchUpdate", "organisation_id"),

		Type: defaults.GetString("BranchUpdate", "type"),

		Version: defaults.GetInt64Ptr("BranchUpdate", "version"),
	}
}

func (m *BranchUpdate) WithAttributes(attributes BranchUpdateAttributes) *BranchUpdate {

	m.Attributes = &attributes

	return m
}

func (m *BranchUpdate) WithoutAttributes() *BranchUpdate {
	m.Attributes = nil
	return m
}

func (m *BranchUpdate) WithID(id strfmt.UUID) *BranchUpdate {

	m.ID = &id

	return m
}

func (m *BranchUpdate) WithoutID() *BranchUpdate {
	m.ID = nil
	return m
}

func (m *BranchUpdate) WithOrganisationID(organisationID strfmt.UUID) *BranchUpdate {

	m.OrganisationID = &organisationID

	return m
}

func (m *BranchUpdate) WithoutOrganisationID() *BranchUpdate {
	m.OrganisationID = nil
	return m
}

func (m *BranchUpdate) WithType(typeVar string) *BranchUpdate {

	m.Type = typeVar

	return m
}

func (m *BranchUpdate) WithVersion(version int64) *BranchUpdate {

	m.Version = &version

	return m
}

func (m *BranchUpdate) WithoutVersion() *BranchUpdate {
	m.Version = nil
	return m
}

// Validate validates this branch update
func (m *BranchUpdate) Validate(formats strfmt.Registry) error {
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

func (m *BranchUpdate) validateAttributes(formats strfmt.Registry) error {

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

func (m *BranchUpdate) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *BranchUpdate) validateOrganisationID(formats strfmt.Registry) error {

	if err := validate.Required("organisation_id", "body", m.OrganisationID); err != nil {
		return err
	}

	if err := validate.FormatOf("organisation_id", "body", "uuid", m.OrganisationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *BranchUpdate) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := validate.Pattern("type", "body", string(m.Type), `^[A-Za-z_]*$`); err != nil {
		return err
	}

	return nil
}

func (m *BranchUpdate) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version); err != nil {
		return err
	}

	if err := validate.MinimumInt("version", "body", int64(*m.Version), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BranchUpdate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BranchUpdate) UnmarshalBinary(b []byte) error {
	var res BranchUpdate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *BranchUpdate) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// BranchUpdateAttributes branch update attributes
// swagger:model BranchUpdateAttributes
type BranchUpdateAttributes struct {

	// acceptance qualifier
	AcceptanceQualifier AcceptanceQualifier `json:"acceptance_qualifier,omitempty"`

	// reference mask
	ReferenceMask ReferenceMask `json:"reference_mask,omitempty"`

	// validation type
	ValidationType BranchValidationType `json:"validation_type,omitempty"`
}

func BranchUpdateAttributesWithDefaults(defaults client.Defaults) *BranchUpdateAttributes {
	return &BranchUpdateAttributes{

		// TODO AcceptanceQualifier: AcceptanceQualifier,

		// TODO ReferenceMask: ReferenceMask,

		// TODO ValidationType: BranchValidationType,

	}
}

func (m *BranchUpdateAttributes) WithAcceptanceQualifier(acceptanceQualifier AcceptanceQualifier) *BranchUpdateAttributes {

	m.AcceptanceQualifier = acceptanceQualifier

	return m
}

func (m *BranchUpdateAttributes) WithReferenceMask(referenceMask ReferenceMask) *BranchUpdateAttributes {

	m.ReferenceMask = referenceMask

	return m
}

func (m *BranchUpdateAttributes) WithValidationType(validationType BranchValidationType) *BranchUpdateAttributes {

	m.ValidationType = validationType

	return m
}

// Validate validates this branch update attributes
func (m *BranchUpdateAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAcceptanceQualifier(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReferenceMask(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValidationType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BranchUpdateAttributes) validateAcceptanceQualifier(formats strfmt.Registry) error {

	if swag.IsZero(m.AcceptanceQualifier) { // not required
		return nil
	}

	if err := m.AcceptanceQualifier.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("attributes" + "." + "acceptance_qualifier")
		}
		return err
	}

	return nil
}

func (m *BranchUpdateAttributes) validateReferenceMask(formats strfmt.Registry) error {

	if swag.IsZero(m.ReferenceMask) { // not required
		return nil
	}

	if err := m.ReferenceMask.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("attributes" + "." + "reference_mask")
		}
		return err
	}

	return nil
}

func (m *BranchUpdateAttributes) validateValidationType(formats strfmt.Registry) error {

	if swag.IsZero(m.ValidationType) { // not required
		return nil
	}

	if err := m.ValidationType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("attributes" + "." + "validation_type")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BranchUpdateAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BranchUpdateAttributes) UnmarshalBinary(b []byte) error {
	var res BranchUpdateAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *BranchUpdateAttributes) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
