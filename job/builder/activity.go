package builder

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"os"
	"os/exec"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/archive"
	"github.com/go-git/go-git/v5"
	"github.com/google/uuid"
)

const imageTag = "docker-image"

func CloneRepo(ctx context.Context, url string) (string, error) {
	clonePath := uuid.NewString()
	if _, err := git.PlainClone(clonePath, false, &git.CloneOptions{
		URL: url,
	}); err != nil {
		return "", err
	}

	return clonePath, nil
}

func RemoveClonedRepo(ctx context.Context, path string) error {
	return os.RemoveAll(path)
}

func DockerBuildImage(ctx context.Context, path string) (string, error) {
	dockerClient, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return "", err
	}
	tar, err := archive.TarWithOptions(path, &archive.TarOptions{})
	if err != nil {
		return "", err
	}
	res, err := dockerClient.ImageBuild(ctx,
		tar,
		types.ImageBuildOptions{
			Dockerfile: "Dockerfile",
			Tags:       []string{imageTag},
			Remove:     true,
		},
	)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	return readResponseLog(res.Body)
}

func KanikoBuildImage(ctx context.Context, path string) (string, error) {
	executorCmd := exec.Command("/kaniko/executor")
	var result bytes.Buffer

	executorCmd.Stdout = &result
	if err := executorCmd.Run(); err != nil {
		return "", err
	}
	return result.String(), nil
}

func readResponseLog(rd io.Reader) (string, error) {
	var log string
	var lastLine string

	scanner := bufio.NewScanner(rd)
	for scanner.Scan() {
		log += scanner.Text()
		lastLine = scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}

	type errorLine struct {
		Error       string `json:"error"`
		ErrorDetail struct {
			Message string `json:"message"`
		} `json:"errorDetail"`
	}

	errLine := &errorLine{}
	json.Unmarshal([]byte(lastLine), errLine)
	if errLine.Error != "" {
		return "", errors.New(errLine.Error)
	}
	return log, nil
}

func buildWaitUntilFinish(res *types.ImageBuildResponse) {
	// the program should process each line of the building command right until
	// the execution ends, otherwise the program will exit before the execution finishes
	scanner := bufio.NewScanner(res.Body)
	for scanner.Scan() {
	}
}
