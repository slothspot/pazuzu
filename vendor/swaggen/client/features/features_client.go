package features

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new features API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for features API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeleteFeaturesName deletes feature

Delete this feature. If another feature depends on this feature nothing will be deleted and an error will be
returned.

*/
func (a *Client) DeleteFeaturesName(params *DeleteFeaturesNameParams) (*DeleteFeaturesNameNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteFeaturesNameParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteFeaturesName",
		Method:             "DELETE",
		PathPattern:        "/features/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteFeaturesNameReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteFeaturesNameNoContent), nil

}

/*
GetDependencies topologicalies sorted list of features

Returns the list of features topologicaly sorted base on the dependencies base on list of top features.

*/
func (a *Client) GetDependencies(params *GetDependenciesParams) (*GetDependenciesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDependenciesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetDependencies",
		Method:             "GET",
		PathPattern:        "/dependencies",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetDependenciesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetDependenciesOK), nil

}

/*
GetFeatures gets features

Returns the features sorted by name.
*/
func (a *Client) GetFeatures(params *GetFeaturesParams) (*GetFeaturesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFeaturesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetFeatures",
		Method:             "GET",
		PathPattern:        "/features",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetFeaturesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetFeaturesOK), nil

}

/*
GetFeaturesName gets single feature

Returns all the feature data without its dependencies.
*/
func (a *Client) GetFeaturesName(params *GetFeaturesNameParams) (*GetFeaturesNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFeaturesNameParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetFeaturesName",
		Method:             "GET",
		PathPattern:        "/features/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetFeaturesNameReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetFeaturesNameOK), nil

}

/*
PostFeatures creates a new feature

Creates a new feature with specified name, docker data, etc. Feature name is checked to be unique.
*/
func (a *Client) PostFeatures(params *PostFeaturesParams) (*PostFeaturesCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostFeaturesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostFeatures",
		Method:             "POST",
		PathPattern:        "/features",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostFeaturesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostFeaturesCreated), nil

}

/*
PutFeaturesName replaces this feature s content

Replaces snippet, test_snippet and meta data for this feature.
*/
func (a *Client) PutFeaturesName(params *PutFeaturesNameParams) (*PutFeaturesNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutFeaturesNameParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutFeaturesName",
		Method:             "PUT",
		PathPattern:        "/features/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PutFeaturesNameReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutFeaturesNameOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
