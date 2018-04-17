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
	"errors"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)


func TestNewBambooClient(t *testing.T) {

	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}

	type arguments struct {
		httpClient *http.Client
		baseUrl    string
		username   string
		password   string
		errorMsg   string
	}

	var invalidArguments = []arguments{
		{nil, "/ABC/BC-122", "username", "password", "bamboo httpClient is not set"},
		{httpClient, "", "username", "password", "bamboo baseUrl is not set"},
		{httpClient, "/ABC/ABC-123", "", "password", "bamboo username is not set"},
		{httpClient, "/ABC/ABC-123", "username", "", "bamboo password is not set"},
	}

	for _, tt := range invalidArguments {
		_, err := NewBambooClient(tt.httpClient, tt.baseUrl, tt.username, tt.password)
		c := strings.Compare(tt.errorMsg, err.Error())
		if c > 0 || c < 0 {
			t.Errorf("Error not correct!. Error: %s", err)
		}
	}
}

func TestDefaultClientTimeoutIsSet(t *testing.T) {
	cl, err := DefaultBambooClient("http://localhost:8085", "username", "password")
	assert.NoErrorf(t, err, "Unexpected Error: %+v", err)
	assert.NotZero(t, cl.(client).httpClient.Timeout)
	assert.Equal(t, 60 * time.Second, cl.(client).httpClient.Timeout)
}

func bambooClient(t *testing.T, url string) BambooClient {
	cl, err := DefaultBambooClient(url, "username", "password")
	assert.NoErrorf(t, err, "Unexpected Error: %+v", err)
	return cl
}

func TestBambooClient_AddLabel(t *testing.T) {
	want := `{"name":"label"}`
	ts := httptest.NewServer(testPostHandler(t, "/rest/api/latest/result/ABC-123/label", want))
	defer ts.Close()

	client := bambooClient(t, ts.URL)

	err := client.AddLabel("label", "ABC-123")
	assert.NoErrorf(t, err, "Unexpected Error: %+v", err)
}

func TestBambooClient_DeleteLabel(t *testing.T) {
	ts := httptest.NewServer(testDeleteHandler(t, "/rest/api/latest/result/ABC-123/label/label"))
	defer ts.Close()

	client := bambooClient(t, ts.URL)

	err := client.DeleteLabel("label", "ABC-123")
	assert.NoErrorf(t, err, "Unexpected Error: %+v", err)
}

func TestBambooClient_StartNextStage(t *testing.T) {
	url := "/rest/api/latest/queue/ABC-123?stage=staging&executeAllStages=false&bamboo.variable.deploySuccess=true"
	ts := httptest.NewServer(testPutHandler(t, url, ""))

	client := bambooClient(t, ts.URL)

	retry := Retry{Count: 0, Retries: 0, RetrySleep: time.Second * 2}
	err := client.StartStage("ABC-123", "staging", map[string]string{"deploySuccess": "true"}, retry)
	assert.NoErrorf(t, err, "Unexpected Error: %+v", err)
}

func TestBambooClient_TriggerPlan(t *testing.T) {
	url := "/rest/api/latest/queue/ABC-DEF&executeAllStages=false"
	ts := httptest.NewServer(testPutHandler(t, url, ""))

	client := bambooClient(t, ts.URL)

	err := client.TriggerPlan("ABC-DEF")
	assert.NoErrorf(t, err, "Unexpected Error: %+v", err)
}

func TestBambooClient_AddComment(t *testing.T) {
	want := `{"content":"some comment"}`
	ts := httptest.NewServer(testPostHandler(t, "/rest/api/latest/result/ABC-123/comment", want))
	defer ts.Close()

	client := bambooClient(t, ts.URL)

	err := client.AddComment("some comment", "ABC-123")
	assert.NoErrorf(t, err, "Unexpected Error: %+v", err)
}

func TestClient_EnablePlan(t *testing.T) {
	url := "/rest/api/latest/plan/ABC-DEF/enable"
	ts := httptest.NewServer(testPostHandler(t, url, ""))
	defer ts.Close()

	client := bambooClient(t, ts.URL)

	err := client.EnablePlan("ABC-DEF")
	assert.NoErrorf(t, err, "Unexpected Error: %+v", err)
}

func TestClient_GetStage(t *testing.T) {
	ts := httptest.NewServer(getHandler(200, resultResponse))
	defer ts.Close()

	client := bambooClient(t, ts.URL)

	lifeCycleState, err := client.GetStage("FLYTE-FLYTEAPI-12", "prod-verify")
	assert.NoError(t, err, "Error Message: %+v", err)
	assert.Equal(t, "NotBuilt", lifeCycleState.LifeCycleState, "Got: %v", lifeCycleState.LifeCycleState)
	assert.Equal(t, "Unknown", lifeCycleState.State, "Got: %v", lifeCycleState.State)
}

func TestClient_GetStageBambooFails(t *testing.T) {
	ts := httptest.NewServer(getHandler(500, resultResponse))
	defer ts.Close()

	client := bambooClient(t, ts.URL)

	errText := errors.New("Some error")
	_, err := client.GetStage("FLYTE-FLYTEAPI-12", "prod-verify")
	assert.Errorf(t, errText, err.Error())
}

