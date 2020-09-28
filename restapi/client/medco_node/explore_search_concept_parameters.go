// Code generated by go-swagger; DO NOT EDIT.

package medco_node

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/ldsec/medco-connector/restapi/models"
)

// NewExploreSearchConceptParams creates a new ExploreSearchConceptParams object
// with the default values initialized.
func NewExploreSearchConceptParams() *ExploreSearchConceptParams {
	var ()
	return &ExploreSearchConceptParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewExploreSearchConceptParamsWithTimeout creates a new ExploreSearchConceptParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewExploreSearchConceptParamsWithTimeout(timeout time.Duration) *ExploreSearchConceptParams {
	var ()
	return &ExploreSearchConceptParams{

		timeout: timeout,
	}
}

// NewExploreSearchConceptParamsWithContext creates a new ExploreSearchConceptParams object
// with the default values initialized, and the ability to set a context for a request
func NewExploreSearchConceptParamsWithContext(ctx context.Context) *ExploreSearchConceptParams {
	var ()
	return &ExploreSearchConceptParams{

		Context: ctx,
	}
}

// NewExploreSearchConceptParamsWithHTTPClient creates a new ExploreSearchConceptParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewExploreSearchConceptParamsWithHTTPClient(client *http.Client) *ExploreSearchConceptParams {
	var ()
	return &ExploreSearchConceptParams{
		HTTPClient: client,
	}
}

/*ExploreSearchConceptParams contains all the parameters to send to the API endpoint
for the explore search concept operation typically these are written to a http.Request
*/
type ExploreSearchConceptParams struct {

	/*SearchConceptRequest
	  MedCo-Explore ontology search concept request.

	*/
	SearchConceptRequest *models.ExploreSearchConcept

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the explore search concept params
func (o *ExploreSearchConceptParams) WithTimeout(timeout time.Duration) *ExploreSearchConceptParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the explore search concept params
func (o *ExploreSearchConceptParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the explore search concept params
func (o *ExploreSearchConceptParams) WithContext(ctx context.Context) *ExploreSearchConceptParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the explore search concept params
func (o *ExploreSearchConceptParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the explore search concept params
func (o *ExploreSearchConceptParams) WithHTTPClient(client *http.Client) *ExploreSearchConceptParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the explore search concept params
func (o *ExploreSearchConceptParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSearchConceptRequest adds the searchConceptRequest to the explore search concept params
func (o *ExploreSearchConceptParams) WithSearchConceptRequest(searchConceptRequest *models.ExploreSearchConcept) *ExploreSearchConceptParams {
	o.SetSearchConceptRequest(searchConceptRequest)
	return o
}

// SetSearchConceptRequest adds the searchConceptRequest to the explore search concept params
func (o *ExploreSearchConceptParams) SetSearchConceptRequest(searchConceptRequest *models.ExploreSearchConcept) {
	o.SearchConceptRequest = searchConceptRequest
}

// WriteToRequest writes these params to a swagger request
func (o *ExploreSearchConceptParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.SearchConceptRequest != nil {
		if err := r.SetBodyParam(o.SearchConceptRequest); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
