// Code generated by go-swagger; DO NOT EDIT.

package mandates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/form3tech-oss/go-form3/pkg/generated/models"
)

// GetMandateReturnSubmissionReader is a Reader for the GetMandateReturnSubmission structure.
type GetMandateReturnSubmissionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMandateReturnSubmissionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetMandateReturnSubmissionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetMandateReturnSubmissionOK creates a GetMandateReturnSubmissionOK with default headers values
func NewGetMandateReturnSubmissionOK() *GetMandateReturnSubmissionOK {
	return &GetMandateReturnSubmissionOK{}
}

/*GetMandateReturnSubmissionOK handles this case with default header values.

Return submission details
*/
type GetMandateReturnSubmissionOK struct {

	//Payload
	*models.MandateReturnSubmissionDetailsResponse
}

func (o *GetMandateReturnSubmissionOK) Error() string {
	return fmt.Sprintf("[GET /transaction/mandates/{id}/returns/{returnId}/submissions/{submissionId}][%d] getMandateReturnSubmissionOK  %+v", 200, o)
}

func (o *GetMandateReturnSubmissionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.MandateReturnSubmissionDetailsResponse = new(models.MandateReturnSubmissionDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.MandateReturnSubmissionDetailsResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}