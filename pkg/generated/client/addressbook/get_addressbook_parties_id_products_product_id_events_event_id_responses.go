// Code generated by go-swagger; DO NOT EDIT.

package addressbook

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/go-form3/v3/pkg/generated/models"
)

// GetAddressbookPartiesIDProductsProductIDEventsEventIDReader is a Reader for the GetAddressbookPartiesIDProductsProductIDEventsEventID structure.
type GetAddressbookPartiesIDProductsProductIDEventsEventIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAddressbookPartiesIDProductsProductIDEventsEventIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAddressbookPartiesIDProductsProductIDEventsEventIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetAddressbookPartiesIDProductsProductIDEventsEventIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetAddressbookPartiesIDProductsProductIDEventsEventIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetAddressbookPartiesIDProductsProductIDEventsEventIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAddressbookPartiesIDProductsProductIDEventsEventIDOK creates a GetAddressbookPartiesIDProductsProductIDEventsEventIDOK with default headers values
func NewGetAddressbookPartiesIDProductsProductIDEventsEventIDOK() *GetAddressbookPartiesIDProductsProductIDEventsEventIDOK {
	return &GetAddressbookPartiesIDProductsProductIDEventsEventIDOK{}
}

/*GetAddressbookPartiesIDProductsProductIDEventsEventIDOK handles this case with default header values.

get contact account response
*/
type GetAddressbookPartiesIDProductsProductIDEventsEventIDOK struct {

	//Payload

	// isStream: false
	*models.PartyProductEventGetResponse
}

func (o *GetAddressbookPartiesIDProductsProductIDEventsEventIDOK) Error() string {
	return fmt.Sprintf("[GET /addressbook/parties/{id}/products/{product_id}/events/{event_id}][%d] getAddressbookPartiesIdProductsProductIdEventsEventIdOK", 200)
}

func (o *GetAddressbookPartiesIDProductsProductIDEventsEventIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.PartyProductEventGetResponse = new(models.PartyProductEventGetResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.PartyProductEventGetResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAddressbookPartiesIDProductsProductIDEventsEventIDBadRequest creates a GetAddressbookPartiesIDProductsProductIDEventsEventIDBadRequest with default headers values
func NewGetAddressbookPartiesIDProductsProductIDEventsEventIDBadRequest() *GetAddressbookPartiesIDProductsProductIDEventsEventIDBadRequest {
	return &GetAddressbookPartiesIDProductsProductIDEventsEventIDBadRequest{}
}

/*GetAddressbookPartiesIDProductsProductIDEventsEventIDBadRequest handles this case with default header values.

Bad Request
*/
type GetAddressbookPartiesIDProductsProductIDEventsEventIDBadRequest struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *GetAddressbookPartiesIDProductsProductIDEventsEventIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /addressbook/parties/{id}/products/{product_id}/events/{event_id}][%d] getAddressbookPartiesIdProductsProductIdEventsEventIdBadRequest", 400)
}

func (o *GetAddressbookPartiesIDProductsProductIDEventsEventIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAddressbookPartiesIDProductsProductIDEventsEventIDForbidden creates a GetAddressbookPartiesIDProductsProductIDEventsEventIDForbidden with default headers values
func NewGetAddressbookPartiesIDProductsProductIDEventsEventIDForbidden() *GetAddressbookPartiesIDProductsProductIDEventsEventIDForbidden {
	return &GetAddressbookPartiesIDProductsProductIDEventsEventIDForbidden{}
}

/*GetAddressbookPartiesIDProductsProductIDEventsEventIDForbidden handles this case with default header values.

Forbidden
*/
type GetAddressbookPartiesIDProductsProductIDEventsEventIDForbidden struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *GetAddressbookPartiesIDProductsProductIDEventsEventIDForbidden) Error() string {
	return fmt.Sprintf("[GET /addressbook/parties/{id}/products/{product_id}/events/{event_id}][%d] getAddressbookPartiesIdProductsProductIdEventsEventIdForbidden", 403)
}

func (o *GetAddressbookPartiesIDProductsProductIDEventsEventIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAddressbookPartiesIDProductsProductIDEventsEventIDNotFound creates a GetAddressbookPartiesIDProductsProductIDEventsEventIDNotFound with default headers values
func NewGetAddressbookPartiesIDProductsProductIDEventsEventIDNotFound() *GetAddressbookPartiesIDProductsProductIDEventsEventIDNotFound {
	return &GetAddressbookPartiesIDProductsProductIDEventsEventIDNotFound{}
}

/*GetAddressbookPartiesIDProductsProductIDEventsEventIDNotFound handles this case with default header values.

Bad Request
*/
type GetAddressbookPartiesIDProductsProductIDEventsEventIDNotFound struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *GetAddressbookPartiesIDProductsProductIDEventsEventIDNotFound) Error() string {
	return fmt.Sprintf("[GET /addressbook/parties/{id}/products/{product_id}/events/{event_id}][%d] getAddressbookPartiesIdProductsProductIdEventsEventIdNotFound", 404)
}

func (o *GetAddressbookPartiesIDProductsProductIDEventsEventIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}
