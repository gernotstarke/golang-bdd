# golang-bdd


Playground to evaluate _Behaviour Driven Development_ tools for Golang:

* Cucumber BDD
* [Godog](https://github.com/cucumber/godog), the official Cucumber tool
* [cucumber-html-reporter](https://www.npmjs.com/package/cucumber-html-reporter)
* (commercial) Cucumber-Studio

The current report should be available [here](test-results/cucumber_report.html)

![feature-linter](https://github.com/gernotstarke/golang-bdd/actions/workflows/feature-linter.yml/badge.svg)
![cucumber](https://github.com/gernotstarke/golang-bdd/actions/workflows/cucumber.yml/badge.svg)
![go_test](https://github.com/gernotstarke/golang-bdd/actions/workflows/go_test.yml/badge.svg)

[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![Go Report Card](https://goreportcard.com/badge/github.com/gernotstarke/golang-bdd)](https://goreportcard.com/report/github.com/gernotstarke/golang-bdd)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=gernotstarke_golang-bdd&metric=alert_status)](https://sonarcloud.io/dashboard?id=gernotstarke_golang-bdd)
[![Maintainability](https://api.codeclimate.com/v1/badges/c481ef8142826f71ff65/maintainability)](https://codeclimate.com/github/gernotstarke/golang-bdd/maintainability)

#
## Installation

### Godog
 
````shell
go get github.com/cucumber/godog/cmd/godog@v0.11.0
````
### Cucumber HTML Reporter

It's written in JavaScript and requires `npm` and `node` to be available on your machine.

```shell
npm install cucumber-html-reporter --save-dev
```

## Usage

I squeezed the required commands into the file `test-and-create-report.sh`
(and will pack them into github actions...)


### Godog

```shell
godog --format cucumber:test-results/cucumber-report.json
```

Notes: 

* godog requires features and scenarios to be written in a `features` directory.
* the `--format` switch can take a file and/or directory name


### Cucumber Report

```shell
node index.js
```

Notes:

* cucumber-html-reporter can handle "Metadata" (see `index.js`), for really interesting data
you should provide it with command line parameters, like commit-id
  
