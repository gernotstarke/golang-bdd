var reporter = require('cucumber-html-reporter');

var options = {
    theme: 'bootstrap',
    jsonFile: 'test-results/cucumber-report.json',
    output: 'test-results/cucumber_report.html',
    reportSuiteAsScenarios: true,
    scenarioTimestamp: true,
    launchReport: true,
    metadata: {
        "App Version":"0.3.2",
        "Test Environment": "STAGING",
        "Browser": "Firefox  88.0 (64-Bit)",
        "Platform": "OSX",
        "Parallel": "Scenarios",
        "Executed": "Remote"
    }
};

reporter.generate(options);

