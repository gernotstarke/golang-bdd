package math_bdd

import (
	"fmt"
	"github.com/cucumber/godog"
)

func iAdd(param2 int) error {
	Add(param2)
	return nil
}

func iEndUpWith(expectedResult int) error {
	if Result != expectedResult {
		return fmt.Errorf("expected %d but actual %d", expectedResult, Result)
	}
	return nil
}

func iStartWith(param1 int) error {
	Param1 = param1
	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^I add (\d+)$`, iAdd)
	ctx.Step(`^I end up with (\d+)$`, iEndUpWith)
	ctx.Step(`^I start with (\d+)$`, iStartWith)
}

