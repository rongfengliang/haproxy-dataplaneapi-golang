// Code generated by go-swagger; DO NOT EDIT.

package h_a_proxy_configuration_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/rongfengliang/haproxy-golang/models"
)

// NewReplaceServerSwitchingRuleParams creates a new ReplaceServerSwitchingRuleParams object
// with the default values initialized.
func NewReplaceServerSwitchingRuleParams() *ReplaceServerSwitchingRuleParams {
	var (
		forceReloadDefault = bool(false)
	)
	return &ReplaceServerSwitchingRuleParams{
		ForceReload: &forceReloadDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewReplaceServerSwitchingRuleParamsWithTimeout creates a new ReplaceServerSwitchingRuleParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReplaceServerSwitchingRuleParamsWithTimeout(timeout time.Duration) *ReplaceServerSwitchingRuleParams {
	var (
		forceReloadDefault = bool(false)
	)
	return &ReplaceServerSwitchingRuleParams{
		ForceReload: &forceReloadDefault,

		timeout: timeout,
	}
}

// NewReplaceServerSwitchingRuleParamsWithContext creates a new ReplaceServerSwitchingRuleParams object
// with the default values initialized, and the ability to set a context for a request
func NewReplaceServerSwitchingRuleParamsWithContext(ctx context.Context) *ReplaceServerSwitchingRuleParams {
	var (
		forceReloadDefault = bool(false)
	)
	return &ReplaceServerSwitchingRuleParams{
		ForceReload: &forceReloadDefault,

		Context: ctx,
	}
}

// NewReplaceServerSwitchingRuleParamsWithHTTPClient creates a new ReplaceServerSwitchingRuleParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReplaceServerSwitchingRuleParamsWithHTTPClient(client *http.Client) *ReplaceServerSwitchingRuleParams {
	var (
		forceReloadDefault = bool(false)
	)
	return &ReplaceServerSwitchingRuleParams{
		ForceReload: &forceReloadDefault,
		HTTPClient:  client,
	}
}

/*ReplaceServerSwitchingRuleParams contains all the parameters to send to the API endpoint
for the replace server switching rule operation typically these are written to a http.Request
*/
type ReplaceServerSwitchingRuleParams struct {

	/*Backend
	  Backend name

	*/
	Backend string
	/*Data*/
	Data *models.ServerSwitchingRule
	/*ForceReload
	  If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration.

	*/
	ForceReload *bool
	/*Index
	  Switching Rule Index

	*/
	Index int64
	/*TransactionID
	  ID of the transaction where we want to add the operation. Cannot be used when version is specified.

	*/
	TransactionID *string
	/*Version
	  Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it's own version.

	*/
	Version *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the replace server switching rule params
func (o *ReplaceServerSwitchingRuleParams) WithTimeout(timeout time.Duration) *ReplaceServerSwitchingRuleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the replace server switching rule params
func (o *ReplaceServerSwitchingRuleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the replace server switching rule params
func (o *ReplaceServerSwitchingRuleParams) WithContext(ctx context.Context) *ReplaceServerSwitchingRuleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the replace server switching rule params
func (o *ReplaceServerSwitchingRuleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the replace server switching rule params
func (o *ReplaceServerSwitchingRuleParams) WithHTTPClient(client *http.Client) *ReplaceServerSwitchingRuleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the replace server switching rule params
func (o *ReplaceServerSwitchingRuleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBackend adds the backend to the replace server switching rule params
func (o *ReplaceServerSwitchingRuleParams) WithBackend(backend string) *ReplaceServerSwitchingRuleParams {
	o.SetBackend(backend)
	return o
}

// SetBackend adds the backend to the replace server switching rule params
func (o *ReplaceServerSwitchingRuleParams) SetBackend(backend string) {
	o.Backend = backend
}

// WithData adds the data to the replace server switching rule params
func (o *ReplaceServerSwitchingRuleParams) WithData(data *models.ServerSwitchingRule) *ReplaceServerSwitchingRuleParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the replace server switching rule params
func (o *ReplaceServerSwitchingRuleParams) SetData(data *models.ServerSwitchingRule) {
	o.Data = data
}

// WithForceReload adds the forceReload to the replace server switching rule params
func (o *ReplaceServerSwitchingRuleParams) WithForceReload(forceReload *bool) *ReplaceServerSwitchingRuleParams {
	o.SetForceReload(forceReload)
	return o
}

// SetForceReload adds the forceReload to the replace server switching rule params
func (o *ReplaceServerSwitchingRuleParams) SetForceReload(forceReload *bool) {
	o.ForceReload = forceReload
}

// WithIndex adds the index to the replace server switching rule params
func (o *ReplaceServerSwitchingRuleParams) WithIndex(index int64) *ReplaceServerSwitchingRuleParams {
	o.SetIndex(index)
	return o
}

// SetIndex adds the index to the replace server switching rule params
func (o *ReplaceServerSwitchingRuleParams) SetIndex(index int64) {
	o.Index = index
}

// WithTransactionID adds the transactionID to the replace server switching rule params
func (o *ReplaceServerSwitchingRuleParams) WithTransactionID(transactionID *string) *ReplaceServerSwitchingRuleParams {
	o.SetTransactionID(transactionID)
	return o
}

// SetTransactionID adds the transactionId to the replace server switching rule params
func (o *ReplaceServerSwitchingRuleParams) SetTransactionID(transactionID *string) {
	o.TransactionID = transactionID
}

// WithVersion adds the version to the replace server switching rule params
func (o *ReplaceServerSwitchingRuleParams) WithVersion(version *int64) *ReplaceServerSwitchingRuleParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the replace server switching rule params
func (o *ReplaceServerSwitchingRuleParams) SetVersion(version *int64) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *ReplaceServerSwitchingRuleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param backend
	qrBackend := o.Backend
	qBackend := qrBackend
	if qBackend != "" {
		if err := r.SetQueryParam("backend", qBackend); err != nil {
			return err
		}
	}

	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	if o.ForceReload != nil {

		// query param force_reload
		var qrForceReload bool
		if o.ForceReload != nil {
			qrForceReload = *o.ForceReload
		}
		qForceReload := swag.FormatBool(qrForceReload)
		if qForceReload != "" {
			if err := r.SetQueryParam("force_reload", qForceReload); err != nil {
				return err
			}
		}

	}

	// path param index
	if err := r.SetPathParam("index", swag.FormatInt64(o.Index)); err != nil {
		return err
	}

	if o.TransactionID != nil {

		// query param transaction_id
		var qrTransactionID string
		if o.TransactionID != nil {
			qrTransactionID = *o.TransactionID
		}
		qTransactionID := qrTransactionID
		if qTransactionID != "" {
			if err := r.SetQueryParam("transaction_id", qTransactionID); err != nil {
				return err
			}
		}

	}

	if o.Version != nil {

		// query param version
		var qrVersion int64
		if o.Version != nil {
			qrVersion = *o.Version
		}
		qVersion := swag.FormatInt64(qrVersion)
		if qVersion != "" {
			if err := r.SetQueryParam("version", qVersion); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
