// This file was auto-generated by Fern from our API Definition.

package testsuites

import (
	context "context"
	fmt "fmt"
	vellumclientgo "github.com/amckinney/vellum-client-go"
	core "github.com/amckinney/vellum-client-go/core"
	http "net/http"
)

type Client struct {
	baseURL string
	caller  *core.Caller
	header  http.Header
}

func NewClient(opts ...core.ClientOption) *Client {
	options := core.NewClientOptions()
	for _, opt := range opts {
		opt(options)
	}
	return &Client{
		baseURL: options.BaseURL,
		caller:  core.NewCaller(options.HTTPClient),
		header:  options.ToHeader(),
	}
}

// Upserts a new test case for a test suite, keying off of the optionally provided test case id.
//
// If an id is provided and has a match, the test case will be updated. If no id is provided or no match
// is found, a new test case will be appended to the end.
//
// Note that a full replacement of the test case is performed, so any fields not provided will be removed
// or overwritten with default values.
//
// A UUID string identifying this test suite.
func (c *Client) UpsertTestSuiteTestCase(ctx context.Context, id string, request *vellumclientgo.TestSuiteTestCaseRequest) (*vellumclientgo.TestSuiteTestCase, error) {
	baseURL := "https://api.vellum.ai"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"v1/test-suites/%v/test-cases", id)

	var response *vellumclientgo.TestSuiteTestCase
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:      endpointURL,
			Method:   http.MethodPost,
			Headers:  c.header,
			Request:  request,
			Response: &response,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Deletes an existing test case for a test suite, keying off of the test case id.
//
// A UUID string identifying this test suite.
// An id identifying the test case that you'd like to delete
func (c *Client) DeleteTestSuiteTestCase(ctx context.Context, id string, testCaseId string) error {
	baseURL := "https://api.vellum.ai"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"v1/test-suites/%v/test-cases/%v", id, testCaseId)

	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:     endpointURL,
			Method:  http.MethodDelete,
			Headers: c.header,
		},
	); err != nil {
		return err
	}
	return nil
}
