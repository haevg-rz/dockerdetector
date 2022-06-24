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
