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
