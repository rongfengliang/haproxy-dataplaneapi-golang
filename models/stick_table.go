// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// StickTable Stick Table
//
// Stick Table Information
//
// swagger:model stick_table
type StickTable struct {

	// fields
	Fields []*StickTableField `json:"fields"`

	// name
	Name string `json:"name,omitempty"`

	// Process number if master-worker mode
	Process *int64 `json:"process,omitempty"`

	// size
	Size *int64 `json:"size,omitempty"`

	// type
	// Enum: [ip ipv6 integer string binary]
	Type string `json:"type,omitempty"`

	// used
	Used *int64 `json:"used,omitempty"`
}

// Validate validates this stick table
func (m *StickTable) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFields(formats); err != nil {
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

func (m *StickTable) validateFields(formats strfmt.Registry) error {

	if swag.IsZero(m.Fields) { // not required
		return nil
	}

	for i := 0; i < len(m.Fields); i++ {
		if swag.IsZero(m.Fields[i]) { // not required
			continue
		}

		if m.Fields[i] != nil {
			if err := m.Fields[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("fields" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var stickTableTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ip","ipv6","integer","string","binary"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		stickTableTypeTypePropEnum = append(stickTableTypeTypePropEnum, v)
	}
}

const (

	// StickTableTypeIP captures enum value "ip"
	StickTableTypeIP string = "ip"

	// StickTableTypeIPV6 captures enum value "ipv6"
	StickTableTypeIPV6 string = "ipv6"

	// StickTableTypeInteger captures enum value "integer"
	StickTableTypeInteger string = "integer"

	// StickTableTypeString captures enum value "string"
	StickTableTypeString string = "string"

	// StickTableTypeBinary captures enum value "binary"
	StickTableTypeBinary string = "binary"
)

// prop value enum
func (m *StickTable) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, stickTableTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *StickTable) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *StickTable) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StickTable) UnmarshalBinary(b []byte) error {
	var res StickTable
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// StickTableField stick table field
//
// swagger:model StickTableField
type StickTableField struct {

	// field
	// Enum: [server_id gpc0 gpc0_rate gpc1 gpc1_rate conn_cnt conn_cur conn_rate sess_cnt sess_rate http_req_cnt http_req_rate http_err_cnt http_err_rate bytes_in_cnt bytes_in_rate bytes_out_cnt bytes_out_rate]
	Field string `json:"field,omitempty"`

	// period
	Period int64 `json:"period,omitempty"`

	// type
	// Enum: [rate counter]
	Type string `json:"type,omitempty"`
}

// Validate validates this stick table field
func (m *StickTableField) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateField(formats); err != nil {
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

var stickTableFieldTypeFieldPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["server_id","gpc0","gpc0_rate","gpc1","gpc1_rate","conn_cnt","conn_cur","conn_rate","sess_cnt","sess_rate","http_req_cnt","http_req_rate","http_err_cnt","http_err_rate","bytes_in_cnt","bytes_in_rate","bytes_out_cnt","bytes_out_rate"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		stickTableFieldTypeFieldPropEnum = append(stickTableFieldTypeFieldPropEnum, v)
	}
}

const (

	// StickTableFieldFieldServerID captures enum value "server_id"
	StickTableFieldFieldServerID string = "server_id"

	// StickTableFieldFieldGpc0 captures enum value "gpc0"
	StickTableFieldFieldGpc0 string = "gpc0"

	// StickTableFieldFieldGpc0Rate captures enum value "gpc0_rate"
	StickTableFieldFieldGpc0Rate string = "gpc0_rate"

	// StickTableFieldFieldGpc1 captures enum value "gpc1"
	StickTableFieldFieldGpc1 string = "gpc1"

	// StickTableFieldFieldGpc1Rate captures enum value "gpc1_rate"
	StickTableFieldFieldGpc1Rate string = "gpc1_rate"

	// StickTableFieldFieldConnCnt captures enum value "conn_cnt"
	StickTableFieldFieldConnCnt string = "conn_cnt"

	// StickTableFieldFieldConnCur captures enum value "conn_cur"
	StickTableFieldFieldConnCur string = "conn_cur"

	// StickTableFieldFieldConnRate captures enum value "conn_rate"
	StickTableFieldFieldConnRate string = "conn_rate"

	// StickTableFieldFieldSessCnt captures enum value "sess_cnt"
	StickTableFieldFieldSessCnt string = "sess_cnt"

	// StickTableFieldFieldSessRate captures enum value "sess_rate"
	StickTableFieldFieldSessRate string = "sess_rate"

	// StickTableFieldFieldHTTPReqCnt captures enum value "http_req_cnt"
	StickTableFieldFieldHTTPReqCnt string = "http_req_cnt"

	// StickTableFieldFieldHTTPReqRate captures enum value "http_req_rate"
	StickTableFieldFieldHTTPReqRate string = "http_req_rate"

	// StickTableFieldFieldHTTPErrCnt captures enum value "http_err_cnt"
	StickTableFieldFieldHTTPErrCnt string = "http_err_cnt"

	// StickTableFieldFieldHTTPErrRate captures enum value "http_err_rate"
	StickTableFieldFieldHTTPErrRate string = "http_err_rate"

	// StickTableFieldFieldBytesInCnt captures enum value "bytes_in_cnt"
	StickTableFieldFieldBytesInCnt string = "bytes_in_cnt"

	// StickTableFieldFieldBytesInRate captures enum value "bytes_in_rate"
	StickTableFieldFieldBytesInRate string = "bytes_in_rate"

	// StickTableFieldFieldBytesOutCnt captures enum value "bytes_out_cnt"
	StickTableFieldFieldBytesOutCnt string = "bytes_out_cnt"

	// StickTableFieldFieldBytesOutRate captures enum value "bytes_out_rate"
	StickTableFieldFieldBytesOutRate string = "bytes_out_rate"
)

// prop value enum
func (m *StickTableField) validateFieldEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, stickTableFieldTypeFieldPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *StickTableField) validateField(formats strfmt.Registry) error {

	if swag.IsZero(m.Field) { // not required
		return nil
	}

	// value enum
	if err := m.validateFieldEnum("field", "body", m.Field); err != nil {
		return err
	}

	return nil
}

var stickTableFieldTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["rate","counter"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		stickTableFieldTypeTypePropEnum = append(stickTableFieldTypeTypePropEnum, v)
	}
}

const (

	// StickTableFieldTypeRate captures enum value "rate"
	StickTableFieldTypeRate string = "rate"

	// StickTableFieldTypeCounter captures enum value "counter"
	StickTableFieldTypeCounter string = "counter"
)

// prop value enum
func (m *StickTableField) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, stickTableFieldTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *StickTableField) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *StickTableField) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StickTableField) UnmarshalBinary(b []byte) error {
	var res StickTableField
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
