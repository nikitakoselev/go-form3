// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/form3tech-oss/go-form3/v4/pkg/client"
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DirectDebitReversalAdmission direct debit reversal admission
// swagger:model DirectDebitReversalAdmission
type DirectDebitReversalAdmission struct {

	// attributes
	// Required: true
	Attributes *DirectDebitReversalAdmissionAttributes `json:"attributes"`

	// created on
	// Format: date-time
	CreatedOn *strfmt.DateTime `json:"created_on,omitempty"`

	// id
	// Required: true
	// Format: uuid
	ID *strfmt.UUID `json:"id"`

	// modified on
	// Format: date-time
	ModifiedOn *strfmt.DateTime `json:"modified_on,omitempty"`

	// organisation id
	// Required: true
	// Format: uuid
	OrganisationID *strfmt.UUID `json:"organisation_id"`

	// relationships
	Relationships *DirectDebitReversalAdmissionRelationships `json:"relationships,omitempty"`

	// type
	// Pattern: ^[A-Za-z_]*$
	Type string `json:"type,omitempty"`

	// version
	// Minimum: 0
	Version *int64 `json:"version,omitempty"`
}

func DirectDebitReversalAdmissionWithDefaults(defaults client.Defaults) *DirectDebitReversalAdmission {
	return &DirectDebitReversalAdmission{

		Attributes: DirectDebitReversalAdmissionAttributesWithDefaults(defaults),

		CreatedOn: defaults.GetStrfmtDateTimePtr("DirectDebitReversalAdmission", "created_on"),

		ID: defaults.GetStrfmtUUIDPtr("DirectDebitReversalAdmission", "id"),

		ModifiedOn: defaults.GetStrfmtDateTimePtr("DirectDebitReversalAdmission", "modified_on"),

		OrganisationID: defaults.GetStrfmtUUIDPtr("DirectDebitReversalAdmission", "organisation_id"),

		Relationships: DirectDebitReversalAdmissionRelationshipsWithDefaults(defaults),

		Type: defaults.GetString("DirectDebitReversalAdmission", "type"),

		Version: defaults.GetInt64Ptr("DirectDebitReversalAdmission", "version"),
	}
}

func (m *DirectDebitReversalAdmission) WithAttributes(attributes DirectDebitReversalAdmissionAttributes) *DirectDebitReversalAdmission {

	m.Attributes = &attributes

	return m
}

func (m *DirectDebitReversalAdmission) WithoutAttributes() *DirectDebitReversalAdmission {
	m.Attributes = nil
	return m
}

func (m *DirectDebitReversalAdmission) WithCreatedOn(createdOn strfmt.DateTime) *DirectDebitReversalAdmission {

	m.CreatedOn = &createdOn

	return m
}

func (m *DirectDebitReversalAdmission) WithoutCreatedOn() *DirectDebitReversalAdmission {
	m.CreatedOn = nil
	return m
}

func (m *DirectDebitReversalAdmission) WithID(id strfmt.UUID) *DirectDebitReversalAdmission {

	m.ID = &id

	return m
}

func (m *DirectDebitReversalAdmission) WithoutID() *DirectDebitReversalAdmission {
	m.ID = nil
	return m
}

func (m *DirectDebitReversalAdmission) WithModifiedOn(modifiedOn strfmt.DateTime) *DirectDebitReversalAdmission {

	m.ModifiedOn = &modifiedOn

	return m
}

func (m *DirectDebitReversalAdmission) WithoutModifiedOn() *DirectDebitReversalAdmission {
	m.ModifiedOn = nil
	return m
}

func (m *DirectDebitReversalAdmission) WithOrganisationID(organisationID strfmt.UUID) *DirectDebitReversalAdmission {

	m.OrganisationID = &organisationID

	return m
}

func (m *DirectDebitReversalAdmission) WithoutOrganisationID() *DirectDebitReversalAdmission {
	m.OrganisationID = nil
	return m
}

