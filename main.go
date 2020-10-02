package main

import (
	"log"

	"github.com/OpenPixelSystems/IAW/lib/handlers/cmdhandler"
	"github.com/OpenPixelSystems/IAW/lib/handlers/filehandler"
	"github.com/OpenPixelSystems/IAW/lib/triggers/ioreader"
	"github.com/OpenPixelSystems/IAW/lib/wrapper"
)

func main() {
	log.Println("Running IAW")

	app := wrapper.CreateNewWrapper("./tests/test_app")

	stdoutTrigger := ioreader.CreateNewReaderTrigger(&app.Stdout, "error")
	app.AddNewTrigger(&stdoutTrigger)

	stderrTrigger := ioreader.CreateNewReaderTrigger(&app.Stderr, "error")
	app.AddNewTrigger(&stderrTrigger)

	tempHandler := filehandler.CreateNewFileHander("/sys/class/hwmon/hwmon0/temp1_input")
	app.AddNewHandler(tempHandler)

	cmdHandler := cmdhandler.CreateNewCmdHandler("dmesg")
	app.AddNewHandler(cmdHandler)

	go app.RunCommand()

	app.StartTriggers()
	for !app.CmdDone {
		app.CheckTriggers()
	}

	app.WaitForHandlers()
}
