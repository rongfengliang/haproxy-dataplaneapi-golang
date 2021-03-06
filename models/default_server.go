// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DefaultServer default server
//
// swagger:model default_server
type DefaultServer struct {

	// check sni
	// Pattern: ^[^\s]+$
	CheckSni string `json:"check-sni,omitempty"`

	// check ssl
	// Enum: [enabled disabled]
	CheckSsl string `json:"check-ssl,omitempty"`

	// downinter
	Downinter *int64 `json:"downinter,omitempty"`

	// fall
	Fall *int64 `json:"fall,omitempty"`

	// fastinter
	Fastinter *int64 `json:"fastinter,omitempty"`

	// init addr
	// Pattern: ^[^\s]+$
	InitAddr string `json:"init-addr,omitempty"`

	// inter
	Inter *int64 `json:"inter,omitempty"`

	// port
	// Maximum: 65535
	// Minimum: 1
	Port *int64 `json:"port,omitempty"`

	// resolve net
	// Pattern: ^[^\s]+$
	ResolveNet string `json:"resolve-net,omitempty"`

	// resolve prefer
	// Pattern: ^[^\s]+$
	ResolvePrefer string `json:"resolve-prefer,omitempty"`

	// resolvers
	// Pattern: ^[^\s]+$
	Resolvers string `json:"resolvers,omitempty"`

	// rise
	Rise *int64 `json:"rise,omitempty"`

	// slowstart
	Slowstart *int64 `json:"slowstart,omitempty"`

	// sni
	// Pattern: ^[^\s]+$
	Sni string `json:"sni,omitempty"`
}

// Validate validates this default server
func (m *DefaultServer) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCheckSni(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCheckSsl(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInitAddr(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePort(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResolveNet(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResolvePrefer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResolvers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSni(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DefaultServer) validateCheckSni(formats strfmt.Registry) error {

	if swag.IsZero(m.CheckSni) { // not required
		return nil
	}

	if err := validate.Pattern("check-sni", "body", string(m.CheckSni), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

var defaultServerTypeCheckSslPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["enabled","disabled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		defaultServerTypeCheckSslPropEnum = append(defaultServerTypeCheckSslPropEnum, v)
	}
}

const (

	// DefaultServerCheckSslEnabled captures enum value "enabled"
	DefaultServerCheckSslEnabled string = "enabled"

	// DefaultServerCheckSslDisabled captures enum value "disabled"
	DefaultServerCheckSslDisabled string = "disabled"
)

// prop value enum
func (m *DefaultServer) validateCheckSslEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, defaultServerTypeCheckSslPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *DefaultServer) validateCheckSsl(formats strfmt.Registry) error {

	if swag.IsZero(m.CheckSsl) { // not required
		return nil
	}

	// value enum
	if err := m.validateCheckSslEnum("check-ssl", "body", m.CheckSsl); err != nil {
		return err
	}

	return nil
}

func (m *DefaultServer) validateInitAddr(formats strfmt.Registry) error {

	if swag.IsZero(m.InitAddr) { // not required
		return nil
	}

	if err := validate.Pattern("init-addr", "body", string(m.InitAddr), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *DefaultServer) validatePort(formats strfmt.Registry) error {

	if swag.IsZero(m.Port) { // not required
		return nil
	}

	if err := validate.MinimumInt("port", "body", int64(*m.Port), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("port", "body", int64(*m.Port), 65535, false); err != nil {
		return err
	}

	return nil
}

func (m *DefaultServer) validateResolveNet(formats strfmt.Registry) error {

	if swag.IsZero(m.ResolveNet) { // not required
		return nil
	}

	if err := validate.Pattern("resolve-net", "body", string(m.ResolveNet), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *DefaultServer) validateResolvePrefer(formats strfmt.Registry) error {

	if swag.IsZero(m.ResolvePrefer) { // not required
		return nil
	}

	if err := validate.Pattern("resolve-prefer", "body", string(m.ResolvePrefer), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *DefaultServer) validateResolvers(formats strfmt.Registry) error {

	if swag.IsZero(m.Resolvers) { // not required
		return nil
	}

	if err := validate.Pattern("resolvers", "body", string(m.Resolvers), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *DefaultServer) validateSni(formats strfmt.Registry) error {

	if swag.IsZero(m.Sni) { // not required
		return nil
	}

	if err := validate.Pattern("sni", "body", string(m.Sni), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DefaultServer) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DefaultServer) UnmarshalBinary(b []byte) error {
	var res DefaultServer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
