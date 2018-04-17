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

var resultResponse = `
{
	"expand": "changes,metadata,plan,vcsRevisions,artifacts,comments,labels,jiraIssues,stages",
	"link": {
		"href": "http://bamboo.com/rest/api/latest/result/FLYTE-FLYTEAPI-12",
		"rel": "self"
	},
	"plan": {
		"shortName": "FLYTEAPI",
		"shortKey": "FLYTEAPI",
		"type": "chain",
		"enabled": true,
		"link": {
			"href": "http://bamboo.com/rest/api/latest/plan/FLYTE-FLYTEAPI",
			"rel": "self"
		},
		"key": "FLYTE-FLYTEAPI",
		"name": "Flyte - FLYTEAPI",
		"planKey": {
			"key": "FLYTE-FLYTEAPI"
		}
	},
	"planName": "FLYTEAPI",
	"projectName": "Flyte",
	"buildResultKey": "FLYTE-FLYTEAPI-12",
	"lifeCycleState": "Finished",
	"id": 293657870,
	"buildStartedTime": "2017-08-09T09:44:23.700-05:00",
	"prettyBuildStartedTime": "Wed, 9 Aug, 09:44 AM",
	"buildCompletedTime": "2017-08-09T12:31:16.040-05:00",
	"buildCompletedDate": "2017-08-09T12:31:16.040-05:00",
	"prettyBuildCompletedTime": "Wed, 9 Aug, 12:31 PM",
	"buildDurationInSeconds": 20614,
	"buildDuration": 20614719,
	"buildDurationDescription": "343 minutes",
	"buildRelativeTime": "4 days ago",
	"vcsRevisionKey": "3323d2932cb383ef7397b065672851b379f336b8",
	"vcsRevisions": {
		"size": 1,
		"start-index": 0,
		"max-result": 1
	},
	"buildTestSummary": "No tests found",
	"successfulTestCount": 0,
	"failedTestCount": 0,
	"quarantinedTestCount": 0,
	"skippedTestCount": 0,
	"continuable": true,
	"onceOff": false,
	"restartable": false,
	"notRunYet": false,
	"finished": true,
	"successful": true,
	"buildReason": "Manual run from the stage: <b>milan-verify</b> by <a href=\"http://bamboo.com/browse/user/bamboo\">Bamboo</a>",
	"reasonSummary": "Manual run from the stage: <b>milan-verify</b> by <a href=\"http://bamboo.com/browse/user/bamboo\">Bamboo</a>",
	"artifacts": {
		"size": 0,
		"start-index": 0,
		"max-result": 0
	},
	"comments": {
		"size": 2,
		"start-index": 0,
		"max-result": 2
	},
	"labels": {
		"size": 2,
		"start-index": 0,
		"max-result": 2
	},
	"jiraIssues": {
		"size": 1,
		"start-index": 0,
		"max-result": 1
	},
	"stages": {
		"size": 10,
		"expand": "stage",
		"start-index": 0,
		"max-result": 10,
		"stage": [{
			"name": "Default Stage",
			"id": 293227916,
			"lifeCycleState": "Finished",
			"state": "Successful",
			"displayClass": "Successful",
			"displayMessage": "1 job passed",
			"collapsedByDefault": true,
			"manual": false,
			"restartable": false,
			"runnable": false
		}, {
			"name": "Branch Validation",
			"description": "",
			"id": 293227917,
			"lifeCycleState": "Finished",
			"state": "Successful",
			"displayClass": "Successful",
			"displayMessage": "1 job passed",
			"collapsedByDefault": true,
			"manual": false,
			"restartable": false,
			"runnable": false
		}, {
			"name": "initiate-deploy-staging",
			"description": "",
			"id": 293227918,
			"lifeCycleState": "Finished",
			"state": "Successful",
			"displayClass": "Successful",
			"displayMessage": "1 job passed",
			"collapsedByDefault": true,
			"manual": true,
			"restartable": false,
			"runnable": false
		}, {
			"name": "staging-verify",
			"description": "",
			"id": 293227919,
			"lifeCycleState": "Finished",
			"state": "Successful",
			"displayClass": "Successful",
			"displayMessage": "1 job passed",
			"collapsedByDefault": true,
			"manual": true,
			"restartable": false,
			"runnable": false
		}, {
			"name": "initiate-deploy-milan",
			"description": "",
			"id": 293227920,
			"lifeCycleState": "Finished",
			"state": "Successful",
			"displayClass": "Successful",
			"displayMessage": "1 job passed",
			"collapsedByDefault": true,
			"manual": true,
			"restartable": false,
			"runnable": false
		}, {
			"name": "milan-verify",
			"description": "",
			"id": 293227921,
			"lifeCycleState": "Finished",
			"state": "Successful",
			"displayClass": "Successful",
			"displayMessage": "1 job passed",
			"collapsedByDefault": true,
			"manual": true,
			"restartable": false,
			"runnable": false
		}, {
			"name": "initiate-deploy-prod-canary",
			"description": "",
			"id": 293227922,
			"lifeCycleState": "NotBuilt",
			"state": "Unknown",
			"displayClass": "NotBuilt",
			"displayMessage": "1 of 1 job not executed",
			"collapsedByDefault": true,
			"manual": true,
			"restartable": false,
			"runnable": true
		}, {
			"name": "prod-canary-verify",
			"description": "",
			"id": 293227923,
			"lifeCycleState": "NotBuilt",
			"state": "Unknown",
			"displayClass": "NotBuilt",
			"displayMessage": "1 of 1 job not executed",
			"collapsedByDefault": true,
			"manual": true,
			"restartable": false,
			"runnable": false
		}, {
			"name": "initiate-deploy-prod",
			"description": "",
			"id": 293227924,
			"lifeCycleState": "NotBuilt",
			"state": "Unknown",
			"displayClass": "NotBuilt",
			"displayMessage": "1 of 1 job not executed",
			"collapsedByDefault": true,
			"manual": true,
			"restartable": false,
			"runnable": false
		}, {
			"name": "prod-verify",
			"description": "",
			"id": 293227925,
			"lifeCycleState": "NotBuilt",
			"state": "Unknown",
			"displayClass": "NotBuilt",
			"displayMessage": "1 of 1 job not executed",
			"collapsedByDefault": true,
			"manual": true,
			"restartable": false,
			"runnable": false
		}]
	},
	"changes": {
		"size": 2,
		"start-index": 0,
		"max-result": 2
	},
	"metadata": {
		"size": 3,
		"start-index": 0,
		"max-result": 3
	},
	"key": "FLYTE-FLYTEAPI-12",
	"planResultKey": {
		"key": "FLYTE-FLYTEAPI-12",
		"entityKey": {
			"key": "FLYTE-FLYTEAPI"
		},
		"resultNumber": 12
	},
	"state": "Successful",
	"buildState": "Successful",
	"number": 12,
	"buildNumber": 12
}
`
