/*
 * Copyright 2018-2019 the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package buildplan_test

import (
	"testing"

	"github.com/buildpack/libbuildpack/buildplan"
	"github.com/buildpack/libbuildpack/internal"
	"github.com/onsi/gomega"
	"github.com/sclevine/spec"
	"github.com/sclevine/spec/report"
)

func TestBuildPlan(t *testing.T) {
	spec.Run(t, "BuildPlan", func(t *testing.T, _ spec.G, it spec.S) {

		g := gomega.NewWithT(t)

		it("unmarshals from os.Stdin", func() {
			console, d := internal.ReplaceConsole(t)
			defer d()

			console.In(t, `[alpha]
  version = "alpha-version"

  [alpha.metadata]
    test-key = "test-value"

[bravo]
`)

			buildPlan := buildplan.BuildPlan{}
			g.Expect(buildPlan.Init()).To(gomega.Succeed())

			g.Expect(buildPlan).To(gomega.Equal(buildplan.BuildPlan{
				"alpha": buildplan.Dependency{
					Version:  "alpha-version",
					Metadata: buildplan.Metadata{"test-key": "test-value"},
				},
				"bravo": buildplan.Dependency{},
			}))
		})

		it("performs a shallow merge", func() {
			b := buildplan.BuildPlan{"test-key-1": buildplan.Dependency{}}

			b.Merge(
				buildplan.BuildPlan{"test-key-2": buildplan.Dependency{}},
				buildplan.BuildPlan{"test-key-3": buildplan.Dependency{}},
			)

			g.Expect(b).To(gomega.Equal(buildplan.BuildPlan{
				"test-key-1": buildplan.Dependency{},
				"test-key-2": buildplan.Dependency{},
				"test-key-3": buildplan.Dependency{},
			}))
		})
	}, spec.Report(report.Terminal{}))
}
