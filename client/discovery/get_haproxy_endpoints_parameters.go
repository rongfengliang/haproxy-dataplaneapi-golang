// Code generated by go-swagger; DO NOT EDIT.

package discovery

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

// NewGetHaproxyEndpointsParams creates a new GetHaproxyEndpointsParams object
// with the default values initialized.
func NewGetHaproxyEndpointsParams() *GetHaproxyEndpointsParams {

	return &GetHaproxyEndpointsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetHaproxyEndpointsParamsWithTimeout creates a new GetHaproxyEndpointsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetHaproxyEndpointsParamsWithTimeout(timeout time.Duration) *GetHaproxyEndpointsParams {

	return &GetHaproxyEndpointsParams{

		timeout: timeout,
	}
}

// NewGetHaproxyEndpointsParamsWithContext creates a new GetHaproxyEndpointsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetHaproxyEndpointsParamsWithContext(ctx context.Context) *GetHaproxyEndpointsParams {

	return &GetHaproxyEndpointsParams{

		Context: ctx,
	}
}

// NewGetHaproxyEndpointsParamsWithHTTPClient creates a new GetHaproxyEndpointsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetHaproxyEndpointsParamsWithHTTPClient(client *http.Client) *GetHaproxyEndpointsParams {

	return &GetHaproxyEndpointsParams{
		HTTPClient: client,
	}
}

/*GetHaproxyEndpointsParams contains all the parameters to send to the API endpoint
for the get haproxy endpoints operation typically these are written to a http.Request
*/
type GetHaproxyEndpointsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get haproxy endpoints params
func (o *GetHaproxyEndpointsParams) WithTimeout(timeout time.Duration) *GetHaproxyEndpointsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get haproxy endpoints params
func (o *GetHaproxyEndpointsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get haproxy endpoints params
func (o *GetHaproxyEndpointsParams) WithContext(ctx context.Context) *GetHaproxyEndpointsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get haproxy endpoints params
func (o *GetHaproxyEndpointsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get haproxy endpoints params
func (o *GetHaproxyEndpointsParams) WithHTTPClient(client *http.Client) *GetHaproxyEndpointsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get haproxy endpoints params
func (o *GetHaproxyEndpointsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetHaproxyEndpointsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
