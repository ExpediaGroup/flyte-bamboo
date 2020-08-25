/*
Copyright (C) 2018 Expedia Group.

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

package command

import (
	"github.com/ExpediaGroup/flyte-bamboo/bamboo"
)

type MockBambooClient struct {
	addLabel    func(build, label string) error
	deleteLabel func(build, label string) error
	startStage  func(build, stage string, vars map[string]string, retry bamboo.Retry) error
	triggerPlan func(plan string) error
	addComment  func(comment, build string) error
	enablePlan  func(plan string) error
	getStage    func(build, stage string) (bamboo.Stage, error)
}

func (m MockBambooClient) AddLabel(build, label string) error {
	return m.addLabel(build, label)
}

func (m MockBambooClient) DeleteLabel(build, label string) error {
	return m.deleteLabel(build, label)
}

func (m MockBambooClient) StartStage(build, stage string, vars map[string]string, retry bamboo.Retry) error {
	return m.startStage(build, stage, vars, retry)
}

func (m MockBambooClient) TriggerPlan(plan string) error {
	return m.triggerPlan(plan)
}

func (m MockBambooClient) AddComment(comment, build string) error {
	return m.addComment(build, comment)
}

func (m MockBambooClient) EnablePlan(plan string) error {
	return m.enablePlan(plan)
}

func (m MockBambooClient) GetStage(build, stage string) (bamboo.Stage, error) {
	return m.getStage(build, stage)
}