func (m *DirectDebitReversalAdmission) WithRelationships(relationships DirectDebitReversalAdmissionRelationships) *DirectDebitReversalAdmission {

	m.Relationships = &relationships

	return m
}

func (m *DirectDebitReversalAdmission) WithoutRelationships() *DirectDebitReversalAdmission {
	m.Relationships = nil
	return m
}

func (m *DirectDebitReversalAdmission) WithType(typeVar string) *DirectDebitReversalAdmission {

	m.Type = typeVar

	return m
}

func (m *DirectDebitReversalAdmission) WithVersion(version int64) *DirectDebitReversalAdmission {

	m.Version = &version

	return m
}

func (m *DirectDebitReversalAdmission) WithoutVersion() *DirectDebitReversalAdmission {
	m.Version = nil
	return m
}

// Validate validates this direct debit reversal admission
func (m *DirectDebitReversalAdmission) Validate(formats strfmt.Registry) error {
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

func (m *DirectDebitReversalAdmission) validateAttributes(formats strfmt.Registry) error {

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

func (m *DirectDebitReversalAdmission) validateCreatedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("created_on", "body", "date-time", m.CreatedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DirectDebitReversalAdmission) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DirectDebitReversalAdmission) validateModifiedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.ModifiedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("modified_on", "body", "date-time", m.ModifiedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DirectDebitReversalAdmission) validateOrganisationID(formats strfmt.Registry) error {

	if err := validate.Required("organisation_id", "body", m.OrganisationID); err != nil {
		return err
	}

	if err := validate.FormatOf("organisation_id", "body", "uuid", m.OrganisationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DirectDebitReversalAdmission) validateRelationships(formats strfmt.Registry) error {

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

func (m *DirectDebitReversalAdmission) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := validate.Pattern("type", "body", string(m.Type), `^[A-Za-z_]*$`); err != nil {
		return err
	}

	return nil
}

func (m *DirectDebitReversalAdmission) validateVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.Version) { // not required
		return nil
	}

	if err := validate.MinimumInt("version", "body", int64(*m.Version), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DirectDebitReversalAdmission) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DirectDebitReversalAdmission) UnmarshalBinary(b []byte) error {
	var res DirectDebitReversalAdmission
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *DirectDebitReversalAdmission) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// DirectDebitReversalAdmissionAttributes direct debit reversal admission attributes
// swagger:model DirectDebitReversalAdmissionAttributes
type DirectDebitReversalAdmissionAttributes struct {

	// scheme status code
	SchemeStatusCode string `json:"scheme_status_code,omitempty"`

	// scheme status code description
	SchemeStatusCodeDescription string `json:"scheme_status_code_description,omitempty"`

	// source gateway
	SourceGateway string `json:"source_gateway,omitempty"`

	// status
	Status DirectDebitReversalAdmissionStatus `json:"status,omitempty"`
}

func DirectDebitReversalAdmissionAttributesWithDefaults(defaults client.Defaults) *DirectDebitReversalAdmissionAttributes {
	return &DirectDebitReversalAdmissionAttributes{

		SchemeStatusCode: defaults.GetString("DirectDebitReversalAdmissionAttributes", "scheme_status_code"),

		SchemeStatusCodeDescription: defaults.GetString("DirectDebitReversalAdmissionAttributes", "scheme_status_code_description"),

		SourceGateway: defaults.GetString("DirectDebitReversalAdmissionAttributes", "source_gateway"),

		// TODO Status: DirectDebitReversalAdmissionStatus,

	}
}

func (m *DirectDebitReversalAdmissionAttributes) WithSchemeStatusCode(schemeStatusCode string) *DirectDebitReversalAdmissionAttributes {

	m.SchemeStatusCode = schemeStatusCode

	return m
}

