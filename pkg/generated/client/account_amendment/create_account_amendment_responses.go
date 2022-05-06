// Code generated by go-swagger; DO NOT EDIT.

package account_amendment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/go-form3/v4/pkg/generated/models"
)

// CreateAccountAmendmentReader is a Reader for the CreateAccountAmendment structure.
type CreateAccountAmendmentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateAccountAmendmentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewCreateAccountAmendmentCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCreateAccountAmendmentBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewCreateAccountAmendmentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewCreateAccountAmendmentForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewCreateAccountAmendmentConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewCreateAccountAmendmentInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 502:
		result := NewCreateAccountAmendmentBadGateway()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewCreateAccountAmendmentServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 504:
		result := NewCreateAccountAmendmentGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateAccountAmendmentCreated creates a CreateAccountAmendmentCreated with default headers values
func NewCreateAccountAmendmentCreated() *CreateAccountAmendmentCreated {
	return &CreateAccountAmendmentCreated{}
}

/*CreateAccountAmendmentCreated handles this case with default header values.

Account Amendment creation response
*/
type CreateAccountAmendmentCreated struct {

	//Payload

	// isStream: false
	*models.AccountAmendmentResponse
}

func (o *CreateAccountAmendmentCreated) Error() string {
	return fmt.Sprintf("[POST /organisation/accountamendments][%d] createAccountAmendmentCreated", 201)
}

func (o *CreateAccountAmendmentCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.AccountAmendmentResponse = new(models.AccountAmendmentResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.AccountAmendmentResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAccountAmendmentBadRequest creates a CreateAccountAmendmentBadRequest with default headers values
func NewCreateAccountAmendmentBadRequest() *CreateAccountAmendmentBadRequest {
	return &CreateAccountAmendmentBadRequest{}
}

/*CreateAccountAmendmentBadRequest handles this case with default header values.

Account Request creation error
*/
type CreateAccountAmendmentBadRequest struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *CreateAccountAmendmentBadRequest) Error() string {
	return fmt.Sprintf("[POST /organisation/accountamendments][%d] createAccountAmendmentBadRequest", 400)
}

func (o *CreateAccountAmendmentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAccountAmendmentUnauthorized creates a CreateAccountAmendmentUnauthorized with default headers values
func NewCreateAccountAmendmentUnauthorized() *CreateAccountAmendmentUnauthorized {
	return &CreateAccountAmendmentUnauthorized{}
}

/*CreateAccountAmendmentUnauthorized handles this case with default header values.

Unauthorized
*/
type CreateAccountAmendmentUnauthorized struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *CreateAccountAmendmentUnauthorized) Error() string {
	return fmt.Sprintf("[POST /organisation/accountamendments][%d] createAccountAmendmentUnauthorized", 401)
}

func (o *CreateAccountAmendmentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAccountAmendmentForbidden creates a CreateAccountAmendmentForbidden with default headers values
func NewCreateAccountAmendmentForbidden() *CreateAccountAmendmentForbidden {
	return &CreateAccountAmendmentForbidden{}
}

/*CreateAccountAmendmentForbidden handles this case with default header values.

Forbidden
*/
type CreateAccountAmendmentForbidden struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *CreateAccountAmendmentForbidden) Error() string {
	return fmt.Sprintf("[POST /organisation/accountamendments][%d] createAccountAmendmentForbidden", 403)
}

func (o *CreateAccountAmendmentForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAccountAmendmentConflict creates a CreateAccountAmendmentConflict with default headers values
func NewCreateAccountAmendmentConflict() *CreateAccountAmendmentConflict {
	return &CreateAccountAmendmentConflict{}
}

/*CreateAccountAmendmentConflict handles this case with default header values.

Conflict
*/
type CreateAccountAmendmentConflict struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *CreateAccountAmendmentConflict) Error() string {
	return fmt.Sprintf("[POST /organisation/accountamendments][%d] createAccountAmendmentConflict", 409)
}

func (o *CreateAccountAmendmentConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAccountAmendmentInternalServerError creates a CreateAccountAmendmentInternalServerError with default headers values
func NewCreateAccountAmendmentInternalServerError() *CreateAccountAmendmentInternalServerError {
	return &CreateAccountAmendmentInternalServerError{}
}

/*CreateAccountAmendmentInternalServerError handles this case with default header values.

Internal Server Error
*/
type CreateAccountAmendmentInternalServerError struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *CreateAccountAmendmentInternalServerError) Error() string {
	return fmt.Sprintf("[POST /organisation/accountamendments][%d] createAccountAmendmentInternalServerError", 500)
}

func (o *CreateAccountAmendmentInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAccountAmendmentBadGateway creates a CreateAccountAmendmentBadGateway with default headers values
func NewCreateAccountAmendmentBadGateway() *CreateAccountAmendmentBadGateway {
	return &CreateAccountAmendmentBadGateway{}
}

/*CreateAccountAmendmentBadGateway handles this case with default header values.

Bad Gateway
*/
type CreateAccountAmendmentBadGateway struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *CreateAccountAmendmentBadGateway) Error() string {
	return fmt.Sprintf("[POST /organisation/accountamendments][%d] createAccountAmendmentBadGateway", 502)
}

func (o *CreateAccountAmendmentBadGateway) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAccountAmendmentServiceUnavailable creates a CreateAccountAmendmentServiceUnavailable with default headers values
func NewCreateAccountAmendmentServiceUnavailable() *CreateAccountAmendmentServiceUnavailable {
	return &CreateAccountAmendmentServiceUnavailable{}
}

/*CreateAccountAmendmentServiceUnavailable handles this case with default header values.

Service unavailable
*/
type CreateAccountAmendmentServiceUnavailable struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *CreateAccountAmendmentServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /organisation/accountamendments][%d] createAccountAmendmentServiceUnavailable", 503)
}

func (o *CreateAccountAmendmentServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAccountAmendmentGatewayTimeout creates a CreateAccountAmendmentGatewayTimeout with default headers values
func NewCreateAccountAmendmentGatewayTimeout() *CreateAccountAmendmentGatewayTimeout {
	return &CreateAccountAmendmentGatewayTimeout{}
}

/*CreateAccountAmendmentGatewayTimeout handles this case with default header values.

Gateway timeout
*/
type CreateAccountAmendmentGatewayTimeout struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *CreateAccountAmendmentGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /organisation/accountamendments][%d] createAccountAmendmentGatewayTimeout", 504)
}

func (o *CreateAccountAmendmentGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}
