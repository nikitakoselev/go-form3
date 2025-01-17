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

// AccountRelationships account relationships
// swagger:model AccountRelationships
type AccountRelationships struct {

	// Events related to the lifecycle of this account
	AccountEvents *RelationshipLinks `json:"account_events,omitempty"`

	// ID of the master account related to this account
	MasterAccount *RelationshipLinks `json:"master_account,omitempty"`
}

func AccountRelationshipsWithDefaults(defaults client.Defaults) *AccountRelationships {
	return &AccountRelationships{

		AccountEvents: RelationshipLinksWithDefaults(defaults),

		MasterAccount: RelationshipLinksWithDefaults(defaults),
	}
}

func (m *AccountRelationships) WithAccountEvents(accountEvents RelationshipLinks) *AccountRelationships {

	m.AccountEvents = &accountEvents

	return m
}

func (m *AccountRelationships) WithoutAccountEvents() *AccountRelationships {
	m.AccountEvents = nil
	return m
}

func (m *AccountRelationships) WithMasterAccount(masterAccount RelationshipLinks) *AccountRelationships {

	m.MasterAccount = &masterAccount

	return m
}

func (m *AccountRelationships) WithoutMasterAccount() *AccountRelationships {
	m.MasterAccount = nil
	return m
}

// Validate validates this account relationships
func (m *AccountRelationships) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccountEvents(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMasterAccount(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountRelationships) validateAccountEvents(formats strfmt.Registry) error {

	if swag.IsZero(m.AccountEvents) { // not required
		return nil
	}

	if m.AccountEvents != nil {
		if err := m.AccountEvents.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("account_events")
			}
			return err
		}
	}

	return nil
}

func (m *AccountRelationships) validateMasterAccount(formats strfmt.Registry) error {

	if swag.IsZero(m.MasterAccount) { // not required
		return nil
	}

	if m.MasterAccount != nil {
		if err := m.MasterAccount.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("master_account")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AccountRelationships) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountRelationships) UnmarshalBinary(b []byte) error {
	var res AccountRelationships
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *AccountRelationships) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
