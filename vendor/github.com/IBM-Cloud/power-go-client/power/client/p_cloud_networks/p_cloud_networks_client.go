// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new p cloud networks API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new p cloud networks API client with basic auth credentials.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - user: user for basic authentication header.
// - password: password for basic authentication header.
func NewClientWithBasicAuth(host, basePath, scheme, user, password string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BasicAuth(user, password)
	return &Client{transport: transport, formats: strfmt.Default}
}

// New creates a new p cloud networks API client with a bearer token for authentication.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - bearerToken: bearer token for Bearer authentication header.
func NewClientWithBearerToken(host, basePath, scheme, bearerToken string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BearerToken(bearerToken)
	return &Client{transport: transport, formats: strfmt.Default}
}

/*
Client for p cloud networks API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption may be used to customize the behavior of Client methods.
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	PcloudNetworksDelete(params *PcloudNetworksDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PcloudNetworksDeleteOK, error)

	PcloudNetworksGet(params *PcloudNetworksGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PcloudNetworksGetOK, error)

	PcloudNetworksGetall(params *PcloudNetworksGetallParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PcloudNetworksGetallOK, error)

	PcloudNetworksPortsDelete(params *PcloudNetworksPortsDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PcloudNetworksPortsDeleteOK, error)

	PcloudNetworksPortsGet(params *PcloudNetworksPortsGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PcloudNetworksPortsGetOK, error)

	PcloudNetworksPortsGetall(params *PcloudNetworksPortsGetallParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PcloudNetworksPortsGetallOK, error)

	PcloudNetworksPortsPost(params *PcloudNetworksPortsPostParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PcloudNetworksPortsPostCreated, error)

	PcloudNetworksPortsPut(params *PcloudNetworksPortsPutParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PcloudNetworksPortsPutOK, error)

	PcloudNetworksPost(params *PcloudNetworksPostParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PcloudNetworksPostOK, *PcloudNetworksPostCreated, error)

	PcloudNetworksPut(params *PcloudNetworksPutParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PcloudNetworksPutOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
PcloudNetworksDelete deletes a network
*/
func (a *Client) PcloudNetworksDelete(params *PcloudNetworksDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PcloudNetworksDeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPcloudNetworksDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "pcloud.networks.delete",
		Method:             "DELETE",
		PathPattern:        "/pcloud/v1/cloud-instances/{cloud_instance_id}/networks/{network_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PcloudNetworksDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PcloudNetworksDeleteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for pcloud.networks.delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PcloudNetworksGet gets a network s current state information
*/
func (a *Client) PcloudNetworksGet(params *PcloudNetworksGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PcloudNetworksGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPcloudNetworksGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "pcloud.networks.get",
		Method:             "GET",
		PathPattern:        "/pcloud/v1/cloud-instances/{cloud_instance_id}/networks/{network_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PcloudNetworksGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PcloudNetworksGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for pcloud.networks.get: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PcloudNetworksGetall gets all networks in this cloud instance
*/
func (a *Client) PcloudNetworksGetall(params *PcloudNetworksGetallParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PcloudNetworksGetallOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPcloudNetworksGetallParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "pcloud.networks.getall",
		Method:             "GET",
		PathPattern:        "/pcloud/v1/cloud-instances/{cloud_instance_id}/networks",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PcloudNetworksGetallReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PcloudNetworksGetallOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for pcloud.networks.getall: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PcloudNetworksPortsDelete deletes a network port
*/
func (a *Client) PcloudNetworksPortsDelete(params *PcloudNetworksPortsDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PcloudNetworksPortsDeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPcloudNetworksPortsDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "pcloud.networks.ports.delete",
		Method:             "DELETE",
		PathPattern:        "/pcloud/v1/cloud-instances/{cloud_instance_id}/networks/{network_id}/ports/{port_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PcloudNetworksPortsDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PcloudNetworksPortsDeleteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for pcloud.networks.ports.delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PcloudNetworksPortsGet gets a port s information
*/
func (a *Client) PcloudNetworksPortsGet(params *PcloudNetworksPortsGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PcloudNetworksPortsGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPcloudNetworksPortsGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "pcloud.networks.ports.get",
		Method:             "GET",
		PathPattern:        "/pcloud/v1/cloud-instances/{cloud_instance_id}/networks/{network_id}/ports/{port_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PcloudNetworksPortsGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PcloudNetworksPortsGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for pcloud.networks.ports.get: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PcloudNetworksPortsGetall gets all ports for this network
*/
func (a *Client) PcloudNetworksPortsGetall(params *PcloudNetworksPortsGetallParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PcloudNetworksPortsGetallOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPcloudNetworksPortsGetallParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "pcloud.networks.ports.getall",
		Method:             "GET",
		PathPattern:        "/pcloud/v1/cloud-instances/{cloud_instance_id}/networks/{network_id}/ports",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PcloudNetworksPortsGetallReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PcloudNetworksPortsGetallOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for pcloud.networks.ports.getall: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PcloudNetworksPortsPost performs port addition deletion and listing
*/
func (a *Client) PcloudNetworksPortsPost(params *PcloudNetworksPortsPostParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PcloudNetworksPortsPostCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPcloudNetworksPortsPostParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "pcloud.networks.ports.post",
		Method:             "POST",
		PathPattern:        "/pcloud/v1/cloud-instances/{cloud_instance_id}/networks/{network_id}/ports",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PcloudNetworksPortsPostReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PcloudNetworksPortsPostCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for pcloud.networks.ports.post: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PcloudNetworksPortsPut updates a port s information
*/
func (a *Client) PcloudNetworksPortsPut(params *PcloudNetworksPortsPutParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PcloudNetworksPortsPutOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPcloudNetworksPortsPutParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "pcloud.networks.ports.put",
		Method:             "PUT",
		PathPattern:        "/pcloud/v1/cloud-instances/{cloud_instance_id}/networks/{network_id}/ports/{port_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PcloudNetworksPortsPutReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PcloudNetworksPortsPutOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for pcloud.networks.ports.put: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PcloudNetworksPost creates a new network
*/
func (a *Client) PcloudNetworksPost(params *PcloudNetworksPostParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PcloudNetworksPostOK, *PcloudNetworksPostCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPcloudNetworksPostParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "pcloud.networks.post",
		Method:             "POST",
		PathPattern:        "/pcloud/v1/cloud-instances/{cloud_instance_id}/networks",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PcloudNetworksPostReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *PcloudNetworksPostOK:
		return value, nil, nil
	case *PcloudNetworksPostCreated:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for p_cloud_networks: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PcloudNetworksPut updates a network
*/
func (a *Client) PcloudNetworksPut(params *PcloudNetworksPutParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PcloudNetworksPutOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPcloudNetworksPutParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "pcloud.networks.put",
		Method:             "PUT",
		PathPattern:        "/pcloud/v1/cloud-instances/{cloud_instance_id}/networks/{network_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PcloudNetworksPutReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PcloudNetworksPutOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for pcloud.networks.put: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
