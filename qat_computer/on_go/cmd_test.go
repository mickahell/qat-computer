package on_go

import (
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	//setUp()
	//log.Println("Do stuff BEFORE the tests!")
	exitVal := m.Run()
	//log.Println("Do stuff AFTER the tests!")
	//tearDown()
	os.Exit(exitVal)
}

func TestBashCMD(t *testing.T) {
	var got string = BashCMD("echo toto")

	if got != "toto" {
		log.Fatalf("got is not what i wanted !")
	}
}
