package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/raymondgitonga/wizard-elixirs/internal/elixircreator"
	"github.com/raymondgitonga/wizard-elixirs/internal/wizardclient"
	"github.com/spf13/cobra"
)

type Root struct {
	surveyPrompt elixircreator.Prompt
	wizardClient wizardclient.Client
}

func NewRoot(surveyPrompt elixircreator.Prompt, wizardClient wizardclient.Client) *Root {
	return &Root{
		surveyPrompt: surveyPrompt,
		wizardClient: wizardClient,
	}
}

func (r *Root) RunElixirCommand() error {
	command := &cobra.Command{
		Use:   "elixir",
		Short: "Command to make elixir",
		RunE: func(cmd *cobra.Command, args []string) error {
			err := r.createElixirs()
			if errors.Is(err, elixircreator.NoIngredientsError{}) {
				fmt.Println("Add at least one ingredient, lets try this again")
				return r.createElixirs()
			}
			if err != nil {
				return fmt.Errorf("error running elixir command: %w", err)
			}
			return nil
		},
	}
	return command.Execute()
}

func (r *Root) createElixirs() error {
	elixirCreator := elixircreator.NewElixirCreator(r.surveyPrompt, r.wizardClient)

	makeElixirs, err := elixirCreator.AskUserToStartMakingElixirs()
	if err != nil {
		return err
	}
	if !makeElixirs {
		fmt.Println("Ok, we can try again next time. Bye!")
		return nil
	}

	chosenIngredients, err := elixirCreator.AskUserToSelectIngredients()
	if err != nil {
		return err
	}

	elixirs, err := elixirCreator.CreateElixirsFromIngredients(chosenIngredients)
	if err != nil {
		return err
	}

	if len(elixirs) < 1 {
		fmt.Println("Oops! No elixirs with that combination")
		return nil
	}

	formattedData, err := json.MarshalIndent(elixirs, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(formattedData))
	return nil
}
