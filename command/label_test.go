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
	"github.com/ExpediaGroup/flyte-bamboo/event"
	"reflect"
	"testing"
)

func TestAddLabelHandler_Success(t *testing.T) {

	mockClient := MockBambooClient{}
	mockClient.addLabel = func(string, string) error { return nil }
	service := NewCommandService(mockClient)

	got := service.AddLabelHandler([]byte(`{"label": "test", "build": "123"}`))

	if got.EventDef.Name != "AddedLabel" {
		t.Errorf("expected event AddedLabel, got %s", got.EventDef)
	}

	expectedPayload := event.AddLabelSuccessPayload{Label: "test", Build: "123"}
	if !reflect.DeepEqual(got.Payload, expectedPayload) {
		t.Errorf("expected payload %+v, got %+v", expectedPayload, got.Payload)
	}
}

func TestDeleteLabelHandler_Success(t *testing.T) {

	mockClient := MockBambooClient{}
	mockClient.deleteLabel = func(string, string) error { return nil }
	service := NewCommandService(mockClient)

	got := service.DeleteLabelHandler([]byte(`{"label": "test","build":"123"}`))

	if got.EventDef.Name != "DeletedLabel" {
		t.Errorf("expected event DeletedLabel, got %s", got.EventDef)
	}

	expectedPayload := event.DeleteLabelSuccessPayload{Label: "test", Build: "123"}
	if !reflect.DeepEqual(got.Payload, expectedPayload) {
		t.Errorf("expected payload %+v, got %+v", expectedPayload, got.Payload)
	}

}

func TestAddLabelHandler_InvalidInput(t *testing.T) {

	mockClient := MockBambooClient{}
	mockClient.addLabel = func(string, string) error { return nil }
	service := NewCommandService(mockClient)

	got := service.AddLabelHandler([]byte(`{{^^}}`))

	if got.EventDef.Name != "FATAL" {
		t.Fatalf("expected event FATAL, got %s", got.EventDef.Name)
	}
}

func TestAddLabelHandler_Failure(t *testing.T) {

	mockClient := MockBambooClient{}
	mockClient.addLabel = func(string, string) error { return errors.New("some crap went up the spout") }
	service := NewCommandService(mockClient)

	got := service.AddLabelHandler([]byte(`{"label": "test","build":"123"}`))

	if got.EventDef.Name != "FailedToAddLabel" {
		t.Fatalf("expected event FATAL, got %s", got.EventDef.Name)
	}
}
