// Code generated by go-swagger; DO NOT EDIT.

package account_identification

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/go-form3/v3/pkg/generated/models"
)

// PostOrganisationAccountsAccountIDIdentificationsReader is a Reader for the PostOrganisationAccountsAccountIDIdentifications structure.
type PostOrganisationAccountsAccountIDIdentificationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostOrganisationAccountsAccountIDIdentificationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostOrganisationAccountsAccountIDIdentificationsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostOrganisationAccountsAccountIDIdentificationsCreated creates a PostOrganisationAccountsAccountIDIdentificationsCreated with default headers values
func NewPostOrganisationAccountsAccountIDIdentificationsCreated() *PostOrganisationAccountsAccountIDIdentificationsCreated {
	return &PostOrganisationAccountsAccountIDIdentificationsCreated{}
}

/*PostOrganisationAccountsAccountIDIdentificationsCreated handles this case with default header values.

Account Identification creation response
*/
type PostOrganisationAccountsAccountIDIdentificationsCreated struct {

	//Payload

	// isStream: false
	*models.AccountIdentificationResponse
}

func (o *PostOrganisationAccountsAccountIDIdentificationsCreated) Error() string {
	return fmt.Sprintf("[POST /organisation/accounts/{account_id}/identifications][%d] postOrganisationAccountsAccountIdIdentificationsCreated", 201)
}

func (o *PostOrganisationAccountsAccountIDIdentificationsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.AccountIdentificationResponse = new(models.AccountIdentificationResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.AccountIdentificationResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}
