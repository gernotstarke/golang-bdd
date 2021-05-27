package main

import (
	"flag"
	"fmt"
	"github.com/cucumber/godog"
	"godog-bdd/math"
	"os"
	"testing"
)

func iAdd(param2 int) error {
	math.Add(param2)
	return nil
}

func iEndUpWith(expectedResult int) error {
	if math.Result != expectedResult {
		return fmt.Errorf("expected %d but actual %d", expectedResult, math.Result)
	}
	return nil
}

func iStartWith(param1 int) error {
	math.Param1 = param1
	return nil
}

func InitializeMathScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^I add (\d+)$`, iAdd)
	ctx.Step(`^I end up with (\d+)$`, iEndUpWith)
	ctx.Step(`^I start with (\d+)$`, iStartWith)
}

func InitializeTestSuite(ctx *godog.TestSuiteContext) {

}

func TestMain(m *testing.M) {
	flag.Parse()


	status := godog.TestSuite{
		Name: "math",
		TestSuiteInitializer: InitializeTestSuite,
		ScenarioInitializer:  InitializeMathScenario,
	}.Run()

	// Optional: Run `testing` package's logic besides godog.
	if st := m.Run(); st > status {
		status = st
	}

	os.Exit(status)
}
