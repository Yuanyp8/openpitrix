// Code generated by go-swagger; DO NOT EDIT.

package access_manager

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"openpitrix.io/openpitrix/test/models"
)

// NewUnbindUserRoleParams creates a new UnbindUserRoleParams object
// with the default values initialized.
func NewUnbindUserRoleParams() *UnbindUserRoleParams {
	var ()
	return &UnbindUserRoleParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUnbindUserRoleParamsWithTimeout creates a new UnbindUserRoleParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUnbindUserRoleParamsWithTimeout(timeout time.Duration) *UnbindUserRoleParams {
	var ()
	return &UnbindUserRoleParams{

		timeout: timeout,
	}
}

// NewUnbindUserRoleParamsWithContext creates a new UnbindUserRoleParams object
// with the default values initialized, and the ability to set a context for a request
func NewUnbindUserRoleParamsWithContext(ctx context.Context) *UnbindUserRoleParams {
	var ()
	return &UnbindUserRoleParams{

		Context: ctx,
	}
}

// NewUnbindUserRoleParamsWithHTTPClient creates a new UnbindUserRoleParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUnbindUserRoleParamsWithHTTPClient(client *http.Client) *UnbindUserRoleParams {
	var ()
	return &UnbindUserRoleParams{
		HTTPClient: client,
	}
}

/*UnbindUserRoleParams contains all the parameters to send to the API endpoint
for the unbind user role operation typically these are written to a http.Request
*/
type UnbindUserRoleParams struct {

	/*Body*/
	Body *models.OpenpitrixUnbindUserRoleRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the unbind user role params
func (o *UnbindUserRoleParams) WithTimeout(timeout time.Duration) *UnbindUserRoleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the unbind user role params
func (o *UnbindUserRoleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the unbind user role params
func (o *UnbindUserRoleParams) WithContext(ctx context.Context) *UnbindUserRoleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the unbind user role params
func (o *UnbindUserRoleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the unbind user role params
func (o *UnbindUserRoleParams) WithHTTPClient(client *http.Client) *UnbindUserRoleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the unbind user role params
func (o *UnbindUserRoleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the unbind user role params
func (o *UnbindUserRoleParams) WithBody(body *models.OpenpitrixUnbindUserRoleRequest) *UnbindUserRoleParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the unbind user role params
func (o *UnbindUserRoleParams) SetBody(body *models.OpenpitrixUnbindUserRoleRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *UnbindUserRoleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
