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

package bamboo

import (
	"encoding/json"
	"fmt"
	"github.com/HotelsDotCom/go-logger"
	"io/ioutil"
	"net/http"
)

type ResultResponse struct {
	Stages Stages `json:"stages"`
}

type Stages struct {
	Stage []Stage `json:"stage"`
}

type Stage struct {
	Name           string `json:"name"`
	LifeCycleState string `json:"lifeCycleState"`
	State          string `json:"state"`
}

func (c client) GetStage(build, stageName string) (Stage, error) {

	var resultResponse ResultResponse

	path := fmt.Sprintf("/rest/api/latest/result/%v.json?expand=stages", build)
	response, err := c.get(path)
	if err != nil {
		err = fmt.Errorf("GET failed: path=%s : build=%s : err=%v", path, build, err)
		logger.Error(err)
		return Stage{}, err
	}

	if response.StatusCode != http.StatusOK {
		err = fmt.Errorf("GET failed: path=%s : build=%s : statusCode=%v", path, build, response.StatusCode)
		logger.Error(err)
		return Stage{}, err
	}

	body, _ := ioutil.ReadAll(response.Body)
	err = json.Unmarshal(body, &resultResponse)
	if err != nil {
		logger.Fatalf("Could not unmarshal body: %v", err)
	}

	return getStageByName(resultResponse.Stages, stageName), nil
}

func getStageByName(stages Stages, name string) Stage {
	for _, s := range stages.Stage {
		if s.Name == name {
			return s
		}
	}
	return Stage{}
}
