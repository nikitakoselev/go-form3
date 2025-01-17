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

// DirectDebitReturnReversalDetailsResponse direct debit return reversal details response
// swagger:model DirectDebitReturnReversalDetailsResponse
type DirectDebitReturnReversalDetailsResponse struct {

	// data
	Data *DirectDebitReturnReversal `json:"data,omitempty"`
}

func DirectDebitReturnReversalDetailsResponseWithDefaults(defaults client.Defaults) *DirectDebitReturnReversalDetailsResponse {
	return &DirectDebitReturnReversalDetailsResponse{

		Data: DirectDebitReturnReversalWithDefaults(defaults),
	}
}

func (m *DirectDebitReturnReversalDetailsResponse) WithData(data DirectDebitReturnReversal) *DirectDebitReturnReversalDetailsResponse {

	m.Data = &data

	return m
}

func (m *DirectDebitReturnReversalDetailsResponse) WithoutData() *DirectDebitReturnReversalDetailsResponse {
	m.Data = nil
	return m
}

// Validate validates this direct debit return reversal details response
func (m *DirectDebitReturnReversalDetailsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DirectDebitReturnReversalDetailsResponse) validateData(formats strfmt.Registry) error {

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
func (m *DirectDebitReturnReversalDetailsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DirectDebitReturnReversalDetailsResponse) UnmarshalBinary(b []byte) error {
	var res DirectDebitReturnReversalDetailsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *DirectDebitReturnReversalDetailsResponse) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
