package container

import (
	"log"
	"qat-cmd/docs"
	"testing"

	"github.com/docker/docker/client"
)

func TestPullImage(t *testing.T) {
	client, err := client.NewClientWithOpts()
	if err != nil {
		log.Fatalln("Unable to create docker client.")
	}

	err = PullImage(client, docs.RegistryImage, docs.RegistryTag)
	if err != nil {
		log.Fatalln("Unable to pull : ", err)
	}
}
