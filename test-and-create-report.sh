#!/bin/zsh

echo on


# COMMIT_ID=$(git rev-parse --verify HEAD)

godog --format cucumber:test-results/cucumber-report.json

# for cucumber-studio you need messages:
# cat report.json | ~/node_modules/.bin/json-to-messages > report.messages

node ./index.js