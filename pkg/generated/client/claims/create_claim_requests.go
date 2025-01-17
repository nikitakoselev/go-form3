// Code generated by go-swagger; DO NOT EDIT.

package claims

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

// Client.CreateClaim creates a new CreateClaimRequest object
// with the default values initialized.
func (c *Client) CreateClaim() *CreateClaimRequest {
	var ()
	return &CreateClaimRequest{

		ClaimCreation: models.ClaimCreationWithDefaults(c.Defaults),

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type CreateClaimRequest struct {

	/*ClaimCreationRequest*/

	*models.ClaimCreation

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client

	transport runtime.ClientTransport
	formats   strfmt.Registry
}

func (o *CreateClaimRequest) FromJson(j string) (*CreateClaimRequest, error) {

	var m models.ClaimCreation
	if err := json.Unmarshal([]byte(j), &m); err != nil {
		return nil, fmt.Errorf("could not unmarshal JSON: %w", err)
	}

	o.ClaimCreation = &m

	return o, nil
}

func (o *CreateClaimRequest) WithClaimCreationRequest(claimCreationRequest models.ClaimCreation) *CreateClaimRequest {

	o.ClaimCreation = &claimCreationRequest

	return o
}

func (o *CreateClaimRequest) WithoutClaimCreationRequest() *CreateClaimRequest {

	o.ClaimCreation = &models.ClaimCreation{}

	return o
}

//////////////////
// WithContext adds the context to the create claim Request
func (o *CreateClaimRequest) WithContext(ctx context.Context) *CreateClaimRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the create claim Request
func (o *CreateClaimRequest) WithHTTPClient(client *http.Client) *CreateClaimRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *CreateClaimRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// ISBODYPARAM
	if o.ClaimCreation != nil {
		if err := r.SetBodyParam(o.ClaimCreation); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
