// Code generated by go-swagger; DO NOT EDIT.

package h_a_proxy_configuration_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/rongfengliang/haproxy-golang/models"
)

// ReplaceHTTPRequestRuleReader is a Reader for the ReplaceHTTPRequestRule structure.
type ReplaceHTTPRequestRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceHTTPRequestRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReplaceHTTPRequestRuleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewReplaceHTTPRequestRuleAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewReplaceHTTPRequestRuleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewReplaceHTTPRequestRuleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewReplaceHTTPRequestRuleDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewReplaceHTTPRequestRuleOK creates a ReplaceHTTPRequestRuleOK with default headers values
func NewReplaceHTTPRequestRuleOK() *ReplaceHTTPRequestRuleOK {
	return &ReplaceHTTPRequestRuleOK{}
}

/*ReplaceHTTPRequestRuleOK handles this case with default header values.

HTTP Request Rule replaced
*/
type ReplaceHTTPRequestRuleOK struct {
	Payload *models.HTTPRequestRule
}

func (o *ReplaceHTTPRequestRuleOK) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/http_request_rules/{index}][%d] replaceHttpRequestRuleOK  %+v", 200, o.Payload)
}

func (o *ReplaceHTTPRequestRuleOK) GetPayload() *models.HTTPRequestRule {
	return o.Payload
}

func (o *ReplaceHTTPRequestRuleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPRequestRule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceHTTPRequestRuleAccepted creates a ReplaceHTTPRequestRuleAccepted with default headers values
func NewReplaceHTTPRequestRuleAccepted() *ReplaceHTTPRequestRuleAccepted {
	return &ReplaceHTTPRequestRuleAccepted{}
}

/*ReplaceHTTPRequestRuleAccepted handles this case with default header values.

Configuration change accepted and reload requested
*/
type ReplaceHTTPRequestRuleAccepted struct {
	/*ID of the requested reload
	 */
	ReloadID string

	Payload *models.HTTPRequestRule
}

func (o *ReplaceHTTPRequestRuleAccepted) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/http_request_rules/{index}][%d] replaceHttpRequestRuleAccepted  %+v", 202, o.Payload)
}

func (o *ReplaceHTTPRequestRuleAccepted) GetPayload() *models.HTTPRequestRule {
	return o.Payload
}

func (o *ReplaceHTTPRequestRuleAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Reload-ID
	o.ReloadID = response.GetHeader("Reload-ID")

	o.Payload = new(models.HTTPRequestRule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceHTTPRequestRuleBadRequest creates a ReplaceHTTPRequestRuleBadRequest with default headers values
func NewReplaceHTTPRequestRuleBadRequest() *ReplaceHTTPRequestRuleBadRequest {
	return &ReplaceHTTPRequestRuleBadRequest{
		ConfigurationVersion: 0,
	}
}

/*ReplaceHTTPRequestRuleBadRequest handles this case with default header values.

Bad request
*/
type ReplaceHTTPRequestRuleBadRequest struct {
	/*Configuration file version
	 */
	ConfigurationVersion int64

	Payload *models.Error
}

func (o *ReplaceHTTPRequestRuleBadRequest) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/http_request_rules/{index}][%d] replaceHttpRequestRuleBadRequest  %+v", 400, o.Payload)
}

func (o *ReplaceHTTPRequestRuleBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *ReplaceHTTPRequestRuleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	configurationVersion, err := swag.ConvertInt64(response.GetHeader("Configuration-Version"))
	if err != nil {
		return errors.InvalidType("Configuration-Version", "header", "int64", response.GetHeader("Configuration-Version"))
	}
	o.ConfigurationVersion = configurationVersion

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceHTTPRequestRuleNotFound creates a ReplaceHTTPRequestRuleNotFound with default headers values
func NewReplaceHTTPRequestRuleNotFound() *ReplaceHTTPRequestRuleNotFound {
	return &ReplaceHTTPRequestRuleNotFound{
		ConfigurationVersion: 0,
	}
}

/*ReplaceHTTPRequestRuleNotFound handles this case with default header values.

The specified resource was not found
*/
type ReplaceHTTPRequestRuleNotFound struct {
	/*Configuration file version
	 */
	ConfigurationVersion int64

	Payload *models.Error
}

func (o *ReplaceHTTPRequestRuleNotFound) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/http_request_rules/{index}][%d] replaceHttpRequestRuleNotFound  %+v", 404, o.Payload)
}

func (o *ReplaceHTTPRequestRuleNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *ReplaceHTTPRequestRuleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	configurationVersion, err := swag.ConvertInt64(response.GetHeader("Configuration-Version"))
	if err != nil {
		return errors.InvalidType("Configuration-Version", "header", "int64", response.GetHeader("Configuration-Version"))
	}
	o.ConfigurationVersion = configurationVersion

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceHTTPRequestRuleDefault creates a ReplaceHTTPRequestRuleDefault with default headers values
func NewReplaceHTTPRequestRuleDefault(code int) *ReplaceHTTPRequestRuleDefault {
	return &ReplaceHTTPRequestRuleDefault{
		_statusCode:          code,
		ConfigurationVersion: 0,
	}
}

/*ReplaceHTTPRequestRuleDefault handles this case with default header values.

General Error
*/
type ReplaceHTTPRequestRuleDefault struct {
	_statusCode int

	/*Configuration file version
	 */
	ConfigurationVersion int64

	Payload *models.Error
}

// Code gets the status code for the replace HTTP request rule default response
func (o *ReplaceHTTPRequestRuleDefault) Code() int {
	return o._statusCode
}

func (o *ReplaceHTTPRequestRuleDefault) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/http_request_rules/{index}][%d] replaceHTTPRequestRule default  %+v", o._statusCode, o.Payload)
}

func (o *ReplaceHTTPRequestRuleDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *ReplaceHTTPRequestRuleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	configurationVersion, err := swag.ConvertInt64(response.GetHeader("Configuration-Version"))
	if err != nil {
		return errors.InvalidType("Configuration-Version", "header", "int64", response.GetHeader("Configuration-Version"))
	}
	o.ConfigurationVersion = configurationVersion

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
