package ioreader

import (
	"bufio"
	"io"
	"log"
	"strings"
)

type ReaderTrigger struct {
	reader  *io.Reader
	scanner *bufio.Scanner
	trigger string
	count   int

	verbose  bool
	triggerd bool
}

func CreateNewReaderTrigger(reader *io.Reader, trigger string, verbose bool) ReaderTrigger {
	log.Println("Creating new IO Reader Trigger")
	var newTrigger ReaderTrigger
	newTrigger.reader = reader
	newTrigger.trigger = trigger
	newTrigger.count = 0
	newTrigger.scanner = bufio.NewScanner(*newTrigger.reader)
	newTrigger.verbose = verbose

	return newTrigger
}

func (r *ReaderTrigger) Triggerd() bool {
	return r.triggerd
}

func (r *ReaderTrigger) Trigger() bool {
	log.Printf("Running trigger on %s", r.trigger)
	for r.scanner.Scan() {
		buffer := r.scanner.Text()
		if r.verbose {
			log.Println(buffer)
		}
		if strings.Contains(buffer, r.trigger) {
			r.triggerd = true
			return true
		}
	}
	return false
}

func (r *ReaderTrigger) ClearTrigger() {
	r.triggerd = false
}
