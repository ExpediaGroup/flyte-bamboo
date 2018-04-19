[![Build Status](https://travis-ci.org/HotelsDotCom/flyte-bamboo.svg?branch=master)](https://travis-ci.org/HotelsDotCom/flyte-bamboo)
[![Docker Stars](https://img.shields.io/docker/stars/hotelsdotcom/flyte-bamboo.svg)](https://hub.docker.com/r/hotelsdotcom/flyte-bamboo)
[![Docker Pulls](https://img.shields.io/docker/pulls/hotelsdotcom/flyte-bamboo.svg)](https://hub.docker.com/r/hotelsdotcom/flyte-bamboo)


## Overview

The Bamboo pack provides the ability to label and delete label from a build

## Build and Run
### Command Line

To build and run from the command line:
* Clone this repo
* Run `dep ensure` (must have [dep](https://github.com/golang/dep) installed )
* Run `go test ./...`
* Run `go build`
* Run `FLYTE_API_URL=http://localhost:8080/ BAMBOO_HOST=http://localhost:8085 BAMBOO_USER=user1 BAMBOO_PASS=pass ./flyte-bamboo`
* Fill in this command with the relevant API url, bamboo host, bamboo user and bamboo password environment variables

### Docker
To build and run from docker
* Run `docker build -t flyte-bamboo .`
* Run `docker run -e FLYTE_API_URL=http://localhost:8080/ -e BAMBOO_HOST=http://localhost:8085 -e BAMBOO_USER=user -e BAMBOO_PASSWORD=password flyte-bamboo`
* All of these environment variables need to be set

### Commands

[AddLabel](docs/label.md)

[DeleteLabel](docs/label.md)

[StartStage](docs/startstage.md)

[EnablePlan](docs/enableplan.md)

[AddComment](docs/comment.md)

[GetLifeCycle](docs/getlifecyclestate.md)