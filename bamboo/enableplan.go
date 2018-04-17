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
	"net/http"
)

func (c client) EnablePlan(plan string) error {

	path := fmt.Sprintf("/rest/api/latest/plan/%v/enable", plan)
	response, err := c.post(path, "")

	if err != nil {
		err = fmt.Errorf("POST failed: path=%s  err=%v", path, err)
		return err
	}

	if response.StatusCode != http.StatusNoContent {
		err = fmt.Errorf("POST failed: path=%s : statusCode=%v", path, response.StatusCode)
		return err
	}

	return nil
}
