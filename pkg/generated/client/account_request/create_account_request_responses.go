// Code generated by go-swagger; DO NOT EDIT.

package account_request

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/go-form3/v3/pkg/generated/models"
)

// CreateAccountRequestReader is a Reader for the CreateAccountRequest structure.
type CreateAccountRequestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateAccountRequestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewCreateAccountRequestCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCreateAccountRequestBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewCreateAccountRequestUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewCreateAccountRequestForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewCreateAccountRequestConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewCreateAccountRequestInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 502:
		result := NewCreateAccountRequestBadGateway()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewCreateAccountRequestServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 504:
		result := NewCreateAccountRequestGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateAccountRequestCreated creates a CreateAccountRequestCreated with default headers values
func NewCreateAccountRequestCreated() *CreateAccountRequestCreated {
	return &CreateAccountRequestCreated{}
}

/*CreateAccountRequestCreated handles this case with default header values.

Account Request creation response
*/
type CreateAccountRequestCreated struct {

	//Payload

	// isStream: false
	*models.AccountRequestResponse
}

func (o *CreateAccountRequestCreated) Error() string {
	return fmt.Sprintf("[POST /organisation/accountrequests][%d] createAccountRequestCreated", 201)
}

func (o *CreateAccountRequestCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.AccountRequestResponse = new(models.AccountRequestResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.AccountRequestResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAccountRequestBadRequest creates a CreateAccountRequestBadRequest with default headers values
func NewCreateAccountRequestBadRequest() *CreateAccountRequestBadRequest {
	return &CreateAccountRequestBadRequest{}
}

/*CreateAccountRequestBadRequest handles this case with default header values.

Account Request creation error
*/
type CreateAccountRequestBadRequest struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *CreateAccountRequestBadRequest) Error() string {
	return fmt.Sprintf("[POST /organisation/accountrequests][%d] createAccountRequestBadRequest", 400)
}

func (o *CreateAccountRequestBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAccountRequestUnauthorized creates a CreateAccountRequestUnauthorized with default headers values
func NewCreateAccountRequestUnauthorized() *CreateAccountRequestUnauthorized {
	return &CreateAccountRequestUnauthorized{}
}

/*CreateAccountRequestUnauthorized handles this case with default header values.

Unauthorized
*/
type CreateAccountRequestUnauthorized struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *CreateAccountRequestUnauthorized) Error() string {
	return fmt.Sprintf("[POST /organisation/accountrequests][%d] createAccountRequestUnauthorized", 401)
}

func (o *CreateAccountRequestUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAccountRequestForbidden creates a CreateAccountRequestForbidden with default headers values
func NewCreateAccountRequestForbidden() *CreateAccountRequestForbidden {
	return &CreateAccountRequestForbidden{}
}

/*CreateAccountRequestForbidden handles this case with default header values.

Forbidden
*/
type CreateAccountRequestForbidden struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *CreateAccountRequestForbidden) Error() string {
	return fmt.Sprintf("[POST /organisation/accountrequests][%d] createAccountRequestForbidden", 403)
}

func (o *CreateAccountRequestForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAccountRequestConflict creates a CreateAccountRequestConflict with default headers values
func NewCreateAccountRequestConflict() *CreateAccountRequestConflict {
	return &CreateAccountRequestConflict{}
}

/*CreateAccountRequestConflict handles this case with default header values.

Conflict
*/
type CreateAccountRequestConflict struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *CreateAccountRequestConflict) Error() string {
	return fmt.Sprintf("[POST /organisation/accountrequests][%d] createAccountRequestConflict", 409)
}

func (o *CreateAccountRequestConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAccountRequestInternalServerError creates a CreateAccountRequestInternalServerError with default headers values
func NewCreateAccountRequestInternalServerError() *CreateAccountRequestInternalServerError {
	return &CreateAccountRequestInternalServerError{}
}

/*CreateAccountRequestInternalServerError handles this case with default header values.

Internal Server Error
*/
type CreateAccountRequestInternalServerError struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *CreateAccountRequestInternalServerError) Error() string {
	return fmt.Sprintf("[POST /organisation/accountrequests][%d] createAccountRequestInternalServerError", 500)
}

func (o *CreateAccountRequestInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAccountRequestBadGateway creates a CreateAccountRequestBadGateway with default headers values
func NewCreateAccountRequestBadGateway() *CreateAccountRequestBadGateway {
	return &CreateAccountRequestBadGateway{}
}

/*CreateAccountRequestBadGateway handles this case with default header values.

Bad Gateway
*/
type CreateAccountRequestBadGateway struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *CreateAccountRequestBadGateway) Error() string {
	return fmt.Sprintf("[POST /organisation/accountrequests][%d] createAccountRequestBadGateway", 502)
}

func (o *CreateAccountRequestBadGateway) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAccountRequestServiceUnavailable creates a CreateAccountRequestServiceUnavailable with default headers values
func NewCreateAccountRequestServiceUnavailable() *CreateAccountRequestServiceUnavailable {
	return &CreateAccountRequestServiceUnavailable{}
}

/*CreateAccountRequestServiceUnavailable handles this case with default header values.

Service unavailable
*/
type CreateAccountRequestServiceUnavailable struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *CreateAccountRequestServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /organisation/accountrequests][%d] createAccountRequestServiceUnavailable", 503)
}

func (o *CreateAccountRequestServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAccountRequestGatewayTimeout creates a CreateAccountRequestGatewayTimeout with default headers values
func NewCreateAccountRequestGatewayTimeout() *CreateAccountRequestGatewayTimeout {
	return &CreateAccountRequestGatewayTimeout{}
}

/*CreateAccountRequestGatewayTimeout handles this case with default header values.

Gateway timeout
*/
type CreateAccountRequestGatewayTimeout struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *CreateAccountRequestGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /organisation/accountrequests][%d] createAccountRequestGatewayTimeout", 504)
}

func (o *CreateAccountRequestGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}
