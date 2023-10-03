package utils

import (
	"log"
	"os"
	"testing"
)

type Utils struct {
	String string
	Int    int  `yaml:"int"`
	Bool   bool `yaml:"bool"`
}

var TestUtils Utils

func TestMain(m *testing.M) {
	//setUp()
	//log.Println("Do stuff BEFORE the tests!")
	exitVal := m.Run()
	//log.Println("Do stuff AFTER the tests!")
	//tearDown()
	os.Exit(exitVal)
}

func TestToJSON(t *testing.T) {
	var got interface{} = ToJSON(TestUtils)

	_, ok := got.(string)
	if ok == false {
		log.Fatalf("got is not what i want !")
	}
}
