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

	"github.com/ldsec/medco/connector/restapi/models"
)

// GetExploreQueryHandlerFunc turns a function with the right signature into a get explore query handler
type GetExploreQueryHandlerFunc func(GetExploreQueryParams, *models.User) middleware.Responder

// Handle executing the request and returning a response
func (fn GetExploreQueryHandlerFunc) Handle(params GetExploreQueryParams, principal *models.User) middleware.Responder {
	return fn(params, principal)
}

// GetExploreQueryHandler interface for that can handle valid get explore query params
type GetExploreQueryHandler interface {
	Handle(GetExploreQueryParams, *models.User) middleware.Responder
}

// NewGetExploreQuery creates a new http.Handler for the get explore query operation
func NewGetExploreQuery(ctx *middleware.Context, handler GetExploreQueryHandler) *GetExploreQuery {
	return &GetExploreQuery{Context: ctx, Handler: handler}
}

/*GetExploreQuery swagger:route GET /node/explore/query/{queryId} medco-node getExploreQuery

Get status and result of a MedCo-Explore query.

*/
type GetExploreQuery struct {
	Context *middleware.Context
	Handler GetExploreQueryHandler
}

func (o *GetExploreQuery) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetExploreQueryParams()

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

// GetExploreQueryDefaultBody get explore query default body
//
// swagger:model GetExploreQueryDefaultBody
type GetExploreQueryDefaultBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this get explore query default body
func (o *GetExploreQueryDefaultBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetExploreQueryDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetExploreQueryDefaultBody) UnmarshalBinary(b []byte) error {
	var res GetExploreQueryDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// GetExploreQueryOKBody get explore query o k body
//
// swagger:model GetExploreQueryOKBody
type GetExploreQueryOKBody struct {

	// id
	ID string `json:"id,omitempty"`

	// query
	Query *models.ExploreQuery `json:"query,omitempty"`

	// result
	Result *models.ExploreQueryResultElement `json:"result,omitempty"`
}

// Validate validates this get explore query o k body
func (o *GetExploreQueryOKBody) Validate(formats strfmt.Registry) error {
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

func (o *GetExploreQueryOKBody) validateQuery(formats strfmt.Registry) error {

	if swag.IsZero(o.Query) { // not required
		return nil
	}

	if o.Query != nil {
		if err := o.Query.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getExploreQueryOK" + "." + "query")
			}
			return err
		}
	}

	return nil
}

func (o *GetExploreQueryOKBody) validateResult(formats strfmt.Registry) error {

	if swag.IsZero(o.Result) { // not required
		return nil
	}

	if o.Result != nil {
		if err := o.Result.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getExploreQueryOK" + "." + "result")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetExploreQueryOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetExploreQueryOKBody) UnmarshalBinary(b []byte) error {
	var res GetExploreQueryOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
