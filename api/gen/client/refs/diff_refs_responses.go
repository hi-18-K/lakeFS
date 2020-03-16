// Code generated by go-swagger; DO NOT EDIT.

package refs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/treeverse/lakefs/api/gen/models"
)

// DiffRefsReader is a Reader for the DiffRefs structure.
type DiffRefsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DiffRefsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDiffRefsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDiffRefsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDiffRefsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDiffRefsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDiffRefsOK creates a DiffRefsOK with default headers values
func NewDiffRefsOK() *DiffRefsOK {
	return &DiffRefsOK{}
}

/*DiffRefsOK handles this case with default header values.

diff between refs
*/
type DiffRefsOK struct {
	Payload *DiffRefsOKBody
}

func (o *DiffRefsOK) Error() string {
	return fmt.Sprintf("[GET /repositories/{repositoryId}/refs/{leftRef}/diff/{rightRef}][%d] diffRefsOK  %+v", 200, o.Payload)
}

func (o *DiffRefsOK) GetPayload() *DiffRefsOKBody {
	return o.Payload
}

func (o *DiffRefsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(DiffRefsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDiffRefsUnauthorized creates a DiffRefsUnauthorized with default headers values
func NewDiffRefsUnauthorized() *DiffRefsUnauthorized {
	return &DiffRefsUnauthorized{}
}

/*DiffRefsUnauthorized handles this case with default header values.

Unauthorized
*/
type DiffRefsUnauthorized struct {
	Payload interface{}
}

func (o *DiffRefsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /repositories/{repositoryId}/refs/{leftRef}/diff/{rightRef}][%d] diffRefsUnauthorized  %+v", 401, o.Payload)
}

func (o *DiffRefsUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *DiffRefsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDiffRefsNotFound creates a DiffRefsNotFound with default headers values
func NewDiffRefsNotFound() *DiffRefsNotFound {
	return &DiffRefsNotFound{}
}

/*DiffRefsNotFound handles this case with default header values.

branch not found
*/
type DiffRefsNotFound struct {
	Payload *models.Error
}

func (o *DiffRefsNotFound) Error() string {
	return fmt.Sprintf("[GET /repositories/{repositoryId}/refs/{leftRef}/diff/{rightRef}][%d] diffRefsNotFound  %+v", 404, o.Payload)
}

func (o *DiffRefsNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DiffRefsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDiffRefsDefault creates a DiffRefsDefault with default headers values
func NewDiffRefsDefault(code int) *DiffRefsDefault {
	return &DiffRefsDefault{
		_statusCode: code,
	}
}

/*DiffRefsDefault handles this case with default header values.

generic error response
*/
type DiffRefsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the diff refs default response
func (o *DiffRefsDefault) Code() int {
	return o._statusCode
}

func (o *DiffRefsDefault) Error() string {
	return fmt.Sprintf("[GET /repositories/{repositoryId}/refs/{leftRef}/diff/{rightRef}][%d] diffRefs default  %+v", o._statusCode, o.Payload)
}

func (o *DiffRefsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DiffRefsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*DiffRefsOKBody diff refs o k body
swagger:model DiffRefsOKBody
*/
type DiffRefsOKBody struct {

	// results
	Results []*models.Diff `json:"results"`
}

// Validate validates this diff refs o k body
func (o *DiffRefsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateResults(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DiffRefsOKBody) validateResults(formats strfmt.Registry) error {

	if swag.IsZero(o.Results) { // not required
		return nil
	}

	for i := 0; i < len(o.Results); i++ {
		if swag.IsZero(o.Results[i]) { // not required
			continue
		}

		if o.Results[i] != nil {
			if err := o.Results[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("diffRefsOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *DiffRefsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DiffRefsOKBody) UnmarshalBinary(b []byte) error {
	var res DiffRefsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}