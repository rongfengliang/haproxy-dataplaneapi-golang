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

// NativeStatStats native stat stats
//
// swagger:model native_stat_stats
type NativeStatStats struct {

	// act
	Act *int64 `json:"act,omitempty"`

	// addr
	Addr string `json:"addr,omitempty"`

	// agent code
	AgentCode *int64 `json:"agent_code,omitempty"`

	// agent desc
	AgentDesc string `json:"agent_desc,omitempty"`

	// agent duration
	AgentDuration *int64 `json:"agent_duration,omitempty"`

	// agent fall
	AgentFall *int64 `json:"agent_fall,omitempty"`

	// agent health
	AgentHealth *int64 `json:"agent_health,omitempty"`

	// agent rise
	AgentRise *int64 `json:"agent_rise,omitempty"`

	// agent status
	// Enum: [UNK INI SOCKERR L40K L4TOUT L4CON L7OK L7STS]
	AgentStatus string `json:"agent_status,omitempty"`

	// algo
	Algo string `json:"algo,omitempty"`

	// bck
	Bck *int64 `json:"bck,omitempty"`

	// bin
	Bin *int64 `json:"bin,omitempty"`

	// bout
	Bout *int64 `json:"bout,omitempty"`

	// check code
	CheckCode *int64 `json:"check_code,omitempty"`

	// check desc
	CheckDesc string `json:"check_desc,omitempty"`

	// check duration
	CheckDuration *int64 `json:"check_duration,omitempty"`

	// check fall
	CheckFall *int64 `json:"check_fall,omitempty"`

	// check health
	CheckHealth *int64 `json:"check_health,omitempty"`

	// check rise
	CheckRise *int64 `json:"check_rise,omitempty"`

	// check status
	// Enum: [UNK INI SOCKERR L40K L4TOUT L4CON L6OK L6TOUT L6RSP L7OK L7OKC L7TOUT L7RSP L7STS]
	CheckStatus string `json:"check_status,omitempty"`

	// chkdown
	Chkdown *int64 `json:"chkdown,omitempty"`

	// chkfail
	Chkfail *int64 `json:"chkfail,omitempty"`

	// cli abrt
	CliAbrt *int64 `json:"cli_abrt,omitempty"`

	// comp byp
	CompByp *int64 `json:"comp_byp,omitempty"`

	// comp in
	CompIn *int64 `json:"comp_in,omitempty"`

	// comp out
	CompOut *int64 `json:"comp_out,omitempty"`

	// comp rsp
	CompRsp *int64 `json:"comp_rsp,omitempty"`

	// conn rate
	ConnRate *int64 `json:"conn_rate,omitempty"`

	// conn rate max
	ConnRateMax *int64 `json:"conn_rate_max,omitempty"`

	// conn tot
	ConnTot *int64 `json:"conn_tot,omitempty"`

	// cookie
	Cookie string `json:"cookie,omitempty"`

	// ctime
	Ctime *int64 `json:"ctime,omitempty"`

	// dcon
	Dcon *int64 `json:"dcon,omitempty"`

	// downtime
	Downtime *int64 `json:"downtime,omitempty"`

	// dreq
	Dreq *int64 `json:"dreq,omitempty"`

	// dresp
	Dresp *int64 `json:"dresp,omitempty"`

	// dses
	Dses *int64 `json:"dses,omitempty"`

	// econ
	Econ *int64 `json:"econ,omitempty"`

	// ereq
	Ereq *int64 `json:"ereq,omitempty"`

	// eresp
	Eresp *int64 `json:"eresp,omitempty"`

	// hanafail
	Hanafail string `json:"hanafail,omitempty"`

	// hrsp 1xx
	Hrsp1xx *int64 `json:"hrsp_1xx,omitempty"`

	// hrsp 2xx
	Hrsp2xx *int64 `json:"hrsp_2xx,omitempty"`

	// hrsp 3xx
	Hrsp3xx *int64 `json:"hrsp_3xx,omitempty"`

	// hrsp 4xx
	Hrsp4xx *int64 `json:"hrsp_4xx,omitempty"`

	// hrsp 5xx
	Hrsp5xx *int64 `json:"hrsp_5xx,omitempty"`

	// hrsp other
	HrspOther *int64 `json:"hrsp_other,omitempty"`

	// iid
	Iid *int64 `json:"iid,omitempty"`

	// intercepted
	Intercepted *int64 `json:"intercepted,omitempty"`

	// lastchg
	Lastchg *int64 `json:"lastchg,omitempty"`

	// lastsess
	Lastsess *int64 `json:"lastsess,omitempty"`

	// lbtot
	Lbtot *int64 `json:"lbtot,omitempty"`

	// mode
	// Enum: [tcp http health unknown]
	Mode string `json:"mode,omitempty"`

	// pid
	Pid *int64 `json:"pid,omitempty"`

	// qcur
	Qcur *int64 `json:"qcur,omitempty"`

	// qlimit
	Qlimit *int64 `json:"qlimit,omitempty"`

	// qmax
	Qmax *int64 `json:"qmax,omitempty"`

	// qtime
	Qtime *int64 `json:"qtime,omitempty"`

	// rate
	Rate *int64 `json:"rate,omitempty"`

	// rate lim
	RateLim *int64 `json:"rate_lim,omitempty"`

	// rate max
	RateMax *int64 `json:"rate_max,omitempty"`

	// req rate
	ReqRate *int64 `json:"req_rate,omitempty"`

	// req rate max
	ReqRateMax *int64 `json:"req_rate_max,omitempty"`

	// req tot
	ReqTot *int64 `json:"req_tot,omitempty"`

	// rtime
	Rtime *int64 `json:"rtime,omitempty"`

	// scur
	Scur *int64 `json:"scur,omitempty"`

	// sid
	Sid *int64 `json:"sid,omitempty"`

	// slim
	Slim *int64 `json:"slim,omitempty"`

	// smax
	Smax *int64 `json:"smax,omitempty"`

	// srv abrt
	SrvAbrt *int64 `json:"srv_abrt,omitempty"`

	// status
	// Enum: [UP DOWN NOLB MAINT no check]
	Status string `json:"status,omitempty"`

	// stot
	Stot *int64 `json:"stot,omitempty"`

	// throttle
	Throttle *int64 `json:"throttle,omitempty"`

	// tracked
	Tracked string `json:"tracked,omitempty"`

	// ttime
	Ttime *int64 `json:"ttime,omitempty"`

	// weight
	Weight *int64 `json:"weight,omitempty"`

	// wredis
	Wredis *int64 `json:"wredis,omitempty"`

	// wretr
	Wretr *int64 `json:"wretr,omitempty"`
}

