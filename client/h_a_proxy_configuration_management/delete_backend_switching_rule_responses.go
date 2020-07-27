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

// DeleteBackendSwitchingRuleReader is a Reader for the DeleteBackendSwitchingRule structure.
type DeleteBackendSwitchingRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteBackendSwitchingRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewDeleteBackendSwitchingRuleAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewDeleteBackendSwitchingRuleNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteBackendSwitchingRuleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteBackendSwitchingRuleDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteBackendSwitchingRuleAccepted creates a DeleteBackendSwitchingRuleAccepted with default headers values
func NewDeleteBackendSwitchingRuleAccepted() *DeleteBackendSwitchingRuleAccepted {
	return &DeleteBackendSwitchingRuleAccepted{}
}

/*DeleteBackendSwitchingRuleAccepted handles this case with default header values.

Configuration change accepted and reload requested
*/
type DeleteBackendSwitchingRuleAccepted struct {
	/*ID of the requested reload
	 */
	ReloadID string
}

func (o *DeleteBackendSwitchingRuleAccepted) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/backend_switching_rules/{index}][%d] deleteBackendSwitchingRuleAccepted ", 202)
}

func (o *DeleteBackendSwitchingRuleAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Reload-ID
	o.ReloadID = response.GetHeader("Reload-ID")

	return nil
}

// NewDeleteBackendSwitchingRuleNoContent creates a DeleteBackendSwitchingRuleNoContent with default headers values
func NewDeleteBackendSwitchingRuleNoContent() *DeleteBackendSwitchingRuleNoContent {
	return &DeleteBackendSwitchingRuleNoContent{}
}

/*DeleteBackendSwitchingRuleNoContent handles this case with default header values.

Backend Switching Rule deleted
*/
type DeleteBackendSwitchingRuleNoContent struct {
}

func (o *DeleteBackendSwitchingRuleNoContent) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/backend_switching_rules/{index}][%d] deleteBackendSwitchingRuleNoContent ", 204)
}

func (o *DeleteBackendSwitchingRuleNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteBackendSwitchingRuleNotFound creates a DeleteBackendSwitchingRuleNotFound with default headers values
func NewDeleteBackendSwitchingRuleNotFound() *DeleteBackendSwitchingRuleNotFound {
	return &DeleteBackendSwitchingRuleNotFound{
		ConfigurationVersion: 0,
	}
}

/*DeleteBackendSwitchingRuleNotFound handles this case with default header values.

The specified resource was not found
*/
type DeleteBackendSwitchingRuleNotFound struct {
	/*Configuration file version
	 */
	ConfigurationVersion int64

	Payload *models.Error
}

func (o *DeleteBackendSwitchingRuleNotFound) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/backend_switching_rules/{index}][%d] deleteBackendSwitchingRuleNotFound  %+v", 404, o.Payload)
}

func (o *DeleteBackendSwitchingRuleNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteBackendSwitchingRuleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteBackendSwitchingRuleDefault creates a DeleteBackendSwitchingRuleDefault with default headers values
func NewDeleteBackendSwitchingRuleDefault(code int) *DeleteBackendSwitchingRuleDefault {
	return &DeleteBackendSwitchingRuleDefault{
		_statusCode:          code,
		ConfigurationVersion: 0,
	}
}

/*DeleteBackendSwitchingRuleDefault handles this case with default header values.

General Error
*/
type DeleteBackendSwitchingRuleDefault struct {
	_statusCode int

	/*Configuration file version
	 */
	ConfigurationVersion int64

	Payload *models.Error
}

// Code gets the status code for the delete backend switching rule default response
func (o *DeleteBackendSwitchingRuleDefault) Code() int {
	return o._statusCode
}

func (o *DeleteBackendSwitchingRuleDefault) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/backend_switching_rules/{index}][%d] deleteBackendSwitchingRule default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteBackendSwitchingRuleDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteBackendSwitchingRuleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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