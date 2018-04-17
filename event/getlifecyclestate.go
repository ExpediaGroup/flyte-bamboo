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

var GetLifeCycleSuccessEventDef = flyte.EventDef{Name: "GetLifeCycleState"}
var GetLifeCycleErrorEventDef = flyte.EventDef{Name: "GetLifeCycleFailed"}

func GetLifeCycleSuccessEvent(build, stage, state, lifecycle string) flyte.Event {
	return flyte.Event{
		EventDef: GetLifeCycleSuccessEventDef,
		Payload:  GetLifeCycleSuccessPayload{stage, lifecycle, state, build},
	}
}

func GetLifeCycleErrorEvent(error, build, stage string) flyte.Event {
	return flyte.Event{
		EventDef: GetLifeCycleErrorEventDef,
		Payload:  GetLifeCycleFailurePayload{stage, build, error},
	}
}

type GetLifeCycleSuccessPayload struct {
	Stage     string `json:"stage"`
	LifeCycle string `json:"lifecycle"`
	State     string `json:"state"`
	Build     string `json:"build"`
}

type GetLifeCycleFailurePayload struct {
	Stage string `json:"stage"`
	Build string `json:"build"`
	Error string `json:"error"`
}
