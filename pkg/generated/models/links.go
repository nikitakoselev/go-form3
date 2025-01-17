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

// Links links
// swagger:model Links
type Links struct {

	// Link to the first resource in the list
	First *string `json:"first,omitempty"`

	// Link to the last resource in the list
	Last *string `json:"last,omitempty"`

	// Link to the next resource in the list
	Next *string `json:"next,omitempty"`

	// Link to the previous resource in the list
	Prev *string `json:"prev,omitempty"`

	// Link to this resource type
	// Required: true
	Self *string `json:"self"`
}

func LinksWithDefaults(defaults client.Defaults) *Links {
	return &Links{

		First: defaults.GetStringPtr("Links", "first"),

		Last: defaults.GetStringPtr("Links", "last"),

		Next: defaults.GetStringPtr("Links", "next"),

		Prev: defaults.GetStringPtr("Links", "prev"),

		Self: defaults.GetStringPtr("Links", "self"),
	}
}

func (m *Links) WithFirst(first string) *Links {

	m.First = &first

	return m
}

func (m *Links) WithoutFirst() *Links {
	m.First = nil
	return m
}

func (m *Links) WithLast(last string) *Links {

	m.Last = &last

	return m
}

func (m *Links) WithoutLast() *Links {
	m.Last = nil
	return m
}

func (m *Links) WithNext(next string) *Links {

	m.Next = &next

	return m
}

func (m *Links) WithoutNext() *Links {
	m.Next = nil
	return m
}

func (m *Links) WithPrev(prev string) *Links {

	m.Prev = &prev

	return m
}

func (m *Links) WithoutPrev() *Links {
	m.Prev = nil
	return m
}

func (m *Links) WithSelf(self string) *Links {

	m.Self = &self

	return m
}

func (m *Links) WithoutSelf() *Links {
	m.Self = nil
	return m
}

// Validate validates this links
func (m *Links) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Links) validateSelf(formats strfmt.Registry) error {

	if err := validate.Required("self", "body", m.Self); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Links) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Links) UnmarshalBinary(b []byte) error {
	var res Links
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *Links) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
