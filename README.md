# dockerdetector

[![Go](https://github.com/dhcgn/dockerdetector/actions/workflows/go.yml/badge.svg)](https://github.com/dhcgn/dockerdetector/actions/workflows/go.yml)
[![codecov](https://codecov.io/gh/dhcgn/dockerdetector/branch/main/graph/badge.svg?token=9CMJ0HZA6B)](https://codecov.io/gh/dhcgn/dockerdetector)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=dhcgn_dockerdetector&metric=alert_status)](https://sonarcloud.io/dashboard?id=dhcgn_dockerdetector)
[![Go Report Card](https://goreportcard.com/badge/github.com/dhcgn/dockerdetector)](https://goreportcard.com/report/github.com/dhcgn/dockerdetector)
[![Go Doc](https://godoc.org/github.com/dhcgn/dockerdetector?status.svg)](http://godoc.org/github.com/dhcgn/dockerdetector)


## Intro

This package use cgroup to determine if we run as a docker container. 
It runs only under linux, windows und macOS will always returning false.

## Usage

`go get -u github.com/dhcgn/dockerdetector`

```go
package main

import (
	"fmt"

	"github.com/dhcgn/dockerdetector"
)

func main() {
	isDocker, err := dockerdetector.IsRunningInContainer()
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("Run in Docker:", isDocker)
}
```

## Run in docker

```
root@110c80b1eb43:/go# go install github.com/dhcgn/dockerdetector/cmd/dockerdetector@latest
root@110c80b1eb43:/go# dockerdetector
Run in Docker: true
```
