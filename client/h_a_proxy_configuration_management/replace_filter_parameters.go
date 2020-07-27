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

// NewReplaceFilterParams creates a new ReplaceFilterParams object
// with the default values initialized.
func NewReplaceFilterParams() *ReplaceFilterParams {
	var (
		forceReloadDefault = bool(false)
	)
	return &ReplaceFilterParams{
		ForceReload: &forceReloadDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewReplaceFilterParamsWithTimeout creates a new ReplaceFilterParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReplaceFilterParamsWithTimeout(timeout time.Duration) *ReplaceFilterParams {
	var (
		forceReloadDefault = bool(false)
	)
	return &ReplaceFilterParams{
		ForceReload: &forceReloadDefault,

		timeout: timeout,
	}
}

// NewReplaceFilterParamsWithContext creates a new ReplaceFilterParams object
// with the default values initialized, and the ability to set a context for a request
func NewReplaceFilterParamsWithContext(ctx context.Context) *ReplaceFilterParams {
	var (
		forceReloadDefault = bool(false)
	)
	return &ReplaceFilterParams{
		ForceReload: &forceReloadDefault,

		Context: ctx,
	}
}

// NewReplaceFilterParamsWithHTTPClient creates a new ReplaceFilterParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReplaceFilterParamsWithHTTPClient(client *http.Client) *ReplaceFilterParams {
	var (
		forceReloadDefault = bool(false)
	)
	return &ReplaceFilterParams{
		ForceReload: &forceReloadDefault,
		HTTPClient:  client,
	}
}

/*ReplaceFilterParams contains all the parameters to send to the API endpoint
for the replace filter operation typically these are written to a http.Request
*/
type ReplaceFilterParams struct {

	/*Data*/
	Data *models.Filter
	/*ForceReload
	  If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration.

	*/
	ForceReload *bool
	/*Index
	  Filter Index

	*/
	Index int64
	/*ParentName
	  Parent name

	*/
	ParentName string
	/*ParentType
	  Parent type

	*/
	ParentType string
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

// WithTimeout adds the timeout to the replace filter params
func (o *ReplaceFilterParams) WithTimeout(timeout time.Duration) *ReplaceFilterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the replace filter params
func (o *ReplaceFilterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the replace filter params
func (o *ReplaceFilterParams) WithContext(ctx context.Context) *ReplaceFilterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the replace filter params
func (o *ReplaceFilterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the replace filter params
func (o *ReplaceFilterParams) WithHTTPClient(client *http.Client) *ReplaceFilterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the replace filter params
func (o *ReplaceFilterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the replace filter params
func (o *ReplaceFilterParams) WithData(data *models.Filter) *ReplaceFilterParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the replace filter params
func (o *ReplaceFilterParams) SetData(data *models.Filter) {
	o.Data = data
}

// WithForceReload adds the forceReload to the replace filter params
func (o *ReplaceFilterParams) WithForceReload(forceReload *bool) *ReplaceFilterParams {
	o.SetForceReload(forceReload)
	return o
}

// SetForceReload adds the forceReload to the replace filter params
func (o *ReplaceFilterParams) SetForceReload(forceReload *bool) {
	o.ForceReload = forceReload
}

// WithIndex adds the index to the replace filter params
func (o *ReplaceFilterParams) WithIndex(index int64) *ReplaceFilterParams {
	o.SetIndex(index)
	return o
}

// SetIndex adds the index to the replace filter params
func (o *ReplaceFilterParams) SetIndex(index int64) {
	o.Index = index
}

// WithParentName adds the parentName to the replace filter params
func (o *ReplaceFilterParams) WithParentName(parentName string) *ReplaceFilterParams {
	o.SetParentName(parentName)
	return o
}

// SetParentName adds the parentName to the replace filter params
func (o *ReplaceFilterParams) SetParentName(parentName string) {
	o.ParentName = parentName
}

// WithParentType adds the parentType to the replace filter params
func (o *ReplaceFilterParams) WithParentType(parentType string) *ReplaceFilterParams {
	o.SetParentType(parentType)
	return o
}

// SetParentType adds the parentType to the replace filter params
func (o *ReplaceFilterParams) SetParentType(parentType string) {
	o.ParentType = parentType
}

// WithTransactionID adds the transactionID to the replace filter params
func (o *ReplaceFilterParams) WithTransactionID(transactionID *string) *ReplaceFilterParams {
	o.SetTransactionID(transactionID)
	return o
}

// SetTransactionID adds the transactionId to the replace filter params
func (o *ReplaceFilterParams) SetTransactionID(transactionID *string) {
	o.TransactionID = transactionID
}

// WithVersion adds the version to the replace filter params
func (o *ReplaceFilterParams) WithVersion(version *int64) *ReplaceFilterParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the replace filter params
func (o *ReplaceFilterParams) SetVersion(version *int64) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *ReplaceFilterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	// query param parent_name
	qrParentName := o.ParentName
	qParentName := qrParentName
	if qParentName != "" {
		if err := r.SetQueryParam("parent_name", qParentName); err != nil {
			return err
		}
	}

	// query param parent_type
	qrParentType := o.ParentType
	qParentType := qrParentType
	if qParentType != "" {
		if err := r.SetQueryParam("parent_type", qParentType); err != nil {
			return err
		}
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