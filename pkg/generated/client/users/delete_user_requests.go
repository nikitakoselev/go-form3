// Code generated by go-swagger; DO NOT EDIT.

package users

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// Client.DeleteUser creates a new DeleteUserRequest object
// with the default values initialized.
func (c *Client) DeleteUser() *DeleteUserRequest {
	var ()
	return &DeleteUserRequest{

		UserID: c.Defaults.GetStrfmtUUID("DeleteUser", "user_id"),

		Version: c.Defaults.GetInt64("DeleteUser", "version"),

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type DeleteUserRequest struct {

	/*UserID      User Id      */

	UserID strfmt.UUID

	/*Version      Version      */

	Version int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client

	transport runtime.ClientTransport
	formats   strfmt.Registry
}

func (o *DeleteUserRequest) FromJson(j string) (*DeleteUserRequest, error) {

	return o, nil
}

func (o *DeleteUserRequest) WithUserID(userID strfmt.UUID) *DeleteUserRequest {

	o.UserID = userID

	return o
}

func (o *DeleteUserRequest) WithVersion(version int64) *DeleteUserRequest {

	o.Version = version

	return o
}

//////////////////
// WithContext adds the context to the delete user Request
func (o *DeleteUserRequest) WithContext(ctx context.Context) *DeleteUserRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the delete user Request
func (o *DeleteUserRequest) WithHTTPClient(client *http.Client) *DeleteUserRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *DeleteUserRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param user_id
	if err := r.SetPathParam("user_id", o.UserID.String()); err != nil {
		return err
	}

	// query param version
	qrVersion := o.Version
	qVersion := swag.FormatInt64(qrVersion)
	if qVersion != "" {
		if err := r.SetQueryParam("version", qVersion); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
