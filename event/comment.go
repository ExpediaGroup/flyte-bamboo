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

package event

import "github.com/HotelsDotCom/flyte-client/flyte"

var AddCommentSuccessEventDef = flyte.EventDef{Name: "CommentAdded"}
var AddCommentErrorEventDef = flyte.EventDef{Name: "FailedToAddComment"}

func NewAddCommentSuccessEvent(build, comment string) flyte.Event {

	return flyte.Event{
		EventDef: AddCommentSuccessEventDef,
		Payload:  AddCommentSuccessPayload{comment, build},
	}
}

func NewAddCommentErrorEvent(error, build, comment string) flyte.Event {
	return flyte.Event{
		EventDef: AddCommentErrorEventDef,
		Payload:  AddLabelFailurePayload{comment, build, error},
	}
}

type AddCommentSuccessPayload struct {
	Comment string `json:"comment"`
	Build   string `json:"build"`
}

type AddCommentFailurePayload struct {
	Comment string `json:"comment"`
	Build   string `json:"build"`
	Error   string `json:"error"`
}
