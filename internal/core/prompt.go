package core

import "github.com/AlecAivazis/survey/v2"

type SurveyPrompt struct{}

type Prompt interface {
	AskOne(p survey.Prompt, response interface{}, opts ...survey.AskOpt) error
}

func NewSurveyPrompt() Prompt {
	return &SurveyPrompt{}
}

func (s SurveyPrompt) AskOne(p survey.Prompt, response interface{}, opts ...survey.AskOpt) error {
	return survey.AskOne(p, response, opts...)
}
