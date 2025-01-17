// Code generated by go-swagger; DO NOT EDIT.

package direct_debits

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/go-form3/v4/pkg/generated/models"
)

// PostTransactionDirectdebitsIDRecallsReader is a Reader for the PostTransactionDirectdebitsIDRecalls structure.
type PostTransactionDirectdebitsIDRecallsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostTransactionDirectdebitsIDRecallsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostTransactionDirectdebitsIDRecallsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPostTransactionDirectdebitsIDRecallsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostTransactionDirectdebitsIDRecallsCreated creates a PostTransactionDirectdebitsIDRecallsCreated with default headers values
func NewPostTransactionDirectdebitsIDRecallsCreated() *PostTransactionDirectdebitsIDRecallsCreated {
	return &PostTransactionDirectdebitsIDRecallsCreated{}
}

/*PostTransactionDirectdebitsIDRecallsCreated handles this case with default header values.

Recall creation response
*/
type PostTransactionDirectdebitsIDRecallsCreated struct {

	//Payload

	// isStream: false
	*models.DirectDebitRecallCreationResponse
}

func (o *PostTransactionDirectdebitsIDRecallsCreated) Error() string {
	return fmt.Sprintf("[POST /transaction/directdebits/{id}/recalls][%d] postTransactionDirectdebitsIdRecallsCreated", 201)
}

func (o *PostTransactionDirectdebitsIDRecallsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.DirectDebitRecallCreationResponse = new(models.DirectDebitRecallCreationResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.DirectDebitRecallCreationResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTransactionDirectdebitsIDRecallsBadRequest creates a PostTransactionDirectdebitsIDRecallsBadRequest with default headers values
func NewPostTransactionDirectdebitsIDRecallsBadRequest() *PostTransactionDirectdebitsIDRecallsBadRequest {
	return &PostTransactionDirectdebitsIDRecallsBadRequest{}
}

/*PostTransactionDirectdebitsIDRecallsBadRequest handles this case with default header values.

Recall creation error
*/
type PostTransactionDirectdebitsIDRecallsBadRequest struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *PostTransactionDirectdebitsIDRecallsBadRequest) Error() string {
	return fmt.Sprintf("[POST /transaction/directdebits/{id}/recalls][%d] postTransactionDirectdebitsIdRecallsBadRequest", 400)
}

func (o *PostTransactionDirectdebitsIDRecallsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}
