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

// PaymentSubmission payment submission
// swagger:model PaymentSubmission
type PaymentSubmission struct {

	// attributes
	Attributes *PaymentSubmissionAttributes `json:"attributes,omitempty"`

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

	// Unique ID of the organisation this resource is created by
	// Required: true
	// Format: uuid
	OrganisationID *strfmt.UUID `json:"organisation_id"`

	// relationships
	Relationships *PaymentSubmissionRelationships `json:"relationships,omitempty"`

	// Name of the resource type
	// Pattern: ^[A-Za-z_]*$
	Type string `json:"type,omitempty"`

	// Version number
	// Minimum: 0
	Version *int64 `json:"version,omitempty"`
}

func PaymentSubmissionWithDefaults(defaults client.Defaults) *PaymentSubmission {
	return &PaymentSubmission{

		Attributes: PaymentSubmissionAttributesWithDefaults(defaults),

		CreatedOn: defaults.GetStrfmtDateTimePtr("PaymentSubmission", "created_on"),

		ID: defaults.GetStrfmtUUIDPtr("PaymentSubmission", "id"),

		ModifiedOn: defaults.GetStrfmtDateTimePtr("PaymentSubmission", "modified_on"),

		OrganisationID: defaults.GetStrfmtUUIDPtr("PaymentSubmission", "organisation_id"),

		Relationships: PaymentSubmissionRelationshipsWithDefaults(defaults),

		Type: defaults.GetString("PaymentSubmission", "type"),

		Version: defaults.GetInt64Ptr("PaymentSubmission", "version"),
	}
}

func (m *PaymentSubmission) WithAttributes(attributes PaymentSubmissionAttributes) *PaymentSubmission {

	m.Attributes = &attributes

	return m
}

func (m *PaymentSubmission) WithoutAttributes() *PaymentSubmission {
	m.Attributes = nil
	return m
}

func (m *PaymentSubmission) WithCreatedOn(createdOn strfmt.DateTime) *PaymentSubmission {

	m.CreatedOn = &createdOn

	return m
}

func (m *PaymentSubmission) WithoutCreatedOn() *PaymentSubmission {
	m.CreatedOn = nil
	return m
}

func (m *PaymentSubmission) WithID(id strfmt.UUID) *PaymentSubmission {

	m.ID = &id

	return m
}

func (m *PaymentSubmission) WithoutID() *PaymentSubmission {
	m.ID = nil
	return m
}

func (m *PaymentSubmission) WithModifiedOn(modifiedOn strfmt.DateTime) *PaymentSubmission {

	m.ModifiedOn = &modifiedOn

	return m
}

func (m *PaymentSubmission) WithoutModifiedOn() *PaymentSubmission {
	m.ModifiedOn = nil
	return m
}

func (m *PaymentSubmission) WithOrganisationID(organisationID strfmt.UUID) *PaymentSubmission {

	m.OrganisationID = &organisationID

	return m
}

func (m *PaymentSubmission) WithoutOrganisationID() *PaymentSubmission {
	m.OrganisationID = nil
	return m
}

func (m *PaymentSubmission) WithRelationships(relationships PaymentSubmissionRelationships) *PaymentSubmission {

	m.Relationships = &relationships

	return m
}

func (m *PaymentSubmission) WithoutRelationships() *PaymentSubmission {
	m.Relationships = nil
	return m
}

func (m *PaymentSubmission) WithType(typeVar string) *PaymentSubmission {

	m.Type = typeVar

	return m
}

func (m *PaymentSubmission) WithVersion(version int64) *PaymentSubmission {

	m.Version = &version

	return m
}

func (m *PaymentSubmission) WithoutVersion() *PaymentSubmission {
	m.Version = nil
	return m
}

