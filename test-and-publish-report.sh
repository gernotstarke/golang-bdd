#!/bin/bash

echo on


COMMIT_ID=$(git rev-parse --verify HEAD)

REPORT_PATH="test-results/cucumber-report"

#godog --format cucumber:${REPORT_PATH}.json

# for cucumber-studio we need to convert json to messages:
# https://www.npmjs.com/package/@cucumber/json-to-messages
#cat ${REPORT_PATH}.json | ~/node_modules/.bin/json-to-messages > ${REPORT_PATH}.messages


# publish report to cucumber-studio with curl:
curl -X POST https://studio.cucumber.io/cucumber_project/results \
     -F messages=${REPORT_PATH}.messages \
      -H "project-access-token: 27877807453784791111842877191917738921804894910830921354" \
      -H "provider: github" \
      -H "repo: gernotstarke/golang-bdd" \
      -H "branch: main" \
      -H "revision: ${COMMIT_ID}"


# now convert to cucumber-report.html
#node ./index.js
