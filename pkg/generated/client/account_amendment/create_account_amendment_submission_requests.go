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

	"github.com/form3tech-oss/go-form3/v4/pkg/generated/models"
)

// Client.CreateAccountAmendmentSubmission creates a new CreateAccountAmendmentSubmissionRequest object
// with the default values initialized.
func (c *Client) CreateAccountAmendmentSubmission() *CreateAccountAmendmentSubmissionRequest {
	var ()
	return &CreateAccountAmendmentSubmissionRequest{

		AccountAmendmentSubmissionCreation: models.AccountAmendmentSubmissionCreationWithDefaults(c.Defaults),

		ID: c.Defaults.GetStrfmtUUID("CreateAccountAmendmentSubmission", "id"),

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type CreateAccountAmendmentSubmissionRequest struct {

	/*AccountAmendmentSubmissionRequest*/

	*models.AccountAmendmentSubmissionCreation

	/*ID      Account Amendment ID      */

	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client

	transport runtime.ClientTransport
	formats   strfmt.Registry
}

func (o *CreateAccountAmendmentSubmissionRequest) FromJson(j string) (*CreateAccountAmendmentSubmissionRequest, error) {

	var m models.AccountAmendmentSubmissionCreation
	if err := json.Unmarshal([]byte(j), &m); err != nil {
		return nil, fmt.Errorf("could not unmarshal JSON: %w", err)
	}

	o.AccountAmendmentSubmissionCreation = &m

	return o, nil
}

func (o *CreateAccountAmendmentSubmissionRequest) WithAccountAmendmentSubmissionRequest(accountAmendmentSubmissionRequest models.AccountAmendmentSubmissionCreation) *CreateAccountAmendmentSubmissionRequest {

	o.AccountAmendmentSubmissionCreation = &accountAmendmentSubmissionRequest

	return o
}

func (o *CreateAccountAmendmentSubmissionRequest) WithoutAccountAmendmentSubmissionRequest() *CreateAccountAmendmentSubmissionRequest {

	o.AccountAmendmentSubmissionCreation = &models.AccountAmendmentSubmissionCreation{}

	return o
}

func (o *CreateAccountAmendmentSubmissionRequest) WithID(id strfmt.UUID) *CreateAccountAmendmentSubmissionRequest {

	o.ID = id

	return o
}

//////////////////
// WithContext adds the context to the create account amendment submission Request
func (o *CreateAccountAmendmentSubmissionRequest) WithContext(ctx context.Context) *CreateAccountAmendmentSubmissionRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the create account amendment submission Request
func (o *CreateAccountAmendmentSubmissionRequest) WithHTTPClient(client *http.Client) *CreateAccountAmendmentSubmissionRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *CreateAccountAmendmentSubmissionRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// ISBODYPARAM
	if o.AccountAmendmentSubmissionCreation != nil {
		if err := r.SetBodyParam(o.AccountAmendmentSubmissionCreation); err != nil {
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
