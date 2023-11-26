package container

import (
	"context"
	"fmt"
	"qat-cmd/docs"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func ExecContainer(client *client.Client, containerID string, command string) error {
	config := types.ExecConfig{
		AttachStderr: true,
		AttachStdout: true,
		Tty:          true,
		Cmd:          strings.Fields(docs.BinaryName + " " + command),
	}

	execCreate, err := client.ContainerExecCreate(context.Background(), containerID, config)
	if err != nil {
		return err
	}

	res, er := client.ContainerExecAttach(
		context.Background(),
		execCreate.ID,
		types.ExecStartCheck{},
	)
	if er != nil {
		return err
	}

	err = client.ContainerExecStart(context.Background(), execCreate.ID, types.ExecStartCheck{})
	if err != nil {
		return err
	}

	for {
		content, _, err := res.Reader.ReadLine()
		if err != nil {
			break
		}
		fmt.Println(string(content))
	}

	return nil

}
