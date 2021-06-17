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

// PartyActorPatchAttributes party actor patch attributes
// swagger:model PartyActorPatchAttributes
type PartyActorPatchAttributes struct {

	// address
	Address []string `json:"address"`

	// city
	City *string `json:"city,omitempty"`

	// contact method
	// Enum: [email]
	ContactMethod *string `json:"contact_method,omitempty"`

	// country
	Country *string `json:"country,omitempty"`

	// customer id
	CustomerID *string `json:"customer_id,omitempty"`

	// email address
	EmailAddress *string `json:"email_address,omitempty"`

	// name
	Name []string `json:"name"`

	// party actor classification
	PartyActorClassification *PatchPartyActorClassification `json:"party_actor_classification,omitempty"`

	// party actor type
	// Enum: [organisation private]
	PartyActorType *string `json:"party_actor_type,omitempty"`

	// post code
	PostCode *string `json:"post_code,omitempty"`

	// private identification
	PrivateIdentification *PatchPartyActorPrivateIdentification `json:"private_identification,omitempty"`
}

func PartyActorPatchAttributesWithDefaults(defaults client.Defaults) *PartyActorPatchAttributes {
	return &PartyActorPatchAttributes{

		Address: make([]string, 0),

		City: defaults.GetStringPtr("PartyActorPatchAttributes", "city"),

		ContactMethod: defaults.GetStringPtr("PartyActorPatchAttributes", "contact_method"),

		Country: defaults.GetStringPtr("PartyActorPatchAttributes", "country"),

		CustomerID: defaults.GetStringPtr("PartyActorPatchAttributes", "customer_id"),

		EmailAddress: defaults.GetStringPtr("PartyActorPatchAttributes", "email_address"),

		Name: make([]string, 0),

		PartyActorClassification: PatchPartyActorClassificationWithDefaults(defaults),

		PartyActorType: defaults.GetStringPtr("PartyActorPatchAttributes", "party_actor_type"),

		PostCode: defaults.GetStringPtr("PartyActorPatchAttributes", "post_code"),

		PrivateIdentification: PatchPartyActorPrivateIdentificationWithDefaults(defaults),
	}
}

func (m *PartyActorPatchAttributes) WithAddress(address []string) *PartyActorPatchAttributes {

	m.Address = address

	return m
}

func (m *PartyActorPatchAttributes) WithCity(city string) *PartyActorPatchAttributes {

	m.City = &city

	return m
}

func (m *PartyActorPatchAttributes) WithoutCity() *PartyActorPatchAttributes {
	m.City = nil
	return m
}

func (m *PartyActorPatchAttributes) WithContactMethod(contactMethod string) *PartyActorPatchAttributes {

	m.ContactMethod = &contactMethod

	return m
}

func (m *PartyActorPatchAttributes) WithoutContactMethod() *PartyActorPatchAttributes {
	m.ContactMethod = nil
	return m
}

func (m *PartyActorPatchAttributes) WithCountry(country string) *PartyActorPatchAttributes {

	m.Country = &country

	return m
}

func (m *PartyActorPatchAttributes) WithoutCountry() *PartyActorPatchAttributes {
	m.Country = nil
	return m
}

func (m *PartyActorPatchAttributes) WithCustomerID(customerID string) *PartyActorPatchAttributes {

	m.CustomerID = &customerID

	return m
}

func (m *PartyActorPatchAttributes) WithoutCustomerID() *PartyActorPatchAttributes {
	m.CustomerID = nil
	return m
}

func (m *PartyActorPatchAttributes) WithEmailAddress(emailAddress string) *PartyActorPatchAttributes {

	m.EmailAddress = &emailAddress

	return m
}

func (m *PartyActorPatchAttributes) WithoutEmailAddress() *PartyActorPatchAttributes {
	m.EmailAddress = nil
	return m
}

func (m *PartyActorPatchAttributes) WithName(name []string) *PartyActorPatchAttributes {

	m.Name = name

	return m
}

func (m *PartyActorPatchAttributes) WithPartyActorClassification(partyActorClassification PatchPartyActorClassification) *PartyActorPatchAttributes {

	m.PartyActorClassification = &partyActorClassification

	return m
}

