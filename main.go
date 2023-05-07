package main

import (
	"github.com/raymondgitonga/wizard-elixirs/cmd"
	"github.com/raymondgitonga/wizard-elixirs/internal/elixircreator"
	"github.com/raymondgitonga/wizard-elixirs/internal/wizardclient"
	"log"
	"net/http"
)

const WizardWorldURL = "https://wizard-world-api.herokuapp.com"

func main() {
	serverPrompt := elixircreator.NewSurveyPrompt()
	wizardClient, err := wizardclient.NewWizardClient(WizardWorldURL, &http.Client{})
	if err != nil {
		log.Fatal(err)
	}

	root := cmd.NewRoot(serverPrompt, wizardClient)

	err = root.RunElixirCommand()
	if err != nil {
		log.Fatal(err)
	}
}
