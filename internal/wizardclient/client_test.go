package wizardclient

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWizardWorldClient_GetIngredients(t *testing.T) {
	testResponse := []byte(`[
        {
            "id": "00eee42e-999c-43ef-bb8d-3f1012275680",
            "name": "Newt spleens"
        }
    ]`)
	tests := []struct {
		name         string
		expectedResp []byte
		expectedErr  error
		httpsStatus  int
		serverError  bool
		parsingError bool
	}{
		{
			name:         "successful call to wizard world api",
			expectedResp: testResponse,
			httpsStatus:  http.StatusOK,
		},
		{
			name:        "failed call to wizard world api",
			serverError: true,
			httpsStatus: http.StatusBadRequest,
		},
		{
			name:         "json decoding error when parsing response",
			expectedResp: []byte(``),
			parsingError: true,
			httpsStatus:  http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wizardServer := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
				rw.WriteHeader(tt.httpsStatus)
				_, err := rw.Write(tt.expectedResp)
				assert.NoError(t, err)
				assert.Equal(t, "/ingredients", req.URL.Path)
			}))
			defer wizardServer.Close()

			wizardClient, err := NewWizardClient(wizardServer.URL, &http.Client{})
			assert.NoError(t, err)

			ingredients, err := wizardClient.GetIngredients()

			if tt.serverError {
				assert.Equal(t, errors.New("error making ingredients call: HTTP-Status 400 Bad Request"), err)
				return
			}

			if tt.parsingError {
				assert.Equal(t, "error decoding ingredients response: unexpected end of JSON input", err.Error())
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, 1, len(ingredients))
			assert.Equal(t, "Newt spleens", ingredients[0].Name)
		})
	}
}

func TestWizardWorldClient_GetElixirs(t *testing.T) {
	testResponse := []byte(`[
    	{
        	"name": "Ageing Potion",
        	"effect": "Ages drinker temporarily",
        	"sideEffects": null,
        	"characteristics": "Green",
        	"difficulty": "Advanced"
    	}
	]`)

	tests := []struct {
		name         string
		expectedResp []byte
		expectedErr  error
		httpsStatus  int
		serverError  bool
		parsingError bool
	}{
		{
			name:         "successful call to wizard world api",
			expectedResp: testResponse,
			httpsStatus:  http.StatusOK,
		},
		{
			name:         "failed call to wizard world api",
			expectedResp: testResponse,
			serverError:  true,
			httpsStatus:  http.StatusInternalServerError,
		},
		{
			name:         "json decoding error when parsing response",
			expectedResp: []byte(``),
			parsingError: true,
			httpsStatus:  http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wizardServer := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
				assert.Equal(t, "/Elixirs", req.URL.Path)

				queryParams := req.URL.Query()
				assert.Contains(t, queryParams, "Ingredient")
				assert.Contains(t, queryParams["Ingredient"], "Newt spleens")
				assert.Contains(t, queryParams["Ingredient"], "Stewed Mandrake")

				rw.WriteHeader(tt.httpsStatus)
				_, err := rw.Write(tt.expectedResp)
				assert.NoError(t, err)
			}))
			defer wizardServer.Close()

			wizardClient, err := NewWizardClient(wizardServer.URL, &http.Client{})
			assert.NoError(t, err)

			elixirs, err := wizardClient.GetElixirs(
				[]string{"Newt spleens", "Stewed Mandrake"},
			)

			if tt.serverError {
				assert.Equal(t, errors.New("error making elixirs call: HTTP-Status 500 Internal Server Error"), err)
				return
			}

			if tt.parsingError {
				assert.Equal(t, "error decoding elixirs response: unexpected end of JSON input", err.Error())
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, 1, len(elixirs))
			assert.Equal(t, "Ageing Potion", elixirs[0].Name)
			assert.Equal(t, "Advanced", elixirs[0].Difficulty)
			assert.Equal(t, "", elixirs[0].SideEffects)
		})
	}
}
