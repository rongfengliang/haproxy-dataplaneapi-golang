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

// GetBackendSwitchingRuleReader is a Reader for the GetBackendSwitchingRule structure.
type GetBackendSwitchingRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBackendSwitchingRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetBackendSwitchingRuleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetBackendSwitchingRuleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetBackendSwitchingRuleDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetBackendSwitchingRuleOK creates a GetBackendSwitchingRuleOK with default headers values
func NewGetBackendSwitchingRuleOK() *GetBackendSwitchingRuleOK {
	return &GetBackendSwitchingRuleOK{}
}

/*GetBackendSwitchingRuleOK handles this case with default header values.

Successful operation
*/
type GetBackendSwitchingRuleOK struct {
	/*Configuration file version
	 */
	ConfigurationVersion int64

	Payload *GetBackendSwitchingRuleOKBody
}

func (o *GetBackendSwitchingRuleOK) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/backend_switching_rules/{index}][%d] getBackendSwitchingRuleOK  %+v", 200, o.Payload)
}

func (o *GetBackendSwitchingRuleOK) GetPayload() *GetBackendSwitchingRuleOKBody {
	return o.Payload
}

func (o *GetBackendSwitchingRuleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	configurationVersion, err := swag.ConvertInt64(response.GetHeader("Configuration-Version"))
	if err != nil {
		return errors.InvalidType("Configuration-Version", "header", "int64", response.GetHeader("Configuration-Version"))
	}
	o.ConfigurationVersion = configurationVersion

	o.Payload = new(GetBackendSwitchingRuleOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBackendSwitchingRuleNotFound creates a GetBackendSwitchingRuleNotFound with default headers values
func NewGetBackendSwitchingRuleNotFound() *GetBackendSwitchingRuleNotFound {
	return &GetBackendSwitchingRuleNotFound{
		ConfigurationVersion: 0,
	}
}

/*GetBackendSwitchingRuleNotFound handles this case with default header values.

The specified resource was not found
*/
type GetBackendSwitchingRuleNotFound struct {
	/*Configuration file version
	 */
	ConfigurationVersion int64

	Payload *models.Error
}

func (o *GetBackendSwitchingRuleNotFound) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/backend_switching_rules/{index}][%d] getBackendSwitchingRuleNotFound  %+v", 404, o.Payload)
}

func (o *GetBackendSwitchingRuleNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetBackendSwitchingRuleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetBackendSwitchingRuleDefault creates a GetBackendSwitchingRuleDefault with default headers values
func NewGetBackendSwitchingRuleDefault(code int) *GetBackendSwitchingRuleDefault {
	return &GetBackendSwitchingRuleDefault{
		_statusCode:          code,
		ConfigurationVersion: 0,
	}
}

/*GetBackendSwitchingRuleDefault handles this case with default header values.

General Error
*/
type GetBackendSwitchingRuleDefault struct {
	_statusCode int

	/*Configuration file version
	 */
	ConfigurationVersion int64

	Payload *models.Error
}

// Code gets the status code for the get backend switching rule default response
func (o *GetBackendSwitchingRuleDefault) Code() int {
	return o._statusCode
}

func (o *GetBackendSwitchingRuleDefault) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/backend_switching_rules/{index}][%d] getBackendSwitchingRule default  %+v", o._statusCode, o.Payload)
}

func (o *GetBackendSwitchingRuleDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetBackendSwitchingRuleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

/*GetBackendSwitchingRuleOKBody get backend switching rule o k body
swagger:model GetBackendSwitchingRuleOKBody
*/
type GetBackendSwitchingRuleOKBody struct {

	// version
	Version int64 `json:"_version,omitempty"`

	// data
	Data *models.BackendSwitchingRule `json:"data,omitempty"`
}

// Validate validates this get backend switching rule o k body
func (o *GetBackendSwitchingRuleOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetBackendSwitchingRuleOKBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getBackendSwitchingRuleOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetBackendSwitchingRuleOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetBackendSwitchingRuleOKBody) UnmarshalBinary(b []byte) error {
	var res GetBackendSwitchingRuleOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
