package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/raymondgitonga/wizard-elixirs/internal/core"
	"github.com/raymondgitonga/wizard-elixirs/internal/wizardclient"
	"github.com/spf13/cobra"
)

type Root struct {
	surveyPrompt core.Prompt
	wizardClient wizardclient.Client
}

func NewRoot(surveyPrompt core.Prompt, wizardClient wizardclient.Client) *Root {
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
			err := r.conjureElixirs()
			if errors.Is(err, core.NoIngredientsError{}) {
				fmt.Println("Add at least one ingredient, lets try this again")
				return r.conjureElixirs()
			}
			if err != nil {
				return fmt.Errorf("error running elixir command: %w", err)
			}
			return nil
		},
	}
	return command.Execute()
}

func (r *Root) conjureElixirs() error {
	conjure := core.NewConjure(r.surveyPrompt, r.wizardClient)
	makeElixirs, err := conjure.AskUserToStartMakingElixirs()
	if err != nil {
		return err
	}
	if !makeElixirs {
		fmt.Println("Ok, we can try again next time. Bye!")
		return nil
	}

	chosenIngredients, err := conjure.AskUserToSelectIngredients()
	if err != nil {
		return err
	}

	elixirs, err := conjure.CreateElixirsFromIngredients(chosenIngredients)
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
