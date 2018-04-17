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

var TriggerPlanSuccessEventDef = flyte.EventDef{Name: "PlanTriggered"}
var TriggerPlanErrorEventDef = flyte.EventDef{Name: "FailedToTriggerPlan"}

func TriggerPlanSuccessEvent(plan string) flyte.Event {

	return flyte.Event{
		EventDef: TriggerPlanSuccessEventDef,
		Payload:  TriggerPlanSuccessPayload{plan},
	}

}

func TriggerPlanFailureEvent(plan, error string) flyte.Event {

	return flyte.Event{
		EventDef: TriggerPlanErrorEventDef,
		Payload:  TriggerPlanFailurePayload{plan, error},
	}

}

type TriggerPlanSuccessPayload struct {
	Plan string `json:"plan"`
}

type TriggerPlanFailurePayload struct {
	Plan  string `json:"plan"`
	Error string `json:"error"`
}
