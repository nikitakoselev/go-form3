// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// PaymentRecallSubmissionResourceType payment recall submission resource type
// swagger:model PaymentRecallSubmissionResourceType
type PaymentRecallSubmissionResourceType string

const (

	// PaymentRecallSubmissionResourceTypeRecallSubmissions captures enum value "recall_submissions"
	PaymentRecallSubmissionResourceTypeRecallSubmissions PaymentRecallSubmissionResourceType = "recall_submissions"
)

// for schema
var paymentRecallSubmissionResourceTypeEnum []interface{}

func init() {
	var res []PaymentRecallSubmissionResourceType
	if err := json.Unmarshal([]byte(`["recall_submissions"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		paymentRecallSubmissionResourceTypeEnum = append(paymentRecallSubmissionResourceTypeEnum, v)
	}
}

func (m PaymentRecallSubmissionResourceType) validatePaymentRecallSubmissionResourceTypeEnum(path, location string, value PaymentRecallSubmissionResourceType) error {
	if err := validate.Enum(path, location, value, paymentRecallSubmissionResourceTypeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this payment recall submission resource type
func (m PaymentRecallSubmissionResourceType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validatePaymentRecallSubmissionResourceTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PaymentRecallSubmissionResourceType) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}