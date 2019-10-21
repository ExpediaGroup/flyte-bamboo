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
	"fmt"
	"github.com/HotelsDotCom/go-logger"
	"net/http"
	"time"
)

func (c client) StartStage(build, stage string, vars map[string]string, retry Retry) (err error) {

	path := fmt.Sprintf("/rest/api/latest/queue/%s?stage=%v&executeAllStages=false", build, stage)
	for k, v := range vars {
		path = fmt.Sprintf("%s&bamboo.variable.%s=%s", path, k, v)
	}

	response, err := c.put(path, "")
	if err != nil {
		err = fmt.Errorf("failed to initiate stage: %s err: %v", path, err)
		logger.Error(err)
		return err
	}

	if response.StatusCode >= 200 && response.StatusCode < 299 {
		logger.Infof("all happy, progressed build: %s to stage: %s\n", build, stage)
		return nil
	}

	if response.StatusCode == http.StatusBadRequest {
		retry.Count++
		if retry.Count >= retry.Retries {
			return fmt.Errorf("failed to progress reached maximum number of retryCount: [%d]", retry.Retries)
		}
		logger.Infof("Stage: %s, is already in progress, not progressing build: %s, to the next stage.", stage, build)
		time.Sleep(retry.RetrySleep)
		return c.StartStage(build, stage, vars, retry)
	}

	return fmt.Errorf("bamboo is returning a %d error. Build: %v,stage: %v", response.StatusCode, build, stage)

}
