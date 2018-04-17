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
	"github.com/HotelsDotCom/flyte-bamboo/event"
	"github.com/HotelsDotCom/flyte-client/flyte"
	"github.com/HotelsDotCom/go-logger"
)

func (c CommandService) AddLabelCommand() flyte.Command {
	return flyte.Command{
		Name:         "AddLabel",
		OutputEvents: []flyte.EventDef{event.AddLabelSuccessEventDef, event.AddLabelErrorEventDef},
		Handler:      c.AddLabelHandler,
	}
}

func (c CommandService) DeleteLabelCommand() flyte.Command {
	return flyte.Command{
		Name:         "DeleteLabel",
		OutputEvents: []flyte.EventDef{event.DeleteLabelSuccessEventDef, event.DeleteLabelErrorEventDef},
		Handler:      c.DeleteLabelHandler,
	}
}

func (c CommandService) AddLabelHandler(input json.RawMessage) flyte.Event {

	var handlerInput struct {
		Label string `json:"label"`
		Build string `json:"build"`
	}

	if err := json.Unmarshal(input, &handlerInput); err != nil {
		err = fmt.Errorf("could not marshal AddLabel command input: %v", err)
		logger.Error(err)
		return flyte.NewFatalEvent(fmt.Sprintf("Input is invalid: %v", err))
	}

	if err := c.bambooClient.AddLabel(handlerInput.Label, handlerInput.Build); err != nil {
		err = fmt.Errorf("Could not apply label: %v\n", err)
		logger.Error(err)
		return event.NewAddLabelErrorEvent(fmt.Sprintf("Fail: %s", err), handlerInput.Build, handlerInput.Label)
	}

	return event.NewAddLabelSuccessEvent(handlerInput.Build, handlerInput.Label)
}

func (c CommandService) DeleteLabelHandler(input json.RawMessage) flyte.Event {

	var handlerInput struct {
		Label string `json:"label"`
		Build string `json:"build"`
	}

	if err := json.Unmarshal(input, &handlerInput); err != nil {
		err = fmt.Errorf("could not marshal DeleteLabel command input: %v", err)
		logger.Error(err)
		return flyte.NewFatalEvent(fmt.Sprintf("Input is invalid: %v", err))
	}

	if err := c.bambooClient.DeleteLabel(handlerInput.Label, handlerInput.Build); err != nil {
		err = fmt.Errorf("could not delete label: %v", err)
		logger.Error(err)
		return event.DeleteLabelErrorEvent(fmt.Sprintf("Fail: %s", err), handlerInput.Build, handlerInput.Label)
	}

	return event.DeleteLabelSuccessEvent(handlerInput.Build, handlerInput.Label)

}
