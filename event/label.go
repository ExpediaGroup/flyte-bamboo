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

var AddLabelSuccessEventDef = flyte.EventDef{Name: "AddedLabel"}
var AddLabelErrorEventDef = flyte.EventDef{Name: "FailedToAddLabel"}
var DeleteLabelSuccessEventDef = flyte.EventDef{Name: "DeletedLabel"}
var DeleteLabelErrorEventDef = flyte.EventDef{Name: "FailedToDeleteLabel"}

func NewAddLabelSuccessEvent(build, label string) flyte.Event {

	return flyte.Event{
		EventDef: AddLabelSuccessEventDef,
		Payload:  AddLabelSuccessPayload{label, build},
	}
}

func NewAddLabelErrorEvent(error, build, label string) flyte.Event {
	return flyte.Event{
		EventDef: AddLabelErrorEventDef,
		Payload:  AddLabelFailurePayload{label, build, error},
	}
}

func DeleteLabelSuccessEvent(build, label string) flyte.Event {

	return flyte.Event{
		EventDef: DeleteLabelSuccessEventDef,
		Payload:  DeleteLabelSuccessPayload{label, build},
	}

}

func DeleteLabelErrorEvent(error, build, label string) flyte.Event {

	return flyte.Event{
		EventDef: DeleteLabelErrorEventDef,
		Payload:  DeleteLabelFailurePayload{label, build, error},
	}

}

type AddLabelSuccessPayload struct {
	Label string `json:"label"`
	Build string `json:"build"`
}

type AddLabelFailurePayload struct {
	Label string `json:"label"`
	Build string `json:"build"`
	Error string `json:"error"`
}

type DeleteLabelSuccessPayload struct {
	Label string `json:"label"`
	Build string `json:"build"`
}

type DeleteLabelFailurePayload struct {
	Label string `json:"label"`
	Build string `json:"build"`
	Error string `json:"error"`
}
