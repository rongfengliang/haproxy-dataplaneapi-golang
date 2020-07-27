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
)

// NewDeleteACLParams creates a new DeleteACLParams object
// with the default values initialized.
func NewDeleteACLParams() *DeleteACLParams {
	var (
		forceReloadDefault = bool(false)
	)
	return &DeleteACLParams{
		ForceReload: &forceReloadDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteACLParamsWithTimeout creates a new DeleteACLParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteACLParamsWithTimeout(timeout time.Duration) *DeleteACLParams {
	var (
		forceReloadDefault = bool(false)
	)
	return &DeleteACLParams{
		ForceReload: &forceReloadDefault,

		timeout: timeout,
	}
}

// NewDeleteACLParamsWithContext creates a new DeleteACLParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteACLParamsWithContext(ctx context.Context) *DeleteACLParams {
	var (
		forceReloadDefault = bool(false)
	)
	return &DeleteACLParams{
		ForceReload: &forceReloadDefault,

		Context: ctx,
	}
}

// NewDeleteACLParamsWithHTTPClient creates a new DeleteACLParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteACLParamsWithHTTPClient(client *http.Client) *DeleteACLParams {
	var (
		forceReloadDefault = bool(false)
	)
	return &DeleteACLParams{
		ForceReload: &forceReloadDefault,
		HTTPClient:  client,
	}
}

/*DeleteACLParams contains all the parameters to send to the API endpoint
for the delete Acl operation typically these are written to a http.Request
*/
type DeleteACLParams struct {

	/*ForceReload
	  If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration.

	*/
	ForceReload *bool
	/*Index
	  ACL line Index

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

// WithTimeout adds the timeout to the delete Acl params
func (o *DeleteACLParams) WithTimeout(timeout time.Duration) *DeleteACLParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete Acl params
func (o *DeleteACLParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete Acl params
func (o *DeleteACLParams) WithContext(ctx context.Context) *DeleteACLParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete Acl params
func (o *DeleteACLParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete Acl params
func (o *DeleteACLParams) WithHTTPClient(client *http.Client) *DeleteACLParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete Acl params
func (o *DeleteACLParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithForceReload adds the forceReload to the delete Acl params
func (o *DeleteACLParams) WithForceReload(forceReload *bool) *DeleteACLParams {
	o.SetForceReload(forceReload)
	return o
}

// SetForceReload adds the forceReload to the delete Acl params
func (o *DeleteACLParams) SetForceReload(forceReload *bool) {
	o.ForceReload = forceReload
}

// WithIndex adds the index to the delete Acl params
func (o *DeleteACLParams) WithIndex(index int64) *DeleteACLParams {
	o.SetIndex(index)
	return o
}

// SetIndex adds the index to the delete Acl params
func (o *DeleteACLParams) SetIndex(index int64) {
	o.Index = index
}

// WithParentName adds the parentName to the delete Acl params
func (o *DeleteACLParams) WithParentName(parentName string) *DeleteACLParams {
	o.SetParentName(parentName)
	return o
}

// SetParentName adds the parentName to the delete Acl params
func (o *DeleteACLParams) SetParentName(parentName string) {
	o.ParentName = parentName
}

// WithParentType adds the parentType to the delete Acl params
func (o *DeleteACLParams) WithParentType(parentType string) *DeleteACLParams {
	o.SetParentType(parentType)
	return o
}

// SetParentType adds the parentType to the delete Acl params
func (o *DeleteACLParams) SetParentType(parentType string) {
	o.ParentType = parentType
}

// WithTransactionID adds the transactionID to the delete Acl params
func (o *DeleteACLParams) WithTransactionID(transactionID *string) *DeleteACLParams {
	o.SetTransactionID(transactionID)
	return o
}

// SetTransactionID adds the transactionId to the delete Acl params
func (o *DeleteACLParams) SetTransactionID(transactionID *string) {
	o.TransactionID = transactionID
}

// WithVersion adds the version to the delete Acl params
func (o *DeleteACLParams) WithVersion(version *int64) *DeleteACLParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the delete Acl params
func (o *DeleteACLParams) SetVersion(version *int64) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteACLParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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
