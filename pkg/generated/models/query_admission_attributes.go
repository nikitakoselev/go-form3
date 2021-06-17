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

// QueryAdmissionAttributes query admission attributes
// swagger:model QueryAdmissionAttributes
type QueryAdmissionAttributes struct {

	// status
	// Required: true
	Status *QueryAdmissionStatus `json:"status"`

	// status reason
	StatusReason string `json:"status_reason,omitempty"`
}

func QueryAdmissionAttributesWithDefaults(defaults client.Defaults) *QueryAdmissionAttributes {
	return &QueryAdmissionAttributes{

		// TODO Status: QueryAdmissionStatus,

		StatusReason: defaults.GetString("QueryAdmissionAttributes", "status_reason"),
	}
}

func (m *QueryAdmissionAttributes) WithStatus(status QueryAdmissionStatus) *QueryAdmissionAttributes {

	m.Status = &status

	return m
}

func (m *QueryAdmissionAttributes) WithoutStatus() *QueryAdmissionAttributes {
	m.Status = nil
	return m
}

func (m *QueryAdmissionAttributes) WithStatusReason(statusReason string) *QueryAdmissionAttributes {

	m.StatusReason = statusReason

	return m
}

// Validate validates this query admission attributes
func (m *QueryAdmissionAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *QueryAdmissionAttributes) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	if m.Status != nil {
		if err := m.Status.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *QueryAdmissionAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *QueryAdmissionAttributes) UnmarshalBinary(b []byte) error {
	var res QueryAdmissionAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *QueryAdmissionAttributes) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