// Validate validates this payment submission
func (m *PaymentSubmission) Validate(formats strfmt.Registry) error {
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

func (m *PaymentSubmission) validateAttributes(formats strfmt.Registry) error {

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

func (m *PaymentSubmission) validateCreatedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("created_on", "body", "date-time", m.CreatedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PaymentSubmission) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PaymentSubmission) validateModifiedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.ModifiedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("modified_on", "body", "date-time", m.ModifiedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PaymentSubmission) validateOrganisationID(formats strfmt.Registry) error {

	if err := validate.Required("organisation_id", "body", m.OrganisationID); err != nil {
		return err
	}

	if err := validate.FormatOf("organisation_id", "body", "uuid", m.OrganisationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PaymentSubmission) validateRelationships(formats strfmt.Registry) error {

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

func (m *PaymentSubmission) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := validate.Pattern("type", "body", string(m.Type), `^[A-Za-z_]*$`); err != nil {
		return err
	}

	return nil
}

func (m *PaymentSubmission) validateVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.Version) { // not required
		return nil
	}

	if err := validate.MinimumInt("version", "body", int64(*m.Version), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PaymentSubmission) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentSubmission) UnmarshalBinary(b []byte) error {
	var res PaymentSubmission
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *PaymentSubmission) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// PaymentSubmissionAttributes payment submission attributes
// swagger:model PaymentSubmissionAttributes
type PaymentSubmissionAttributes struct {

	// Time a payment was released from being held due to a limit breach
	// Read Only: true
	// Format: date-time
	LimitBreachEndDatetime *strfmt.DateTime `json:"limit_breach_end_datetime,omitempty"`

	// Start time a payment was held due to a limit breach
	// Read Only: true
	// Format: date-time
	LimitBreachStartDatetime *strfmt.DateTime `json:"limit_breach_start_datetime,omitempty"`

	// Details of the account to which funds are redirected (if applicable)
	RedirectedAccountNumber string `json:"redirected_account_number,omitempty"`

	// Details of the bank to which funds are redirected (if applicable)
	RedirectedBankID string `json:"redirected_bank_id,omitempty"`

	// Scheme-specific status (if submission has been submitted to a scheme)
	SchemeStatusCode string `json:"scheme_status_code,omitempty"`

	// [Description](http://api-docs.form3.tech/api.html#enumerations-scheme-status-codes-for-bacs) of `scheme_status_code`
	SchemeStatusCodeDescription string `json:"scheme_status_code_description,omitempty"`

	// Cycle in which the payment will be settled
	// Minimum: 0
	SettlementCycle *int64 `json:"settlement_cycle,omitempty"`

	// Date that the payment will be settled
	// Format: date
	SettlementDate *strfmt.Date `json:"settlement_date,omitempty"`

	// status
	Status PaymentSubmissionStatus `json:"status,omitempty"`

	// Description of the submission status
	StatusReason string `json:"status_reason,omitempty"`

	// Date of the submission
	// Read Only: true
	// Format: date-time
	SubmissionDatetime strfmt.DateTime `json:"submission_datetime,omitempty"`

	// Time the request was received by Form3. Used to compute the total transaction time of a payment.
	// Read Only: true
	// Format: date-time
	TransactionStartDatetime strfmt.DateTime `json:"transaction_start_datetime,omitempty"`
}

func PaymentSubmissionAttributesWithDefaults(defaults client.Defaults) *PaymentSubmissionAttributes {
	return &PaymentSubmissionAttributes{

		LimitBreachEndDatetime: defaults.GetStrfmtDateTimePtr("PaymentSubmissionAttributes", "limit_breach_end_datetime"),

		LimitBreachStartDatetime: defaults.GetStrfmtDateTimePtr("PaymentSubmissionAttributes", "limit_breach_start_datetime"),

		RedirectedAccountNumber: defaults.GetString("PaymentSubmissionAttributes", "redirected_account_number"),

		RedirectedBankID: defaults.GetString("PaymentSubmissionAttributes", "redirected_bank_id"),

		SchemeStatusCode: defaults.GetString("PaymentSubmissionAttributes", "scheme_status_code"),

		SchemeStatusCodeDescription: defaults.GetString("PaymentSubmissionAttributes", "scheme_status_code_description"),

		SettlementCycle: defaults.GetInt64Ptr("PaymentSubmissionAttributes", "settlement_cycle"),

		SettlementDate: defaults.GetStrfmtDatePtr("PaymentSubmissionAttributes", "settlement_date"),

		// TODO Status: PaymentSubmissionStatus,

		StatusReason: defaults.GetString("PaymentSubmissionAttributes", "status_reason"),

		SubmissionDatetime: defaults.GetStrfmtDateTime("PaymentSubmissionAttributes", "submission_datetime"),

		TransactionStartDatetime: defaults.GetStrfmtDateTime("PaymentSubmissionAttributes", "transaction_start_datetime"),
	}
}

func (m *PaymentSubmissionAttributes) WithLimitBreachEndDatetime(limitBreachEndDatetime strfmt.DateTime) *PaymentSubmissionAttributes {

	m.LimitBreachEndDatetime = &limitBreachEndDatetime

	return m
}

func (m *PaymentSubmissionAttributes) WithoutLimitBreachEndDatetime() *PaymentSubmissionAttributes {
	m.LimitBreachEndDatetime = nil
	return m
}

func (m *PaymentSubmissionAttributes) WithLimitBreachStartDatetime(limitBreachStartDatetime strfmt.DateTime) *PaymentSubmissionAttributes {

	m.LimitBreachStartDatetime = &limitBreachStartDatetime

	return m
}

func (m *PaymentSubmissionAttributes) WithoutLimitBreachStartDatetime() *PaymentSubmissionAttributes {
	m.LimitBreachStartDatetime = nil
	return m
}

