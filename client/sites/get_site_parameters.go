// Code generated by go-swagger; DO NOT EDIT.

package sites

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

// NewGetSiteParams creates a new GetSiteParams object
// with the default values initialized.
func NewGetSiteParams() *GetSiteParams {
	var ()
	return &GetSiteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSiteParamsWithTimeout creates a new GetSiteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSiteParamsWithTimeout(timeout time.Duration) *GetSiteParams {
	var ()
	return &GetSiteParams{

		timeout: timeout,
	}
}

// NewGetSiteParamsWithContext creates a new GetSiteParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSiteParamsWithContext(ctx context.Context) *GetSiteParams {
	var ()
	return &GetSiteParams{

		Context: ctx,
	}
}

// NewGetSiteParamsWithHTTPClient creates a new GetSiteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSiteParamsWithHTTPClient(client *http.Client) *GetSiteParams {
	var ()
	return &GetSiteParams{
		HTTPClient: client,
	}
}

/*GetSiteParams contains all the parameters to send to the API endpoint
for the get site operation typically these are written to a http.Request
*/
type GetSiteParams struct {

	/*Name
	  Site frontend name

	*/
	Name string
	/*TransactionID
	  ID of the transaction where we want to add the operation. Cannot be used when version is specified.

	*/
	TransactionID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get site params
func (o *GetSiteParams) WithTimeout(timeout time.Duration) *GetSiteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get site params
func (o *GetSiteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get site params
func (o *GetSiteParams) WithContext(ctx context.Context) *GetSiteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get site params
func (o *GetSiteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get site params
func (o *GetSiteParams) WithHTTPClient(client *http.Client) *GetSiteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get site params
func (o *GetSiteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the get site params
func (o *GetSiteParams) WithName(name string) *GetSiteParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get site params
func (o *GetSiteParams) SetName(name string) {
	o.Name = name
}

// WithTransactionID adds the transactionID to the get site params
func (o *GetSiteParams) WithTransactionID(transactionID *string) *GetSiteParams {
	o.SetTransactionID(transactionID)
	return o
}

// SetTransactionID adds the transactionId to the get site params
func (o *GetSiteParams) SetTransactionID(transactionID *string) {
	o.TransactionID = transactionID
}

// WriteToRequest writes these params to a swagger request
func (o *GetSiteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
