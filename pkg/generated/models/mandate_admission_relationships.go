// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/form3tech-oss/go-form3/v3/pkg/client"
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// MandateAdmissionRelationships mandate admission relationships
// swagger:model MandateAdmissionRelationships
type MandateAdmissionRelationships struct {

	// mandate
	Mandate *MandateAdmissionRelationshipsMandate `json:"mandate,omitempty"`
}

func MandateAdmissionRelationshipsWithDefaults(defaults client.Defaults) *MandateAdmissionRelationships {
	return &MandateAdmissionRelationships{

		Mandate: MandateAdmissionRelationshipsMandateWithDefaults(defaults),
	}
}

func (m *MandateAdmissionRelationships) WithMandate(mandate MandateAdmissionRelationshipsMandate) *MandateAdmissionRelationships {

	m.Mandate = &mandate

	return m
}

func (m *MandateAdmissionRelationships) WithoutMandate() *MandateAdmissionRelationships {
	m.Mandate = nil
	return m
}

// Validate validates this mandate admission relationships
func (m *MandateAdmissionRelationships) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMandate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MandateAdmissionRelationships) validateMandate(formats strfmt.Registry) error {

	if swag.IsZero(m.Mandate) { // not required
		return nil
	}

	if m.Mandate != nil {
		if err := m.Mandate.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mandate")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MandateAdmissionRelationships) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MandateAdmissionRelationships) UnmarshalBinary(b []byte) error {
	var res MandateAdmissionRelationships
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *MandateAdmissionRelationships) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// MandateAdmissionRelationshipsMandate mandate admission relationships mandate
// swagger:model MandateAdmissionRelationshipsMandate
type MandateAdmissionRelationshipsMandate struct {

	// data
	Data []*Mandate `json:"data"`
}

func MandateAdmissionRelationshipsMandateWithDefaults(defaults client.Defaults) *MandateAdmissionRelationshipsMandate {
	return &MandateAdmissionRelationshipsMandate{

		Data: make([]*Mandate, 0),
	}
}

func (m *MandateAdmissionRelationshipsMandate) WithData(data []*Mandate) *MandateAdmissionRelationshipsMandate {

	m.Data = data

	return m
}

// Validate validates this mandate admission relationships mandate
func (m *MandateAdmissionRelationshipsMandate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MandateAdmissionRelationshipsMandate) validateData(formats strfmt.Registry) error {

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
					return ve.ValidateName("mandate" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *MandateAdmissionRelationshipsMandate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MandateAdmissionRelationshipsMandate) UnmarshalBinary(b []byte) error {
	var res MandateAdmissionRelationshipsMandate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *MandateAdmissionRelationshipsMandate) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
