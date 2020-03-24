// Code generated by go-swagger; DO NOT EDIT.

package payments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/form3tech-oss/go-form3/pkg/generated/models"
)

// GetPaymentAdmissionReader is a Reader for the GetPaymentAdmission structure.
type GetPaymentAdmissionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPaymentAdmissionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPaymentAdmissionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPaymentAdmissionOK creates a GetPaymentAdmissionOK with default headers values
func NewGetPaymentAdmissionOK() *GetPaymentAdmissionOK {
	return &GetPaymentAdmissionOK{}
}

/*GetPaymentAdmissionOK handles this case with default header values.

Admission details
*/
type GetPaymentAdmissionOK struct {

	//Payload

	// isStream: false
	*models.PaymentAdmissionDetailsResponse
}

func (o *GetPaymentAdmissionOK) Error() string {
	return fmt.Sprintf("[GET /transaction/payments/{id}/admissions/{admissionId}][%d] getPaymentAdmissionOK", 200)
}

func (o *GetPaymentAdmissionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.PaymentAdmissionDetailsResponse = new(models.PaymentAdmissionDetailsResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.PaymentAdmissionDetailsResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}