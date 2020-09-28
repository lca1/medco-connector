// Code generated by go-swagger; DO NOT EDIT.

package medco_node

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/ldsec/medco-connector/restapi/models"
)

// ExploreSearchModifierReader is a Reader for the ExploreSearchModifier structure.
type ExploreSearchModifierReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExploreSearchModifierReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExploreSearchModifierOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewExploreSearchModifierDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExploreSearchModifierOK creates a ExploreSearchModifierOK with default headers values
func NewExploreSearchModifierOK() *ExploreSearchModifierOK {
	return &ExploreSearchModifierOK{}
}

/*ExploreSearchModifierOK handles this case with default header values.

MedCo-Explore search modifier query response.
*/
type ExploreSearchModifierOK struct {
	Payload *ExploreSearchModifierOKBody
}

func (o *ExploreSearchModifierOK) Error() string {
	return fmt.Sprintf("[POST /node/explore/search/modifier][%d] exploreSearchModifierOK  %+v", 200, o.Payload)
}

func (o *ExploreSearchModifierOK) GetPayload() *ExploreSearchModifierOKBody {
	return o.Payload
}

func (o *ExploreSearchModifierOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ExploreSearchModifierOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExploreSearchModifierDefault creates a ExploreSearchModifierDefault with default headers values
func NewExploreSearchModifierDefault(code int) *ExploreSearchModifierDefault {
	return &ExploreSearchModifierDefault{
		_statusCode: code,
	}
}

/*ExploreSearchModifierDefault handles this case with default header values.

Error response.
*/
type ExploreSearchModifierDefault struct {
	_statusCode int

	Payload *ExploreSearchModifierDefaultBody
}

// Code gets the status code for the explore search modifier default response
func (o *ExploreSearchModifierDefault) Code() int {
	return o._statusCode
}

func (o *ExploreSearchModifierDefault) Error() string {
	return fmt.Sprintf("[POST /node/explore/search/modifier][%d] exploreSearchModifier default  %+v", o._statusCode, o.Payload)
}

func (o *ExploreSearchModifierDefault) GetPayload() *ExploreSearchModifierDefaultBody {
	return o.Payload
}

func (o *ExploreSearchModifierDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ExploreSearchModifierDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ExploreSearchModifierDefaultBody explore search modifier default body
swagger:model ExploreSearchModifierDefaultBody
*/
type ExploreSearchModifierDefaultBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this explore search modifier default body
func (o *ExploreSearchModifierDefaultBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ExploreSearchModifierDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ExploreSearchModifierDefaultBody) UnmarshalBinary(b []byte) error {
	var res ExploreSearchModifierDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ExploreSearchModifierOKBody explore search modifier o k body
swagger:model ExploreSearchModifierOKBody
*/
type ExploreSearchModifierOKBody struct {

	// results
	Results []*models.ExploreSearchResultElement `json:"results"`

	// search
	Search *models.ExploreSearchModifier `json:"search,omitempty"`
}

// Validate validates this explore search modifier o k body
func (o *ExploreSearchModifierOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateResults(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSearch(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ExploreSearchModifierOKBody) validateResults(formats strfmt.Registry) error {

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
					return ve.ValidateName("exploreSearchModifierOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *ExploreSearchModifierOKBody) validateSearch(formats strfmt.Registry) error {

	if swag.IsZero(o.Search) { // not required
		return nil
	}

	if o.Search != nil {
		if err := o.Search.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("exploreSearchModifierOK" + "." + "search")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ExploreSearchModifierOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ExploreSearchModifierOKBody) UnmarshalBinary(b []byte) error {
	var res ExploreSearchModifierOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
