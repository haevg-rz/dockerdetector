package dockerdetector

import (
	"bufio"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"runtime"
	"strings"
)

// IsRunningInContainer check in cgroup if your are running in a docker container
func IsRunningInContainer() (bool, error) {
	if runtime.GOOS != "linux" {
		return false, errors.New("Works only with os linux")
	}

	file, err := os.DirFS("/proc/self").Open("cgroup")
	if err != nil {
		return false, err
	}
	defer file.Close()

	isDocker, _, err := isRunningInContainer(file)

	return isDocker, err
}

func CreateID() (string, error) {
	if runtime.GOOS != "linux" {
		return "", errors.New("Works only with os linux")
	}

	file, err := os.DirFS("/proc/self").Open("cgroup")
	if err != nil {
		return "", err
	}
	defer file.Close()

	return createID(file)
}

func createID(file fs.File) (string, error) {
	isDocker, id, err := isRunningInContainer(file)
	if err != nil {
		return "", err
	}

	if !isDocker {
		return "", errors.New("Not a docker container")
	}

	h := sha256.New()
	h.Write([]byte(id))

	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

func CreateProtectedID(salt string) (string, error) {
	if runtime.GOOS != "linux" {
		return "", errors.New("Works only with os linux")
	}

	file, err := os.DirFS("/proc/self").Open("cgroup")
	if err != nil {
		return "", err
	}
	defer file.Close()

	return createProtectedID(salt, file)
}

func createProtectedID(salt string, file fs.File) (string, error) {
	id, err := createID(file)
	if err != nil {
		return "", err
	}

	return protect(salt, id), nil
}

func protect(appID, key string) string {
	mac := hmac.New(sha256.New, []byte(key))
	mac.Write([]byte(appID))
	return hex.EncodeToString(mac.Sum(nil))
}

func isRunningInContainer(file fs.File) (bool, string, error) {
	r := bufio.NewReader(file)

	var line string
	var err error
	for {
		line, err = r.ReadString('\n')
		if err != nil && err != io.EOF {
			break
		}

		if strings.Contains(line, "docker") {
			split := strings.Split(line, "/")
			lastSegment := split[len(split)-1]
			return true, strings.TrimSpace(lastSegment), nil
		}

		if err != nil {
			break
		}
	}

	if err != io.EOF {
		return false, "", err
	}

	return false, "", nil
}
