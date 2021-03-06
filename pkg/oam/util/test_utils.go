package util

import (
	"encoding/json"

	"github.com/onsi/gomega/format"
	"github.com/onsi/gomega/types"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"sigs.k8s.io/yaml"

	"github.com/oam-dev/kubevela/apis/core.oam.dev/v1alpha2"
)

// JSONMarshal returns the JSON encoding
func JSONMarshal(o interface{}) []byte {
	j, _ := json.Marshal(o)
	return j
}

// AlreadyExistMatcher matches the error to be already exist
type AlreadyExistMatcher struct {
}

// Match matches error.
func (matcher AlreadyExistMatcher) Match(actual interface{}) (success bool, err error) {
	if actual == nil {
		return false, nil
	}
	actualError := actual.(error)
	return apierrors.IsAlreadyExists(actualError), nil
}

// FailureMessage builds an error message.
func (matcher AlreadyExistMatcher) FailureMessage(actual interface{}) (message string) {
	return format.Message(actual, "to be already exist")
}

// NegatedFailureMessage builds an error message.
func (matcher AlreadyExistMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	return format.Message(actual, "not to be already exist")
}

// NotFoundMatcher matches the error to be not found.
type NotFoundMatcher struct {
}

// Match matches the api error.
func (matcher NotFoundMatcher) Match(actual interface{}) (success bool, err error) {
	if actual == nil {
		return false, nil
	}
	actualError := actual.(error)
	return apierrors.IsNotFound(actualError), nil
}

// FailureMessage builds an error message.
func (matcher NotFoundMatcher) FailureMessage(actual interface{}) (message string) {
	return format.Message(actual, "to be not found")
}

// NegatedFailureMessage builds an error message.
func (matcher NotFoundMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	return format.Message(actual, "not to be not found")
}

// BeEquivalentToError matches the error to take care of nil.
func BeEquivalentToError(expected error) types.GomegaMatcher {
	return &ErrorMatcher{
		ExpectedError: expected,
	}
}

// ErrorMatcher matches errors.
type ErrorMatcher struct {
	ExpectedError error
}

// Match matches an error.
func (matcher ErrorMatcher) Match(actual interface{}) (success bool, err error) {
	if actual == nil {
		return matcher.ExpectedError == nil, nil
	}
	actualError := actual.(error)
	return actualError.Error() == matcher.ExpectedError.Error(), nil
}

// FailureMessage builds an error message.
//nolint:errorlint
// TODO(roywang) use errors.As() instead of type assertion on error
func (matcher ErrorMatcher) FailureMessage(actual interface{}) (message string) {
	actualError, actualOK := actual.(error)
	expectedError, expectedOK := matcher.ExpectedError.(error)

	if actualOK && expectedOK {
		return format.MessageWithDiff(actualError.Error(), "to equal", expectedError.Error())
	}

	if actualOK && !expectedOK {
		return format.Message(actualError.Error(), "to equal", expectedError.Error())
	}

	if !actualOK && expectedOK {
		return format.Message(actual, "to equal", expectedError.Error())
	}

	return format.Message(actual, "to equal", expectedError)
}

// NegatedFailureMessage builds an error message.
//nolint:errorlint
// TODO(roywang) use errors.As() instead of type assertion on error
func (matcher ErrorMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	actualError, actualOK := actual.(error)
	expectedError, expectedOK := matcher.ExpectedError.(error)

	if actualOK && expectedOK {
		return format.MessageWithDiff(actualError.Error(), "not to equal", expectedError.Error())
	}

	if actualOK && !expectedOK {
		return format.Message(actualError.Error(), "not to equal", expectedError.Error())
	}

	if !actualOK && expectedOK {
		return format.Message(actual, "not to equal", expectedError.Error())
	}

	return format.Message(actual, "not to equal", expectedError)
}

// UnMarshalStringToWorkloadDefinition parse a string to a workloadDefinition object
func UnMarshalStringToWorkloadDefinition(s string) (*v1alpha2.WorkloadDefinition, error) {
	obj := &v1alpha2.WorkloadDefinition{}
	_body, err := yaml.YAMLToJSON([]byte(s))
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(_body, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

// UnMarshalStringToTraitDefinition parse a string to a traitDefinition object
func UnMarshalStringToTraitDefinition(s string) (*v1alpha2.TraitDefinition, error) {
	obj := &v1alpha2.TraitDefinition{}
	_body, err := yaml.YAMLToJSON([]byte(s))
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(_body, obj); err != nil {
		return nil, err
	}
	return obj, nil
}
