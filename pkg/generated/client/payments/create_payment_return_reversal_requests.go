// Code generated by go-swagger; DO NOT EDIT.

package payments

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/go-form3/v4/pkg/generated/models"
)

// Client.CreatePaymentReturnReversal creates a new CreatePaymentReturnReversalRequest object
// with the default values initialized.
func (c *Client) CreatePaymentReturnReversal() *CreatePaymentReturnReversalRequest {
	var ()
	return &CreatePaymentReturnReversalRequest{

		ReturnReversalCreation: models.ReturnReversalCreationWithDefaults(c.Defaults),

		ID: c.Defaults.GetStrfmtUUID("CreatePaymentReturnReversal", "id"),

		ReturnID: c.Defaults.GetStrfmtUUID("CreatePaymentReturnReversal", "returnId"),

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type CreatePaymentReturnReversalRequest struct {

	/*ReturnReversalCreationRequest*/

	*models.ReturnReversalCreation

	/*ID      Payment Id      */

	ID strfmt.UUID

	/*ReturnID      Return Id      */

	ReturnID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client

	transport runtime.ClientTransport
	formats   strfmt.Registry
}

func (o *CreatePaymentReturnReversalRequest) FromJson(j string) (*CreatePaymentReturnReversalRequest, error) {

	var m models.ReturnReversalCreation
	if err := json.Unmarshal([]byte(j), &m); err != nil {
		return nil, fmt.Errorf("could not unmarshal JSON: %w", err)
	}

	o.ReturnReversalCreation = &m

	return o, nil
}

func (o *CreatePaymentReturnReversalRequest) WithReturnReversalCreationRequest(returnReversalCreationRequest models.ReturnReversalCreation) *CreatePaymentReturnReversalRequest {

	o.ReturnReversalCreation = &returnReversalCreationRequest

	return o
}

func (o *CreatePaymentReturnReversalRequest) WithoutReturnReversalCreationRequest() *CreatePaymentReturnReversalRequest {

	o.ReturnReversalCreation = &models.ReturnReversalCreation{}

	return o
}

func (o *CreatePaymentReturnReversalRequest) WithID(id strfmt.UUID) *CreatePaymentReturnReversalRequest {

	o.ID = id

	return o
}

func (o *CreatePaymentReturnReversalRequest) WithReturnID(returnID strfmt.UUID) *CreatePaymentReturnReversalRequest {

	o.ReturnID = returnID

	return o
}

//////////////////
// WithContext adds the context to the create payment return reversal Request
func (o *CreatePaymentReturnReversalRequest) WithContext(ctx context.Context) *CreatePaymentReturnReversalRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the create payment return reversal Request
func (o *CreatePaymentReturnReversalRequest) WithHTTPClient(client *http.Client) *CreatePaymentReturnReversalRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *CreatePaymentReturnReversalRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// ISBODYPARAM
	if o.ReturnReversalCreation != nil {
		if err := r.SetBodyParam(o.ReturnReversalCreation); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	// path param returnId
	if err := r.SetPathParam("returnId", o.ReturnID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