func (m *PartyActorPatchAttributes) WithoutPartyActorClassification() *PartyActorPatchAttributes {
	m.PartyActorClassification = nil
	return m
}

func (m *PartyActorPatchAttributes) WithPartyActorType(partyActorType string) *PartyActorPatchAttributes {

	m.PartyActorType = &partyActorType

	return m
}

func (m *PartyActorPatchAttributes) WithoutPartyActorType() *PartyActorPatchAttributes {
	m.PartyActorType = nil
	return m
}

func (m *PartyActorPatchAttributes) WithPostCode(postCode string) *PartyActorPatchAttributes {

	m.PostCode = &postCode

	return m
}

func (m *PartyActorPatchAttributes) WithoutPostCode() *PartyActorPatchAttributes {
	m.PostCode = nil
	return m
}

func (m *PartyActorPatchAttributes) WithPrivateIdentification(privateIdentification PatchPartyActorPrivateIdentification) *PartyActorPatchAttributes {

	m.PrivateIdentification = &privateIdentification

	return m
}

func (m *PartyActorPatchAttributes) WithoutPrivateIdentification() *PartyActorPatchAttributes {
	m.PrivateIdentification = nil
	return m
}

// Validate validates this party actor patch attributes
func (m *PartyActorPatchAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContactMethod(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePartyActorClassification(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePartyActorType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrivateIdentification(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var partyActorPatchAttributesTypeContactMethodPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["email"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		partyActorPatchAttributesTypeContactMethodPropEnum = append(partyActorPatchAttributesTypeContactMethodPropEnum, v)
	}
}

const (

	// PartyActorPatchAttributesContactMethodEmail captures enum value "email"
	PartyActorPatchAttributesContactMethodEmail string = "email"
)

// prop value enum
func (m *PartyActorPatchAttributes) validateContactMethodEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, partyActorPatchAttributesTypeContactMethodPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PartyActorPatchAttributes) validateContactMethod(formats strfmt.Registry) error {

	if swag.IsZero(m.ContactMethod) { // not required
		return nil
	}

	// value enum
	if err := m.validateContactMethodEnum("contact_method", "body", *m.ContactMethod); err != nil {
		return err
	}

	return nil
}

func (m *PartyActorPatchAttributes) validatePartyActorClassification(formats strfmt.Registry) error {

	if swag.IsZero(m.PartyActorClassification) { // not required
		return nil
	}

	if m.PartyActorClassification != nil {
		if err := m.PartyActorClassification.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("party_actor_classification")
			}
			return err
		}
	}

	return nil
}

var partyActorPatchAttributesTypePartyActorTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["organisation","private"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		partyActorPatchAttributesTypePartyActorTypePropEnum = append(partyActorPatchAttributesTypePartyActorTypePropEnum, v)
	}
}

const (

	// PartyActorPatchAttributesPartyActorTypeOrganisation captures enum value "organisation"
	PartyActorPatchAttributesPartyActorTypeOrganisation string = "organisation"

	// PartyActorPatchAttributesPartyActorTypePrivate captures enum value "private"
	PartyActorPatchAttributesPartyActorTypePrivate string = "private"
)

// prop value enum
func (m *PartyActorPatchAttributes) validatePartyActorTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, partyActorPatchAttributesTypePartyActorTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PartyActorPatchAttributes) validatePartyActorType(formats strfmt.Registry) error {

	if swag.IsZero(m.PartyActorType) { // not required
		return nil
	}

	// value enum
	if err := m.validatePartyActorTypeEnum("party_actor_type", "body", *m.PartyActorType); err != nil {
		return err
	}

	return nil
}

func (m *PartyActorPatchAttributes) validatePrivateIdentification(formats strfmt.Registry) error {

	if swag.IsZero(m.PrivateIdentification) { // not required
		return nil
	}

	if m.PrivateIdentification != nil {
		if err := m.PrivateIdentification.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("private_identification")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PartyActorPatchAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PartyActorPatchAttributes) UnmarshalBinary(b []byte) error {
	var res PartyActorPatchAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *PartyActorPatchAttributes) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
