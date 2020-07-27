// Code generated by go-swagger; DO NOT EDIT.

package stick_table

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

// GetStickTablesReader is a Reader for the GetStickTables structure.
type GetStickTablesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetStickTablesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetStickTablesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetStickTablesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetStickTablesOK creates a GetStickTablesOK with default headers values
func NewGetStickTablesOK() *GetStickTablesOK {
	return &GetStickTablesOK{}
}

/*GetStickTablesOK handles this case with default header values.

Successful operation
*/
type GetStickTablesOK struct {
	Payload models.StickTables
}

func (o *GetStickTablesOK) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/runtime/stick_tables][%d] getStickTablesOK  %+v", 200, o.Payload)
}

func (o *GetStickTablesOK) GetPayload() models.StickTables {
	return o.Payload
}

func (o *GetStickTablesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStickTablesDefault creates a GetStickTablesDefault with default headers values
func NewGetStickTablesDefault(code int) *GetStickTablesDefault {
	return &GetStickTablesDefault{
		_statusCode:          code,
		ConfigurationVersion: 0,
	}
}

/*GetStickTablesDefault handles this case with default header values.

General Error
*/
type GetStickTablesDefault struct {
	_statusCode int

	/*Configuration file version
	 */
	ConfigurationVersion int64

	Payload *models.Error
}

// Code gets the status code for the get stick tables default response
func (o *GetStickTablesDefault) Code() int {
	return o._statusCode
}

func (o *GetStickTablesDefault) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/runtime/stick_tables][%d] getStickTables default  %+v", o._statusCode, o.Payload)
}

func (o *GetStickTablesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetStickTablesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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