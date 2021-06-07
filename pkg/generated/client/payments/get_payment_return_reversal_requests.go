// Code generated by go-swagger; DO NOT EDIT.

package payments

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// Client.GetPaymentReturnReversal creates a new GetPaymentReturnReversalRequest object
// with the default values initialized.
func (c *Client) GetPaymentReturnReversal() *GetPaymentReturnReversalRequest {
	var ()
	return &GetPaymentReturnReversalRequest{

		ID: c.Defaults.GetStrfmtUUID("GetPaymentReturnReversal", "id"),

		ReturnID: c.Defaults.GetStrfmtUUID("GetPaymentReturnReversal", "returnId"),

		ReversalID: c.Defaults.GetStrfmtUUID("GetPaymentReturnReversal", "reversalId"),

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type GetPaymentReturnReversalRequest struct {

	/*ID      Payment Id      */

	ID strfmt.UUID

	/*ReturnID      Return Id      */

	ReturnID strfmt.UUID

	/*ReversalID      Reversal Id      */

	ReversalID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client

	transport runtime.ClientTransport
	formats   strfmt.Registry
}

func (o *GetPaymentReturnReversalRequest) FromJson(j string) (*GetPaymentReturnReversalRequest, error) {

	return o, nil
}

func (o *GetPaymentReturnReversalRequest) WithID(id strfmt.UUID) *GetPaymentReturnReversalRequest {

	o.ID = id

	return o
}

func (o *GetPaymentReturnReversalRequest) WithReturnID(returnID strfmt.UUID) *GetPaymentReturnReversalRequest {

	o.ReturnID = returnID

	return o
}

func (o *GetPaymentReturnReversalRequest) WithReversalID(reversalID strfmt.UUID) *GetPaymentReturnReversalRequest {

	o.ReversalID = reversalID

	return o
}

//////////////////
// WithContext adds the context to the get payment return reversal Request
func (o *GetPaymentReturnReversalRequest) WithContext(ctx context.Context) *GetPaymentReturnReversalRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the get payment return reversal Request
func (o *GetPaymentReturnReversalRequest) WithHTTPClient(client *http.Client) *GetPaymentReturnReversalRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *GetPaymentReturnReversalRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	// path param returnId
	if err := r.SetPathParam("returnId", o.ReturnID.String()); err != nil {
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
