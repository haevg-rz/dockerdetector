# dockerdetector

[![Go](https://github.com/haevg-rz/dockerdetector/actions/workflows/go.yml/badge.svg)](https://github.com/haevg-rz/dockerdetector/actions/workflows/go.yml)
[![codecov](https://codecov.io/gh/haevg-rz/dockerdetector/branch/main/graph/badge.svg?token=9CMJ0HZA6B)](https://codecov.io/gh/haevg-rz/dockerdetector)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=haevg-rz_dockerdetector&metric=alert_status)](https://sonarcloud.io/dashboard?id=haevg-rz_dockerdetector)
[![Go Report Card](https://goreportcard.com/badge/github.com/haevg-rz/dockerdetector)](https://goreportcard.com/report/github.com/haevg-rz/dockerdetector)
[![Go Doc](https://godoc.org/github.com/haevg-rz/dockerdetector?status.svg)](http://godoc.org/github.com/haevg-rz/dockerdetector)
[![Known Vulnerabilities](https://snyk.io/test/github/haevg-rz/dockerdetector/badge.svg?targetFile=go.mod)](https://snyk.io/test/github/haevg-rz/dockerdetector?targetFile=go.mod)

## Intro

This package use cgroup to determine if we run as a docker container. 
It runs only under linux, windows und macOS will always returning false.

This package can also creates a static ID (Docker ID) and a salted Id (Docker Protected ID) encoded as a hex string.
Used cryptographic primitives are SHA256 and HMAC-SHA256.

## Usage

`go get -u github.com/haevg-rz/dockerdetector`

```go
package main

import (
	"fmt"

	"github.com/haevg-rz/dockerdetector"
)

func main() {
	isDocker, err := dockerdetector.IsRunningInContainer()
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("Run in Docker:", isDocker)

	if isDocker {
		id, err := dockerdetector.CreateIDFromDocker()
		if err != nil {
			fmt.Println("Error:", err)
		}
		fmt.Println("Docker ID:", id)
	}

	if isDocker {
		id, err := dockerdetector.CreateProtectedFromDockerID("My Salt")
		if err != nil {
			fmt.Println("Error:", err)
		}
		fmt.Println("Docker Protected ID:", id)
	}
}

```

## Run in docker

```
$ docker run -i -t --rm golang

root@110c80b1eb43:/go# go install github.com/haevg-rz/dockerdetector/cmd/dockerdetector@latest
root@110c80b1eb43:/go# dockerdetector
Run in Docker: true
Docker ID: 3ae2bacb803925e4e1be937b8e4609d138abcd6cf61165d9a57a48823107ad56
Docker Protected ID: 4dbeebcbcf17bf073343360ee3db67f7fb31214d661d61a2a3b03abe83c9ac3c
```
