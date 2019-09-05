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

// GetMandatesHealthReader is a Reader for the GetMandatesHealth structure.
type GetMandatesHealthReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMandatesHealthReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetMandatesHealthOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetMandatesHealthOK creates a GetMandatesHealthOK with default headers values
func NewGetMandatesHealthOK() *GetMandatesHealthOK {
	return &GetMandatesHealthOK{}
}

/*GetMandatesHealthOK handles this case with default header values.

Mandate service health
*/
type GetMandatesHealthOK struct {

	//Payload
	*models.Health
}

func (o *GetMandatesHealthOK) Error() string {
	return fmt.Sprintf("[GET /transaction/mandates/health][%d] getMandatesHealthOK  %+v", 200, o)
}

func (o *GetMandatesHealthOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Health = new(models.Health)

	// response payload
	if err := consumer.Consume(response.Body(), o.Health); err != nil && err != io.EOF {
		return err
	}

	return nil
}