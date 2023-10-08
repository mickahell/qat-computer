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

func TestStructCMDs(t *testing.T) {
	var got interface{} = BashCMD("echo toto")

	_, ok := got.(*Cmds)

	if ok == false {
		log.Fatalln("got is not a Cmds !")
	}
}

func TestBashCMD(t *testing.T) {
	var got *Cmds = BashCMD("echo toto")

	if got.Stdout != "toto" {
		log.Fatalf("got is not what i wanted !")
	}
}
