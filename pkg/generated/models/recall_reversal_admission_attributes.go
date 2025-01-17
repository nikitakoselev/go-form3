// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"

	"github.com/form3tech-oss/go-form3/v4/pkg/client"
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// RecallReversalAdmissionAttributes recall reversal admission attributes
// swagger:model RecallReversalAdmissionAttributes
type RecallReversalAdmissionAttributes struct {

	// Scheme-specific status code
	SchemeStatusCode string `json:"scheme_status_code,omitempty"`

	// source gateway
	SourceGateway string `json:"source_gateway,omitempty"`
}

func RecallReversalAdmissionAttributesWithDefaults(defaults client.Defaults) *RecallReversalAdmissionAttributes {
	return &RecallReversalAdmissionAttributes{

		SchemeStatusCode: defaults.GetString("RecallReversalAdmissionAttributes", "scheme_status_code"),

		SourceGateway: defaults.GetString("RecallReversalAdmissionAttributes", "source_gateway"),
	}
}

func (m *RecallReversalAdmissionAttributes) WithSchemeStatusCode(schemeStatusCode string) *RecallReversalAdmissionAttributes {

	m.SchemeStatusCode = schemeStatusCode

	return m
}

func (m *RecallReversalAdmissionAttributes) WithSourceGateway(sourceGateway string) *RecallReversalAdmissionAttributes {

	m.SourceGateway = sourceGateway

	return m
}

// Validate validates this recall reversal admission attributes
func (m *RecallReversalAdmissionAttributes) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RecallReversalAdmissionAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RecallReversalAdmissionAttributes) UnmarshalBinary(b []byte) error {
	var res RecallReversalAdmissionAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *RecallReversalAdmissionAttributes) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
