Feature: Syntactic Checks on Feature files
  In order to develop the system
  As a developer
  I want regular syntax checks of our feature files

  Rules:
  - See https://www.npmjs.com/package/gherkin-lint#configuration-file

  Scenario: Creating an regular given-when-then scenario
    Given there is a syntactically correct scenario called "developer"
    When I check the syntax of this scenario
    Then I should have 0 errors
