# golang-bdd


Playground to evaluate _Behaviour Driven Development_ tools for Golang:

* Cucumber BDD
* [godog](https://github.com/cucumber/godog), the official Cucumber tool
* [cucumber-html-reporter](https://www.npmjs.com/package/cucumber-html-reporter)
* (commercial) Cucumber-Studio

The current report should be available [here](test-results/cucumber_report.html)

![cucumber](https://github.com/gernotstarke/golang-bdd/actions/workflows/cucumber.yml/badge.svg)
![go_test](https://github.com/gernotstarke/golang-bdd/actions/workflows/go_test.yml/badge.svg)


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
  
