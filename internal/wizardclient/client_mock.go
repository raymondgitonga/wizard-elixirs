// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package wizardclient

import (
	"sync"
)

// Ensure, that ClientMock does implement Client.
// If this is not the case, regenerate this file with moq.
var _ Client = &ClientMock{}

// ClientMock is a mock implementation of Client.
//
//	func TestSomethingThatUsesClient(t *testing.T) {
//
//		// make and configure a mocked Client
//		mockedClient := &ClientMock{
//			GetElixirsFunc: func(ingredients []string) ([]Elixir, error) {
//				panic("mock out the GetElixirs method")
//			},
//			GetIngredientsFunc: func() ([]Ingredient, error) {
//				panic("mock out the GetIngredients method")
//			},
//		}
//
//		// use mockedClient in code that requires Client
//		// and then make assertions.
//
//	}
type ClientMock struct {
	// GetElixirsFunc mocks the GetElixirs method.
	GetElixirsFunc func(ingredients []string) ([]Elixir, error)

	// GetIngredientsFunc mocks the GetIngredients method.
	GetIngredientsFunc func() ([]Ingredient, error)

	// calls tracks calls to the methods.
	calls struct {
		// GetElixirs holds details about calls to the GetElixirs method.
		GetElixirs []struct {
			// Ingredients is the ingredients argument value.
			Ingredients []string
		}
		// GetIngredients holds details about calls to the GetIngredients method.
		GetIngredients []struct {
		}
	}
	lockGetElixirs     sync.RWMutex
	lockGetIngredients sync.RWMutex
}

// GetElixirs calls GetElixirsFunc.
func (mock *ClientMock) GetElixirs(ingredients []string) ([]Elixir, error) {
	if mock.GetElixirsFunc == nil {
		panic("ClientMock.GetElixirsFunc: method is nil but Client.GetElixirs was just called")
	}
	callInfo := struct {
		Ingredients []string
	}{
		Ingredients: ingredients,
	}
	mock.lockGetElixirs.Lock()
	mock.calls.GetElixirs = append(mock.calls.GetElixirs, callInfo)
	mock.lockGetElixirs.Unlock()
	return mock.GetElixirsFunc(ingredients)
}

// GetElixirsCalls gets all the calls that were made to GetElixirs.
// Check the length with:
//
//	len(mockedClient.GetElixirsCalls())
func (mock *ClientMock) GetElixirsCalls() []struct {
	Ingredients []string
} {
	var calls []struct {
		Ingredients []string
	}
	mock.lockGetElixirs.RLock()
	calls = mock.calls.GetElixirs
	mock.lockGetElixirs.RUnlock()
	return calls
}

// GetIngredients calls GetIngredientsFunc.
func (mock *ClientMock) GetIngredients() ([]Ingredient, error) {
	if mock.GetIngredientsFunc == nil {
		panic("ClientMock.GetIngredientsFunc: method is nil but Client.GetIngredients was just called")
	}
	callInfo := struct {
	}{}
	mock.lockGetIngredients.Lock()
	mock.calls.GetIngredients = append(mock.calls.GetIngredients, callInfo)
	mock.lockGetIngredients.Unlock()
	return mock.GetIngredientsFunc()
}

// GetIngredientsCalls gets all the calls that were made to GetIngredients.
// Check the length with:
//
//	len(mockedClient.GetIngredientsCalls())
func (mock *ClientMock) GetIngredientsCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGetIngredients.RLock()
	calls = mock.calls.GetIngredients
	mock.lockGetIngredients.RUnlock()
	return calls
}