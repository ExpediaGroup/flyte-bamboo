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
	"reflect"
	"github.com/HotelsDotCom/flyte-bamboo/event"
	"testing"
)

func TestEnablePlanHandler_Success(t *testing.T) {

	mockClient := MockBambooClient{}
	mockClient.enablePlan = func(string) error { return nil }
	service := NewCommandService(mockClient)

	got := service.EnablePlanHandler([]byte(`{"plan":"ABC-DEF"}`))

	if got.EventDef.Name != "PlanEnabled" {
		t.Errorf("expected event PlanEnabled, got %s", got.EventDef)
	}

	expectedPayload := event.EnablePlanSuccessPayload{Plan: "ABC-DEF"}
	if !reflect.DeepEqual(got.Payload, expectedPayload) {
		t.Errorf("expected payload %+v, got %+v", expectedPayload, got.Payload)
	}
}

func TestEnablePlanHandler_InvalidInput(t *testing.T) {

	mockClient := MockBambooClient{}
	mockClient.enablePlan = func(string) error { return nil }
	service := NewCommandService(mockClient)

	got := service.EnablePlanHandler([]byte(`{"plan"}`))

	if got.EventDef.Name != "FATAL" {
		t.Errorf("expected event EnablePlanSuccess, got %s", got.EventDef)
	}

}

func TestEnablePlanHandler_Failure(t *testing.T) {

	mockClient := MockBambooClient{}
	mockClient.enablePlan = func(string) error { return errors.New("Couldn't add a stupid comment") }
	service := NewCommandService(mockClient)

	got := service.EnablePlanHandler([]byte(`{"plan":"ABC-DEF"}`))

	if got.EventDef.Name != "FailedToEnablePlan" {
		t.Errorf("expected event FailedToEnablePlan, got %s", got.EventDef)
	}

}
