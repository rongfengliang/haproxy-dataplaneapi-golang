// Code generated by go-swagger; DO NOT EDIT.

package configuration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new configuration API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for configuration API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetHAProxyConfiguration(params *GetHAProxyConfigurationParams, authInfo runtime.ClientAuthInfoWriter) (*GetHAProxyConfigurationOK, error)

	PostHAProxyConfiguration(params *PostHAProxyConfigurationParams, authInfo runtime.ClientAuthInfoWriter) (*PostHAProxyConfigurationCreated, *PostHAProxyConfigurationAccepted, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetHAProxyConfiguration returns h a proxy configuration

  Returns HAProxy configuration file in plain text
*/
func (a *Client) GetHAProxyConfiguration(params *GetHAProxyConfigurationParams, authInfo runtime.ClientAuthInfoWriter) (*GetHAProxyConfigurationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetHAProxyConfigurationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getHAProxyConfiguration",
		Method:             "GET",
		PathPattern:        "/services/haproxy/configuration/raw",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetHAProxyConfigurationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetHAProxyConfigurationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetHAProxyConfigurationDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  PostHAProxyConfiguration pushes new haproxy configuration

  Push a new haproxy configuration file in plain text
*/
func (a *Client) PostHAProxyConfiguration(params *PostHAProxyConfigurationParams, authInfo runtime.ClientAuthInfoWriter) (*PostHAProxyConfigurationCreated, *PostHAProxyConfigurationAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostHAProxyConfigurationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postHAProxyConfiguration",
		Method:             "POST",
		PathPattern:        "/services/haproxy/configuration/raw",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"text/plain"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostHAProxyConfigurationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *PostHAProxyConfigurationCreated:
		return value, nil, nil
	case *PostHAProxyConfigurationAccepted:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostHAProxyConfigurationDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
