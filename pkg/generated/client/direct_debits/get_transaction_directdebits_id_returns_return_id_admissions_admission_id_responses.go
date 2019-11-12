// Code generated by go-swagger; DO NOT EDIT.

package direct_debits

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/form3tech-oss/go-form3/pkg/generated/models"
)

// GetTransactionDirectdebitsIDReturnsReturnIDAdmissionsAdmissionIDReader is a Reader for the GetTransactionDirectdebitsIDReturnsReturnIDAdmissionsAdmissionID structure.
type GetTransactionDirectdebitsIDReturnsReturnIDAdmissionsAdmissionIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTransactionDirectdebitsIDReturnsReturnIDAdmissionsAdmissionIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetTransactionDirectdebitsIDReturnsReturnIDAdmissionsAdmissionIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetTransactionDirectdebitsIDReturnsReturnIDAdmissionsAdmissionIDOK creates a GetTransactionDirectdebitsIDReturnsReturnIDAdmissionsAdmissionIDOK with default headers values
func NewGetTransactionDirectdebitsIDReturnsReturnIDAdmissionsAdmissionIDOK() *GetTransactionDirectdebitsIDReturnsReturnIDAdmissionsAdmissionIDOK {
	return &GetTransactionDirectdebitsIDReturnsReturnIDAdmissionsAdmissionIDOK{}
}

/*GetTransactionDirectdebitsIDReturnsReturnIDAdmissionsAdmissionIDOK handles this case with default header values.

Reversal admission details
*/
type GetTransactionDirectdebitsIDReturnsReturnIDAdmissionsAdmissionIDOK struct {

	//Payload

	// isStream: false
	*models.DirectDebitReturnAdmissionDetailsResponse
}

func (o *GetTransactionDirectdebitsIDReturnsReturnIDAdmissionsAdmissionIDOK) Error() string {
	return fmt.Sprintf("[GET /transaction/directdebits/{id}/returns/{returnId}/admissions/{admissionId}][%d] getTransactionDirectdebitsIdReturnsReturnIdAdmissionsAdmissionIdOK  %+v", 200, o)
}

func (o *GetTransactionDirectdebitsIDReturnsReturnIDAdmissionsAdmissionIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.DirectDebitReturnAdmissionDetailsResponse = new(models.DirectDebitReturnAdmissionDetailsResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.DirectDebitReturnAdmissionDetailsResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}