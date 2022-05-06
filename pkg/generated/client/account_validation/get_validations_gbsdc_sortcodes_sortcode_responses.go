// Code generated by go-swagger; DO NOT EDIT.

package account_validation

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/go-form3/v4/pkg/generated/models"
)

// GetValidationsGbsdcSortcodesSortcodeReader is a Reader for the GetValidationsGbsdcSortcodesSortcode structure.
type GetValidationsGbsdcSortcodesSortcodeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetValidationsGbsdcSortcodesSortcodeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetValidationsGbsdcSortcodesSortcodeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetValidationsGbsdcSortcodesSortcodeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetValidationsGbsdcSortcodesSortcodeOK creates a GetValidationsGbsdcSortcodesSortcodeOK with default headers values
func NewGetValidationsGbsdcSortcodesSortcodeOK() *GetValidationsGbsdcSortcodesSortcodeOK {
	return &GetValidationsGbsdcSortcodesSortcodeOK{}
}

/*GetValidationsGbsdcSortcodesSortcodeOK handles this case with default header values.

Sort code details
*/
type GetValidationsGbsdcSortcodesSortcodeOK struct {

	//Payload

	// isStream: false
	*models.SortCodeDetailsResponse
}

func (o *GetValidationsGbsdcSortcodesSortcodeOK) Error() string {
	return fmt.Sprintf("[GET /validations/gbsdc/sortcodes/{sortcode}][%d] getValidationsGbsdcSortcodesSortcodeOK", 200)
}

func (o *GetValidationsGbsdcSortcodesSortcodeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.SortCodeDetailsResponse = new(models.SortCodeDetailsResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.SortCodeDetailsResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetValidationsGbsdcSortcodesSortcodeBadRequest creates a GetValidationsGbsdcSortcodesSortcodeBadRequest with default headers values
func NewGetValidationsGbsdcSortcodesSortcodeBadRequest() *GetValidationsGbsdcSortcodesSortcodeBadRequest {
	return &GetValidationsGbsdcSortcodesSortcodeBadRequest{}
}

/*GetValidationsGbsdcSortcodesSortcodeBadRequest handles this case with default header values.

Validation failed
*/
type GetValidationsGbsdcSortcodesSortcodeBadRequest struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *GetValidationsGbsdcSortcodesSortcodeBadRequest) Error() string {
	return fmt.Sprintf("[GET /validations/gbsdc/sortcodes/{sortcode}][%d] getValidationsGbsdcSortcodesSortcodeBadRequest", 400)
}

func (o *GetValidationsGbsdcSortcodesSortcodeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}
