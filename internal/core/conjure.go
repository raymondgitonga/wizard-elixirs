package core

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/raymondgitonga/wizard-elixirs/internal/wizardclient"
)

type NoIngredientsError struct{}

type Conjure struct {
	prompt Prompt
	client wizardclient.Client
}

func (n NoIngredientsError) Error() string {
	return "you must add at least one ingredient"
}

func NewConjure(prompt Prompt, client wizardclient.Client) *Conjure {
	return &Conjure{
		prompt: prompt,
		client: client,
	}
}

func (e *Conjure) AskUserToStartMakingElixirs() (bool, error) {
	var makeElixirs bool
	prompt := &survey.Confirm{Message: "Welcome to the wizarding world, ready to make elixirs?"}
	err := e.prompt.AskOne(prompt, &makeElixirs)
	return makeElixirs, err
}

func (e *Conjure) AskUserToSelectIngredients() ([]string, error) {
	ingredients, err := e.client.GetIngredients()
	if err != nil {
		return nil, err
	}

	ingredientOptions := e.createIngredientOptions(ingredients)
	prompt := &survey.MultiSelect{
		Message: "Choose your ingredients:",
		Options: ingredientOptions,
	}

	var chosenIngredients []string
	err = e.prompt.AskOne(prompt, &chosenIngredients)
	if err != nil {
		return nil, err
	}

	if len(chosenIngredients) < 1 {
		return nil, NoIngredientsError{}
	}
	return chosenIngredients, err
}

func (e *Conjure) CreateElixirsFromIngredients(chosenIngredients []string) ([]wizardclient.Elixir, error) {
	elixirs, err := e.client.GetElixirs(chosenIngredients)
	if err != nil {
		return nil, err
	}
	return elixirs, nil
}

func (e *Conjure) createIngredientOptions(ingredients []wizardclient.Ingredient) []string {
	var ingredientOptions []string
	for i := range ingredients {
		ingredientOptions = append(ingredientOptions, ingredients[i].Name)
	}
	return ingredientOptions
}
