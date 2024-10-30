package main

import "github.com/cucumber/godog"

type apiFeature struct{}

func (api *apiFeature) iTranslateItTo(arg1 string) error {
	return godog.ErrPending
}

func (api *apiFeature) theResponseShouldBe(arg1 string) error {
	return godog.ErrPending
}

func (api *apiFeature) theWord(arg1 string) error {
	return godog.ErrPending
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	api := &apiFeature{}

	ctx.Step(`^I translate it to "([^"]*)"$`, api.iTranslateItTo)
	ctx.Step(`^the response should be "([^"]*)"$`, api.theResponseShouldBe)
	ctx.Step(`^the word "([^"]*)"$`, api.theWord)
}
