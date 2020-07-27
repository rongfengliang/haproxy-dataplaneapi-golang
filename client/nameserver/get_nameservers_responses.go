// Code generated by go-swagger; DO NOT EDIT.

package nameserver

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/rongfengliang/haproxy-golang/models"
)

// GetNameserversReader is a Reader for the GetNameservers structure.
type GetNameserversReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNameserversReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNameserversOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetNameserversDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetNameserversOK creates a GetNameserversOK with default headers values
func NewGetNameserversOK() *GetNameserversOK {
	return &GetNameserversOK{}
}

/*GetNameserversOK handles this case with default header values.

Successful operation
*/
type GetNameserversOK struct {
	/*Configuration file version
	 */
	ConfigurationVersion int64

	Payload *GetNameserversOKBody
}

func (o *GetNameserversOK) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/nameservers][%d] getNameserversOK  %+v", 200, o.Payload)
}

func (o *GetNameserversOK) GetPayload() *GetNameserversOKBody {
	return o.Payload
}

func (o *GetNameserversOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	configurationVersion, err := swag.ConvertInt64(response.GetHeader("Configuration-Version"))
	if err != nil {
		return errors.InvalidType("Configuration-Version", "header", "int64", response.GetHeader("Configuration-Version"))
	}
	o.ConfigurationVersion = configurationVersion

	o.Payload = new(GetNameserversOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNameserversDefault creates a GetNameserversDefault with default headers values
func NewGetNameserversDefault(code int) *GetNameserversDefault {
	return &GetNameserversDefault{
		_statusCode:          code,
		ConfigurationVersion: 0,
	}
}

/*GetNameserversDefault handles this case with default header values.

General Error
*/
type GetNameserversDefault struct {
	_statusCode int

	/*Configuration file version
	 */
	ConfigurationVersion int64

	Payload *models.Error
}

// Code gets the status code for the get nameservers default response
func (o *GetNameserversDefault) Code() int {
	return o._statusCode
}

func (o *GetNameserversDefault) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/nameservers][%d] getNameservers default  %+v", o._statusCode, o.Payload)
}

func (o *GetNameserversDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetNameserversDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

/*GetNameserversOKBody get nameservers o k body
swagger:model GetNameserversOKBody
*/
type GetNameserversOKBody struct {

	// version
	Version int64 `json:"_version,omitempty"`

	// data
	// Required: true
	Data models.Nameservers `json:"data"`
}

// Validate validates this get nameservers o k body
func (o *GetNameserversOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetNameserversOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("getNameserversOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	if err := o.Data.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("getNameserversOK" + "." + "data")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetNameserversOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetNameserversOKBody) UnmarshalBinary(b []byte) error {
	var res GetNameserversOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
