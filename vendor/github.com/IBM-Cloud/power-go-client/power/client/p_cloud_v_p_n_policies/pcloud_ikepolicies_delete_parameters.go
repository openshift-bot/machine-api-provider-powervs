// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_v_p_n_policies

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
)

// NewPcloudIkepoliciesDeleteParams creates a new PcloudIkepoliciesDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPcloudIkepoliciesDeleteParams() *PcloudIkepoliciesDeleteParams {
	return &PcloudIkepoliciesDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPcloudIkepoliciesDeleteParamsWithTimeout creates a new PcloudIkepoliciesDeleteParams object
// with the ability to set a timeout on a request.
func NewPcloudIkepoliciesDeleteParamsWithTimeout(timeout time.Duration) *PcloudIkepoliciesDeleteParams {
	return &PcloudIkepoliciesDeleteParams{
		timeout: timeout,
	}
}

// NewPcloudIkepoliciesDeleteParamsWithContext creates a new PcloudIkepoliciesDeleteParams object
// with the ability to set a context for a request.
func NewPcloudIkepoliciesDeleteParamsWithContext(ctx context.Context) *PcloudIkepoliciesDeleteParams {
	return &PcloudIkepoliciesDeleteParams{
		Context: ctx,
	}
}

// NewPcloudIkepoliciesDeleteParamsWithHTTPClient creates a new PcloudIkepoliciesDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewPcloudIkepoliciesDeleteParamsWithHTTPClient(client *http.Client) *PcloudIkepoliciesDeleteParams {
	return &PcloudIkepoliciesDeleteParams{
		HTTPClient: client,
	}
}

/*
PcloudIkepoliciesDeleteParams contains all the parameters to send to the API endpoint

	for the pcloud ikepolicies delete operation.

	Typically these are written to a http.Request.
*/
type PcloudIkepoliciesDeleteParams struct {

	/* CloudInstanceID.

	   Cloud Instance ID of a PCloud Instance
	*/
	CloudInstanceID string

	/* IkePolicyID.

	   ID of a IKE Policy
	*/
	IkePolicyID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the pcloud ikepolicies delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PcloudIkepoliciesDeleteParams) WithDefaults() *PcloudIkepoliciesDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the pcloud ikepolicies delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PcloudIkepoliciesDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the pcloud ikepolicies delete params
func (o *PcloudIkepoliciesDeleteParams) WithTimeout(timeout time.Duration) *PcloudIkepoliciesDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pcloud ikepolicies delete params
func (o *PcloudIkepoliciesDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pcloud ikepolicies delete params
func (o *PcloudIkepoliciesDeleteParams) WithContext(ctx context.Context) *PcloudIkepoliciesDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pcloud ikepolicies delete params
func (o *PcloudIkepoliciesDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pcloud ikepolicies delete params
func (o *PcloudIkepoliciesDeleteParams) WithHTTPClient(client *http.Client) *PcloudIkepoliciesDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pcloud ikepolicies delete params
func (o *PcloudIkepoliciesDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCloudInstanceID adds the cloudInstanceID to the pcloud ikepolicies delete params
func (o *PcloudIkepoliciesDeleteParams) WithCloudInstanceID(cloudInstanceID string) *PcloudIkepoliciesDeleteParams {
	o.SetCloudInstanceID(cloudInstanceID)
	return o
}

// SetCloudInstanceID adds the cloudInstanceId to the pcloud ikepolicies delete params
func (o *PcloudIkepoliciesDeleteParams) SetCloudInstanceID(cloudInstanceID string) {
	o.CloudInstanceID = cloudInstanceID
}

// WithIkePolicyID adds the ikePolicyID to the pcloud ikepolicies delete params
func (o *PcloudIkepoliciesDeleteParams) WithIkePolicyID(ikePolicyID string) *PcloudIkepoliciesDeleteParams {
	o.SetIkePolicyID(ikePolicyID)
	return o
}

// SetIkePolicyID adds the ikePolicyId to the pcloud ikepolicies delete params
func (o *PcloudIkepoliciesDeleteParams) SetIkePolicyID(ikePolicyID string) {
	o.IkePolicyID = ikePolicyID
}

// WriteToRequest writes these params to a swagger request
func (o *PcloudIkepoliciesDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cloud_instance_id
	if err := r.SetPathParam("cloud_instance_id", o.CloudInstanceID); err != nil {
		return err
	}

	// path param ike_policy_id
	if err := r.SetPathParam("ike_policy_id", o.IkePolicyID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
