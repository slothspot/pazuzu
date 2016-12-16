package features

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetAPIFeaturesParams creates a new GetAPIFeaturesParams object
// with the default values initialized.
func NewGetAPIFeaturesParams() *GetAPIFeaturesParams {
	var ()
	return &GetAPIFeaturesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAPIFeaturesParamsWithTimeout creates a new GetAPIFeaturesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAPIFeaturesParamsWithTimeout(timeout time.Duration) *GetAPIFeaturesParams {
	var ()
	return &GetAPIFeaturesParams{

		timeout: timeout,
	}
}

// NewGetAPIFeaturesParamsWithContext creates a new GetAPIFeaturesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAPIFeaturesParamsWithContext(ctx context.Context) *GetAPIFeaturesParams {
	var ()
	return &GetAPIFeaturesParams{

		Context: ctx,
	}
}

/*GetAPIFeaturesParams contains all the parameters to send to the API endpoint
for the get API features operation typically these are written to a http.Request
*/
type GetAPIFeaturesParams struct {

	/*Name
	  value, that must present in feature name.

	*/
	Name []string
	/*Sorted
	  flag to indicate if output should be sorted

	*/
	Sorted *int64

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the get API features params
func (o *GetAPIFeaturesParams) WithTimeout(timeout time.Duration) *GetAPIFeaturesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get API features params
func (o *GetAPIFeaturesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get API features params
func (o *GetAPIFeaturesParams) WithContext(ctx context.Context) *GetAPIFeaturesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get API features params
func (o *GetAPIFeaturesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithName adds the name to the get API features params
func (o *GetAPIFeaturesParams) WithName(name []string) *GetAPIFeaturesParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get API features params
func (o *GetAPIFeaturesParams) SetName(name []string) {
	o.Name = name
}

// WithSorted adds the sorted to the get API features params
func (o *GetAPIFeaturesParams) WithSorted(sorted *int64) *GetAPIFeaturesParams {
	o.SetSorted(sorted)
	return o
}

// SetSorted adds the sorted to the get API features params
func (o *GetAPIFeaturesParams) SetSorted(sorted *int64) {
	o.Sorted = sorted
}

// WriteToRequest writes these params to a swagger request
func (o *GetAPIFeaturesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	valuesName := o.Name

	joinedName := swag.JoinByFormat(valuesName, "")
	// query array param name
	if err := r.SetQueryParam("name", joinedName...); err != nil {
		return err
	}

	if o.Sorted != nil {

		// query param sorted
		var qrSorted int64
		if o.Sorted != nil {
			qrSorted = *o.Sorted
		}
		qSorted := swag.FormatInt64(qrSorted)
		if qSorted != "" {
			if err := r.SetQueryParam("sorted", qSorted); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
