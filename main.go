package main

import (
	"log"

	"./lib/triggers/reader"
	"./lib/wrapper"
)

func main() {
	log.Println("Running IAW")

	app := wrapper.CreateNewWrapper("./tests/test_app")

	stdoutTrigger := ioreader.CreateNewReaderTrigger(&app.Stdout, "error")
	app.AddNewTrigger(&stdoutTrigger)

	stderrTrigger := ioreader.CreateNewReaderTrigger(&app.Stderr, "error")
	app.AddNewTrigger(&stderrTrigger)

	go app.RunCommand()

	app.StartTriggers()
	for !app.CmdDone {
		app.CheckTriggers()
	}

}
