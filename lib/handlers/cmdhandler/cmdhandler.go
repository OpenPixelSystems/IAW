package cmdhandler

import (
	"io/ioutil"
	"log"
	"os/exec"
	"path"
	"sync"
)

type CmdHandler struct {
	cmdStr  string
	argsStr []string
}

func CreateNewCmdHandler(cmd string, args ...string) *CmdHandler {
	var newCh CmdHandler
	newCh.cmdStr = cmd
	newCh.argsStr = args

	return &newCh
}

func (c *CmdHandler) RunHandler(wg *sync.WaitGroup, outdir string) {
	defer wg.Done()
	cmd := exec.Command(c.cmdStr, c.argsStr...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Failed to run cmd: %s", err.Error())
	}

	dest := path.Base(c.cmdStr)
	fullDest := path.Clean(outdir + "/" + dest)

	log.Printf("Storing output of %s to %s", c.cmdStr, fullDest)
	err = ioutil.WriteFile(fullDest, output, 0644)
	if err != nil {
		log.Fatalf("Failed to write file: %s", fullDest)
	}
}
