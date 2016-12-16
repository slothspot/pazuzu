package features

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteAPIFeaturesIDParams creates a new DeleteAPIFeaturesIDParams object
// with the default values initialized.
func NewDeleteAPIFeaturesIDParams() *DeleteAPIFeaturesIDParams {
	var ()
	return &DeleteAPIFeaturesIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteAPIFeaturesIDParamsWithTimeout creates a new DeleteAPIFeaturesIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteAPIFeaturesIDParamsWithTimeout(timeout time.Duration) *DeleteAPIFeaturesIDParams {
	var ()
	return &DeleteAPIFeaturesIDParams{

		timeout: timeout,
	}
}

// NewDeleteAPIFeaturesIDParamsWithContext creates a new DeleteAPIFeaturesIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteAPIFeaturesIDParamsWithContext(ctx context.Context) *DeleteAPIFeaturesIDParams {
	var ()
	return &DeleteAPIFeaturesIDParams{

		Context: ctx,
	}
}

/*DeleteAPIFeaturesIDParams contains all the parameters to send to the API endpoint
for the delete API features ID operation typically these are written to a http.Request
*/
type DeleteAPIFeaturesIDParams struct {

	/*ID
	  Name of feature

	*/
	ID string

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the delete API features ID params
func (o *DeleteAPIFeaturesIDParams) WithTimeout(timeout time.Duration) *DeleteAPIFeaturesIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete API features ID params
func (o *DeleteAPIFeaturesIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete API features ID params
func (o *DeleteAPIFeaturesIDParams) WithContext(ctx context.Context) *DeleteAPIFeaturesIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete API features ID params
func (o *DeleteAPIFeaturesIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithID adds the id to the delete API features ID params
func (o *DeleteAPIFeaturesIDParams) WithID(id string) *DeleteAPIFeaturesIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete API features ID params
func (o *DeleteAPIFeaturesIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteAPIFeaturesIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
