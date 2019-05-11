package oscmd

import (
	"github.com/openshift/service-ca-operator/tools/junitreport/pkg/builder"
	"github.com/openshift/service-ca-operator/tools/junitreport/pkg/parser"
	"github.com/openshift/service-ca-operator/tools/junitreport/pkg/parser/stack"
)

func NewParser(builder builder.TestSuitesBuilder, stream bool) parser.TestOutputParser {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return stack.NewParser(builder, newTestDataParser(), newTestSuiteDataParser(), stream)
}
