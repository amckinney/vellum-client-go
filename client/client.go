// This file was auto-generated by Fern from our API Definition.

package client

import (
	bytes "bytes"
	context "context"
	json "encoding/json"
	errors "errors"
	vellumclientgo "github.com/amckinney/vellum-client-go"
	core "github.com/amckinney/vellum-client-go/core"
	deployments "github.com/amckinney/vellum-client-go/deployments"
	documentindexes "github.com/amckinney/vellum-client-go/documentindexes"
	documents "github.com/amckinney/vellum-client-go/documents"
	modelversions "github.com/amckinney/vellum-client-go/modelversions"
	registeredprompts "github.com/amckinney/vellum-client-go/registeredprompts"
	sandboxes "github.com/amckinney/vellum-client-go/sandboxes"
	testsuites "github.com/amckinney/vellum-client-go/testsuites"
	io "io"
	http "net/http"
)

type Client struct {
	baseURL string
	caller  *core.Caller
	header  http.Header

	Deployments       *deployments.Client
	DocumentIndexes   *documentindexes.Client
	Documents         *documents.Client
	ModelVersions     *modelversions.Client
	RegisteredPrompts *registeredprompts.Client
	Sandboxes         *sandboxes.Client
	TestSuites        *testsuites.Client
}

func NewClient(opts ...core.ClientOption) *Client {
	options := core.NewClientOptions()
	for _, opt := range opts {
		opt(options)
	}
	return &Client{
		baseURL:           options.BaseURL,
		caller:            core.NewCaller(options.HTTPClient),
		header:            options.ToHeader(),
		Deployments:       deployments.NewClient(opts...),
		DocumentIndexes:   documentindexes.NewClient(opts...),
		Documents:         documents.NewClient(opts...),
		ModelVersions:     modelversions.NewClient(opts...),
		RegisteredPrompts: registeredprompts.NewClient(opts...),
		Sandboxes:         sandboxes.NewClient(opts...),
		TestSuites:        testsuites.NewClient(opts...),
	}
}

// Executes a deployed Workflow and streams back its results.
func (c *Client) ExecuteWorkflowStream(ctx context.Context, request *vellumclientgo.ExecuteWorkflowStreamRequest) (*core.Stream[vellumclientgo.WorkflowStreamEvent], error) {
	baseURL := "https://predict.vellum.ai"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := baseURL + "/" + "v1/execute-workflow-stream"

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 400:
			value := new(vellumclientgo.BadRequestError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 404:
			value := new(vellumclientgo.NotFoundError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 500:
			value := new(vellumclientgo.InternalServerError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		}
		return apiError
	}

	streamer := core.NewStreamer[vellumclientgo.WorkflowStreamEvent](c.caller)
	return streamer.Stream(
		ctx,
		&core.StreamParams{
			URL:          endpointURL,
			Method:       http.MethodPost,
			Headers:      c.header,
			Request:      request,
			ErrorDecoder: errorDecoder,
		},
	)
}

// Generate a completion using a previously defined deployment.
//
// **Note:** Uses a base url of `https://predict.vellum.ai`.
func (c *Client) Generate(ctx context.Context, request *vellumclientgo.GenerateBodyRequest) (*vellumclientgo.GenerateResponse, error) {
	baseURL := "https://predict.vellum.ai"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := baseURL + "/" + "v1/generate"

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 400:
			value := new(vellumclientgo.BadRequestError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 403:
			value := new(vellumclientgo.ForbiddenError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 404:
			value := new(vellumclientgo.NotFoundError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 500:
			value := new(vellumclientgo.InternalServerError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		}
		return apiError
	}

	var response *vellumclientgo.GenerateResponse
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:          endpointURL,
			Method:       http.MethodPost,
			Headers:      c.header,
			Request:      request,
			Response:     &response,
			ErrorDecoder: errorDecoder,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Generate a stream of completions using a previously defined deployment.
//
// **Note:** Uses a base url of `https://predict.vellum.ai`.
func (c *Client) GenerateStream(ctx context.Context, request *vellumclientgo.GenerateStreamBodyRequest) (*core.Stream[vellumclientgo.GenerateStreamResponse], error) {
	baseURL := "https://predict.vellum.ai"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := baseURL + "/" + "v1/generate-stream"

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 400:
			value := new(vellumclientgo.BadRequestError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 403:
			value := new(vellumclientgo.ForbiddenError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 404:
			value := new(vellumclientgo.NotFoundError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 500:
			value := new(vellumclientgo.InternalServerError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		}
		return apiError
	}

	streamer := core.NewStreamer[vellumclientgo.GenerateStreamResponse](c.caller)
	return streamer.Stream(
		ctx,
		&core.StreamParams{
			URL:          endpointURL,
			Method:       http.MethodPost,
			Headers:      c.header,
			Request:      request,
			ErrorDecoder: errorDecoder,
		},
	)
}

// Perform a search against a document index.
//
// **Note:** Uses a base url of `https://predict.vellum.ai`.
func (c *Client) Search(ctx context.Context, request *vellumclientgo.SearchRequestBodyRequest) (*vellumclientgo.SearchResponse, error) {
	baseURL := "https://predict.vellum.ai"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := baseURL + "/" + "v1/search"

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 400:
			value := new(vellumclientgo.BadRequestError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 404:
			value := new(vellumclientgo.NotFoundError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 500:
			value := new(vellumclientgo.InternalServerError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		}
		return apiError
	}

	var response *vellumclientgo.SearchResponse
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:          endpointURL,
			Method:       http.MethodPost,
			Headers:      c.header,
			Request:      request,
			Response:     &response,
			ErrorDecoder: errorDecoder,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Used to submit feedback regarding the quality of previously generated completions.
//
// **Note:** Uses a base url of `https://predict.vellum.ai`.
func (c *Client) SubmitCompletionActuals(ctx context.Context, request *vellumclientgo.SubmitCompletionActualsRequest) error {
	baseURL := "https://predict.vellum.ai"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := baseURL + "/" + "v1/submit-completion-actuals"

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 400:
			value := new(vellumclientgo.BadRequestError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 404:
			value := new(vellumclientgo.NotFoundError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 500:
			value := new(vellumclientgo.InternalServerError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		}
		return apiError
	}

	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:          endpointURL,
			Method:       http.MethodPost,
			Headers:      c.header,
			Request:      request,
			ErrorDecoder: errorDecoder,
		},
	); err != nil {
		return err
	}
	return nil
}

// Used to submit feedback regarding the quality of previous workflow execution and its outputs.
//
// **Note:** Uses a base url of `https://predict.vellum.ai`.
func (c *Client) SubmitWorkflowExecutionActuals(ctx context.Context, request *vellumclientgo.SubmitWorkflowExecutionActualsRequest) error {
	baseURL := "https://predict.vellum.ai"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := baseURL + "/" + "v1/submit-workflow-execution-actuals"

	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:     endpointURL,
			Method:  http.MethodPost,
			Headers: c.header,
			Request: request,
		},
	); err != nil {
		return err
	}
	return nil
}
