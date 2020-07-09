// Code generated by go-swagger; DO NOT EDIT.

package addressbook

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/form3tech-oss/go-form3/pkg/generated/models"
)

// GetAddressbookPartiesIDActorsPartyActorIDReader is a Reader for the GetAddressbookPartiesIDActorsPartyActorID structure.
type GetAddressbookPartiesIDActorsPartyActorIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAddressbookPartiesIDActorsPartyActorIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAddressbookPartiesIDActorsPartyActorIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewGetAddressbookPartiesIDActorsPartyActorIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetAddressbookPartiesIDActorsPartyActorIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAddressbookPartiesIDActorsPartyActorIDOK creates a GetAddressbookPartiesIDActorsPartyActorIDOK with default headers values
func NewGetAddressbookPartiesIDActorsPartyActorIDOK() *GetAddressbookPartiesIDActorsPartyActorIDOK {
	return &GetAddressbookPartiesIDActorsPartyActorIDOK{}
}

/*GetAddressbookPartiesIDActorsPartyActorIDOK handles this case with default header values.

get reponse
*/
type GetAddressbookPartiesIDActorsPartyActorIDOK struct {

	//Payload

	// isStream: false
	*models.PartyActorGetResponse
}

func (o *GetAddressbookPartiesIDActorsPartyActorIDOK) Error() string {
	return fmt.Sprintf("[GET /addressbook/parties/{id}/actors/{party_actor_id}][%d] getAddressbookPartiesIdActorsPartyActorIdOK", 200)
}

func (o *GetAddressbookPartiesIDActorsPartyActorIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.PartyActorGetResponse = new(models.PartyActorGetResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.PartyActorGetResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAddressbookPartiesIDActorsPartyActorIDForbidden creates a GetAddressbookPartiesIDActorsPartyActorIDForbidden with default headers values
func NewGetAddressbookPartiesIDActorsPartyActorIDForbidden() *GetAddressbookPartiesIDActorsPartyActorIDForbidden {
	return &GetAddressbookPartiesIDActorsPartyActorIDForbidden{}
}

/*GetAddressbookPartiesIDActorsPartyActorIDForbidden handles this case with default header values.

Forbidden
*/
type GetAddressbookPartiesIDActorsPartyActorIDForbidden struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *GetAddressbookPartiesIDActorsPartyActorIDForbidden) Error() string {
	return fmt.Sprintf("[GET /addressbook/parties/{id}/actors/{party_actor_id}][%d] getAddressbookPartiesIdActorsPartyActorIdForbidden", 403)
}

func (o *GetAddressbookPartiesIDActorsPartyActorIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAddressbookPartiesIDActorsPartyActorIDNotFound creates a GetAddressbookPartiesIDActorsPartyActorIDNotFound with default headers values
func NewGetAddressbookPartiesIDActorsPartyActorIDNotFound() *GetAddressbookPartiesIDActorsPartyActorIDNotFound {
	return &GetAddressbookPartiesIDActorsPartyActorIDNotFound{}
}

/*GetAddressbookPartiesIDActorsPartyActorIDNotFound handles this case with default header values.

Record not found
*/
type GetAddressbookPartiesIDActorsPartyActorIDNotFound struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *GetAddressbookPartiesIDActorsPartyActorIDNotFound) Error() string {
	return fmt.Sprintf("[GET /addressbook/parties/{id}/actors/{party_actor_id}][%d] getAddressbookPartiesIdActorsPartyActorIdNotFound", 404)
}

func (o *GetAddressbookPartiesIDActorsPartyActorIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}
