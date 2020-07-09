// Code generated by go-swagger; DO NOT EDIT.

package query_api

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// Client.GetTransactionQueriesQueryID creates a new GetTransactionQueriesQueryIDRequest object
// with the default values initialized.
func (c *Client) GetTransactionQueriesQueryID() *GetTransactionQueriesQueryIDRequest {
	var ()
	return &GetTransactionQueriesQueryIDRequest{

		QueryID: c.Defaults.GetStrfmtUUID("GetTransactionQueriesQueryID", "query_id"),

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type GetTransactionQueriesQueryIDRequest struct {

	/*QueryID      Query ID      */

	QueryID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client

	transport runtime.ClientTransport
	formats   strfmt.Registry
}

func (o *GetTransactionQueriesQueryIDRequest) FromJson(j string) *GetTransactionQueriesQueryIDRequest {

	return o
}

func (o *GetTransactionQueriesQueryIDRequest) WithQueryID(queryID strfmt.UUID) *GetTransactionQueriesQueryIDRequest {

	o.QueryID = queryID

	return o
}

//////////////////
// WithContext adds the context to the get transaction queries query ID Request
func (o *GetTransactionQueriesQueryIDRequest) WithContext(ctx context.Context) *GetTransactionQueriesQueryIDRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the get transaction queries query ID Request
func (o *GetTransactionQueriesQueryIDRequest) WithHTTPClient(client *http.Client) *GetTransactionQueriesQueryIDRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *GetTransactionQueriesQueryIDRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param query_id
	if err := r.SetPathParam("query_id", o.QueryID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
