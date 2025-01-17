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

// Balance balance
// swagger:model Balance
type Balance struct {

	// attributes
	Attributes *BalanceAttributes `json:"attributes,omitempty"`

	// Unique resource ID
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// Unique ID of the organisation this resource is created by
	// Format: uuid
	OrganisationID strfmt.UUID `json:"organisation_id,omitempty"`

	// Name of the resource type
	Type string `json:"type,omitempty"`

	// Version number
	// Minimum: 0
	Version *int64 `json:"version,omitempty"`
}

func BalanceWithDefaults(defaults client.Defaults) *Balance {
	return &Balance{

		Attributes: BalanceAttributesWithDefaults(defaults),

		ID: defaults.GetStrfmtUUID("Balance", "id"),

		OrganisationID: defaults.GetStrfmtUUID("Balance", "organisation_id"),

		Type: defaults.GetString("Balance", "type"),

		Version: defaults.GetInt64Ptr("Balance", "version"),
	}
}

func (m *Balance) WithAttributes(attributes BalanceAttributes) *Balance {

	m.Attributes = &attributes

	return m
}

func (m *Balance) WithoutAttributes() *Balance {
	m.Attributes = nil
	return m
}

func (m *Balance) WithID(id strfmt.UUID) *Balance {

	m.ID = id

	return m
}

func (m *Balance) WithOrganisationID(organisationID strfmt.UUID) *Balance {

	m.OrganisationID = organisationID

	return m
}

func (m *Balance) WithType(typeVar string) *Balance {

	m.Type = typeVar

	return m
}

func (m *Balance) WithVersion(version int64) *Balance {

	m.Version = &version

	return m
}

func (m *Balance) WithoutVersion() *Balance {
	m.Version = nil
	return m
}

// Validate validates this balance
func (m *Balance) Validate(formats strfmt.Registry) error {
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

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Balance) validateAttributes(formats strfmt.Registry) error {

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

func (m *Balance) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Balance) validateOrganisationID(formats strfmt.Registry) error {

	if swag.IsZero(m.OrganisationID) { // not required
		return nil
	}

	if err := validate.FormatOf("organisation_id", "body", "uuid", m.OrganisationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Balance) validateVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.Version) { // not required
		return nil
	}

	if err := validate.MinimumInt("version", "body", int64(*m.Version), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Balance) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Balance) UnmarshalBinary(b []byte) error {
	var res Balance
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *Balance) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
