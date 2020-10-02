package wrapper

import (
	"io"
	"log"
	"os"
	"os/exec"
	"path"
	"sync"
	"time"

	"github.com/OpenPixelSystems/IAW/lib/handlers"
	"github.com/OpenPixelSystems/IAW/lib/triggers"
)

type Wrapper struct {
	cmdStr  string
	argsStr []string
	Cmd     *exec.Cmd

	Stdout io.Reader
	Stderr io.Reader

	triggers []trigger.Trigger
	handlers []handler.Handler

	wg sync.WaitGroup

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

func (w *Wrapper) AddNewHandler(nh handler.Handler) {
	w.handlers = append(w.handlers, nh)
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

func (w *Wrapper) RunHandlers() {

	outDir := "output/" + time.Now().Format("2006-01-02_15:04:05") + "-" + path.Base(w.cmdStr)
	os.MkdirAll(outDir, 0755)
	for _, h := range w.handlers {
		w.wg.Add(1)
		go h.RunHandler(&w.wg, outDir)
	}
}

func (w *Wrapper) CheckTriggers() bool {
	for _, t := range w.triggers {
		if t.Triggerd() {
			log.Println("Triggerd!!")
			t.ClearTrigger()
			w.RunHandlers()

			go t.Trigger() //Re-arm trigger that finished
			return true
		}
	}
	return false
}

func (w *Wrapper) RunCommand() {
	log.Printf("Running: %s %s", w.cmdStr, w.argsStr)
	if err := w.Cmd.Run(); err != nil {
		log.Fatalf("Wrapper: Failed to run command: %s", err.Error())
	}
	w.CmdDone = true
}

func (w *Wrapper) WaitForHandlers() {
	w.wg.Wait()
}
