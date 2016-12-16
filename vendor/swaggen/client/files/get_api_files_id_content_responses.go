package files

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"swaggen/models"
)

// GetAPIFilesIDContentReader is a Reader for the GetAPIFilesIDContent structure.
type GetAPIFilesIDContentReader struct {
	formats strfmt.Registry
	writer  io.Writer
}

// ReadResponse reads a server response into the received o.
func (o *GetAPIFilesIDContentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAPIFilesIDContentOK(o.writer)
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetAPIFilesIDContentDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewGetAPIFilesIDContentOK creates a GetAPIFilesIDContentOK with default headers values
func NewGetAPIFilesIDContentOK(writer io.Writer) *GetAPIFilesIDContentOK {
	return &GetAPIFilesIDContentOK{
		Payload: writer,
	}
}

/*GetAPIFilesIDContentOK handles this case with default header values.

OK
*/
type GetAPIFilesIDContentOK struct {
	Payload io.Writer
}

func (o *GetAPIFilesIDContentOK) Error() string {
	return fmt.Sprintf("[GET /api/files/{id}/content][%d] getApiFilesIdContentOK  %+v", 200, o.Payload)
}

func (o *GetAPIFilesIDContentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIFilesIDContentDefault creates a GetAPIFilesIDContentDefault with default headers values
func NewGetAPIFilesIDContentDefault(code int) *GetAPIFilesIDContentDefault {
	return &GetAPIFilesIDContentDefault{
		_statusCode: code,
	}
}

/*GetAPIFilesIDContentDefault handles this case with default header values.

Request can not be fulfilled.
*/
type GetAPIFilesIDContentDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get API files ID content default response
func (o *GetAPIFilesIDContentDefault) Code() int {
	return o._statusCode
}

func (o *GetAPIFilesIDContentDefault) Error() string {
	return fmt.Sprintf("[GET /api/files/{id}/content][%d] GetAPIFilesIDContent default  %+v", o._statusCode, o.Payload)
}

func (o *GetAPIFilesIDContentDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
