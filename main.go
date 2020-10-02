package main

import (
	"log"

	"./lib/handlers/filehandler"
	"./lib/triggers/ioreader"
	"./lib/wrapper"
)

func main() {
	log.Println("Running IAW")

	app := wrapper.CreateNewWrapper("./tests/test_app")

	stdoutTrigger := ioreader.CreateNewReaderTrigger(&app.Stdout, "error")
	app.AddNewTrigger(&stdoutTrigger)

	stderrTrigger := ioreader.CreateNewReaderTrigger(&app.Stderr, "error")
	app.AddNewTrigger(&stderrTrigger)

	tempHandler := filehandler.CreateNewFileHander("main.go")
	app.AddNewHandler(tempHandler)

	go app.RunCommand()

	app.StartTriggers()
	for !app.CmdDone {
		app.CheckTriggers()
	}

}
