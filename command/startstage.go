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
	"encoding/json"
	"fmt"
	"github.com/HotelsDotCom/flyte-bamboo/bamboo"
	"github.com/HotelsDotCom/flyte-bamboo/event"
	"github.com/HotelsDotCom/flyte-client/flyte"
	"github.com/HotelsDotCom/go-logger"
	"time"
)

type startStageInput struct {
	Build    string            `json:"build"`
	Stage    string            `json:"stage"`
	Vars     map[string]string `json:"vars"`
	MaxRetry int               `json:"maxRetry"`
	Delay    time.Duration     `json:"delay"`
}

func (c CommandService) StartStageCommand() flyte.Command {
	return flyte.Command{
		Name:         "StartStage",
		OutputEvents: []flyte.EventDef{event.StartStageSuccessEventDef, event.StartStageErrorEventDef},
		Handler:      c.StartStageHandler,
	}

}
func (c CommandService) StartStageHandler(rawInput json.RawMessage) flyte.Event {

	var input startStageInput

	if err := json.Unmarshal(rawInput, &input); err != nil {
		err = fmt.Errorf("could not marshal startStage command input: %v\n", err)
		logger.Error(err)
		return flyte.NewFatalEvent(fmt.Sprintf("input is invalid: %v", err))
	}

	retry := bamboo.Retry{Count: 0, Retries: input.MaxRetry, RetrySleep: input.Delay * time.Second}

	if err := c.bambooClient.StartStage(input.Build, input.Stage, input.Vars, retry); err != nil {
		logger.Error(err)
		return event.StartStageFailureEvent(fmt.Sprintf("Fail: %s", err), input.Stage, input.Vars, err)
	}

	return event.StartStageSuccessEvent(input.Build, input.Stage, input.Vars)
}
