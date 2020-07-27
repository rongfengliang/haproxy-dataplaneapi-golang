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
)

// NewGetStickRulesParams creates a new GetStickRulesParams object
// with the default values initialized.
func NewGetStickRulesParams() *GetStickRulesParams {
	var ()
	return &GetStickRulesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetStickRulesParamsWithTimeout creates a new GetStickRulesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetStickRulesParamsWithTimeout(timeout time.Duration) *GetStickRulesParams {
	var ()
	return &GetStickRulesParams{

		timeout: timeout,
	}
}

// NewGetStickRulesParamsWithContext creates a new GetStickRulesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetStickRulesParamsWithContext(ctx context.Context) *GetStickRulesParams {
	var ()
	return &GetStickRulesParams{

		Context: ctx,
	}
}

// NewGetStickRulesParamsWithHTTPClient creates a new GetStickRulesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetStickRulesParamsWithHTTPClient(client *http.Client) *GetStickRulesParams {
	var ()
	return &GetStickRulesParams{
		HTTPClient: client,
	}
}

/*GetStickRulesParams contains all the parameters to send to the API endpoint
for the get stick rules operation typically these are written to a http.Request
*/
type GetStickRulesParams struct {

	/*Backend
	  Backend name

	*/
	Backend string
	/*TransactionID
	  ID of the transaction where we want to add the operation. Cannot be used when version is specified.

	*/
	TransactionID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get stick rules params
func (o *GetStickRulesParams) WithTimeout(timeout time.Duration) *GetStickRulesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get stick rules params
func (o *GetStickRulesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get stick rules params
func (o *GetStickRulesParams) WithContext(ctx context.Context) *GetStickRulesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get stick rules params
func (o *GetStickRulesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get stick rules params
func (o *GetStickRulesParams) WithHTTPClient(client *http.Client) *GetStickRulesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get stick rules params
func (o *GetStickRulesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBackend adds the backend to the get stick rules params
func (o *GetStickRulesParams) WithBackend(backend string) *GetStickRulesParams {
	o.SetBackend(backend)
	return o
}

// SetBackend adds the backend to the get stick rules params
func (o *GetStickRulesParams) SetBackend(backend string) {
	o.Backend = backend
}

// WithTransactionID adds the transactionID to the get stick rules params
func (o *GetStickRulesParams) WithTransactionID(transactionID *string) *GetStickRulesParams {
	o.SetTransactionID(transactionID)
	return o
}

// SetTransactionID adds the transactionId to the get stick rules params
func (o *GetStickRulesParams) SetTransactionID(transactionID *string) {
	o.TransactionID = transactionID
}

// WriteToRequest writes these params to a swagger request
func (o *GetStickRulesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}