package cmd

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"github.com/openshift/service-ca-operator/tools/junitreport/pkg/api"
)

func Summarize(input io.Reader) (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var testSuites api.TestSuites
	if err := xml.NewDecoder(input).Decode(&testSuites); err != nil {
		return "", err
	}
	var summary bytes.Buffer
	var numTests, numFailed, numSkipped uint
	var duration float64
	for _, testSuite := range testSuites.Suites {
		numTests += testSuite.NumTests
		numFailed += testSuite.NumFailed
		numSkipped += testSuite.NumSkipped
		duration += testSuite.Duration
	}
	verb := "were"
	if numSkipped == 1 {
		verb = "was"
	}
	summary.WriteString(fmt.Sprintf("Of %d tests executed in %.3fs, %d succeeded, %d failed, and %d %s skipped.\n\n", numTests, duration, (numTests - numFailed - numSkipped), numFailed, numSkipped, verb))
	for _, testSuite := range testSuites.Suites {
		summarizeTests(testSuite, &summary)
	}
	return summary.String(), nil
}
func summarizeTests(testSuite *api.TestSuite, summary *bytes.Buffer) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, testCase := range testSuite.TestCases {
		if testCase.FailureOutput != nil {
			summary.WriteString(fmt.Sprintf("In suite %q, test case %q failed:\n%s\n\n", testSuite.Name, testCase.Name, testCase.FailureOutput.Output))
		}
		if testCase.SkipMessage != nil {
			summary.WriteString(fmt.Sprintf("In suite %q, test case %q was skipped:\n%s\n\n", testSuite.Name, testCase.Name, testCase.SkipMessage.Message))
		}
	}
	for _, childSuite := range testSuite.Children {
		summarizeTests(childSuite, summary)
	}
}
