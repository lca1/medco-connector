// Code generated by go-swagger; DO NOT EDIT.

package genomic_annotations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetVariantsParams creates a new GetVariantsParams object
// with the default values initialized.
func NewGetVariantsParams() GetVariantsParams {

	var (
		// initialize parameters with default values

		encryptedDefault = bool(true)
	)

	return GetVariantsParams{
		Encrypted: &encryptedDefault,
	}
}

// GetVariantsParams contains all the bound params for the get variants operation
// typically these are obtained from a http.Request
//
// swagger:parameters getVariants
type GetVariantsParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Genomic annotation name.
	  Required: true
	  Pattern: ^\w+$
	  In: path
	*/
	Annotation string
	/*Request pre-encrypted variant identifiers (defaults to true).
	  In: query
	  Default: true
	*/
	Encrypted *bool
	/*Genomic annotation value.
	  Required: true
	  Min Length: 1
	  In: path
	*/
	Value string
	/*Genomic annotation zygosity, if null defaults to all.
	  In: query
	*/
	Zygosity []string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetVariantsParams() beforehand.
func (o *GetVariantsParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	rAnnotation, rhkAnnotation, _ := route.Params.GetOK("annotation")
	if err := o.bindAnnotation(rAnnotation, rhkAnnotation, route.Formats); err != nil {
		res = append(res, err)
	}

	qEncrypted, qhkEncrypted, _ := qs.GetOK("encrypted")
	if err := o.bindEncrypted(qEncrypted, qhkEncrypted, route.Formats); err != nil {
		res = append(res, err)
	}

	rValue, rhkValue, _ := route.Params.GetOK("value")
	if err := o.bindValue(rValue, rhkValue, route.Formats); err != nil {
		res = append(res, err)
	}

	qZygosity, qhkZygosity, _ := qs.GetOK("zygosity")
	if err := o.bindZygosity(qZygosity, qhkZygosity, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindAnnotation binds and validates parameter Annotation from path.
func (o *GetVariantsParams) bindAnnotation(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.Annotation = raw

	if err := o.validateAnnotation(formats); err != nil {
		return err
	}

	return nil
}

// validateAnnotation carries on validations for parameter Annotation
func (o *GetVariantsParams) validateAnnotation(formats strfmt.Registry) error {

	if err := validate.Pattern("annotation", "path", o.Annotation, `^\w+$`); err != nil {
		return err
	}

	return nil
}

// bindEncrypted binds and validates parameter Encrypted from query.
func (o *GetVariantsParams) bindEncrypted(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewGetVariantsParams()
		return nil
	}

	value, err := swag.ConvertBool(raw)
	if err != nil {
		return errors.InvalidType("encrypted", "query", "bool", raw)
	}
	o.Encrypted = &value

	return nil
}

// bindValue binds and validates parameter Value from path.
func (o *GetVariantsParams) bindValue(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.Value = raw

	if err := o.validateValue(formats); err != nil {
		return err
	}

	return nil
}

// validateValue carries on validations for parameter Value
func (o *GetVariantsParams) validateValue(formats strfmt.Registry) error {

	if err := validate.MinLength("value", "path", o.Value, 1); err != nil {
		return err
	}

	return nil
}

// bindZygosity binds and validates array parameter Zygosity from query.
//
// Arrays are parsed according to CollectionFormat: "" (defaults to "csv" when empty).
func (o *GetVariantsParams) bindZygosity(rawData []string, hasKey bool, formats strfmt.Registry) error {

	var qvZygosity string
	if len(rawData) > 0 {
		qvZygosity = rawData[len(rawData)-1]
	}

	// CollectionFormat:
	zygosityIC := swag.SplitByFormat(qvZygosity, "")
	if len(zygosityIC) == 0 {
		return nil
	}

	var zygosityIR []string
	for i, zygosityIV := range zygosityIC {
		zygosityI := zygosityIV

		if err := validate.Enum(fmt.Sprintf("%s.%v", "zygosity", i), "query", zygosityI, []interface{}{"heterozygous", "homozygous", "unknown"}); err != nil {
			return err
		}

		zygosityIR = append(zygosityIR, zygosityI)
	}

	o.Zygosity = zygosityIR

	return nil
}
