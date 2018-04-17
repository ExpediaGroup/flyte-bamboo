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

func (c client) AddLabel(newLabel, build string) error {

	args := map[string]string{"name": newLabel}
	path := fmt.Sprintf("/rest/api/latest/result/%v/label", build)
	request, err := c.post(path, args)

	if err != nil {
		err = fmt.Errorf("POST failed: path=%s : label=%s : err=%v", path, newLabel, err)
		return err
	}

	if request.StatusCode != http.StatusNoContent {
		err = fmt.Errorf("POST failed: path=%s : label=%s : statusCode=%v", path, newLabel, request.StatusCode)
		return err
	}

	return nil
}

func (c client) DeleteLabel(deleteLabel, build string) error {

	path := fmt.Sprintf("/rest/api/latest/result/%v/label/%v", build, deleteLabel)
	request, err := c.delete(path, "")

	if err != nil {
		err = fmt.Errorf("DELETE failed: path=%s : label=%s : err=%v", path, deleteLabel, err)
		return err
	}

	if request.StatusCode != http.StatusNoContent {
		err = fmt.Errorf("DELETE failed: path=%s : label=%s : statusCode=%v", path, deleteLabel, request.StatusCode)
		return err
	}

	return nil

}
