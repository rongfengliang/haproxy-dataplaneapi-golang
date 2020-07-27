// Code generated by go-swagger; DO NOT EDIT.

package peer

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

// NewCreatePeerParams creates a new CreatePeerParams object
// with the default values initialized.
func NewCreatePeerParams() *CreatePeerParams {
	var (
		forceReloadDefault = bool(false)
	)
	return &CreatePeerParams{
		ForceReload: &forceReloadDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewCreatePeerParamsWithTimeout creates a new CreatePeerParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreatePeerParamsWithTimeout(timeout time.Duration) *CreatePeerParams {
	var (
		forceReloadDefault = bool(false)
	)
	return &CreatePeerParams{
		ForceReload: &forceReloadDefault,

		timeout: timeout,
	}
}

// NewCreatePeerParamsWithContext creates a new CreatePeerParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreatePeerParamsWithContext(ctx context.Context) *CreatePeerParams {
	var (
		forceReloadDefault = bool(false)
	)
	return &CreatePeerParams{
		ForceReload: &forceReloadDefault,

		Context: ctx,
	}
}

// NewCreatePeerParamsWithHTTPClient creates a new CreatePeerParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreatePeerParamsWithHTTPClient(client *http.Client) *CreatePeerParams {
	var (
		forceReloadDefault = bool(false)
	)
	return &CreatePeerParams{
		ForceReload: &forceReloadDefault,
		HTTPClient:  client,
	}
}

/*CreatePeerParams contains all the parameters to send to the API endpoint
for the create peer operation typically these are written to a http.Request
*/
type CreatePeerParams struct {

	/*Data*/
	Data *models.PeerSection
	/*ForceReload
	  If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration.

	*/
	ForceReload *bool
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

// WithTimeout adds the timeout to the create peer params
func (o *CreatePeerParams) WithTimeout(timeout time.Duration) *CreatePeerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create peer params
func (o *CreatePeerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create peer params
func (o *CreatePeerParams) WithContext(ctx context.Context) *CreatePeerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create peer params
func (o *CreatePeerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create peer params
func (o *CreatePeerParams) WithHTTPClient(client *http.Client) *CreatePeerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create peer params
func (o *CreatePeerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the create peer params
func (o *CreatePeerParams) WithData(data *models.PeerSection) *CreatePeerParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the create peer params
func (o *CreatePeerParams) SetData(data *models.PeerSection) {
	o.Data = data
}

// WithForceReload adds the forceReload to the create peer params
func (o *CreatePeerParams) WithForceReload(forceReload *bool) *CreatePeerParams {
	o.SetForceReload(forceReload)
	return o
}

// SetForceReload adds the forceReload to the create peer params
func (o *CreatePeerParams) SetForceReload(forceReload *bool) {
	o.ForceReload = forceReload
}

// WithTransactionID adds the transactionID to the create peer params
func (o *CreatePeerParams) WithTransactionID(transactionID *string) *CreatePeerParams {
	o.SetTransactionID(transactionID)
	return o
}

// SetTransactionID adds the transactionId to the create peer params
func (o *CreatePeerParams) SetTransactionID(transactionID *string) {
	o.TransactionID = transactionID
}

// WithVersion adds the version to the create peer params
func (o *CreatePeerParams) WithVersion(version *int64) *CreatePeerParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the create peer params
func (o *CreatePeerParams) SetVersion(version *int64) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *CreatePeerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
