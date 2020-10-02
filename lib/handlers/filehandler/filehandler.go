package filehandler

import (
	"io/ioutil"
	"log"
	"path"
	"sync"
)

type FileHandler struct {
	file string
}

func CreateNewFileHander(filepath string) *FileHandler {
	var newFh FileHandler
	newFh.file = filepath
	return &newFh
}

func (f *FileHandler) copyFile(outdir string) {
	input, err := ioutil.ReadFile(f.file)
	if err != nil {
		log.Fatal(err)
		return
	}

	dest := path.Base(f.file)
	fullDest := path.Clean(outdir + "/" + dest)

	log.Printf("Copying to %s", fullDest)
	err = ioutil.WriteFile(fullDest, input, 0644)
	if err != nil {
		log.Fatalf("Error creating: %s", dest, err)
		return
	}
}

func (f *FileHandler) RunHandler(wg *sync.WaitGroup, outdir string) {
	defer wg.Done()
	f.copyFile(outdir)
}