func (m *PaymentSubmissionAttributes) WithRedirectedAccountNumber(redirectedAccountNumber string) *PaymentSubmissionAttributes {

	m.RedirectedAccountNumber = redirectedAccountNumber

	return m
}

func (m *PaymentSubmissionAttributes) WithRedirectedBankID(redirectedBankID string) *PaymentSubmissionAttributes {

	m.RedirectedBankID = redirectedBankID

	return m
}

func (m *PaymentSubmissionAttributes) WithSchemeStatusCode(schemeStatusCode string) *PaymentSubmissionAttributes {

	m.SchemeStatusCode = schemeStatusCode

	return m
}

func (m *PaymentSubmissionAttributes) WithSchemeStatusCodeDescription(schemeStatusCodeDescription string) *PaymentSubmissionAttributes {

	m.SchemeStatusCodeDescription = schemeStatusCodeDescription

	return m
}

func (m *PaymentSubmissionAttributes) WithSettlementCycle(settlementCycle int64) *PaymentSubmissionAttributes {

	m.SettlementCycle = &settlementCycle

	return m
}

func (m *PaymentSubmissionAttributes) WithoutSettlementCycle() *PaymentSubmissionAttributes {
	m.SettlementCycle = nil
	return m
}

func (m *PaymentSubmissionAttributes) WithSettlementDate(settlementDate strfmt.Date) *PaymentSubmissionAttributes {

	m.SettlementDate = &settlementDate

	return m
}

func (m *PaymentSubmissionAttributes) WithoutSettlementDate() *PaymentSubmissionAttributes {
	m.SettlementDate = nil
	return m
}

func (m *PaymentSubmissionAttributes) WithStatus(status PaymentSubmissionStatus) *PaymentSubmissionAttributes {

	m.Status = status

	return m
}

func (m *PaymentSubmissionAttributes) WithStatusReason(statusReason string) *PaymentSubmissionAttributes {

	m.StatusReason = statusReason

	return m
}

func (m *PaymentSubmissionAttributes) WithSubmissionDatetime(submissionDatetime strfmt.DateTime) *PaymentSubmissionAttributes {

	m.SubmissionDatetime = submissionDatetime

	return m
}

func (m *PaymentSubmissionAttributes) WithTransactionStartDatetime(transactionStartDatetime strfmt.DateTime) *PaymentSubmissionAttributes {

	m.TransactionStartDatetime = transactionStartDatetime

	return m
}

// Validate validates this payment submission attributes
func (m *PaymentSubmissionAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLimitBreachEndDatetime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLimitBreachStartDatetime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSettlementCycle(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSettlementDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubmissionDatetime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTransactionStartDatetime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PaymentSubmissionAttributes) validateLimitBreachEndDatetime(formats strfmt.Registry) error {

	if swag.IsZero(m.LimitBreachEndDatetime) { // not required
		return nil
	}

	if err := validate.FormatOf("attributes"+"."+"limit_breach_end_datetime", "body", "date-time", m.LimitBreachEndDatetime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PaymentSubmissionAttributes) validateLimitBreachStartDatetime(formats strfmt.Registry) error {

	if swag.IsZero(m.LimitBreachStartDatetime) { // not required
		return nil
	}

	if err := validate.FormatOf("attributes"+"."+"limit_breach_start_datetime", "body", "date-time", m.LimitBreachStartDatetime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PaymentSubmissionAttributes) validateSettlementCycle(formats strfmt.Registry) error {

	if swag.IsZero(m.SettlementCycle) { // not required
		return nil
	}

	if err := validate.MinimumInt("attributes"+"."+"settlement_cycle", "body", int64(*m.SettlementCycle), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *PaymentSubmissionAttributes) validateSettlementDate(formats strfmt.Registry) error {

	if swag.IsZero(m.SettlementDate) { // not required
		return nil
	}

	if err := validate.FormatOf("attributes"+"."+"settlement_date", "body", "date", m.SettlementDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PaymentSubmissionAttributes) validateStatus(formats strfmt.Registry) error {

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

func (m *PaymentSubmissionAttributes) validateSubmissionDatetime(formats strfmt.Registry) error {

	if swag.IsZero(m.SubmissionDatetime) { // not required
		return nil
	}

	if err := validate.FormatOf("attributes"+"."+"submission_datetime", "body", "date-time", m.SubmissionDatetime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PaymentSubmissionAttributes) validateTransactionStartDatetime(formats strfmt.Registry) error {

	if swag.IsZero(m.TransactionStartDatetime) { // not required
		return nil
	}

	if err := validate.FormatOf("attributes"+"."+"transaction_start_datetime", "body", "date-time", m.TransactionStartDatetime.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PaymentSubmissionAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentSubmissionAttributes) UnmarshalBinary(b []byte) error {
	var res PaymentSubmissionAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *PaymentSubmissionAttributes) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
