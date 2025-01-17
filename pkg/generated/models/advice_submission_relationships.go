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
)

// AdviceSubmissionRelationships advice submission relationships
// swagger:model AdviceSubmissionRelationships
type AdviceSubmissionRelationships struct {

	// advice
	Advice *RelationshipAdvices `json:"advice,omitempty"`

	// payment
	Payment *RelationshipPayments `json:"payment,omitempty"`
}

func AdviceSubmissionRelationshipsWithDefaults(defaults client.Defaults) *AdviceSubmissionRelationships {
	return &AdviceSubmissionRelationships{

		Advice: RelationshipAdvicesWithDefaults(defaults),

		Payment: RelationshipPaymentsWithDefaults(defaults),
	}
}

func (m *AdviceSubmissionRelationships) WithAdvice(advice RelationshipAdvices) *AdviceSubmissionRelationships {

	m.Advice = &advice

	return m
}

func (m *AdviceSubmissionRelationships) WithoutAdvice() *AdviceSubmissionRelationships {
	m.Advice = nil
	return m
}

func (m *AdviceSubmissionRelationships) WithPayment(payment RelationshipPayments) *AdviceSubmissionRelationships {

	m.Payment = &payment

	return m
}

func (m *AdviceSubmissionRelationships) WithoutPayment() *AdviceSubmissionRelationships {
	m.Payment = nil
	return m
}

// Validate validates this advice submission relationships
func (m *AdviceSubmissionRelationships) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdvice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePayment(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AdviceSubmissionRelationships) validateAdvice(formats strfmt.Registry) error {

	if swag.IsZero(m.Advice) { // not required
		return nil
	}

	if m.Advice != nil {
		if err := m.Advice.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("advice")
			}
			return err
		}
	}

	return nil
}

func (m *AdviceSubmissionRelationships) validatePayment(formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *AdviceSubmissionRelationships) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AdviceSubmissionRelationships) UnmarshalBinary(b []byte) error {
	var res AdviceSubmissionRelationships
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *AdviceSubmissionRelationships) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
