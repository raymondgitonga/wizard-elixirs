package core

import (
	"errors"
	"github.com/AlecAivazis/survey/v2"
	"github.com/raymondgitonga/wizard-elixirs/internal/wizardclient"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConjure_AskUserToStartMakingElixirs(t *testing.T) {
	tests := []struct {
		name      string
		prompt    Prompt
		wantError bool
	}{
		{
			name: "Success_UserStartsMakingElixirs",
			prompt: &PromptMock{AskOneFunc: func(p survey.Prompt, response interface{}, opts ...survey.AskOpt) error {
				makeElixirs := response.(*bool)
				*makeElixirs = true
				return nil
			}},
			wantError: false,
		},
		{
			name: "Fail_PromptError",
			prompt: &PromptMock{AskOneFunc: func(p survey.Prompt, response interface{}, opts ...survey.AskOpt) error {
				return errors.New("error making elixir")
			}},
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conjure := NewConjure(tt.prompt, &wizardclient.ClientMock{})

			makeElixirs, err := conjure.AskUserToStartMakingElixirs()

			if tt.wantError {
				assert.Error(t, err)
				assert.Equal(t, errors.New("error making elixir"), err)
				return
			}
			assert.NoError(t, err)
			assert.True(t, makeElixirs)
		})
	}
}

func TestConjure_AskUserToSelectIngredients(t *testing.T) {
	tests := []struct {
		name                  string
		prompt                Prompt
		wizardClient          wizardclient.Client
		wantPromptError       bool
		wantWizardClientError bool
	}{
		{
			name: "Success_SelectingIngredients",
			prompt: &PromptMock{AskOneFunc: func(p survey.Prompt, response interface{}, opts ...survey.AskOpt) error {
				chosenIngredients := response.(*[]string)
				*chosenIngredients = []string{"Newt spleens"}
				return nil
			}},
			wizardClient: &wizardclient.ClientMock{GetIngredientsFunc: func() ([]wizardclient.Ingredient, error) {
				return []wizardclient.Ingredient{
					{
						Name: "Newt spleens",
					},
				}, nil
			}},
			wantPromptError:       false,
			wantWizardClientError: false,
		},
		{
			name: "Fail_NoIngredientSelected",
			prompt: &PromptMock{AskOneFunc: func(p survey.Prompt, response interface{}, opts ...survey.AskOpt) error {
				return nil
			}},
			wizardClient: &wizardclient.ClientMock{GetIngredientsFunc: func() ([]wizardclient.Ingredient, error) {
				return []wizardclient.Ingredient{
					{
						Name: "Newt spleens",
					},
				}, nil
			}},
			wantPromptError:       true,
			wantWizardClientError: false,
		},
		{
			name: "Fail_PromptError",
			prompt: &PromptMock{AskOneFunc: func(p survey.Prompt, response interface{}, opts ...survey.AskOpt) error {
				return errors.New("something went wrong")
			}},
			wizardClient: &wizardclient.ClientMock{GetIngredientsFunc: func() ([]wizardclient.Ingredient, error) {
				return []wizardclient.Ingredient{
					{
						Name: "Newt spleens",
					},
				}, nil
			}},
			wantPromptError:       true,
			wantWizardClientError: false,
		},
		{
			name: "Fail_WizardClientError",
			prompt: &PromptMock{AskOneFunc: func(p survey.Prompt, response interface{}, opts ...survey.AskOpt) error {
				return nil
			}},
			wizardClient: &wizardclient.ClientMock{GetIngredientsFunc: func() ([]wizardclient.Ingredient, error) {
				return nil, errors.New("something went wrong")
			}},
			wantPromptError:       false,
			wantWizardClientError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conjure := NewConjure(tt.prompt, tt.wizardClient)
			ingredients, err := conjure.AskUserToSelectIngredients()

			if tt.wantPromptError {
				assert.Error(t, err)
				return
			}

			if tt.wantWizardClientError {
				assert.Error(t, err)
				assert.Equal(t, errors.New("something went wrong"), err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, 1, len(ingredients))
			assert.Equal(t, "Newt spleens", ingredients[0])
		})
	}
}

func TestConjure_MakeElixir(t *testing.T) {
	tests := []struct {
		name         string
		wizardClient wizardclient.Client
		wantError    bool
	}{
		{
			name: "Success_ElixirCreated",
			wizardClient: &wizardclient.ClientMock{GetElixirsFunc: func(ingredients []string) ([]wizardclient.Elixir, error) {
				return []wizardclient.Elixir{
					{
						Name:            "Ageing Potion",
						Effect:          "Ages drinker temporarily",
						SideEffects:     "",
						Characteristics: "Green",
						Difficulty:      "Advanced",
					},
				}, nil
			}},
			wantError: false,
		},
		{
			name: "Fail_WizardClientError",
			wizardClient: &wizardclient.ClientMock{GetElixirsFunc: func(ingredients []string) ([]wizardclient.Elixir, error) {
				return nil, errors.New("failed to get elixirs")
			}},
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conjure := NewConjure(&PromptMock{}, tt.wizardClient)
			elixir, err := conjure.CreateElixirsFromIngredients([]string{"Newt spleens"})

			if tt.wantError {
				assert.Error(t, err)
				assert.Equal(t, errors.New("failed to get elixirs"), err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, 1, len(elixir))
			assert.Equal(t, "Ageing Potion", elixir[0].Name)
		})
	}
}
