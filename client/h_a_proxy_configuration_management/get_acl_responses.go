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

// GetACLReader is a Reader for the GetACL structure.
type GetACLReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetACLReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetACLOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetACLNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetACLDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetACLOK creates a GetACLOK with default headers values
func NewGetACLOK() *GetACLOK {
	return &GetACLOK{}
}

/*GetACLOK handles this case with default header values.

Successful operation
*/
type GetACLOK struct {
	/*Configuration file version
	 */
	ConfigurationVersion int64

	Payload *GetACLOKBody
}

func (o *GetACLOK) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/acls/{index}][%d] getAclOK  %+v", 200, o.Payload)
}

func (o *GetACLOK) GetPayload() *GetACLOKBody {
	return o.Payload
}

func (o *GetACLOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	configurationVersion, err := swag.ConvertInt64(response.GetHeader("Configuration-Version"))
	if err != nil {
		return errors.InvalidType("Configuration-Version", "header", "int64", response.GetHeader("Configuration-Version"))
	}
	o.ConfigurationVersion = configurationVersion

	o.Payload = new(GetACLOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetACLNotFound creates a GetACLNotFound with default headers values
func NewGetACLNotFound() *GetACLNotFound {
	return &GetACLNotFound{
		ConfigurationVersion: 0,
	}
}

/*GetACLNotFound handles this case with default header values.

The specified resource was not found
*/
type GetACLNotFound struct {
	/*Configuration file version
	 */
	ConfigurationVersion int64

	Payload *models.Error
}

func (o *GetACLNotFound) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/acls/{index}][%d] getAclNotFound  %+v", 404, o.Payload)
}

func (o *GetACLNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetACLNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetACLDefault creates a GetACLDefault with default headers values
func NewGetACLDefault(code int) *GetACLDefault {
	return &GetACLDefault{
		_statusCode:          code,
		ConfigurationVersion: 0,
	}
}

/*GetACLDefault handles this case with default header values.

General Error
*/
type GetACLDefault struct {
	_statusCode int

	/*Configuration file version
	 */
	ConfigurationVersion int64

	Payload *models.Error
}

// Code gets the status code for the get Acl default response
func (o *GetACLDefault) Code() int {
	return o._statusCode
}

func (o *GetACLDefault) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/acls/{index}][%d] getAcl default  %+v", o._statusCode, o.Payload)
}

func (o *GetACLDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetACLDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

/*GetACLOKBody get ACL o k body
swagger:model GetACLOKBody
*/
type GetACLOKBody struct {

	// version
	Version int64 `json:"_version,omitempty"`

	// data
	Data *models.ACL `json:"data,omitempty"`
}

// Validate validates this get ACL o k body
func (o *GetACLOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetACLOKBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getAclOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetACLOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetACLOKBody) UnmarshalBinary(b []byte) error {
	var res GetACLOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
