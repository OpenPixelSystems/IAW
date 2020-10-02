package wrapper

import (
	"io"
	"log"
	"os/exec"

	"../triggers"
)

type Wrapper struct {
	cmdStr  string
	argsStr []string
	Cmd     *exec.Cmd

	Stdout io.Reader
	Stderr io.Reader

	triggers []trigger.Trigger

	CmdDone bool
}

func CreateNewWrapper(cmd string, args ...string) *Wrapper {
	var newWrapper Wrapper
	newWrapper.cmdStr = cmd
	newWrapper.argsStr = args
	newWrapper.Cmd = exec.Command(newWrapper.cmdStr, newWrapper.argsStr...)
	newWrapper.CmdDone = false

	newWrapper.Stdout, _ = newWrapper.Cmd.StdoutPipe()
	newWrapper.Stderr, _ = newWrapper.Cmd.StderrPipe()

	return &newWrapper
}

func (w *Wrapper) AddNewTrigger(nt trigger.Trigger) {
	w.triggers = append(w.triggers, nt)
}

func (w *Wrapper) StartTriggers() {
	/* Spawn triggers */
	for _, t := range w.triggers {
		go t.Trigger()
	}
}

func (w *Wrapper) CheckTriggers() bool {
	for _, t := range w.triggers {
		if t.Triggerd() {
			log.Println("Triggerd!!")
			t.ClearTrigger()
			return true
		}
	}
	return false
}

func (w *Wrapper) RunCommand() {
	log.Println("Running!")
	if err := w.Cmd.Run(); err != nil {
		log.Fatalf("Failed to run command: %s", err)
	}
	w.CmdDone = true
}
