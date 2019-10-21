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
	"github.com/HotelsDotCom/flyte-bamboo/bamboo"
	"github.com/HotelsDotCom/flyte-bamboo/event"
	"reflect"
	"testing"
)

func TestStartStageHandler_Success(t *testing.T) {

	mockClient := MockBambooClient{}
	mockClient.startStage = func(string, string, map[string]string, bamboo.Retry) error { return nil }
	service := NewCommandService(mockClient)

	got := service.StartStageHandler([]byte(`{"build":"ABC-123","delay":2,"maxRetry":3,"stage":"initiate-staging","vars":{"deployStatus":"success"}}`))

	if got.EventDef.Name != "StartStageSuccess" {
		t.Fatalf("expected event NextStageInitiated, got %s", got.EventDef)
	}

	expectedPayload := event.StartStageSuccessPayload{Build: "ABC-123", Stage: "initiate-staging", Vars: map[string]string{"deployStatus": "success"}}

	if !reflect.DeepEqual(got.Payload, expectedPayload) {
		t.Fatalf("expected payload %+v, got %+v", expectedPayload, got.Payload)
	}
}

func TestStartStageHandler_InvalidInput(t *testing.T) {

	mockClient := MockBambooClient{}
	mockClient.startStage = func(string, string, map[string]string, bamboo.Retry) error { return nil }
	service := NewCommandService(mockClient)

	got := service.StartStageHandler([]byte(`{"build"}`))

	if got.EventDef.Name != "FATAL" {
		t.Fatalf("expected event FATAL, got %s", got.EventDef.Name)
	}
}

func TestStartStageHandler_Failure(t *testing.T) {

	mockClient := MockBambooClient{}
	mockClient.startStage = func(string, string, map[string]string, bamboo.Retry) error {
		return errors.New("some crap went up the spout")
	}

	service := NewCommandService(mockClient)

	got := service.StartStageHandler([]byte(`{"build":"ABC-123","stage":"initiate-staging","vars":{"deployStatus":"success"}}`))

	if got.EventDef.Name != "FailedToStartStage" {
		t.Fatalf("expected event FailedToStartStage, got %s", got.EventDef.Name)
	}
}
