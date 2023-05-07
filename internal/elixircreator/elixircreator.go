package elixircreator

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/raymondgitonga/wizard-elixirs/internal/wizardclient"
)

type NoIngredientsError struct{}

type ElixirCreator struct {
	prompt Prompt
	client wizardclient.Client
}

func (n NoIngredientsError) Error() string {
	return "you must add at least one ingredient"
}

func NewElixirCreator(prompt Prompt, client wizardclient.Client) *ElixirCreator {
	return &ElixirCreator{
		prompt: prompt,
		client: client,
	}
}

func (e *ElixirCreator) AskUserToStartMakingElixirs() (bool, error) {
	var makeElixirs bool
	prompt := &survey.Confirm{Message: "Welcome to the wizarding world, ready to make elixirs?"}
	err := e.prompt.AskOne(prompt, &makeElixirs)
	return makeElixirs, err
}

func (e *ElixirCreator) AskUserToSelectIngredients() ([]string, error) {
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

func (e *ElixirCreator) CreateElixirsFromIngredients(chosenIngredients []string) ([]wizardclient.Elixir, error) {
	elixirs, err := e.client.GetElixirs(chosenIngredients)
	if err != nil {
		return nil, err
	}
	return elixirs, nil
}

func (e *ElixirCreator) createIngredientOptions(ingredients []wizardclient.Ingredient) []string {
	var ingredientOptions []string
	for i := range ingredients {
		ingredientOptions = append(ingredientOptions, ingredients[i].Name)
	}
	return ingredientOptions
}
