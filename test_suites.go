// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
	core "github.com/amckinney/vellum-client-go/core"
)

type EvaluationParamsRequest struct {
	// The target value to compare the LLM output against. Typically what you expect or desire the LLM output to be.
	Target *string `json:"target,omitempty"`

	_rawJSON json.RawMessage
}

func (e *EvaluationParamsRequest) UnmarshalJSON(data []byte) error {
	type unmarshaler EvaluationParamsRequest
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*e = EvaluationParamsRequest(value)
	e._rawJSON = json.RawMessage(data)
	return nil
}

func (e *EvaluationParamsRequest) String() string {
	if len(e._rawJSON) > 0 {
		if value, err := core.StringifyJSON(e._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(e); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", e)
}

type TestSuiteTestCase struct {
	// The id of the test case to update. If none is provided, an id will be generated and a new test case will be appended.
	TestCaseId *string `json:"test_case_id,omitempty"`
	// A human-friendly label for the test case.
	Label *string `json:"label,omitempty"`
	// Key/value pairs for each input variable that the Test Suite expects.
	InputValues map[string]interface{} `json:"input_values,omitempty"`
	// Parameters to use when evaluating the test case, specific to the test suite's evaluation metric.
	EvaluationParams *EvaluationParams `json:"evaluation_params,omitempty"`

	_rawJSON json.RawMessage
}

func (t *TestSuiteTestCase) UnmarshalJSON(data []byte) error {
	type unmarshaler TestSuiteTestCase
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*t = TestSuiteTestCase(value)
	t._rawJSON = json.RawMessage(data)
	return nil
}

func (t *TestSuiteTestCase) String() string {
	if len(t._rawJSON) > 0 {
		if value, err := core.StringifyJSON(t._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(t); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", t)
}

type TestSuiteTestCaseRequest struct {
	// The id of the test case to update. If none is provided, an id will be generated and a new test case will be appended.
	TestCaseId *string `json:"test_case_id,omitempty"`
	// A human-friendly label for the test case.
	Label *string `json:"label,omitempty"`
	// Key/value pairs for each input variable that the Test Suite expects.
	InputValues map[string]interface{} `json:"input_values,omitempty"`
	// Parameters to use when evaluating the test case, specific to the test suite's evaluation metric.
	EvaluationParams *EvaluationParamsRequest `json:"evaluation_params,omitempty"`
}
