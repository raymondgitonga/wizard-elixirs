package elixircreator

import "github.com/AlecAivazis/survey/v2"

type SurveyPrompt struct{}

// Prompt Abstracted the third-party library to enable mocking and testing of interactions with our code
type Prompt interface {
	AskOne(p survey.Prompt, response interface{}, opts ...survey.AskOpt) error
}

func NewSurveyPrompt() Prompt {
	return &SurveyPrompt{}
}

func (s SurveyPrompt) AskOne(p survey.Prompt, response interface{}, opts ...survey.AskOpt) error {
	return survey.AskOne(p, response, opts...)
}
