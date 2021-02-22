package dockerdetector

import (
	"bufio"
	"io"
	"io/fs"
	"os"
	"runtime"
	"strings"
)

// IsRunningInContainer check in cgroup if your are running in a docker container
func IsRunningInContainer() (bool, error) {
	if runtime.GOOS != "linux" {
		return false, nil
	}

	file, err := os.DirFS("/proc/self").Open("cgroup")
	if err != nil {
		return false, err
	}
	defer file.Close()

	return isRunningInContainer(file)
}

func isRunningInContainer(file fs.File) (bool, error) {
	r := bufio.NewReader(file)

	var line string
	var err error
	for {
		line, err = r.ReadString('\n')
		if err != nil && err != io.EOF {
			break
		}

		if strings.Contains(line, "docker") {
			return true, nil
		}

		if err != nil {
			break
		}
	}

	if err != io.EOF {
		return false, err
	}

	return false, nil
}
