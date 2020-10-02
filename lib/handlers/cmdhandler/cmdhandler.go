package cmdhandler

import (
	"io/ioutil"
	"log"
	"os/exec"
	"path"
)

type CmdHandler struct {
	cmdStr  string
	argsStr []string

	Cmd *exec.Cmd
}

func CreateNewCmdHandler(cmd string, args ...string) *CmdHandler {
	var newCh CmdHandler
	newCh.cmdStr = cmd
	newCh.argsStr = args
	newCh.Cmd = exec.Command(newCh.cmdStr, newCh.argsStr...)

	return &newCh
}

func (c *CmdHandler) RunHandler(outdir string) {
	output, err := c.Cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Failed to run cmd: %s", err)
	}

	dest := path.Base(c.cmdStr)
	fullDest := path.Clean(outdir + "/" + dest)

	err = ioutil.WriteFile(fullDest, output, 0644)
	if err != nil {
		log.Fatalf("Failed to write file: %s", fullDest)
	}
}
