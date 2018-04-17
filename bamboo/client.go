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
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
	"crypto/tls"
)

type Retry struct {
	Count      int
	Retries    int
	RetrySleep time.Duration
}

type BambooClient interface {
	StartStage(build, stage string, vars map[string]string, retry Retry) error
	AddLabel(label, build string) error
	DeleteLabel(label, build string) error
	TriggerPlan(plan string) error
	AddComment(comment, build string) error
	EnablePlan(planName string) error
	GetStage(build, stageName string) (Stage, error)
}

func NewBambooClient(httpClient *http.Client, baseUrl, user, pass string) (BambooClient, error) {

	if httpClient == nil {
		return nil, errors.New("bamboo httpClient is not set")
	}

	if strings.TrimSpace(baseUrl) == "" {
		return nil, errors.New("bamboo baseUrl is not set")
	}

	parsedBaseUrl, err := url.Parse(baseUrl)
	if err != nil {
		return nil, errors.New("couldn't parse Url")
	}

	if strings.TrimSpace(user) == "" {
		return nil, errors.New("bamboo username is not set")
	}

	if strings.TrimSpace(pass) == "" {
		return nil, errors.New("bamboo password is not set")
	}

	return client{
		baseUrl:    parsedBaseUrl,
		User:       user,
		Pass:       pass,
		httpClient: httpClient,
	}, nil
}

func DefaultBambooClient(baseUrl, username, password string) (BambooClient, error) {
	tr := &http.Transport{TLSClientConfig:&tls.Config{InsecureSkipVerify:true}}
	httpClient := &http.Client{Timeout: 60 * time.Second, Transport:tr}
	return NewBambooClient(httpClient, baseUrl, username, password)
}

type client struct {
	baseUrl    *url.URL
	User       string
	Pass       string
	httpClient *http.Client
}

func (c *client) put(path string, body interface{}) (*http.Response, error) {

	b, err := json.Marshal(body)
	urlStr := fmt.Sprintf("%s%s", c.baseUrl, path)

	if err != nil {
		return nil, fmt.Errorf("cannot marshal body '%+v': %v", body, err)
	}

	req, err := http.NewRequest(http.MethodPut, urlStr, bytes.NewBuffer(b))
	if err != nil {
		return nil, fmt.Errorf("cannot create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.SetBasicAuth(c.User, c.Pass)

	return c.httpClient.Do(req)
}

func (c *client) delete(path string, body interface{}) (*http.Response, error) {

	b, err := json.Marshal(body)
	urlStr := fmt.Sprintf("%s%s", c.baseUrl, path)

	if err != nil {
		return nil, fmt.Errorf("cannot marshal body '%+v': %v", body, err)
	}

	req, err := http.NewRequest(http.MethodDelete, urlStr, bytes.NewBuffer(b))
	if err != nil {
		return nil, fmt.Errorf("cannot create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.SetBasicAuth(c.User, c.Pass)

	return c.httpClient.Do(req)
}

func (c *client) post(path string, body interface{}) (*http.Response, error) {

	b, err := json.Marshal(body)
	urlStr := fmt.Sprintf("%s%s", c.baseUrl, path)

	if err != nil {
		return nil, fmt.Errorf("cannot marshal body '%+v': %v", body, err)
	}

	req, err := http.NewRequest(http.MethodPost, urlStr, bytes.NewBuffer(b))
	if err != nil {
		return nil, fmt.Errorf("cannot create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.SetBasicAuth(c.User, c.Pass)

	return c.httpClient.Do(req)
}

func (c *client) get(path string) (*http.Response, error) {

	urlStr := fmt.Sprintf("%s%s", c.baseUrl, path)
	req, err := http.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")
	req.SetBasicAuth(c.User, c.Pass)
	return c.httpClient.Do(req)
}

func (c *client) getUrl(path string) string {
	url := c.baseUrl
	url.Path = path
	return url.String()
}
