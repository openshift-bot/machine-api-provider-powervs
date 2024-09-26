// Code generated by go-swagger; DO NOT EDIT.

package networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/power-go-client/power/models"
)

// NewV1NetworksNetworkInterfacesPostParams creates a new V1NetworksNetworkInterfacesPostParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewV1NetworksNetworkInterfacesPostParams() *V1NetworksNetworkInterfacesPostParams {
	return &V1NetworksNetworkInterfacesPostParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewV1NetworksNetworkInterfacesPostParamsWithTimeout creates a new V1NetworksNetworkInterfacesPostParams object
// with the ability to set a timeout on a request.
func NewV1NetworksNetworkInterfacesPostParamsWithTimeout(timeout time.Duration) *V1NetworksNetworkInterfacesPostParams {
	return &V1NetworksNetworkInterfacesPostParams{
		timeout: timeout,
	}
}

// NewV1NetworksNetworkInterfacesPostParamsWithContext creates a new V1NetworksNetworkInterfacesPostParams object
// with the ability to set a context for a request.
func NewV1NetworksNetworkInterfacesPostParamsWithContext(ctx context.Context) *V1NetworksNetworkInterfacesPostParams {
	return &V1NetworksNetworkInterfacesPostParams{
		Context: ctx,
	}
}

// NewV1NetworksNetworkInterfacesPostParamsWithHTTPClient creates a new V1NetworksNetworkInterfacesPostParams object
// with the ability to set a custom HTTPClient for a request.
func NewV1NetworksNetworkInterfacesPostParamsWithHTTPClient(client *http.Client) *V1NetworksNetworkInterfacesPostParams {
	return &V1NetworksNetworkInterfacesPostParams{
		HTTPClient: client,
	}
}

/*
V1NetworksNetworkInterfacesPostParams contains all the parameters to send to the API endpoint

	for the v1 networks network interfaces post operation.

	Typically these are written to a http.Request.
*/
type V1NetworksNetworkInterfacesPostParams struct {

	/* Body.

	   Create a Network Interface
	*/
	Body *models.NetworkInterfaceCreate

	/* NetworkID.

	   Network ID
	*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the v1 networks network interfaces post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *V1NetworksNetworkInterfacesPostParams) WithDefaults() *V1NetworksNetworkInterfacesPostParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the v1 networks network interfaces post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *V1NetworksNetworkInterfacesPostParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the v1 networks network interfaces post params
func (o *V1NetworksNetworkInterfacesPostParams) WithTimeout(timeout time.Duration) *V1NetworksNetworkInterfacesPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 networks network interfaces post params
func (o *V1NetworksNetworkInterfacesPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 networks network interfaces post params
func (o *V1NetworksNetworkInterfacesPostParams) WithContext(ctx context.Context) *V1NetworksNetworkInterfacesPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 networks network interfaces post params
func (o *V1NetworksNetworkInterfacesPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 networks network interfaces post params
func (o *V1NetworksNetworkInterfacesPostParams) WithHTTPClient(client *http.Client) *V1NetworksNetworkInterfacesPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 networks network interfaces post params
func (o *V1NetworksNetworkInterfacesPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 networks network interfaces post params
func (o *V1NetworksNetworkInterfacesPostParams) WithBody(body *models.NetworkInterfaceCreate) *V1NetworksNetworkInterfacesPostParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 networks network interfaces post params
func (o *V1NetworksNetworkInterfacesPostParams) SetBody(body *models.NetworkInterfaceCreate) {
	o.Body = body
}

// WithNetworkID adds the networkID to the v1 networks network interfaces post params
func (o *V1NetworksNetworkInterfacesPostParams) WithNetworkID(networkID string) *V1NetworksNetworkInterfacesPostParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the v1 networks network interfaces post params
func (o *V1NetworksNetworkInterfacesPostParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *V1NetworksNetworkInterfacesPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param network_id
	if err := r.SetPathParam("network_id", o.NetworkID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
