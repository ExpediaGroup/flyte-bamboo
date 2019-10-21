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
	"github.com/HotelsDotCom/flyte-bamboo/event"
	"reflect"
	"testing"
)

func TestTriggerPlanHandler_Success(t *testing.T) {

	mockClient := MockBambooClient{}
	mockClient.triggerPlan = func(string) error { return nil }
	service := NewCommandService(mockClient)

	got := service.TriggerPlanHandler([]byte(`{"plan":"ABC-123"}`))

	if got.EventDef.Name != "PlanTriggered" {
		t.Fatalf("expected event PlanTriggered, got %s", got.EventDef.Name)
	}

	expectedPayload := event.TriggerPlanSuccessPayload{Plan: "ABC-123"}

	if !reflect.DeepEqual(got.Payload, expectedPayload) {
		t.Fatalf("expected payload %+v, got %+v", expectedPayload, got.Payload)
	}
}

func TestTriggerPlanHandler_InvalidInput(t *testing.T) {

	mockClient := MockBambooClient{}
	mockClient.triggerPlan = func(string) error { return nil }
	service := NewCommandService(mockClient)

	got := service.TriggerPlanHandler([]byte(`{""}`))

	if got.EventDef.Name != "FATAL" {
		t.Fatalf("expected event FATAL, got %s", got.EventDef.Name)
	}
}

func TestTriggerPlanHandler_Failure(t *testing.T) {

	mockClient := MockBambooClient{}
	mockClient.triggerPlan = func(string) error { return errors.New("some crap went up the spout") }
	service := NewCommandService(mockClient)

	got := service.TriggerPlanHandler([]byte(`{"plan":"ABC-123"}`))

	if got.EventDef.Name != "FailedToTriggerPlan" {
		t.Fatalf("expected event FailedToTriggerPlan, got %s", got.EventDef.Name)
	}
}