func TestClient_GetStageWrongLifeCycleState(t *testing.T) {
	ts := httptest.NewServer(getHandler(200, resultResponse))
	defer ts.Close()

	client := bambooClient(t, ts.URL)

	_, err := client.GetStage("FLYTE-FLYTEAPI-12", "prod-verify")
	assert.NoError(t, err)
	assert.NotEqual(t, "Built", resultResponse)
}

func TestClient_DelayedStageRun_Handles400(t *testing.T) {

	// given
	tries := 0
	ts := httptest.NewServer(func() http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			tries++
			if tries > 3 {
				w.WriteHeader(200)
				w.Write([]byte(""))
				return
			}
			w.WriteHeader(400)
			w.Write([]byte(""))
		}
	}())

	defer ts.Close()
	client := bambooClient(t, ts.URL)

	// when
	retry := Retry{Count: 0, Retries: 5, RetrySleep: time.Second * 2}
	err := client.StartStage("ABC-123", "staging", map[string]string{"deploySuccess": "true"}, retry)

	// then
	assert.NoError(t, err)
}

func TestClient_StartNextStage_FailsAfterMaxNumberOFRetries(t *testing.T) {

	// given

	ts := httptest.NewServer(postHandler2(400, ""))
	client := bambooClient(t, ts.URL)
	defer ts.Close()

	//when
	retry := Retry{Count: 0, Retries: 4, RetrySleep: time.Second * 2}
	err := client.StartStage("ABC-123", "staging", map[string]string{"deploySuccess": "true"}, retry)

	//then
	errString := errors.New("failed to progress reached maximum number of retryCount: [4]")
	assert.EqualError(t, errString, err.Error())
}

// test two different failure codes.
func TestClient_DelayedStageRun_Handles500(t *testing.T) {
	ts := httptest.NewServer(postHandler2(500, ""))
	client := bambooClient(t, ts.URL)
	defer ts.Close()

	retry := Retry{Count: 0, Retries: 4, RetrySleep: time.Second * 2}
	errText := errors.New("bamboo is returning a 500 error. Build: ABC-123,stage: staging")

	err := client.StartStage("ABC-123", "staging", map[string]string{"deploySuccess": "true"},retry)
	assert.Errorf(t, err, errText.Error())
}

func TestClient_DelayedStageRun_Handles404(t *testing.T) {
	ts := httptest.NewServer(postHandler2(404, ""))
	defer ts.Close()

	client := bambooClient(t, ts.URL)

	errString := errors.New("bamboo is returning a 404 error. Build: ABC-123,stage: staging")
	retry := Retry{Count: 0, Retries: 4, RetrySleep: time.Second * 2}

	err := client.StartStage("ABC-123", "staging", map[string]string{"deploySuccess": "true"}, retry)
	assert.Errorf(t, errString, err.Error())
}

func postHandler2(responseCode int, responseBody string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(responseCode)
		w.Write([]byte(responseBody))
	}
}

func testPostHandler(t *testing.T, url, want interface{}) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.RequestURI != url {
			t.Errorf("Page Not Found\ngot  = %s\nwant = %s", r.RequestURI, url)
			w.WriteHeader(http.StatusNotFound)
		} else if r.Method != http.MethodPost {

			w.WriteHeader(http.StatusBadRequest)
		} else {
			b, err := ioutil.ReadAll(r.Body)
			if err != nil {
				t.Error(err)
			}
			got := string(b)

			// OK, this looks weird but needed to handle empty body
			if want == "" {
				want = "\"\""
			}

			if got != want {
				w.WriteHeader(http.StatusBadRequest)
				t.Errorf("Bad request\ngot  = %s\nwant = %s", got, want)
			} else {
				w.WriteHeader(http.StatusNoContent)
			}
		}
	}
}

func testPutHandler(t *testing.T, url, want interface{}) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.RequestURI != url {
			t.Errorf("Page Not Found\ngot  = %s\nwant = %s", r.RequestURI, url)
			w.WriteHeader(http.StatusNotFound)
		} else if r.Method != http.MethodPut {

			w.WriteHeader(http.StatusBadRequest)
		} else {
			b, err := ioutil.ReadAll(r.Body)
			if err != nil {
				t.Error(err)
			}

			got := string(b)

			// OK, this looks weird but needed to handle empty body
			if want == "" {
				want = "\"\""
			}

			if got != want {
				w.WriteHeader(http.StatusBadRequest)
				t.Errorf("Bad request\ngot  = %s\nwant = %s", got, want)
			} else {
				w.WriteHeader(http.StatusOK)
			}
		}
	}
}

func getHandler(responseCode int, responseBody string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(responseCode)
		w.Write([]byte(responseBody))
	}
}

func testDeleteHandler(t *testing.T, url string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.RequestURI != url {
			t.Errorf("Page Not Found%v", r.RequestURI)
			w.WriteHeader(http.StatusNotFound)
		} else if r.Method != http.MethodDelete {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			w.WriteHeader(http.StatusNoContent)
		}
	}
}
