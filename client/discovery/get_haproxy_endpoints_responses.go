// Code generated by go-swagger; DO NOT EDIT.

package discovery

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

// GetHaproxyEndpointsReader is a Reader for the GetHaproxyEndpoints structure.
type GetHaproxyEndpointsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetHaproxyEndpointsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetHaproxyEndpointsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetHaproxyEndpointsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetHaproxyEndpointsOK creates a GetHaproxyEndpointsOK with default headers values
func NewGetHaproxyEndpointsOK() *GetHaproxyEndpointsOK {
	return &GetHaproxyEndpointsOK{}
}

/*GetHaproxyEndpointsOK handles this case with default header values.

Success
*/
type GetHaproxyEndpointsOK struct {
	Payload models.Endpoints
}

func (o *GetHaproxyEndpointsOK) Error() string {
	return fmt.Sprintf("[GET /services/haproxy][%d] getHaproxyEndpointsOK  %+v", 200, o.Payload)
}

func (o *GetHaproxyEndpointsOK) GetPayload() models.Endpoints {
	return o.Payload
}

func (o *GetHaproxyEndpointsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetHaproxyEndpointsDefault creates a GetHaproxyEndpointsDefault with default headers values
func NewGetHaproxyEndpointsDefault(code int) *GetHaproxyEndpointsDefault {
	return &GetHaproxyEndpointsDefault{
		_statusCode:          code,
		ConfigurationVersion: 0,
	}
}

/*GetHaproxyEndpointsDefault handles this case with default header values.

General Error
*/
type GetHaproxyEndpointsDefault struct {
	_statusCode int

	/*Configuration file version
	 */
	ConfigurationVersion int64

	Payload *models.Error
}

// Code gets the status code for the get haproxy endpoints default response
func (o *GetHaproxyEndpointsDefault) Code() int {
	return o._statusCode
}

func (o *GetHaproxyEndpointsDefault) Error() string {
	return fmt.Sprintf("[GET /services/haproxy][%d] getHaproxyEndpoints default  %+v", o._statusCode, o.Payload)
}

func (o *GetHaproxyEndpointsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetHaproxyEndpointsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
