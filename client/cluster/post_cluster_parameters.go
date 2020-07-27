// Code generated by go-swagger; DO NOT EDIT.

package cluster

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

// NewPostClusterParams creates a new PostClusterParams object
// with the default values initialized.
func NewPostClusterParams() *PostClusterParams {
	var ()
	return &PostClusterParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostClusterParamsWithTimeout creates a new PostClusterParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostClusterParamsWithTimeout(timeout time.Duration) *PostClusterParams {
	var ()
	return &PostClusterParams{

		timeout: timeout,
	}
}

// NewPostClusterParamsWithContext creates a new PostClusterParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostClusterParamsWithContext(ctx context.Context) *PostClusterParams {
	var ()
	return &PostClusterParams{

		Context: ctx,
	}
}

// NewPostClusterParamsWithHTTPClient creates a new PostClusterParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostClusterParamsWithHTTPClient(client *http.Client) *PostClusterParams {
	var ()
	return &PostClusterParams{
		HTTPClient: client,
	}
}

/*PostClusterParams contains all the parameters to send to the API endpoint
for the post cluster operation typically these are written to a http.Request
*/
type PostClusterParams struct {

	/*Configuration
	  In case of moving to single mode do we keep or clean configuration

	*/
	Configuration *string
	/*Data*/
	Data *models.ClusterSettings
	/*Version
	  Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it's own version.

	*/
	Version *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post cluster params
func (o *PostClusterParams) WithTimeout(timeout time.Duration) *PostClusterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post cluster params
func (o *PostClusterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post cluster params
func (o *PostClusterParams) WithContext(ctx context.Context) *PostClusterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post cluster params
func (o *PostClusterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post cluster params
func (o *PostClusterParams) WithHTTPClient(client *http.Client) *PostClusterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post cluster params
func (o *PostClusterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConfiguration adds the configuration to the post cluster params
func (o *PostClusterParams) WithConfiguration(configuration *string) *PostClusterParams {
	o.SetConfiguration(configuration)
	return o
}

// SetConfiguration adds the configuration to the post cluster params
func (o *PostClusterParams) SetConfiguration(configuration *string) {
	o.Configuration = configuration
}

// WithData adds the data to the post cluster params
func (o *PostClusterParams) WithData(data *models.ClusterSettings) *PostClusterParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the post cluster params
func (o *PostClusterParams) SetData(data *models.ClusterSettings) {
	o.Data = data
}

// WithVersion adds the version to the post cluster params
func (o *PostClusterParams) WithVersion(version *int64) *PostClusterParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the post cluster params
func (o *PostClusterParams) SetVersion(version *int64) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *PostClusterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Configuration != nil {

		// query param configuration
		var qrConfiguration string
		if o.Configuration != nil {
			qrConfiguration = *o.Configuration
		}
		qConfiguration := qrConfiguration
		if qConfiguration != "" {
			if err := r.SetQueryParam("configuration", qConfiguration); err != nil {
				return err
			}
		}

	}

	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
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
