// Code generated by go-swagger; DO NOT EDIT.

package account_amendment

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

// Client.CreateAccountAmendment creates a new CreateAccountAmendmentRequest object
// with the default values initialized.
func (c *Client) CreateAccountAmendment() *CreateAccountAmendmentRequest {
	var ()
	return &CreateAccountAmendmentRequest{

		AccountAmendmentCreation: models.AccountAmendmentCreationWithDefaults(c.Defaults),

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type CreateAccountAmendmentRequest struct {

	/*AccountAmendmentCreationBody*/

	*models.AccountAmendmentCreation

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client

	transport runtime.ClientTransport
	formats   strfmt.Registry
}

func (o *CreateAccountAmendmentRequest) FromJson(j string) (*CreateAccountAmendmentRequest, error) {

	var m models.AccountAmendmentCreation
	if err := json.Unmarshal([]byte(j), &m); err != nil {
		return nil, fmt.Errorf("could not unmarshal JSON: %w", err)
	}

	o.AccountAmendmentCreation = &m

	return o, nil
}

func (o *CreateAccountAmendmentRequest) WithAccountAmendmentCreationBody(accountAmendmentCreationBody models.AccountAmendmentCreation) *CreateAccountAmendmentRequest {

	o.AccountAmendmentCreation = &accountAmendmentCreationBody

	return o
}

func (o *CreateAccountAmendmentRequest) WithoutAccountAmendmentCreationBody() *CreateAccountAmendmentRequest {

	o.AccountAmendmentCreation = &models.AccountAmendmentCreation{}

	return o
}

//////////////////
// WithContext adds the context to the create account amendment Request
func (o *CreateAccountAmendmentRequest) WithContext(ctx context.Context) *CreateAccountAmendmentRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the create account amendment Request
func (o *CreateAccountAmendmentRequest) WithHTTPClient(client *http.Client) *CreateAccountAmendmentRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *CreateAccountAmendmentRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// ISBODYPARAM
	if o.AccountAmendmentCreation != nil {
		if err := r.SetBodyParam(o.AccountAmendmentCreation); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}