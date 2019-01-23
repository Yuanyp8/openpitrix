// Code generated by go-swagger; DO NOT EDIT.

package app_vendor_manager

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"openpitrix.io/openpitrix/test/models"
)

// SubmitVendorVerifyInfoReader is a Reader for the SubmitVendorVerifyInfo structure.
type SubmitVendorVerifyInfoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SubmitVendorVerifyInfoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSubmitVendorVerifyInfoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSubmitVendorVerifyInfoOK creates a SubmitVendorVerifyInfoOK with default headers values
func NewSubmitVendorVerifyInfoOK() *SubmitVendorVerifyInfoOK {
	return &SubmitVendorVerifyInfoOK{}
}

/*SubmitVendorVerifyInfoOK handles this case with default header values.

A successful response.
*/
type SubmitVendorVerifyInfoOK struct {
	Payload *models.OpenpitrixSubmitVendorVerifyInfoResponse
}

func (o *SubmitVendorVerifyInfoOK) Error() string {
	return fmt.Sprintf("[POST /v1/app_vendors][%d] submitVendorVerifyInfoOK  %+v", 200, o.Payload)
}

func (o *SubmitVendorVerifyInfoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OpenpitrixSubmitVendorVerifyInfoResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
