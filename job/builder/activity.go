package builder

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/archive"
	"github.com/go-git/go-git/v5"
)

func CloneRepo(ctx context.Context) (string, error) {
	const clonePath = "/hello-world"
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

	dockerRegistryUserId := "Gromitmugs"
	opts := types.ImageBuildOptions{
		Dockerfile: "Dockerfile",
		Tags:       []string{dockerRegistryUserId + clonePath},
		Remove:     true,
	}
	res, err := dockerClient.ImageBuild(ctx, tar, opts)
	if err != nil {
		return err
	}
	fmt.Println(res.Body)

	defer res.Body.Close()
	if err != nil {
		return nil
	}

	return nil
}
