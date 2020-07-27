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

// DeleteACLReader is a Reader for the DeleteACL structure.
type DeleteACLReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteACLReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewDeleteACLAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewDeleteACLNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteACLNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteACLDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteACLAccepted creates a DeleteACLAccepted with default headers values
func NewDeleteACLAccepted() *DeleteACLAccepted {
	return &DeleteACLAccepted{}
}

/*DeleteACLAccepted handles this case with default header values.

Configuration change accepted and reload requested
*/
type DeleteACLAccepted struct {
	/*ID of the requested reload
	 */
	ReloadID string
}

func (o *DeleteACLAccepted) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/acls/{index}][%d] deleteAclAccepted ", 202)
}

func (o *DeleteACLAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Reload-ID
	o.ReloadID = response.GetHeader("Reload-ID")

	return nil
}

// NewDeleteACLNoContent creates a DeleteACLNoContent with default headers values
func NewDeleteACLNoContent() *DeleteACLNoContent {
	return &DeleteACLNoContent{}
}

/*DeleteACLNoContent handles this case with default header values.

ACL line deleted
*/
type DeleteACLNoContent struct {
}

func (o *DeleteACLNoContent) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/acls/{index}][%d] deleteAclNoContent ", 204)
}

func (o *DeleteACLNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteACLNotFound creates a DeleteACLNotFound with default headers values
func NewDeleteACLNotFound() *DeleteACLNotFound {
	return &DeleteACLNotFound{
		ConfigurationVersion: 0,
	}
}

/*DeleteACLNotFound handles this case with default header values.

The specified resource was not found
*/
type DeleteACLNotFound struct {
	/*Configuration file version
	 */
	ConfigurationVersion int64

	Payload *models.Error
}

func (o *DeleteACLNotFound) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/acls/{index}][%d] deleteAclNotFound  %+v", 404, o.Payload)
}

func (o *DeleteACLNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteACLNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteACLDefault creates a DeleteACLDefault with default headers values
func NewDeleteACLDefault(code int) *DeleteACLDefault {
	return &DeleteACLDefault{
		_statusCode:          code,
		ConfigurationVersion: 0,
	}
}

/*DeleteACLDefault handles this case with default header values.

General Error
*/
type DeleteACLDefault struct {
	_statusCode int

	/*Configuration file version
	 */
	ConfigurationVersion int64

	Payload *models.Error
}

// Code gets the status code for the delete Acl default response
func (o *DeleteACLDefault) Code() int {
	return o._statusCode
}

func (o *DeleteACLDefault) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/acls/{index}][%d] deleteAcl default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteACLDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteACLDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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