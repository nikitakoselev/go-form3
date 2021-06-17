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

// ContactAttributes contact attributes
// swagger:model ContactAttributes
type ContactAttributes struct {

	// address
	Address []string `json:"address"`

	// city
	City string `json:"city,omitempty"`

	// contact method
	ContactMethod string `json:"contact_method,omitempty"`

	// contact type
	// Enum: [organisation private]
	ContactType string `json:"contact_type,omitempty"`

	// country
	Country string `json:"country,omitempty"`

	// district
	District string `json:"district,omitempty"`

	// email address
	EmailAddress string `json:"email_address,omitempty"`

	// identification
	Identification []string `json:"identification,omitempty"`

	// identification type
	IdentificationType string `json:"identification_type,omitempty"`

	// name
	Name []string `json:"name"`

	// post code
	PostCode string `json:"post_code,omitempty"`

	// province
	Province string `json:"province,omitempty"`

	// telephone number
	TelephoneNumber string `json:"telephone_number,omitempty"`
}

func ContactAttributesWithDefaults(defaults client.Defaults) *ContactAttributes {
	return &ContactAttributes{

		Address: make([]string, 0),

		City: defaults.GetString("ContactAttributes", "city"),

		ContactMethod: defaults.GetString("ContactAttributes", "contact_method"),

		ContactType: defaults.GetString("ContactAttributes", "contact_type"),

		Country: defaults.GetString("ContactAttributes", "country"),

		District: defaults.GetString("ContactAttributes", "district"),

		EmailAddress: defaults.GetString("ContactAttributes", "email_address"),

		Identification: make([]string, 0),

		IdentificationType: defaults.GetString("ContactAttributes", "identification_type"),

		Name: make([]string, 0),

		PostCode: defaults.GetString("ContactAttributes", "post_code"),

		Province: defaults.GetString("ContactAttributes", "province"),

		TelephoneNumber: defaults.GetString("ContactAttributes", "telephone_number"),
	}
}

func (m *ContactAttributes) WithAddress(address []string) *ContactAttributes {

	m.Address = address

	return m
}

func (m *ContactAttributes) WithCity(city string) *ContactAttributes {

	m.City = city

	return m
}

func (m *ContactAttributes) WithContactMethod(contactMethod string) *ContactAttributes {

	m.ContactMethod = contactMethod

	return m
}

func (m *ContactAttributes) WithContactType(contactType string) *ContactAttributes {

	m.ContactType = contactType

	return m
}

func (m *ContactAttributes) WithCountry(country string) *ContactAttributes {

	m.Country = country

	return m
}

func (m *ContactAttributes) WithDistrict(district string) *ContactAttributes {

	m.District = district

	return m
}

func (m *ContactAttributes) WithEmailAddress(emailAddress string) *ContactAttributes {

	m.EmailAddress = emailAddress

	return m
}

func (m *ContactAttributes) WithIdentification(identification []string) *ContactAttributes {

	m.Identification = identification

	return m
}

func (m *ContactAttributes) WithIdentificationType(identificationType string) *ContactAttributes {

	m.IdentificationType = identificationType

	return m
}

func (m *ContactAttributes) WithName(name []string) *ContactAttributes {

	m.Name = name

	return m
}

func (m *ContactAttributes) WithPostCode(postCode string) *ContactAttributes {

	m.PostCode = postCode

	return m
}

func (m *ContactAttributes) WithProvince(province string) *ContactAttributes {

	m.Province = province

	return m
}

func (m *ContactAttributes) WithTelephoneNumber(telephoneNumber string) *ContactAttributes {

	m.TelephoneNumber = telephoneNumber

	return m
}

// Validate validates this contact attributes
func (m *ContactAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContactType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var contactAttributesTypeContactTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["organisation","private"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		contactAttributesTypeContactTypePropEnum = append(contactAttributesTypeContactTypePropEnum, v)
	}
}

const (

	// ContactAttributesContactTypeOrganisation captures enum value "organisation"
	ContactAttributesContactTypeOrganisation string = "organisation"

	// ContactAttributesContactTypePrivate captures enum value "private"
	ContactAttributesContactTypePrivate string = "private"
)

// prop value enum
func (m *ContactAttributes) validateContactTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, contactAttributesTypeContactTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ContactAttributes) validateContactType(formats strfmt.Registry) error {

	if swag.IsZero(m.ContactType) { // not required
		return nil
	}

	// value enum
	if err := m.validateContactTypeEnum("contact_type", "body", m.ContactType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ContactAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContactAttributes) UnmarshalBinary(b []byte) error {
	var res ContactAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *ContactAttributes) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
