package docs

import (
	"log"
	"testing"
)

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