// Validate validates this native stat stats
func (m *NativeStatStats) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAgentStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCheckStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var nativeStatStatsTypeAgentStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["UNK","INI","SOCKERR","L40K","L4TOUT","L4CON","L7OK","L7STS"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		nativeStatStatsTypeAgentStatusPropEnum = append(nativeStatStatsTypeAgentStatusPropEnum, v)
	}
}

const (

	// NativeStatStatsAgentStatusUNK captures enum value "UNK"
	NativeStatStatsAgentStatusUNK string = "UNK"

	// NativeStatStatsAgentStatusINI captures enum value "INI"
	NativeStatStatsAgentStatusINI string = "INI"

	// NativeStatStatsAgentStatusSOCKERR captures enum value "SOCKERR"
	NativeStatStatsAgentStatusSOCKERR string = "SOCKERR"

	// NativeStatStatsAgentStatusL40K captures enum value "L40K"
	NativeStatStatsAgentStatusL40K string = "L40K"

	// NativeStatStatsAgentStatusL4TOUT captures enum value "L4TOUT"
	NativeStatStatsAgentStatusL4TOUT string = "L4TOUT"

	// NativeStatStatsAgentStatusL4CON captures enum value "L4CON"
	NativeStatStatsAgentStatusL4CON string = "L4CON"

	// NativeStatStatsAgentStatusL7OK captures enum value "L7OK"
	NativeStatStatsAgentStatusL7OK string = "L7OK"

	// NativeStatStatsAgentStatusL7STS captures enum value "L7STS"
	NativeStatStatsAgentStatusL7STS string = "L7STS"
)

// prop value enum
func (m *NativeStatStats) validateAgentStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, nativeStatStatsTypeAgentStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *NativeStatStats) validateAgentStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.AgentStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateAgentStatusEnum("agent_status", "body", m.AgentStatus); err != nil {
		return err
	}

	return nil
}

var nativeStatStatsTypeCheckStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["UNK","INI","SOCKERR","L40K","L4TOUT","L4CON","L6OK","L6TOUT","L6RSP","L7OK","L7OKC","L7TOUT","L7RSP","L7STS"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		nativeStatStatsTypeCheckStatusPropEnum = append(nativeStatStatsTypeCheckStatusPropEnum, v)
	}
}

