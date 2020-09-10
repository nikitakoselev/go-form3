// Code generated by go-swagger; DO NOT EDIT.

package direct_debits

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/go-form3/v2/pkg/generated/models"
)

// Client.PostTransactionDirectdebitsIDRecalls creates a new PostTransactionDirectdebitsIDRecallsRequest object
// with the default values initialized.
func (c *Client) PostTransactionDirectdebitsIDRecalls() *PostTransactionDirectdebitsIDRecallsRequest {
	var ()
	return &PostTransactionDirectdebitsIDRecallsRequest{

		DirectDebitRecallCreation: models.DirectDebitRecallCreationWithDefaults(c.Defaults),

		ID: c.Defaults.GetStrfmtUUID("PostTransactionDirectdebitsIDRecalls", "id"),

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type PostTransactionDirectdebitsIDRecallsRequest struct {

	/*RecallCreationRequest*/

	*models.DirectDebitRecallCreation

	/*ID      Direct Debit Id      */

	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client

	transport runtime.ClientTransport
	formats   strfmt.Registry
}

func (o *PostTransactionDirectdebitsIDRecallsRequest) FromJson(j string) *PostTransactionDirectdebitsIDRecallsRequest {

	var m models.DirectDebitRecallCreation
	if err := json.Unmarshal([]byte(j), &m); err != nil {
		log.Fatal(err)
	}

	o.DirectDebitRecallCreation = &m

	return o
}

func (o *PostTransactionDirectdebitsIDRecallsRequest) WithRecallCreationRequest(recallCreationRequest models.DirectDebitRecallCreation) *PostTransactionDirectdebitsIDRecallsRequest {

	o.DirectDebitRecallCreation = &recallCreationRequest

	return o
}

func (o *PostTransactionDirectdebitsIDRecallsRequest) WithoutRecallCreationRequest() *PostTransactionDirectdebitsIDRecallsRequest {

	o.DirectDebitRecallCreation = &models.DirectDebitRecallCreation{}

	return o
}

func (o *PostTransactionDirectdebitsIDRecallsRequest) WithID(id strfmt.UUID) *PostTransactionDirectdebitsIDRecallsRequest {

	o.ID = id

	return o
}

//////////////////
// WithContext adds the context to the post transaction directdebits ID recalls Request
func (o *PostTransactionDirectdebitsIDRecallsRequest) WithContext(ctx context.Context) *PostTransactionDirectdebitsIDRecallsRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the post transaction directdebits ID recalls Request
func (o *PostTransactionDirectdebitsIDRecallsRequest) WithHTTPClient(client *http.Client) *PostTransactionDirectdebitsIDRecallsRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *PostTransactionDirectdebitsIDRecallsRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// ISBODYPARAM
	if o.DirectDebitRecallCreation != nil {
		if err := r.SetBodyParam(o.DirectDebitRecallCreation); err != nil {
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
