// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/form3tech-oss/go-form3/v4/pkg/client"
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PaymentAdmissionBeneficiaryAccount payment admission beneficiary account
// swagger:model PaymentAdmissionBeneficiaryAccount
type PaymentAdmissionBeneficiaryAccount struct {

	// attributes
	Attributes *PaymentAdmissionBeneficiaryAccountAttributes `json:"attributes,omitempty"`

	// id
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// Name of the resource type
	// Pattern: ^[A-Za-z_]*$
	Type string `json:"type,omitempty"`
}

func PaymentAdmissionBeneficiaryAccountWithDefaults(defaults client.Defaults) *PaymentAdmissionBeneficiaryAccount {
	return &PaymentAdmissionBeneficiaryAccount{

		Attributes: PaymentAdmissionBeneficiaryAccountAttributesWithDefaults(defaults),

		ID: defaults.GetStrfmtUUID("PaymentAdmissionBeneficiaryAccount", "id"),

		Type: defaults.GetString("PaymentAdmissionBeneficiaryAccount", "type"),
	}
}

func (m *PaymentAdmissionBeneficiaryAccount) WithAttributes(attributes PaymentAdmissionBeneficiaryAccountAttributes) *PaymentAdmissionBeneficiaryAccount {

	m.Attributes = &attributes

	return m
}

func (m *PaymentAdmissionBeneficiaryAccount) WithoutAttributes() *PaymentAdmissionBeneficiaryAccount {
	m.Attributes = nil
	return m
}

func (m *PaymentAdmissionBeneficiaryAccount) WithID(id strfmt.UUID) *PaymentAdmissionBeneficiaryAccount {

	m.ID = id

	return m
}

func (m *PaymentAdmissionBeneficiaryAccount) WithType(typeVar string) *PaymentAdmissionBeneficiaryAccount {

	m.Type = typeVar

	return m
}

// Validate validates this payment admission beneficiary account
func (m *PaymentAdmissionBeneficiaryAccount) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttributes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PaymentAdmissionBeneficiaryAccount) validateAttributes(formats strfmt.Registry) error {

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

func (m *PaymentAdmissionBeneficiaryAccount) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PaymentAdmissionBeneficiaryAccount) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := validate.Pattern("type", "body", string(m.Type), `^[A-Za-z_]*$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PaymentAdmissionBeneficiaryAccount) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentAdmissionBeneficiaryAccount) UnmarshalBinary(b []byte) error {
	var res PaymentAdmissionBeneficiaryAccount
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *PaymentAdmissionBeneficiaryAccount) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// PaymentAdmissionBeneficiaryAccountAttributes payment admission beneficiary account attributes
// swagger:model PaymentAdmissionBeneficiaryAccountAttributes
type PaymentAdmissionBeneficiaryAccountAttributes struct {

	// All purpose list of key-value pairs specific data stored on the associated beneficiary account.
	// Max Items: 5
	UserDefinedData []*UserDefinedData `json:"user_defined_data"`
}

func PaymentAdmissionBeneficiaryAccountAttributesWithDefaults(defaults client.Defaults) *PaymentAdmissionBeneficiaryAccountAttributes {
	return &PaymentAdmissionBeneficiaryAccountAttributes{

		UserDefinedData: make([]*UserDefinedData, 0),
	}
}

func (m *PaymentAdmissionBeneficiaryAccountAttributes) WithUserDefinedData(userDefinedData []*UserDefinedData) *PaymentAdmissionBeneficiaryAccountAttributes {

	m.UserDefinedData = userDefinedData

	return m
}

// Validate validates this payment admission beneficiary account attributes
func (m *PaymentAdmissionBeneficiaryAccountAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateUserDefinedData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PaymentAdmissionBeneficiaryAccountAttributes) validateUserDefinedData(formats strfmt.Registry) error {

	if swag.IsZero(m.UserDefinedData) { // not required
		return nil
	}

	iUserDefinedDataSize := int64(len(m.UserDefinedData))

	if err := validate.MaxItems("attributes"+"."+"user_defined_data", "body", iUserDefinedDataSize, 5); err != nil {
		return err
	}

	for i := 0; i < len(m.UserDefinedData); i++ {
		if swag.IsZero(m.UserDefinedData[i]) { // not required
			continue
		}

		if m.UserDefinedData[i] != nil {
			if err := m.UserDefinedData[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("attributes" + "." + "user_defined_data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PaymentAdmissionBeneficiaryAccountAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentAdmissionBeneficiaryAccountAttributes) UnmarshalBinary(b []byte) error {
	var res PaymentAdmissionBeneficiaryAccountAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *PaymentAdmissionBeneficiaryAccountAttributes) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
