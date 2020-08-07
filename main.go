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

package main

import (
	"github.com/ExpediaGroup/flyte-bamboo/bamboo"
	"github.com/ExpediaGroup/flyte-bamboo/command"
	"github.com/HotelsDotCom/flyte-client/client"
	"github.com/HotelsDotCom/flyte-client/flyte"
	"github.com/HotelsDotCom/go-logger"
	"net/url"
	"os"
	"time"
)

func main() {

	bambooClient, err := newBambooClient()
	if err != nil {
		logger.Fatalf("Could not initialise Bamboo Client: %v", err)
	}

	commandService := command.NewCommandService(bambooClient)

	packDef := flyte.PackDef{
		Name:    "Bamboo",
		HelpURL: getUrl("https://github.com/ExpediaGroup/flyte-bamboo/blob/master/README.md"),
		Commands: []flyte.Command{
			commandService.AddLabelCommand(),
			commandService.DeleteLabelCommand(),
			commandService.StartStageCommand(),
			commandService.TriggerPlanCommand(),
			commandService.EnablePlanCommand(),
			commandService.AddCommentCommand(),
			commandService.GetLifecycleState()},
	}

	p := flyte.NewPack(packDef, client.NewClient(hostUrl(), 10*time.Second))
	p.Start()
	select {}
}

func newBambooClient() (bamboo.BambooClient, error) {
	return bamboo.DefaultBambooClient(os.Getenv("BAMBOO_HOST"), os.Getenv("BAMBOO_USER"), os.Getenv("BAMBOO_PASS"))
}

func hostUrl() *url.URL {
	if os.Getenv("FLYTE_API_URL") == "" {
		logger.Fatalf("FLYTE_API_URL environment variable is not set")
	}

	return getUrl(os.Getenv("FLYTE_API_URL"))
}

func getUrl(rawUrl string) *url.URL {
	url, err := url.Parse(rawUrl)
	if err != nil {
		logger.Fatalf("%s is not a valid url", rawUrl)
	}
	return url
}
