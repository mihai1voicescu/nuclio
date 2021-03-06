// +build test_integration
// +build test_local

/*
Copyright 2017 The Nuclio Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package test

import (
	"testing"

	"github.com/nuclio/nuclio/pkg/processor/build/runtime/test/suite"

	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	buildsuite.TestSuite
	runtime string
}

func (suite *TestSuite) SetupSuite() {
	suite.TestSuite.SetupSuite()

	suite.TestSuite.RuntimeSuite = suite
	suite.TestSuite.ArchivePattern = "python"
	suite.Runtime = suite.runtime
}

func (suite *TestSuite) GetFunctionInfo(functionName string) buildsuite.FunctionInfo {
	functionInfo := buildsuite.FunctionInfo{
		Runtime: suite.runtime,
	}

	switch functionName {

	case "reverser":
		functionInfo.Path = []string{suite.GetTestFunctionsDir(), "common", "reverser", "python", "reverser.py"}
		functionInfo.Handler = "reverser:handler"

	case "json-parser-with-function-config":
		functionInfo.Path = []string{suite.GetTestFunctionsDir(), "common", "json-parser-with-function-config", "python"}

	case "json-parser-with-inline-function-config":
		functionInfo.Path = []string{suite.GetTestFunctionsDir(), "common", "json-parser-with-inline-function-config", "python", "parser.py"}

	case "invalid-inline-config":
		functionInfo.Path = []string{suite.GetTestFunctionsDir(), "common", "invalid-inline-config", "python", "parser.py"}

	case "long-initialization":
		functionInfo.Path = []string{suite.GetTestFunctionsDir(), "common", "long-initialization", "python", "sleepy.py"}

	case "context-init-fail":
		functionInfo.Path = []string{suite.GetTestFunctionsDir(), "common", "context-init-fail", "python", "contextinitfail.py"}

	default:
		suite.Logger.InfoWith("Test skipped", "functionName", functionName)

		functionInfo.Skip = true
	}

	return functionInfo
}

func TestIntegrationSuite(t *testing.T) {
	if testing.Short() {
		return
	}

	for _, testCase := range []struct {
		runtimeName string
	}{
		{runtimeName: "python:3.6"},
		{runtimeName: "python:3.7"},
		{runtimeName: "python:3.8"},
	} {
		t.Run(testCase.runtimeName, func(t *testing.T) {
			testSuite := new(TestSuite)
			testSuite.runtime = testCase.runtimeName
			suite.Run(t, testSuite)
		})
	}
}
