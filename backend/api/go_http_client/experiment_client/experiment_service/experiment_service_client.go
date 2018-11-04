// Copyright 2018 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by go-swagger; DO NOT EDIT.

package experiment_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new experiment service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for experiment service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreateExperiment create experiment API
*/
func (a *Client) CreateExperiment(params *CreateExperimentParams, authInfo runtime.ClientAuthInfoWriter) (*CreateExperimentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateExperimentParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateExperiment",
		Method:             "POST",
		PathPattern:        "/apis/v1beta1/experiments",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateExperimentReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateExperimentOK), nil

}

/*
GetExperiment get experiment API
*/
func (a *Client) GetExperiment(params *GetExperimentParams, authInfo runtime.ClientAuthInfoWriter) (*GetExperimentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetExperimentParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetExperiment",
		Method:             "GET",
		PathPattern:        "/apis/v1beta1/experiments/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetExperimentReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetExperimentOK), nil

}

/*
ListExperiments list experiments API
*/
func (a *Client) ListExperiments(params *ListExperimentsParams, authInfo runtime.ClientAuthInfoWriter) (*ListExperimentsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListExperimentsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListExperiments",
		Method:             "GET",
		PathPattern:        "/apis/v1beta1/experiments",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListExperimentsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListExperimentsOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
