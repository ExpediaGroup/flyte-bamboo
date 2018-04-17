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

func TestAddCommentHandler_Success(t *testing.T) {

	mockClient := MockBambooClient{}
	mockClient.addComment = func(string, string) error { return nil }
	service := NewCommandService(mockClient)

	got := service.AddCommentHandler([]byte(`{"comment": "test", "build": "123"}`))

	if got.EventDef.Name != "CommentAdded" {
		t.Errorf("expected event CommentAdded, got %s", got.EventDef.Name)
	}

	expectedPayload := event.AddCommentSuccessPayload{Comment: "test", Build: "123"}
	if !reflect.DeepEqual(got.Payload, expectedPayload) {
		t.Errorf("expected payload %+v, got %+v", expectedPayload, got.Payload)
	}
}

func TestAddCommentHandler_InvalidInput(t *testing.T) {

	mockClient := MockBambooClient{}
	mockClient.addComment = func(string, string) error { return nil }
	service := NewCommandService(mockClient)

	got := service.AddCommentHandler([]byte(`{""}`))

	if got.EventDef.Name != "FATAL" {
		t.Fatalf("expected event FATAL, got %s", got.EventDef.Name)
	}

}

func TestAddCommentHandler_Failure(t *testing.T) {

	mockClient := MockBambooClient{}
	mockClient.addComment = func(string, string) error { return errors.New("couldn't add a stupid comment") }
	service := NewCommandService(mockClient)

	got := service.AddCommentHandler([]byte(`{"comment": "test", "build": "123"}`))

	if got.EventDef.Name != "FailedToAddComment" {
		t.Fatalf("expected event FailedToAddComment, got %s", got.EventDef.Name)
	}

}
