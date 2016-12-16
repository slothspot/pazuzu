package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// File file
// swagger:model File
type File struct {

	// href to resource with file content.
	ContentHref string `json:"content_href,omitempty"`

	// Unique identifier of a file.
	ID string `json:"id,omitempty"`

	// Filename as it appears in search results and Dockerfile fragments.
	Name string `json:"name,omitempty"`
}

// Validate validates this file
func (m *File) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
