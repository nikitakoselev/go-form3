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

	"github.com/form3tech-oss/go-form3/v2/pkg/generated/models"
)

// Client.CreatePaymentAdvice creates a new CreatePaymentAdviceRequest object
// with the default values initialized.
func (c *Client) CreatePaymentAdvice() *CreatePaymentAdviceRequest {
	var ()
	return &CreatePaymentAdviceRequest{

		AdviceCreation: models.AdviceCreationWithDefaults(c.Defaults),

		ID: c.Defaults.GetStrfmtUUID("CreatePaymentAdvice", "id"),

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type CreatePaymentAdviceRequest struct {

	/*AdviceCreationRequest*/

	*models.AdviceCreation

	/*ID      Payment Id      */

	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client

	transport runtime.ClientTransport
	formats   strfmt.Registry
}

func (o *CreatePaymentAdviceRequest) FromJson(j string) (*CreatePaymentAdviceRequest, error) {

	var m models.AdviceCreation
	if err := json.Unmarshal([]byte(j), &m); err != nil {
		return nil, fmt.Errorf("could not unmarshal JSON: %w", err)
	}

	o.AdviceCreation = &m

	return o, nil
}

func (o *CreatePaymentAdviceRequest) WithAdviceCreationRequest(adviceCreationRequest models.AdviceCreation) *CreatePaymentAdviceRequest {

	o.AdviceCreation = &adviceCreationRequest

	return o
}

func (o *CreatePaymentAdviceRequest) WithoutAdviceCreationRequest() *CreatePaymentAdviceRequest {

	o.AdviceCreation = &models.AdviceCreation{}

	return o
}

func (o *CreatePaymentAdviceRequest) WithID(id strfmt.UUID) *CreatePaymentAdviceRequest {

	o.ID = id

	return o
}

//////////////////
// WithContext adds the context to the create payment advice Request
func (o *CreatePaymentAdviceRequest) WithContext(ctx context.Context) *CreatePaymentAdviceRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the create payment advice Request
func (o *CreatePaymentAdviceRequest) WithHTTPClient(client *http.Client) *CreatePaymentAdviceRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *CreatePaymentAdviceRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// ISBODYPARAM
	if o.AdviceCreation != nil {
		if err := r.SetBodyParam(o.AdviceCreation); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
