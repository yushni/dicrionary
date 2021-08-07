// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Translation translation
//
// swagger:model Translation
type Translation struct {

	// samples
	// Required: true
	Samples []string `json:"samples"`

	// transcription
	// Required: true
	Transcription *string `json:"transcription"`

	// translation
	// Required: true
	Translation *string `json:"translation"`
}

// Validate validates this translation
func (m *Translation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSamples(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTranscription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTranslation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Translation) validateSamples(formats strfmt.Registry) error {

	if err := validate.Required("samples", "body", m.Samples); err != nil {
		return err
	}

	return nil
}

func (m *Translation) validateTranscription(formats strfmt.Registry) error {

	if err := validate.Required("transcription", "body", m.Transcription); err != nil {
		return err
	}

	return nil
}

func (m *Translation) validateTranslation(formats strfmt.Registry) error {

	if err := validate.Required("translation", "body", m.Translation); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this translation based on context it is used
func (m *Translation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Translation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Translation) UnmarshalBinary(b []byte) error {
	var res Translation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
