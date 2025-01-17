// Code generated by go-swagger; DO NOT EDIT.

package payments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/go-form3/v4/pkg/generated/models"
)

// GetTransactionPaymentsIDAdmissionsAdmissionIDTasksTaskIDReader is a Reader for the GetTransactionPaymentsIDAdmissionsAdmissionIDTasksTaskID structure.
type GetTransactionPaymentsIDAdmissionsAdmissionIDTasksTaskIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTransactionPaymentsIDAdmissionsAdmissionIDTasksTaskIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetTransactionPaymentsIDAdmissionsAdmissionIDTasksTaskIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetTransactionPaymentsIDAdmissionsAdmissionIDTasksTaskIDOK creates a GetTransactionPaymentsIDAdmissionsAdmissionIDTasksTaskIDOK with default headers values
func NewGetTransactionPaymentsIDAdmissionsAdmissionIDTasksTaskIDOK() *GetTransactionPaymentsIDAdmissionsAdmissionIDTasksTaskIDOK {
	return &GetTransactionPaymentsIDAdmissionsAdmissionIDTasksTaskIDOK{}
}

/*GetTransactionPaymentsIDAdmissionsAdmissionIDTasksTaskIDOK handles this case with default header values.

Payment Admission Task Details details
*/
type GetTransactionPaymentsIDAdmissionsAdmissionIDTasksTaskIDOK struct {

	//Payload

	// isStream: false
	*models.PaymentAdmissionTaskDetailsResponse
}

func (o *GetTransactionPaymentsIDAdmissionsAdmissionIDTasksTaskIDOK) Error() string {
	return fmt.Sprintf("[GET /transaction/payments/{id}/admissions/{admissionId}/tasks/{taskId}][%d] getTransactionPaymentsIdAdmissionsAdmissionIdTasksTaskIdOK", 200)
}

func (o *GetTransactionPaymentsIDAdmissionsAdmissionIDTasksTaskIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.PaymentAdmissionTaskDetailsResponse = new(models.PaymentAdmissionTaskDetailsResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.PaymentAdmissionTaskDetailsResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}
