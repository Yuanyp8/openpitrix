// Code generated by go-swagger; DO NOT EDIT.

package cluster_manager

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDescribeSubnetsParams creates a new DescribeSubnetsParams object
// with the default values initialized.
func NewDescribeSubnetsParams() *DescribeSubnetsParams {
	var ()
	return &DescribeSubnetsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDescribeSubnetsParamsWithTimeout creates a new DescribeSubnetsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDescribeSubnetsParamsWithTimeout(timeout time.Duration) *DescribeSubnetsParams {
	var ()
	return &DescribeSubnetsParams{

		timeout: timeout,
	}
}

// NewDescribeSubnetsParamsWithContext creates a new DescribeSubnetsParams object
// with the default values initialized, and the ability to set a context for a request
func NewDescribeSubnetsParamsWithContext(ctx context.Context) *DescribeSubnetsParams {
	var ()
	return &DescribeSubnetsParams{

		Context: ctx,
	}
}

// NewDescribeSubnetsParamsWithHTTPClient creates a new DescribeSubnetsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDescribeSubnetsParamsWithHTTPClient(client *http.Client) *DescribeSubnetsParams {
	var ()
	return &DescribeSubnetsParams{
		HTTPClient: client,
	}
}

/*DescribeSubnetsParams contains all the parameters to send to the API endpoint
for the describe subnets operation typically these are written to a http.Request
*/
type DescribeSubnetsParams struct {

	/*AdvancedParam*/
	AdvancedParam []string
	/*Limit*/
	Limit *int64
	/*Offset*/
	Offset *int64
	/*RuntimeID*/
	RuntimeID *string
	/*SubnetID*/
	SubnetID []string
	/*SubnetTypeValue
	  The uint32 value.

	*/
	SubnetTypeValue *int64
	/*Zone*/
	Zone []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the describe subnets params
func (o *DescribeSubnetsParams) WithTimeout(timeout time.Duration) *DescribeSubnetsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the describe subnets params
func (o *DescribeSubnetsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the describe subnets params
func (o *DescribeSubnetsParams) WithContext(ctx context.Context) *DescribeSubnetsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the describe subnets params
func (o *DescribeSubnetsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the describe subnets params
func (o *DescribeSubnetsParams) WithHTTPClient(client *http.Client) *DescribeSubnetsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the describe subnets params
func (o *DescribeSubnetsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAdvancedParam adds the advancedParam to the describe subnets params
func (o *DescribeSubnetsParams) WithAdvancedParam(advancedParam []string) *DescribeSubnetsParams {
	o.SetAdvancedParam(advancedParam)
	return o
}

// SetAdvancedParam adds the advancedParam to the describe subnets params
func (o *DescribeSubnetsParams) SetAdvancedParam(advancedParam []string) {
	o.AdvancedParam = advancedParam
}

// WithLimit adds the limit to the describe subnets params
func (o *DescribeSubnetsParams) WithLimit(limit *int64) *DescribeSubnetsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the describe subnets params
func (o *DescribeSubnetsParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the describe subnets params
func (o *DescribeSubnetsParams) WithOffset(offset *int64) *DescribeSubnetsParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the describe subnets params
func (o *DescribeSubnetsParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithRuntimeID adds the runtimeID to the describe subnets params
func (o *DescribeSubnetsParams) WithRuntimeID(runtimeID *string) *DescribeSubnetsParams {
	o.SetRuntimeID(runtimeID)
	return o
}

// SetRuntimeID adds the runtimeId to the describe subnets params
func (o *DescribeSubnetsParams) SetRuntimeID(runtimeID *string) {
	o.RuntimeID = runtimeID
}

// WithSubnetID adds the subnetID to the describe subnets params
func (o *DescribeSubnetsParams) WithSubnetID(subnetID []string) *DescribeSubnetsParams {
	o.SetSubnetID(subnetID)
	return o
}

// SetSubnetID adds the subnetId to the describe subnets params
func (o *DescribeSubnetsParams) SetSubnetID(subnetID []string) {
	o.SubnetID = subnetID
}

// WithSubnetTypeValue adds the subnetTypeValue to the describe subnets params
func (o *DescribeSubnetsParams) WithSubnetTypeValue(subnetTypeValue *int64) *DescribeSubnetsParams {
	o.SetSubnetTypeValue(subnetTypeValue)
	return o
}

// SetSubnetTypeValue adds the subnetTypeValue to the describe subnets params
func (o *DescribeSubnetsParams) SetSubnetTypeValue(subnetTypeValue *int64) {
	o.SubnetTypeValue = subnetTypeValue
}

// WithZone adds the zone to the describe subnets params
func (o *DescribeSubnetsParams) WithZone(zone []string) *DescribeSubnetsParams {
	o.SetZone(zone)
	return o
}

// SetZone adds the zone to the describe subnets params
func (o *DescribeSubnetsParams) SetZone(zone []string) {
	o.Zone = zone
}

// WriteToRequest writes these params to a swagger request
func (o *DescribeSubnetsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	valuesAdvancedParam := o.AdvancedParam

	joinedAdvancedParam := swag.JoinByFormat(valuesAdvancedParam, "multi")
	// query array param advanced_param
	if err := r.SetQueryParam("advanced_param", joinedAdvancedParam...); err != nil {
		return err
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	if o.RuntimeID != nil {

		// query param runtime_id
		var qrRuntimeID string
		if o.RuntimeID != nil {
			qrRuntimeID = *o.RuntimeID
		}
		qRuntimeID := qrRuntimeID
		if qRuntimeID != "" {
			if err := r.SetQueryParam("runtime_id", qRuntimeID); err != nil {
				return err
			}
		}

	}

	valuesSubnetID := o.SubnetID

	joinedSubnetID := swag.JoinByFormat(valuesSubnetID, "multi")
	// query array param subnet_id
	if err := r.SetQueryParam("subnet_id", joinedSubnetID...); err != nil {
		return err
	}

	if o.SubnetTypeValue != nil {

		// query param subnet_type.value
		var qrSubnetTypeValue int64
		if o.SubnetTypeValue != nil {
			qrSubnetTypeValue = *o.SubnetTypeValue
		}
		qSubnetTypeValue := swag.FormatInt64(qrSubnetTypeValue)
		if qSubnetTypeValue != "" {
			if err := r.SetQueryParam("subnet_type.value", qSubnetTypeValue); err != nil {
				return err
			}
		}

	}

	valuesZone := o.Zone

	joinedZone := swag.JoinByFormat(valuesZone, "multi")
	// query array param zone
	if err := r.SetQueryParam("zone", joinedZone...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
