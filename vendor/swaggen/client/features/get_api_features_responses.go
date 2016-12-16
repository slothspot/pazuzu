package features

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"swaggen/models"
)

// GetAPIFeaturesReader is a Reader for the GetAPIFeatures structure.
type GetAPIFeaturesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPIFeaturesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAPIFeaturesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetAPIFeaturesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewGetAPIFeaturesOK creates a GetAPIFeaturesOK with default headers values
func NewGetAPIFeaturesOK() *GetAPIFeaturesOK {
	return &GetAPIFeaturesOK{}
}

/*GetAPIFeaturesOK handles this case with default header values.

An array of features
*/
type GetAPIFeaturesOK struct {
	Payload []*models.Feature
}

func (o *GetAPIFeaturesOK) Error() string {
	return fmt.Sprintf("[GET /api/features][%d] getApiFeaturesOK  %+v", 200, o.Payload)
}

func (o *GetAPIFeaturesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIFeaturesDefault creates a GetAPIFeaturesDefault with default headers values
func NewGetAPIFeaturesDefault(code int) *GetAPIFeaturesDefault {
	return &GetAPIFeaturesDefault{
		_statusCode: code,
	}
}

/*GetAPIFeaturesDefault handles this case with default header values.

Unexpected error
*/
type GetAPIFeaturesDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get API features default response
func (o *GetAPIFeaturesDefault) Code() int {
	return o._statusCode
}

func (o *GetAPIFeaturesDefault) Error() string {
	return fmt.Sprintf("[GET /api/features][%d] GetAPIFeatures default  %+v", o._statusCode, o.Payload)
}

func (o *GetAPIFeaturesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
