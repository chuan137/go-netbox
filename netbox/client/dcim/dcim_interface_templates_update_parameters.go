// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2018 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package dcim

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/chuan137/go-netbox/netbox/models"
)

// NewDcimInterfaceTemplatesUpdateParams creates a new DcimInterfaceTemplatesUpdateParams object
// with the default values initialized.
func NewDcimInterfaceTemplatesUpdateParams() *DcimInterfaceTemplatesUpdateParams {
	var ()
	return &DcimInterfaceTemplatesUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDcimInterfaceTemplatesUpdateParamsWithTimeout creates a new DcimInterfaceTemplatesUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDcimInterfaceTemplatesUpdateParamsWithTimeout(timeout time.Duration) *DcimInterfaceTemplatesUpdateParams {
	var ()
	return &DcimInterfaceTemplatesUpdateParams{

		timeout: timeout,
	}
}

// NewDcimInterfaceTemplatesUpdateParamsWithContext creates a new DcimInterfaceTemplatesUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewDcimInterfaceTemplatesUpdateParamsWithContext(ctx context.Context) *DcimInterfaceTemplatesUpdateParams {
	var ()
	return &DcimInterfaceTemplatesUpdateParams{

		Context: ctx,
	}
}

// NewDcimInterfaceTemplatesUpdateParamsWithHTTPClient creates a new DcimInterfaceTemplatesUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDcimInterfaceTemplatesUpdateParamsWithHTTPClient(client *http.Client) *DcimInterfaceTemplatesUpdateParams {
	var ()
	return &DcimInterfaceTemplatesUpdateParams{
		HTTPClient: client,
	}
}

/*DcimInterfaceTemplatesUpdateParams contains all the parameters to send to the API endpoint
for the dcim interface templates update operation typically these are written to a http.Request
*/
type DcimInterfaceTemplatesUpdateParams struct {

	/*Data*/
	Data *models.WritableInterfaceTemplate
	/*ID
	  A unique integer value identifying this interface template.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dcim interface templates update params
func (o *DcimInterfaceTemplatesUpdateParams) WithTimeout(timeout time.Duration) *DcimInterfaceTemplatesUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim interface templates update params
func (o *DcimInterfaceTemplatesUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim interface templates update params
func (o *DcimInterfaceTemplatesUpdateParams) WithContext(ctx context.Context) *DcimInterfaceTemplatesUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim interface templates update params
func (o *DcimInterfaceTemplatesUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim interface templates update params
func (o *DcimInterfaceTemplatesUpdateParams) WithHTTPClient(client *http.Client) *DcimInterfaceTemplatesUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim interface templates update params
func (o *DcimInterfaceTemplatesUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim interface templates update params
func (o *DcimInterfaceTemplatesUpdateParams) WithData(data *models.WritableInterfaceTemplate) *DcimInterfaceTemplatesUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim interface templates update params
func (o *DcimInterfaceTemplatesUpdateParams) SetData(data *models.WritableInterfaceTemplate) {
	o.Data = data
}

// WithID adds the id to the dcim interface templates update params
func (o *DcimInterfaceTemplatesUpdateParams) WithID(id int64) *DcimInterfaceTemplatesUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim interface templates update params
func (o *DcimInterfaceTemplatesUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimInterfaceTemplatesUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
