package builder

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/archive"
	"github.com/go-git/go-git/v5"
)

func CloneRepo(ctx context.Context) (string, error) {
	const clonePath = "cloned-repo"
	_, err := git.PlainClone(clonePath, false, &git.CloneOptions{
		URL: "https://github.com/Gromitmugs/hello-world-docker",
	})
	if err != nil {
		return "", err
	}
	return clonePath, nil
}

func BuildImage(ctx context.Context, clonePath string) error {
	dockerClient, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return err
	}
	tar, err := archive.TarWithOptions(clonePath, &archive.TarOptions{})
	if err != nil {
		return err
	}

	dockerRegistryUserId := "gromitmugs"
	opts := types.ImageBuildOptions{
		Dockerfile: "Dockerfile",
		Tags:       []string{dockerRegistryUserId + clonePath},
		Remove:     true,
	}
	res, err := dockerClient.ImageBuild(ctx, tar, opts)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	// the program should process each line of the building command right until
	// the execution ends, otherwise the program will exit before the execution finishes
	scanner := bufio.NewScanner(res.Body)
	for scanner.Scan() {
	}

	return nil
}

func printDockerResp(rd io.Reader) error {
	var lastLine string
	type ErrorLine struct {
		Error       string `json:"error"`
		ErrorDetail struct {
			Message string `json:"message"`
		} `json:"errorDetail"`
	}
	scanner := bufio.NewScanner(rd)
	for scanner.Scan() {
		lastLine = scanner.Text()
		fmt.Println(scanner.Text())
	}

	errLine := &ErrorLine{}
	json.Unmarshal([]byte(lastLine), errLine)
	if errLine.Error != "" {
		return errors.New(errLine.Error)
	}

	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}
