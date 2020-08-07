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
	"errors"
	"github.com/ExpediaGroup/flyte-bamboo/bamboo"
	"github.com/ExpediaGroup/flyte-bamboo/event"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetLifeCycle_Success(t *testing.T) {

	mockClient := MockBambooClient{}
	mockClient.getStage = func(build, stage string) (bamboo.Stage, error) {
		return bamboo.Stage{
			Name:           "prod-verify",
			LifeCycleState: "NotBuilt",
			State:          "Unknown"}, nil
	}

	service := NewCommandService(mockClient)
	got := service.GetLifecycleStateHandler([]byte(`{"build":"ABC-DEF-123","stage":"prod-verify"}`))

	assert.Equal(t, got.EventDef.Name, "GetLifeCycleState")
	expectedPayload := event.GetLifeCycleSuccessPayload{Build: "ABC-DEF-123", Stage: "prod-verify", State: "Unknown", LifeCycle: "NotBuilt"}
	assert.Equal(t, got.Payload, expectedPayload)

}

func TestGetLifeCycle_InvalidInput(t *testing.T) {
	mockClient := MockBambooClient{}
	mockClient.getStage = func(build, stage string) (bamboo.Stage, error) { return bamboo.Stage{}, nil }
	service := NewCommandService(mockClient)
	got := service.GetLifecycleStateHandler([]byte(`{"build}`))
	assert.Equal(t, got.EventDef.Name, "FATAL")
}

func TestGetLifeCycle_Failure(t *testing.T) {
	mockClient := MockBambooClient{}
	mockClient.getStage = func(build, stage string) (bamboo.Stage, error) {
		return bamboo.Stage{}, errors.New("could not get lifeCycleState")
	}
	service := NewCommandService(mockClient)
	got := service.GetLifecycleStateHandler([]byte(`{"build":"ABC-DEF"}`))
	assert.Equal(t, got.EventDef.Name, "GetLifeCycleFailed")
}