const (

	// NativeStatStatsCheckStatusUNK captures enum value "UNK"
	NativeStatStatsCheckStatusUNK string = "UNK"

	// NativeStatStatsCheckStatusINI captures enum value "INI"
	NativeStatStatsCheckStatusINI string = "INI"

	// NativeStatStatsCheckStatusSOCKERR captures enum value "SOCKERR"
	NativeStatStatsCheckStatusSOCKERR string = "SOCKERR"

	// NativeStatStatsCheckStatusL40K captures enum value "L40K"
	NativeStatStatsCheckStatusL40K string = "L40K"

	// NativeStatStatsCheckStatusL4TOUT captures enum value "L4TOUT"
	NativeStatStatsCheckStatusL4TOUT string = "L4TOUT"

	// NativeStatStatsCheckStatusL4CON captures enum value "L4CON"
	NativeStatStatsCheckStatusL4CON string = "L4CON"

	// NativeStatStatsCheckStatusL6OK captures enum value "L6OK"
	NativeStatStatsCheckStatusL6OK string = "L6OK"

	// NativeStatStatsCheckStatusL6TOUT captures enum value "L6TOUT"
	NativeStatStatsCheckStatusL6TOUT string = "L6TOUT"

	// NativeStatStatsCheckStatusL6RSP captures enum value "L6RSP"
	NativeStatStatsCheckStatusL6RSP string = "L6RSP"

	// NativeStatStatsCheckStatusL7OK captures enum value "L7OK"
	NativeStatStatsCheckStatusL7OK string = "L7OK"

	// NativeStatStatsCheckStatusL7OKC captures enum value "L7OKC"
	NativeStatStatsCheckStatusL7OKC string = "L7OKC"

	// NativeStatStatsCheckStatusL7TOUT captures enum value "L7TOUT"
	NativeStatStatsCheckStatusL7TOUT string = "L7TOUT"

	// NativeStatStatsCheckStatusL7RSP captures enum value "L7RSP"
	NativeStatStatsCheckStatusL7RSP string = "L7RSP"

	// NativeStatStatsCheckStatusL7STS captures enum value "L7STS"
	NativeStatStatsCheckStatusL7STS string = "L7STS"
)

// prop value enum
func (m *NativeStatStats) validateCheckStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, nativeStatStatsTypeCheckStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *NativeStatStats) validateCheckStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.CheckStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateCheckStatusEnum("check_status", "body", m.CheckStatus); err != nil {
		return err
	}

	return nil
}

var nativeStatStatsTypeModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["tcp","http","health","unknown"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		nativeStatStatsTypeModePropEnum = append(nativeStatStatsTypeModePropEnum, v)
	}
}

const (

	// NativeStatStatsModeTCP captures enum value "tcp"
	NativeStatStatsModeTCP string = "tcp"

	// NativeStatStatsModeHTTP captures enum value "http"
	NativeStatStatsModeHTTP string = "http"

	// NativeStatStatsModeHealth captures enum value "health"
	NativeStatStatsModeHealth string = "health"

	// NativeStatStatsModeUnknown captures enum value "unknown"
	NativeStatStatsModeUnknown string = "unknown"
)

// prop value enum
func (m *NativeStatStats) validateModeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, nativeStatStatsTypeModePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *NativeStatStats) validateMode(formats strfmt.Registry) error {

	if swag.IsZero(m.Mode) { // not required
		return nil
	}

	// value enum
	if err := m.validateModeEnum("mode", "body", m.Mode); err != nil {
		return err
	}

	return nil
}

var nativeStatStatsTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["UP","DOWN","NOLB","MAINT","no check"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		nativeStatStatsTypeStatusPropEnum = append(nativeStatStatsTypeStatusPropEnum, v)
	}
}

const (

	// NativeStatStatsStatusUP captures enum value "UP"
	NativeStatStatsStatusUP string = "UP"

	// NativeStatStatsStatusDOWN captures enum value "DOWN"
	NativeStatStatsStatusDOWN string = "DOWN"

	// NativeStatStatsStatusNOLB captures enum value "NOLB"
	NativeStatStatsStatusNOLB string = "NOLB"

	// NativeStatStatsStatusMAINT captures enum value "MAINT"
	NativeStatStatsStatusMAINT string = "MAINT"

	// NativeStatStatsStatusNoCheck captures enum value "no check"
	NativeStatStatsStatusNoCheck string = "no check"
)

// prop value enum
func (m *NativeStatStats) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, nativeStatStatsTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *NativeStatStats) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NativeStatStats) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NativeStatStats) UnmarshalBinary(b []byte) error {
	var res NativeStatStats
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
