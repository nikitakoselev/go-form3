// Code generated by go-swagger; DO NOT EDIT.

package payments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/go-form3/v2/pkg/generated/models"
)

// GetTransactionPaymentsIDAdmissionsAdmissionIDTasksReader is a Reader for the GetTransactionPaymentsIDAdmissionsAdmissionIDTasks structure.
type GetTransactionPaymentsIDAdmissionsAdmissionIDTasksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTransactionPaymentsIDAdmissionsAdmissionIDTasksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetTransactionPaymentsIDAdmissionsAdmissionIDTasksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetTransactionPaymentsIDAdmissionsAdmissionIDTasksOK creates a GetTransactionPaymentsIDAdmissionsAdmissionIDTasksOK with default headers values
func NewGetTransactionPaymentsIDAdmissionsAdmissionIDTasksOK() *GetTransactionPaymentsIDAdmissionsAdmissionIDTasksOK {
	return &GetTransactionPaymentsIDAdmissionsAdmissionIDTasksOK{}
}

/*GetTransactionPaymentsIDAdmissionsAdmissionIDTasksOK handles this case with default header values.

List of Task Details
*/
type GetTransactionPaymentsIDAdmissionsAdmissionIDTasksOK struct {

	//Payload

	// isStream: false
	*models.PaymentAdmissionTaskListResponse
}

func (o *GetTransactionPaymentsIDAdmissionsAdmissionIDTasksOK) Error() string {
	return fmt.Sprintf("[GET /transaction/payments/{id}/admissions/{admissionId}/tasks][%d] getTransactionPaymentsIdAdmissionsAdmissionIdTasksOK", 200)
}

func (o *GetTransactionPaymentsIDAdmissionsAdmissionIDTasksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.PaymentAdmissionTaskListResponse = new(models.PaymentAdmissionTaskListResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.PaymentAdmissionTaskListResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}