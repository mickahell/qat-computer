package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"qat-cmd/container"
	"qat-cmd/docs"
	"qat-cmd/flags"

	"github.com/docker/docker/client"
)

func main() {
	flags.StartOptions()
	client, err := client.NewClientWithOpts()
	if err != nil {
		fmt.Printf("Unable to create docker client: %s", err)
		os.Exit(1)
	}

	if flags.Pullflag {
		err := container.PullImage(client, flags.FullImageNameflag, flags.ImageTagflag)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
	}

	if flags.Initflag {
		path, err := os.Getwd()
		if err != nil {
			log.Fatalln("Abs path doesn't exist !")
			os.Exit(1)
		}

		container_id, err := container.ListContainer(client, docs.ImageName)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
		if container_id == nil {
			container_id, err = container.RunContainer(
				client,
				docs.RegistryImage,
				docs.ContainerName,
				filepath.Join(path, flags.ComputeVolumeflag),
				filepath.Join(path, flags.ConfVolumeflag),
			)
			if err != nil {
				log.Println(err)
			}
		}
		err = container.StartContainer(client, *container_id)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
	}

	if flags.Stopflag {
		err = container.StopAndRemoveContainer(client, docs.ContainerName, flags.StopRemoveflag)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
	}

	if flags.Computerflag {
		container_id, err := container.ListContainer(client, docs.ImageName)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
		err = container.ExecContainer(client, *container_id, flags.ComputerArgsflag)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
	}

}
