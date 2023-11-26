package container

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

// Init container
func RunContainer(
	client *client.Client,
	imagename string,
	containername string,
	computevolume string,
	confvolume string,
) (*string, error) {
	binds := []string{
		"/sys/fs/cgroup:/sys/fs/cgroup:rw",
		computevolume + ":/etc/qat-computer/compute",
	}

	if confvolume != "" {
		binds = []string{
			"/sys/fs/cgroup:/sys/fs/cgroup:rw",
			confvolume + ":/etc/qat-computer/conf:rw",
			computevolume + ":/etc/qat-computer/compute:rw",
		}
	}
	// Configured hostConfig:
	// https://godoc.org/github.com/docker/docker/api/types/container#HostConfig
	hostConfig := &container.HostConfig{
		LogConfig: container.LogConfig{
			Type:   "json-file",
			Config: map[string]string{},
		},
		Privileged: true,
		Cgroup:     "hosts",
		Binds:      binds,
	}

	// Configuration
	// https://godoc.org/github.com/docker/docker/api/types/container#Config
	config := &container.Config{
		Image:       imagename,
		Hostname:    fmt.Sprintf("%s-qatcmd", containername),
		Healthcheck: &container.HealthConfig{},
	}

	// Creating the actual container.
	cont, err := client.ContainerCreate(
		context.Background(),
		config,
		hostConfig,
		nil,
		nil,
		containername,
	)

	if err != nil {
		return nil, err
	}

	return &cont.ID, nil
}

func StartContainer(client *client.Client, cont_id string) error {
	err := client.ContainerStart(context.Background(), cont_id, types.ContainerStartOptions{})
	if err != nil {
		return err
	}
	log.Printf("Container %s is starting...", cont_id[:12])
	json, err := client.ContainerInspect(context.Background(), cont_id)
	if err != nil {
		return err
	}

	for json.State.Health.Status == "starting" {
		time.Sleep(5 * time.Second)
		json, err = client.ContainerInspect(context.Background(), cont_id)
		if err != nil {
			return err
		}
	}
	if json.State.Health.Status == "healthy" {
		log.Printf("Container %s is healthy", cont_id[:12])
	} else {
		log.Printf("Container %s is not healthy", cont_id[:12])
	}

	return nil
}

// Stop and remove a container
func StopAndRemoveContainer(client *client.Client, containername string, remove bool) error {
	if err := client.ContainerStop(context.Background(), containername, container.StopOptions{}); err != nil {
		log.Printf("Unable to stop container %s", containername)
		return err
	}

	if remove {
		removeOptions := types.ContainerRemoveOptions{
			RemoveVolumes: true,
			Force:         true,
		}

		if err := client.ContainerRemove(context.Background(), containername, removeOptions); err != nil {
			log.Printf("Unable to remove container")
			return err
		}
	}

	return nil
}

func ListContainer(client *client.Client, imagename string) (*string, error) {

	containers, err := client.ContainerList(
		context.Background(),
		types.ContainerListOptions{All: true},
	)

	if len(containers) > 0 {
		for _, container := range containers {
			if strings.Contains(container.Image, imagename) {
				return &container.ID, nil
			}
		}
	}
	return nil, err
}
