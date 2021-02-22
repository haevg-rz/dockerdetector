# dockerdetector

[![Go](https://github.com/dhcgn/dockerdetector/actions/workflows/go.yml/badge.svg)](https://github.com/dhcgn/dockerdetector/actions/workflows/go.yml)
[![codecov](https://codecov.io/gh/dhcgn/dockerdetector/branch/main/graph/badge.svg?token=9CMJ0HZA6B)](https://codecov.io/gh/dhcgn/dockerdetector)

## Intro

This package use cgroup to determine if we run as a docker container. 
It runs only under linux, windows und macOS will always returning false.

## Run in docker

```
root@110c80b1eb43:/go# go install github.com/dhcgn/dockerdetector/cmd/dockerdetector@latest
root@110c80b1eb43:/go# dockerdetector
Run in Docker: true
```
