package container

import (
	"context"
	"io"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func PullImage(client *client.Client, imagename string, imagetag string) error {
	imageNameTag := imagename + ":" + imagetag
	reader, err := client.ImagePull(context.Background(), imageNameTag, types.ImagePullOptions{})
	if err != nil {
		return err
	}
	_, err = io.Copy(os.Stdout, reader)
	if err != nil {
		return err
	}

	return nil

}
