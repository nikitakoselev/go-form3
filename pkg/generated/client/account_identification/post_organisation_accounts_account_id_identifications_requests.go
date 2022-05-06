// Code generated by go-swagger; DO NOT EDIT.

package account_identification

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

// Client.PostOrganisationAccountsAccountIDIdentifications creates a new PostOrganisationAccountsAccountIDIdentificationsRequest object
// with the default values initialized.
func (c *Client) PostOrganisationAccountsAccountIDIdentifications() *PostOrganisationAccountsAccountIDIdentificationsRequest {
	var ()
	return &PostOrganisationAccountsAccountIDIdentificationsRequest{

		AccountIdentificationRequest: models.AccountIdentificationRequestWithDefaults(c.Defaults),

		AccountID: c.Defaults.GetStrfmtUUID("PostOrganisationAccountsAccountIDIdentifications", "account_id"),

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type PostOrganisationAccountsAccountIDIdentificationsRequest struct {

	/*AccountIdentificationCreationRequest*/

	*models.AccountIdentificationRequest

	/*AccountID      Account Id to which this identification relates to      */

	AccountID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client

	transport runtime.ClientTransport
	formats   strfmt.Registry
}

func (o *PostOrganisationAccountsAccountIDIdentificationsRequest) FromJson(j string) (*PostOrganisationAccountsAccountIDIdentificationsRequest, error) {

	var m models.AccountIdentificationRequest
	if err := json.Unmarshal([]byte(j), &m); err != nil {
		return nil, fmt.Errorf("could not unmarshal JSON: %w", err)
	}

	o.AccountIdentificationRequest = &m

	return o, nil
}

func (o *PostOrganisationAccountsAccountIDIdentificationsRequest) WithAccountIdentificationCreationRequest(accountIdentificationCreationRequest models.AccountIdentificationRequest) *PostOrganisationAccountsAccountIDIdentificationsRequest {

	o.AccountIdentificationRequest = &accountIdentificationCreationRequest

	return o
}

func (o *PostOrganisationAccountsAccountIDIdentificationsRequest) WithoutAccountIdentificationCreationRequest() *PostOrganisationAccountsAccountIDIdentificationsRequest {

	o.AccountIdentificationRequest = &models.AccountIdentificationRequest{}

	return o
}

func (o *PostOrganisationAccountsAccountIDIdentificationsRequest) WithAccountID(accountID strfmt.UUID) *PostOrganisationAccountsAccountIDIdentificationsRequest {

	o.AccountID = accountID

	return o
}

//////////////////
// WithContext adds the context to the post organisation accounts account ID identifications Request
func (o *PostOrganisationAccountsAccountIDIdentificationsRequest) WithContext(ctx context.Context) *PostOrganisationAccountsAccountIDIdentificationsRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the post organisation accounts account ID identifications Request
func (o *PostOrganisationAccountsAccountIDIdentificationsRequest) WithHTTPClient(client *http.Client) *PostOrganisationAccountsAccountIDIdentificationsRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *PostOrganisationAccountsAccountIDIdentificationsRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// ISBODYPARAM
	if o.AccountIdentificationRequest != nil {
		if err := r.SetBodyParam(o.AccountIdentificationRequest); err != nil {
			return err
		}
	}

	// path param account_id
	if err := r.SetPathParam("account_id", o.AccountID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
