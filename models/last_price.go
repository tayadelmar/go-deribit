// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
)

// LastPrice The price for the last trade
// swagger:model last_price
type LastPrice float64

// Validate validates this last price
func (m LastPrice) Validate(formats strfmt.Registry) error {
	return nil
}