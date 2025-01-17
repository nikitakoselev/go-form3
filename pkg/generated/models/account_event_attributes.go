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

// AccountEventAttributes account event attributes
// swagger:model AccountEventAttributes
type AccountEventAttributes struct {

	// ID of the account this event relates to
	// Required: true
	// Format: uuid
	AccountID *strfmt.UUID `json:"account_id"`

	// date time
	// Required: true
	// Format: date-time
	DateTime *strfmt.DateTime `json:"date_time"`

	// Contains the event description
	// Enum: [pending failed confirmed]
	Description string `json:"description,omitempty"`

	// Failure reason. Should only be present when description is failed
	Reason string `json:"reason,omitempty"`

	// Contains the routing status
	// Required: true
	// Enum: [unroutable routable deleted]
	RoutingStatus *string `json:"routing_status"`

	// Contains the event status
	// Required: true
	// Enum: [pending failed confirmed]
	Status *string `json:"status"`
}

func AccountEventAttributesWithDefaults(defaults client.Defaults) *AccountEventAttributes {
	return &AccountEventAttributes{

		AccountID: defaults.GetStrfmtUUIDPtr("AccountEventAttributes", "account_id"),

		DateTime: defaults.GetStrfmtDateTimePtr("AccountEventAttributes", "date_time"),

		Description: defaults.GetString("AccountEventAttributes", "description"),

		Reason: defaults.GetString("AccountEventAttributes", "reason"),

		RoutingStatus: defaults.GetStringPtr("AccountEventAttributes", "routing_status"),

		Status: defaults.GetStringPtr("AccountEventAttributes", "status"),
	}
}

func (m *AccountEventAttributes) WithAccountID(accountID strfmt.UUID) *AccountEventAttributes {

	m.AccountID = &accountID

	return m
}

func (m *AccountEventAttributes) WithoutAccountID() *AccountEventAttributes {
	m.AccountID = nil
	return m
}

func (m *AccountEventAttributes) WithDateTime(dateTime strfmt.DateTime) *AccountEventAttributes {

	m.DateTime = &dateTime

	return m
}

func (m *AccountEventAttributes) WithoutDateTime() *AccountEventAttributes {
	m.DateTime = nil
	return m
}

func (m *AccountEventAttributes) WithDescription(description string) *AccountEventAttributes {

	m.Description = description

	return m
}

func (m *AccountEventAttributes) WithReason(reason string) *AccountEventAttributes {

	m.Reason = reason

	return m
}

func (m *AccountEventAttributes) WithRoutingStatus(routingStatus string) *AccountEventAttributes {

	m.RoutingStatus = &routingStatus

	return m
}

func (m *AccountEventAttributes) WithoutRoutingStatus() *AccountEventAttributes {
	m.RoutingStatus = nil
	return m
}

func (m *AccountEventAttributes) WithStatus(status string) *AccountEventAttributes {

	m.Status = &status

	return m
}

func (m *AccountEventAttributes) WithoutStatus() *AccountEventAttributes {
	m.Status = nil
	return m
}

// Validate validates this account event attributes
func (m *AccountEventAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccountID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRoutingStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountEventAttributes) validateAccountID(formats strfmt.Registry) error {

	if err := validate.Required("account_id", "body", m.AccountID); err != nil {
		return err
	}

	if err := validate.FormatOf("account_id", "body", "uuid", m.AccountID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *AccountEventAttributes) validateDateTime(formats strfmt.Registry) error {

	if err := validate.Required("date_time", "body", m.DateTime); err != nil {
		return err
	}

	if err := validate.FormatOf("date_time", "body", "date-time", m.DateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

var accountEventAttributesTypeDescriptionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["pending","failed","confirmed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		accountEventAttributesTypeDescriptionPropEnum = append(accountEventAttributesTypeDescriptionPropEnum, v)
	}
}

const (

	// AccountEventAttributesDescriptionPending captures enum value "pending"
	AccountEventAttributesDescriptionPending string = "pending"

	// AccountEventAttributesDescriptionFailed captures enum value "failed"
	AccountEventAttributesDescriptionFailed string = "failed"

	// AccountEventAttributesDescriptionConfirmed captures enum value "confirmed"
	AccountEventAttributesDescriptionConfirmed string = "confirmed"
)

// prop value enum
func (m *AccountEventAttributes) validateDescriptionEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, accountEventAttributesTypeDescriptionPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *AccountEventAttributes) validateDescription(formats strfmt.Registry) error {

	if swag.IsZero(m.Description) { // not required
		return nil
	}

	// value enum
	if err := m.validateDescriptionEnum("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

var accountEventAttributesTypeRoutingStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["unroutable","routable","deleted"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		accountEventAttributesTypeRoutingStatusPropEnum = append(accountEventAttributesTypeRoutingStatusPropEnum, v)
	}
}

const (

	// AccountEventAttributesRoutingStatusUnroutable captures enum value "unroutable"
	AccountEventAttributesRoutingStatusUnroutable string = "unroutable"

	// AccountEventAttributesRoutingStatusRoutable captures enum value "routable"
	AccountEventAttributesRoutingStatusRoutable string = "routable"

	// AccountEventAttributesRoutingStatusDeleted captures enum value "deleted"
	AccountEventAttributesRoutingStatusDeleted string = "deleted"
)

// prop value enum
func (m *AccountEventAttributes) validateRoutingStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, accountEventAttributesTypeRoutingStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *AccountEventAttributes) validateRoutingStatus(formats strfmt.Registry) error {

	if err := validate.Required("routing_status", "body", m.RoutingStatus); err != nil {
		return err
	}

	// value enum
	if err := m.validateRoutingStatusEnum("routing_status", "body", *m.RoutingStatus); err != nil {
		return err
	}

	return nil
}

var accountEventAttributesTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["pending","failed","confirmed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		accountEventAttributesTypeStatusPropEnum = append(accountEventAttributesTypeStatusPropEnum, v)
	}
}

const (

	// AccountEventAttributesStatusPending captures enum value "pending"
	AccountEventAttributesStatusPending string = "pending"

	// AccountEventAttributesStatusFailed captures enum value "failed"
	AccountEventAttributesStatusFailed string = "failed"

	// AccountEventAttributesStatusConfirmed captures enum value "confirmed"
	AccountEventAttributesStatusConfirmed string = "confirmed"
)

// prop value enum
func (m *AccountEventAttributes) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, accountEventAttributesTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *AccountEventAttributes) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", *m.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AccountEventAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountEventAttributes) UnmarshalBinary(b []byte) error {
	var res AccountEventAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *AccountEventAttributes) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
