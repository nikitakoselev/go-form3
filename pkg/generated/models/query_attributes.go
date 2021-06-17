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

// QueryAttributes query attributes
// swagger:model QueryAttributes
type QueryAttributes struct {

	// auto handled
	// Required: true
	AutoHandled *bool `json:"auto_handled"`

	// message id
	MessageID string `json:"message_id,omitempty"`

	// processing date
	// Format: date
	ProcessingDate *strfmt.Date `json:"processing_date,omitempty"`

	// query type
	// Required: true
	// Enum: [claim_non_receipt modify_payment]
	QueryType *string `json:"query_type"`

	// scheme transaction id
	SchemeTransactionID string `json:"scheme_transaction_id,omitempty"`

	// status
	// Enum: [pending closed]
	Status string `json:"status,omitempty"`
}

func QueryAttributesWithDefaults(defaults client.Defaults) *QueryAttributes {
	return &QueryAttributes{

		AutoHandled: defaults.GetBoolPtr("QueryAttributes", "auto_handled"),

		MessageID: defaults.GetString("QueryAttributes", "message_id"),

		ProcessingDate: defaults.GetStrfmtDatePtr("QueryAttributes", "processing_date"),

		QueryType: defaults.GetStringPtr("QueryAttributes", "query_type"),

		SchemeTransactionID: defaults.GetString("QueryAttributes", "scheme_transaction_id"),

		Status: defaults.GetString("QueryAttributes", "status"),
	}
}

func (m *QueryAttributes) WithAutoHandled(autoHandled bool) *QueryAttributes {

	m.AutoHandled = &autoHandled

	return m
}

func (m *QueryAttributes) WithoutAutoHandled() *QueryAttributes {
	m.AutoHandled = nil
	return m
}

func (m *QueryAttributes) WithMessageID(messageID string) *QueryAttributes {

	m.MessageID = messageID

	return m
}

func (m *QueryAttributes) WithProcessingDate(processingDate strfmt.Date) *QueryAttributes {

	m.ProcessingDate = &processingDate

	return m
}

func (m *QueryAttributes) WithoutProcessingDate() *QueryAttributes {
	m.ProcessingDate = nil
	return m
}

func (m *QueryAttributes) WithQueryType(queryType string) *QueryAttributes {

	m.QueryType = &queryType

	return m
}

func (m *QueryAttributes) WithoutQueryType() *QueryAttributes {
	m.QueryType = nil
	return m
}

func (m *QueryAttributes) WithSchemeTransactionID(schemeTransactionID string) *QueryAttributes {

	m.SchemeTransactionID = schemeTransactionID

	return m
}

func (m *QueryAttributes) WithStatus(status string) *QueryAttributes {

	m.Status = status

	return m
}

// Validate validates this query attributes
func (m *QueryAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAutoHandled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProcessingDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQueryType(formats); err != nil {
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

func (m *QueryAttributes) validateAutoHandled(formats strfmt.Registry) error {

	if err := validate.Required("auto_handled", "body", m.AutoHandled); err != nil {
		return err
	}

	return nil
}

func (m *QueryAttributes) validateProcessingDate(formats strfmt.Registry) error {

	if swag.IsZero(m.ProcessingDate) { // not required
		return nil
	}

	if err := validate.FormatOf("processing_date", "body", "date", m.ProcessingDate.String(), formats); err != nil {
		return err
	}

	return nil
}

var queryAttributesTypeQueryTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["claim_non_receipt","modify_payment"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		queryAttributesTypeQueryTypePropEnum = append(queryAttributesTypeQueryTypePropEnum, v)
	}
}

const (

	// QueryAttributesQueryTypeClaimNonReceipt captures enum value "claim_non_receipt"
	QueryAttributesQueryTypeClaimNonReceipt string = "claim_non_receipt"

	// QueryAttributesQueryTypeModifyPayment captures enum value "modify_payment"
	QueryAttributesQueryTypeModifyPayment string = "modify_payment"
)

// prop value enum
func (m *QueryAttributes) validateQueryTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, queryAttributesTypeQueryTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *QueryAttributes) validateQueryType(formats strfmt.Registry) error {

	if err := validate.Required("query_type", "body", m.QueryType); err != nil {
		return err
	}

	// value enum
	if err := m.validateQueryTypeEnum("query_type", "body", *m.QueryType); err != nil {
		return err
	}

	return nil
}

var queryAttributesTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["pending","closed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		queryAttributesTypeStatusPropEnum = append(queryAttributesTypeStatusPropEnum, v)
	}
}

const (

	// QueryAttributesStatusPending captures enum value "pending"
	QueryAttributesStatusPending string = "pending"

	// QueryAttributesStatusClosed captures enum value "closed"
	QueryAttributesStatusClosed string = "closed"
)

// prop value enum
func (m *QueryAttributes) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, queryAttributesTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *QueryAttributes) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *QueryAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *QueryAttributes) UnmarshalBinary(b []byte) error {
	var res QueryAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *QueryAttributes) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
