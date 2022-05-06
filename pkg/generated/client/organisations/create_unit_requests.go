// Code generated by go-swagger; DO NOT EDIT.

package organisations

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

// Client.CreateUnit creates a new CreateUnitRequest object
// with the default values initialized.
func (c *Client) CreateUnit() *CreateUnitRequest {
	var ()
	return &CreateUnitRequest{

		OrganisationCreation: models.OrganisationCreationWithDefaults(c.Defaults),

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type CreateUnitRequest struct {

	/*OrganisationCreationRequest*/

	*models.OrganisationCreation

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client

	transport runtime.ClientTransport
	formats   strfmt.Registry
}

func (o *CreateUnitRequest) FromJson(j string) (*CreateUnitRequest, error) {

	var m models.OrganisationCreation
	if err := json.Unmarshal([]byte(j), &m); err != nil {
		return nil, fmt.Errorf("could not unmarshal JSON: %w", err)
	}

	o.OrganisationCreation = &m

	return o, nil
}

func (o *CreateUnitRequest) WithOrganisationCreationRequest(organisationCreationRequest models.OrganisationCreation) *CreateUnitRequest {

	o.OrganisationCreation = &organisationCreationRequest

	return o
}

func (o *CreateUnitRequest) WithoutOrganisationCreationRequest() *CreateUnitRequest {

	o.OrganisationCreation = &models.OrganisationCreation{}

	return o
}

//////////////////
// WithContext adds the context to the create unit Request
func (o *CreateUnitRequest) WithContext(ctx context.Context) *CreateUnitRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the create unit Request
func (o *CreateUnitRequest) WithHTTPClient(client *http.Client) *CreateUnitRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *CreateUnitRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// ISBODYPARAM
	if o.OrganisationCreation != nil {
		if err := r.SetBodyParam(o.OrganisationCreation); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
