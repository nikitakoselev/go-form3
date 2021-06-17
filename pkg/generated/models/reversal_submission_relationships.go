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
)

// ReversalSubmissionRelationships reversal submission relationships
// swagger:model ReversalSubmissionRelationships
type ReversalSubmissionRelationships struct {

	// payment
	Payment *RelationshipPayments `json:"payment,omitempty"`

	// reversal
	Reversal *RelationshipReversals `json:"reversal,omitempty"`

	// validations
	Validations *RelationshipLinks `json:"validations,omitempty"`
}

func ReversalSubmissionRelationshipsWithDefaults(defaults client.Defaults) *ReversalSubmissionRelationships {
	return &ReversalSubmissionRelationships{

		Payment: RelationshipPaymentsWithDefaults(defaults),

		Reversal: RelationshipReversalsWithDefaults(defaults),

		Validations: RelationshipLinksWithDefaults(defaults),
	}
}

func (m *ReversalSubmissionRelationships) WithPayment(payment RelationshipPayments) *ReversalSubmissionRelationships {

	m.Payment = &payment

	return m
}

func (m *ReversalSubmissionRelationships) WithoutPayment() *ReversalSubmissionRelationships {
	m.Payment = nil
	return m
}

func (m *ReversalSubmissionRelationships) WithReversal(reversal RelationshipReversals) *ReversalSubmissionRelationships {

	m.Reversal = &reversal

	return m
}

func (m *ReversalSubmissionRelationships) WithoutReversal() *ReversalSubmissionRelationships {
	m.Reversal = nil
	return m
}

func (m *ReversalSubmissionRelationships) WithValidations(validations RelationshipLinks) *ReversalSubmissionRelationships {

	m.Validations = &validations

	return m
}

func (m *ReversalSubmissionRelationships) WithoutValidations() *ReversalSubmissionRelationships {
	m.Validations = nil
	return m
}

// Validate validates this reversal submission relationships
func (m *ReversalSubmissionRelationships) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePayment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReversal(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValidations(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReversalSubmissionRelationships) validatePayment(formats strfmt.Registry) error {

	if swag.IsZero(m.Payment) { // not required
		return nil
	}

	if m.Payment != nil {
		if err := m.Payment.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("payment")
			}
			return err
		}
	}

	return nil
}

func (m *ReversalSubmissionRelationships) validateReversal(formats strfmt.Registry) error {

	if swag.IsZero(m.Reversal) { // not required
		return nil
	}

	if m.Reversal != nil {
		if err := m.Reversal.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("reversal")
			}
			return err
		}
	}

	return nil
}

func (m *ReversalSubmissionRelationships) validateValidations(formats strfmt.Registry) error {

	if swag.IsZero(m.Validations) { // not required
		return nil
	}

	if m.Validations != nil {
		if err := m.Validations.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("validations")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReversalSubmissionRelationships) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReversalSubmissionRelationships) UnmarshalBinary(b []byte) error {
	var res ReversalSubmissionRelationships
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *ReversalSubmissionRelationships) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
