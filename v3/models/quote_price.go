// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
)

// QuotePrice Price in quote currency
// swagger:model quote_price
type QuotePrice float64

// Validate validates this quote price
func (m QuotePrice) Validate(formats strfmt.Registry) error {
	return nil
}