func (m *DirectDebitReversalAdmissionAttributes) WithSchemeStatusCodeDescription(schemeStatusCodeDescription string) *DirectDebitReversalAdmissionAttributes {

	m.SchemeStatusCodeDescription = schemeStatusCodeDescription

	return m
}

func (m *DirectDebitReversalAdmissionAttributes) WithSourceGateway(sourceGateway string) *DirectDebitReversalAdmissionAttributes {

	m.SourceGateway = sourceGateway

	return m
}

func (m *DirectDebitReversalAdmissionAttributes) WithStatus(status DirectDebitReversalAdmissionStatus) *DirectDebitReversalAdmissionAttributes {

	m.Status = status

	return m
}

// Validate validates this direct debit reversal admission attributes
func (m *DirectDebitReversalAdmissionAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DirectDebitReversalAdmissionAttributes) validateStatus(formats strfmt.Registry) error {

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
func (m *DirectDebitReversalAdmissionAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DirectDebitReversalAdmissionAttributes) UnmarshalBinary(b []byte) error {
	var res DirectDebitReversalAdmissionAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *DirectDebitReversalAdmissionAttributes) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// DirectDebitReversalAdmissionRelationships direct debit reversal admission relationships
// swagger:model DirectDebitReversalAdmissionRelationships
type DirectDebitReversalAdmissionRelationships struct {

	// direct debit
	DirectDebit *DirectDebitReversalAdmissionRelationshipsDirectDebit `json:"direct_debit,omitempty"`

	// direct debit reversal
	DirectDebitReversal *DirectDebitReversalAdmissionRelationshipsDirectDebitReversal `json:"direct_debit_reversal,omitempty"`
}

func DirectDebitReversalAdmissionRelationshipsWithDefaults(defaults client.Defaults) *DirectDebitReversalAdmissionRelationships {
	return &DirectDebitReversalAdmissionRelationships{

		DirectDebit: DirectDebitReversalAdmissionRelationshipsDirectDebitWithDefaults(defaults),

		DirectDebitReversal: DirectDebitReversalAdmissionRelationshipsDirectDebitReversalWithDefaults(defaults),
	}
}

func (m *DirectDebitReversalAdmissionRelationships) WithDirectDebit(directDebit DirectDebitReversalAdmissionRelationshipsDirectDebit) *DirectDebitReversalAdmissionRelationships {

	m.DirectDebit = &directDebit

	return m
}

func (m *DirectDebitReversalAdmissionRelationships) WithoutDirectDebit() *DirectDebitReversalAdmissionRelationships {
	m.DirectDebit = nil
	return m
}

func (m *DirectDebitReversalAdmissionRelationships) WithDirectDebitReversal(directDebitReversal DirectDebitReversalAdmissionRelationshipsDirectDebitReversal) *DirectDebitReversalAdmissionRelationships {

	m.DirectDebitReversal = &directDebitReversal

	return m
}

func (m *DirectDebitReversalAdmissionRelationships) WithoutDirectDebitReversal() *DirectDebitReversalAdmissionRelationships {
	m.DirectDebitReversal = nil
	return m
}

// Validate validates this direct debit reversal admission relationships
func (m *DirectDebitReversalAdmissionRelationships) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDirectDebit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDirectDebitReversal(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DirectDebitReversalAdmissionRelationships) validateDirectDebit(formats strfmt.Registry) error {

	if swag.IsZero(m.DirectDebit) { // not required
		return nil
	}

	if m.DirectDebit != nil {
		if err := m.DirectDebit.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("relationships" + "." + "direct_debit")
			}
			return err
		}
	}

	return nil
}

