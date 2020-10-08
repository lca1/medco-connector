// Code generated by go-swagger; DO NOT EDIT.

package medco_node

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/ldsec/medco/connector/restapi/models"
)

// ExploreQueryHandlerFunc turns a function with the right signature into a explore query handler
type ExploreQueryHandlerFunc func(ExploreQueryParams, *models.User) middleware.Responder

// Handle executing the request and returning a response
func (fn ExploreQueryHandlerFunc) Handle(params ExploreQueryParams, principal *models.User) middleware.Responder {
	return fn(params, principal)
}

// ExploreQueryHandler interface for that can handle valid explore query params
type ExploreQueryHandler interface {
	Handle(ExploreQueryParams, *models.User) middleware.Responder
}

// NewExploreQuery creates a new http.Handler for the explore query operation
func NewExploreQuery(ctx *middleware.Context, handler ExploreQueryHandler) *ExploreQuery {
	return &ExploreQuery{Context: ctx, Handler: handler}
}

/*ExploreQuery swagger:route POST /node/explore/query medco-node exploreQuery

MedCo-Explore query to the node.

*/
type ExploreQuery struct {
	Context *middleware.Context
	Handler ExploreQueryHandler
}

func (o *ExploreQuery) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewExploreQueryParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal *models.User
	if uprinc != nil {
		principal = uprinc.(*models.User) // this is really a models.User, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// ExploreQueryBody explore query body
//
// swagger:model ExploreQueryBody
type ExploreQueryBody struct {

	// id
	// Pattern: ^[\w:-]+$
	ID string `json:"id,omitempty"`

	// query
	Query *models.ExploreQuery `json:"query,omitempty"`
}

// Validate validates this explore query body
func (o *ExploreQueryBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateQuery(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ExploreQueryBody) validateID(formats strfmt.Registry) error {

	if swag.IsZero(o.ID) { // not required
		return nil
	}

	if err := validate.Pattern("queryRequest"+"."+"id", "body", string(o.ID), `^[\w:-]+$`); err != nil {
		return err
	}

	return nil
}

func (o *ExploreQueryBody) validateQuery(formats strfmt.Registry) error {

	if swag.IsZero(o.Query) { // not required
		return nil
	}

	if o.Query != nil {
		if err := o.Query.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("queryRequest" + "." + "query")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ExploreQueryBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ExploreQueryBody) UnmarshalBinary(b []byte) error {
	var res ExploreQueryBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// ExploreQueryDefaultBody explore query default body
//
// swagger:model ExploreQueryDefaultBody
type ExploreQueryDefaultBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this explore query default body
func (o *ExploreQueryDefaultBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ExploreQueryDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ExploreQueryDefaultBody) UnmarshalBinary(b []byte) error {
	var res ExploreQueryDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// ExploreQueryOKBody explore query o k body
//
// swagger:model ExploreQueryOKBody
type ExploreQueryOKBody struct {

	// id
	ID string `json:"id,omitempty"`

	// query
	Query *models.ExploreQuery `json:"query,omitempty"`

	// result
	Result *models.ExploreQueryResultElement `json:"result,omitempty"`
}

// Validate validates this explore query o k body
func (o *ExploreQueryOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateQuery(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ExploreQueryOKBody) validateQuery(formats strfmt.Registry) error {

	if swag.IsZero(o.Query) { // not required
		return nil
	}

	if o.Query != nil {
		if err := o.Query.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("exploreQueryOK" + "." + "query")
			}
			return err
		}
	}

	return nil
}

func (o *ExploreQueryOKBody) validateResult(formats strfmt.Registry) error {

	if swag.IsZero(o.Result) { // not required
		return nil
	}

	if o.Result != nil {
		if err := o.Result.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("exploreQueryOK" + "." + "result")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ExploreQueryOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ExploreQueryOKBody) UnmarshalBinary(b []byte) error {
	var res ExploreQueryOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}