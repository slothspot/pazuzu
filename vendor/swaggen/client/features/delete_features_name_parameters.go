package features

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteFeaturesNameParams creates a new DeleteFeaturesNameParams object
// with the default values initialized.
func NewDeleteFeaturesNameParams() *DeleteFeaturesNameParams {
	var ()
	return &DeleteFeaturesNameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteFeaturesNameParamsWithTimeout creates a new DeleteFeaturesNameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteFeaturesNameParamsWithTimeout(timeout time.Duration) *DeleteFeaturesNameParams {
	var ()
	return &DeleteFeaturesNameParams{

		timeout: timeout,
	}
}

// NewDeleteFeaturesNameParamsWithContext creates a new DeleteFeaturesNameParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteFeaturesNameParamsWithContext(ctx context.Context) *DeleteFeaturesNameParams {
	var ()
	return &DeleteFeaturesNameParams{

		Context: ctx,
	}
}

/*DeleteFeaturesNameParams contains all the parameters to send to the API endpoint
for the delete features name operation typically these are written to a http.Request
*/
type DeleteFeaturesNameParams struct {

	/*Name
	  The feature name.

	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete features name params
func (o *DeleteFeaturesNameParams) WithTimeout(timeout time.Duration) *DeleteFeaturesNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete features name params
func (o *DeleteFeaturesNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete features name params
func (o *DeleteFeaturesNameParams) WithContext(ctx context.Context) *DeleteFeaturesNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete features name params
func (o *DeleteFeaturesNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithName adds the name to the delete features name params
func (o *DeleteFeaturesNameParams) WithName(name string) *DeleteFeaturesNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the delete features name params
func (o *DeleteFeaturesNameParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteFeaturesNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}