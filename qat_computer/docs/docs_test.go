package docs

import (
	"log"
	"testing"
)

func TestDrawStart(t *testing.T) {
	DrawStart()
}

func TestGetVersion(t *testing.T) {
	var version interface{} = GetVersion()

	_, ok := version.(string)

	if ok == false {
		log.Fatalln("version is not a string !")
	}

	if version == "" {
		log.Fatalln("version is empty !")
	}
}

func TestGetAppName(t *testing.T) {
	var got interface{} = GetAppName()

	_, ok := got.(string)

	if ok == false {
		log.Fatalln("app name is not a string !")
	}

	if got == "" {
		log.Fatalln("app name is empty !")
	}
}

func TestGetConfigPathMan(t *testing.T) {
	var got interface{} = GetConfigPathMan()

	_, ok := got.(string)

	if ok == false {
		log.Fatalln("man config path is not a string !")
	}
}

func TestGetLogLevelMan(t *testing.T) {
	var got interface{} = GetLogLevelMan()

	_, ok := got.(string)

	if ok == false {
		log.Fatalln("man log level is not a string !")
	}
}

func TestGetComputePathMan(t *testing.T) {
	var got interface{} = GetComputePathMan()

	_, ok := got.(string)

	if ok == false {
		log.Fatalln("man compute path is not a string !")
	}
}
