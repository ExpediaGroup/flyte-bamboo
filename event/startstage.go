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

var StartStageSuccessEventDef = flyte.EventDef{Name: "StartStageSuccess"}
var StartStageErrorEventDef = flyte.EventDef{Name: "FailedToStartStage"}

func StartStageSuccessEvent(build, stage string, vars map[string]string) flyte.Event {

	return flyte.Event{
		EventDef: StartStageSuccessEventDef,
		Payload:  StartStageSuccessPayload{build, stage, vars},
	}

}

func StartStageFailureEvent(build, stage string, vars map[string]string, error error) flyte.Event {

	return flyte.Event{
		EventDef: StartStageErrorEventDef,
		Payload:  StartStageFailurePayload{build, stage, vars, error.Error()},
	}

}

type StartStageSuccessPayload struct {
	Build string            `json:"build"`
	Stage string            `json:"stage"`
	Vars  map[string]string `json:"vars"`
}

type StartStageFailurePayload struct {
	Build string            `json:"build"`
	Stage string            `json:"stage"`
	Vars  map[string]string `json:"vars"`
	Error string            `json:"error"`
}
