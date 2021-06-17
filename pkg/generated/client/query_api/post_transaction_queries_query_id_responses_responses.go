// Code generated by go-swagger; DO NOT EDIT.

package query_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/go-form3/v3/pkg/generated/models"
)

// PostTransactionQueriesQueryIDResponsesReader is a Reader for the PostTransactionQueriesQueryIDResponses structure.
type PostTransactionQueriesQueryIDResponsesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostTransactionQueriesQueryIDResponsesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostTransactionQueriesQueryIDResponsesCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPostTransactionQueriesQueryIDResponsesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewPostTransactionQueriesQueryIDResponsesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 502:
		result := NewPostTransactionQueriesQueryIDResponsesBadGateway()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostTransactionQueriesQueryIDResponsesCreated creates a PostTransactionQueriesQueryIDResponsesCreated with default headers values
func NewPostTransactionQueriesQueryIDResponsesCreated() *PostTransactionQueriesQueryIDResponsesCreated {
	return &PostTransactionQueriesQueryIDResponsesCreated{}
}

/*PostTransactionQueriesQueryIDResponsesCreated handles this case with default header values.

creation response
*/
type PostTransactionQueriesQueryIDResponsesCreated struct {

	//Payload

	// isStream: false
	*models.QueryResponseResponse
}

func (o *PostTransactionQueriesQueryIDResponsesCreated) Error() string {
	return fmt.Sprintf("[POST /transaction/queries/{query_id}/responses][%d] postTransactionQueriesQueryIdResponsesCreated", 201)
}

func (o *PostTransactionQueriesQueryIDResponsesCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.QueryResponseResponse = new(models.QueryResponseResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.QueryResponseResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTransactionQueriesQueryIDResponsesBadRequest creates a PostTransactionQueriesQueryIDResponsesBadRequest with default headers values
func NewPostTransactionQueriesQueryIDResponsesBadRequest() *PostTransactionQueriesQueryIDResponsesBadRequest {
	return &PostTransactionQueriesQueryIDResponsesBadRequest{}
}

/*PostTransactionQueriesQueryIDResponsesBadRequest handles this case with default header values.

bad request
*/
type PostTransactionQueriesQueryIDResponsesBadRequest struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *PostTransactionQueriesQueryIDResponsesBadRequest) Error() string {
	return fmt.Sprintf("[POST /transaction/queries/{query_id}/responses][%d] postTransactionQueriesQueryIdResponsesBadRequest", 400)
}

func (o *PostTransactionQueriesQueryIDResponsesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTransactionQueriesQueryIDResponsesForbidden creates a PostTransactionQueriesQueryIDResponsesForbidden with default headers values
func NewPostTransactionQueriesQueryIDResponsesForbidden() *PostTransactionQueriesQueryIDResponsesForbidden {
	return &PostTransactionQueriesQueryIDResponsesForbidden{}
}

/*PostTransactionQueriesQueryIDResponsesForbidden handles this case with default header values.

forbidden
*/
type PostTransactionQueriesQueryIDResponsesForbidden struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *PostTransactionQueriesQueryIDResponsesForbidden) Error() string {
	return fmt.Sprintf("[POST /transaction/queries/{query_id}/responses][%d] postTransactionQueriesQueryIdResponsesForbidden", 403)
}

func (o *PostTransactionQueriesQueryIDResponsesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTransactionQueriesQueryIDResponsesBadGateway creates a PostTransactionQueriesQueryIDResponsesBadGateway with default headers values
func NewPostTransactionQueriesQueryIDResponsesBadGateway() *PostTransactionQueriesQueryIDResponsesBadGateway {
	return &PostTransactionQueriesQueryIDResponsesBadGateway{}
}

/*PostTransactionQueriesQueryIDResponsesBadGateway handles this case with default header values.

gateway issue
*/
type PostTransactionQueriesQueryIDResponsesBadGateway struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *PostTransactionQueriesQueryIDResponsesBadGateway) Error() string {
	return fmt.Sprintf("[POST /transaction/queries/{query_id}/responses][%d] postTransactionQueriesQueryIdResponsesBadGateway", 502)
}

func (o *PostTransactionQueriesQueryIDResponsesBadGateway) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}
