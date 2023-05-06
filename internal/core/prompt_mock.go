// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package core

import (
	"github.com/AlecAivazis/survey/v2"
	"sync"
)

// Ensure, that PromptMock does implement Prompt.
// If this is not the case, regenerate this file with moq.
var _ Prompt = &PromptMock{}

// PromptMock is a mock implementation of Prompt.
//
//	func TestSomethingThatUsesPrompt(t *testing.T) {
//
//		// make and configure a mocked Prompt
//		mockedPrompt := &PromptMock{
//			AskOneFunc: func(p survey.Prompt, response interface{}, opts ...survey.AskOpt) error {
//				panic("mock out the AskOne method")
//			},
//		}
//
//		// use mockedPrompt in code that requires Prompt
//		// and then make assertions.
//
//	}
type PromptMock struct {
	// AskOneFunc mocks the AskOne method.
	AskOneFunc func(p survey.Prompt, response interface{}, opts ...survey.AskOpt) error

	// calls tracks calls to the methods.
	calls struct {
		// AskOne holds details about calls to the AskOne method.
		AskOne []struct {
			// P is the p argument value.
			P survey.Prompt
			// Response is the response argument value.
			Response interface{}
			// Opts is the opts argument value.
			Opts []survey.AskOpt
		}
	}
	lockAskOne sync.RWMutex
}

// AskOne calls AskOneFunc.
func (mock *PromptMock) AskOne(p survey.Prompt, response interface{}, opts ...survey.AskOpt) error {
	if mock.AskOneFunc == nil {
		panic("PromptMock.AskOneFunc: method is nil but Prompt.AskOne was just called")
	}
	callInfo := struct {
		P        survey.Prompt
		Response interface{}
		Opts     []survey.AskOpt
	}{
		P:        p,
		Response: response,
		Opts:     opts,
	}
	mock.lockAskOne.Lock()
	mock.calls.AskOne = append(mock.calls.AskOne, callInfo)
	mock.lockAskOne.Unlock()
	return mock.AskOneFunc(p, response, opts...)
}

// AskOneCalls gets all the calls that were made to AskOne.
// Check the length with:
//
//	len(mockedPrompt.AskOneCalls())
func (mock *PromptMock) AskOneCalls() []struct {
	P        survey.Prompt
	Response interface{}
	Opts     []survey.AskOpt
} {
	var calls []struct {
		P        survey.Prompt
		Response interface{}
		Opts     []survey.AskOpt
	}
	mock.lockAskOne.RLock()
	calls = mock.calls.AskOne
	mock.lockAskOne.RUnlock()
	return calls
}
