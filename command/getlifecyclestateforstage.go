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
	"github.com/ExpediaGroup/flyte-bamboo/event"
	"github.com/HotelsDotCom/flyte-client/flyte"
	"github.com/HotelsDotCom/go-logger"
)

func (c CommandService) GetLifecycleState() flyte.Command {
	return flyte.Command{
		Name:         "GetLifecycleState",
		OutputEvents: []flyte.EventDef{event.GetLifeCycleSuccessEventDef, event.GetLifeCycleErrorEventDef},
		Handler:      c.GetLifecycleStateHandler,
	}
}

func (c CommandService) GetLifecycleStateHandler(input json.RawMessage) flyte.Event {

	var handlerInput struct {
		Build string `json:"build"`
		Stage string `json:"stage"`
	}

	if err := json.Unmarshal(input, &handlerInput); err != nil {
		err = fmt.Errorf("could not unmarshal GetStage command input: %v", err)
		logger.Error(err)
		return flyte.NewFatalEvent(fmt.Sprintf("Input is invalid: %v", err))
	}

	stage, err := c.bambooClient.GetStage(handlerInput.Build, handlerInput.Stage)
	if err != nil {
		err = fmt.Errorf("could not get life cycle of build: %v\n", err)
		return event.GetLifeCycleErrorEvent(err.Error(), handlerInput.Build, handlerInput.Stage)
		logger.Errorf(err.Error())
	}

	return event.GetLifeCycleSuccessEvent(handlerInput.Build, handlerInput.Stage, stage.State, stage.LifeCycleState)
}
