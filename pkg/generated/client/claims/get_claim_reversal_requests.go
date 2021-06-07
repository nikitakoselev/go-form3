// Code generated by go-swagger; DO NOT EDIT.

package claims

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// Client.GetClaimReversal creates a new GetClaimReversalRequest object
// with the default values initialized.
func (c *Client) GetClaimReversal() *GetClaimReversalRequest {
	var ()
	return &GetClaimReversalRequest{

		ID: c.Defaults.GetStrfmtUUID("GetClaimReversal", "id"),

		ReversalID: c.Defaults.GetStrfmtUUID("GetClaimReversal", "reversalId"),

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type GetClaimReversalRequest struct {

	/*ID      Claim Id      */

	ID strfmt.UUID

	/*ReversalID      Claim Reversal Id      */

	ReversalID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client

	transport runtime.ClientTransport
	formats   strfmt.Registry
}

func (o *GetClaimReversalRequest) FromJson(j string) (*GetClaimReversalRequest, error) {

	return o, nil
}

func (o *GetClaimReversalRequest) WithID(id strfmt.UUID) *GetClaimReversalRequest {

	o.ID = id

	return o
}

func (o *GetClaimReversalRequest) WithReversalID(reversalID strfmt.UUID) *GetClaimReversalRequest {

	o.ReversalID = reversalID

	return o
}

//////////////////
// WithContext adds the context to the get claim reversal Request
func (o *GetClaimReversalRequest) WithContext(ctx context.Context) *GetClaimReversalRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the get claim reversal Request
func (o *GetClaimReversalRequest) WithHTTPClient(client *http.Client) *GetClaimReversalRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *GetClaimReversalRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	// path param reversalId
	if err := r.SetPathParam("reversalId", o.ReversalID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
