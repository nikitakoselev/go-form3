// Code generated by go-swagger; DO NOT EDIT.

package a_c_e

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/go-form3/v3/pkg/generated/models"
)

// ListAcesReader is a Reader for the ListAces structure.
type ListAcesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListAcesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListAcesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListAcesOK creates a ListAcesOK with default headers values
func NewListAcesOK() *ListAcesOK {
	return &ListAcesOK{}
}

/*ListAcesOK handles this case with default header values.

List of ACE details
*/
type ListAcesOK struct {

	//Payload

	// isStream: false
	*models.AceDetailsListResponse
}

func (o *ListAcesOK) Error() string {
	return fmt.Sprintf("[GET /security/roles/{role_id}/aces][%d] listAcesOK", 200)
}

func (o *ListAcesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.AceDetailsListResponse = new(models.AceDetailsListResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.AceDetailsListResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}
