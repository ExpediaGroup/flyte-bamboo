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

func (c client) AddComment(newComment, build string) error {

	args := map[string]string{"content": newComment}

	path := fmt.Sprintf("/rest/api/latest/result/%v/comment", build)
	response, err := c.post(path, args)

	if err != nil {
		err = fmt.Errorf("POST failed: path=%s : comment=%s : err=%v", path, newComment, err)
		return err
	}

	if response.StatusCode != http.StatusNoContent {
		err = fmt.Errorf("POST failed: path=%s : comment=%s : statusCode=%v", path, newComment, response)
		return err
	}

	return nil
}
