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

// PartyActorAmendment party actor amendment
// swagger:model PartyActorAmendment
type PartyActorAmendment struct {

	// data
	Data *PartyActorUpdate `json:"data,omitempty"`
}

func PartyActorAmendmentWithDefaults(defaults client.Defaults) *PartyActorAmendment {
	return &PartyActorAmendment{

		Data: PartyActorUpdateWithDefaults(defaults),
	}
}

func (m *PartyActorAmendment) WithData(data PartyActorUpdate) *PartyActorAmendment {

	m.Data = &data

	return m
}

func (m *PartyActorAmendment) WithoutData() *PartyActorAmendment {
	m.Data = nil
	return m
}

// Validate validates this party actor amendment
func (m *PartyActorAmendment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PartyActorAmendment) validateData(formats strfmt.Registry) error {

	if swag.IsZero(m.Data) { // not required
		return nil
	}

	if m.Data != nil {
		if err := m.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PartyActorAmendment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PartyActorAmendment) UnmarshalBinary(b []byte) error {
	var res PartyActorAmendment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *PartyActorAmendment) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
