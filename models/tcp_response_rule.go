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

// TCPResponseRule TCP Response Rule
//
// HAProxy TCP Response Rule configuration (corresponds to tcp-response)
//
// swagger:model tcp_response_rule
type TCPResponseRule struct {

	// action
	// Enum: [accept reject lua]
	Action string `json:"action,omitempty"`

	// cond
	// Enum: [if unless]
	Cond string `json:"cond,omitempty"`

	// cond test
	CondTest string `json:"cond_test,omitempty"`

	// index
	// Required: true
	Index *int64 `json:"index"`

	// lua action
	// Pattern: ^[^\s]+$
	LuaAction string `json:"lua_action,omitempty"`

	// lua params
	LuaParams string `json:"lua_params,omitempty"`

	// timeout
	Timeout *int64 `json:"timeout,omitempty"`

	// type
	// Required: true
	// Enum: [content inspect-delay]
	Type string `json:"type"`
}

// Validate validates this tcp response rule
func (m *TCPResponseRule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCond(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIndex(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLuaAction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var tcpResponseRuleTypeActionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["accept","reject","lua"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		tcpResponseRuleTypeActionPropEnum = append(tcpResponseRuleTypeActionPropEnum, v)
	}
}

const (

	// TCPResponseRuleActionAccept captures enum value "accept"
	TCPResponseRuleActionAccept string = "accept"

	// TCPResponseRuleActionReject captures enum value "reject"
	TCPResponseRuleActionReject string = "reject"

	// TCPResponseRuleActionLua captures enum value "lua"
	TCPResponseRuleActionLua string = "lua"
)

// prop value enum
func (m *TCPResponseRule) validateActionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, tcpResponseRuleTypeActionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *TCPResponseRule) validateAction(formats strfmt.Registry) error {

	if swag.IsZero(m.Action) { // not required
		return nil
	}

	// value enum
	if err := m.validateActionEnum("action", "body", m.Action); err != nil {
		return err
	}

	return nil
}

var tcpResponseRuleTypeCondPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["if","unless"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		tcpResponseRuleTypeCondPropEnum = append(tcpResponseRuleTypeCondPropEnum, v)
	}
}

const (

	// TCPResponseRuleCondIf captures enum value "if"
	TCPResponseRuleCondIf string = "if"

	// TCPResponseRuleCondUnless captures enum value "unless"
	TCPResponseRuleCondUnless string = "unless"
)

// prop value enum
func (m *TCPResponseRule) validateCondEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, tcpResponseRuleTypeCondPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *TCPResponseRule) validateCond(formats strfmt.Registry) error {

	if swag.IsZero(m.Cond) { // not required
		return nil
	}

	// value enum
	if err := m.validateCondEnum("cond", "body", m.Cond); err != nil {
		return err
	}

	return nil
}

func (m *TCPResponseRule) validateIndex(formats strfmt.Registry) error {

	if err := validate.Required("index", "body", m.Index); err != nil {
		return err
	}

	return nil
}

func (m *TCPResponseRule) validateLuaAction(formats strfmt.Registry) error {

	if swag.IsZero(m.LuaAction) { // not required
		return nil
	}

	if err := validate.Pattern("lua_action", "body", string(m.LuaAction), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

var tcpResponseRuleTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["content","inspect-delay"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		tcpResponseRuleTypeTypePropEnum = append(tcpResponseRuleTypeTypePropEnum, v)
	}
}

const (

	// TCPResponseRuleTypeContent captures enum value "content"
	TCPResponseRuleTypeContent string = "content"

	// TCPResponseRuleTypeInspectDelay captures enum value "inspect-delay"
	TCPResponseRuleTypeInspectDelay string = "inspect-delay"
)

// prop value enum
func (m *TCPResponseRule) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, tcpResponseRuleTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *TCPResponseRule) validateType(formats strfmt.Registry) error {

	if err := validate.RequiredString("type", "body", string(m.Type)); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TCPResponseRule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TCPResponseRule) UnmarshalBinary(b []byte) error {
	var res TCPResponseRule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}