func (m *DirectDebitReversalAdmissionRelationships) validateDirectDebitReversal(formats strfmt.Registry) error {

	if swag.IsZero(m.DirectDebitReversal) { // not required
		return nil
	}

	if m.DirectDebitReversal != nil {
		if err := m.DirectDebitReversal.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("relationships" + "." + "direct_debit_reversal")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DirectDebitReversalAdmissionRelationships) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DirectDebitReversalAdmissionRelationships) UnmarshalBinary(b []byte) error {
	var res DirectDebitReversalAdmissionRelationships
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *DirectDebitReversalAdmissionRelationships) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// DirectDebitReversalAdmissionRelationshipsDirectDebit direct debit reversal admission relationships direct debit
// swagger:model DirectDebitReversalAdmissionRelationshipsDirectDebit
type DirectDebitReversalAdmissionRelationshipsDirectDebit struct {

	// data
	Data []*DirectDebit `json:"data"`
}

func DirectDebitReversalAdmissionRelationshipsDirectDebitWithDefaults(defaults client.Defaults) *DirectDebitReversalAdmissionRelationshipsDirectDebit {
	return &DirectDebitReversalAdmissionRelationshipsDirectDebit{

		Data: make([]*DirectDebit, 0),
	}
}

func (m *DirectDebitReversalAdmissionRelationshipsDirectDebit) WithData(data []*DirectDebit) *DirectDebitReversalAdmissionRelationshipsDirectDebit {

	m.Data = data

	return m
}

// Validate validates this direct debit reversal admission relationships direct debit
func (m *DirectDebitReversalAdmissionRelationshipsDirectDebit) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DirectDebitReversalAdmissionRelationshipsDirectDebit) validateData(formats strfmt.Registry) error {

	if swag.IsZero(m.Data) { // not required
		return nil
	}

	for i := 0; i < len(m.Data); i++ {
		if swag.IsZero(m.Data[i]) { // not required
			continue
		}

		if m.Data[i] != nil {
			if err := m.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("relationships" + "." + "direct_debit" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DirectDebitReversalAdmissionRelationshipsDirectDebit) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DirectDebitReversalAdmissionRelationshipsDirectDebit) UnmarshalBinary(b []byte) error {
	var res DirectDebitReversalAdmissionRelationshipsDirectDebit
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *DirectDebitReversalAdmissionRelationshipsDirectDebit) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// DirectDebitReversalAdmissionRelationshipsDirectDebitReversal direct debit reversal admission relationships direct debit reversal
// swagger:model DirectDebitReversalAdmissionRelationshipsDirectDebitReversal
type DirectDebitReversalAdmissionRelationshipsDirectDebitReversal struct {

	// data
	Data []*DirectDebitReversal `json:"data"`
}

func DirectDebitReversalAdmissionRelationshipsDirectDebitReversalWithDefaults(defaults client.Defaults) *DirectDebitReversalAdmissionRelationshipsDirectDebitReversal {
	return &DirectDebitReversalAdmissionRelationshipsDirectDebitReversal{

		Data: make([]*DirectDebitReversal, 0),
	}
}

func (m *DirectDebitReversalAdmissionRelationshipsDirectDebitReversal) WithData(data []*DirectDebitReversal) *DirectDebitReversalAdmissionRelationshipsDirectDebitReversal {

	m.Data = data

	return m
}

// Validate validates this direct debit reversal admission relationships direct debit reversal
func (m *DirectDebitReversalAdmissionRelationshipsDirectDebitReversal) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DirectDebitReversalAdmissionRelationshipsDirectDebitReversal) validateData(formats strfmt.Registry) error {

	if swag.IsZero(m.Data) { // not required
		return nil
	}

	for i := 0; i < len(m.Data); i++ {
		if swag.IsZero(m.Data[i]) { // not required
			continue
		}

		if m.Data[i] != nil {
			if err := m.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("relationships" + "." + "direct_debit_reversal" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DirectDebitReversalAdmissionRelationshipsDirectDebitReversal) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DirectDebitReversalAdmissionRelationshipsDirectDebitReversal) UnmarshalBinary(b []byte) error {
	var res DirectDebitReversalAdmissionRelationshipsDirectDebitReversal
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *DirectDebitReversalAdmissionRelationshipsDirectDebitReversal) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
