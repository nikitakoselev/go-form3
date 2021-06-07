// Code generated by go-swagger; DO NOT EDIT.

package a_c_e

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// Client.DeleteAce creates a new DeleteAceRequest object
// with the default values initialized.
func (c *Client) DeleteAce() *DeleteAceRequest {
	var ()
	return &DeleteAceRequest{

		AceID: c.Defaults.GetStrfmtUUID("DeleteAce", "ace_id"),

		RoleID: c.Defaults.GetStrfmtUUID("DeleteAce", "role_id"),

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type DeleteAceRequest struct {

	/*AceID      Ace Id      */

	AceID strfmt.UUID

	/*RoleID      Role Id      */

	RoleID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client

	transport runtime.ClientTransport
	formats   strfmt.Registry
}

func (o *DeleteAceRequest) FromJson(j string) (*DeleteAceRequest, error) {

	return o, nil
}

func (o *DeleteAceRequest) WithAceID(aceID strfmt.UUID) *DeleteAceRequest {

	o.AceID = aceID

	return o
}

func (o *DeleteAceRequest) WithRoleID(roleID strfmt.UUID) *DeleteAceRequest {

	o.RoleID = roleID

	return o
}

//////////////////
// WithContext adds the context to the delete ace Request
func (o *DeleteAceRequest) WithContext(ctx context.Context) *DeleteAceRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the delete ace Request
func (o *DeleteAceRequest) WithHTTPClient(client *http.Client) *DeleteAceRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *DeleteAceRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param ace_id
	if err := r.SetPathParam("ace_id", o.AceID.String()); err != nil {
		return err
	}

	// path param role_id
	if err := r.SetPathParam("role_id", o.RoleID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
