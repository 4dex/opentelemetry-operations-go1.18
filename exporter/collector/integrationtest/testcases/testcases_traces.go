// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package testcases

import "github.com/4dex/opentelemetry-operations-go/exporter/collector"

var TracesTestCases = []TestCase{
	{
		Name:                 "Basic traces",
		OTLPInputFixturePath: "testdata/fixtures/traces/traces_basic.json",
		ExpectFixturePath:    "testdata/fixtures/traces/traces_basic_expected.json",
	},
	{
		Name:                 "Custom User Agent",
		OTLPInputFixturePath: "testdata/fixtures/traces/traces_basic.json",
		ExpectFixturePath:    "testdata/fixtures/traces/traces_user_agent_expected.json",
		ConfigureCollector: func(cfg *collector.Config) {
			cfg.UserAgent = "custom-user-agent {{version}}"
		},
	},
}
