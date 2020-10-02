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

	triggerd bool
}

func CreateNewReaderTrigger(reader *io.Reader, trigger string) ReaderTrigger {
	log.Println("Creating new Reader Trigger")
	var newTrigger ReaderTrigger
	newTrigger.reader = reader
	newTrigger.trigger = trigger
	newTrigger.count = 0
	newTrigger.scanner = bufio.NewScanner(*newTrigger.reader)

	return newTrigger
}

func (r *ReaderTrigger) Triggerd() bool {
	return r.triggerd
}

func (r *ReaderTrigger) Trigger() bool {
	log.Printf("Running trigger on %s", r.trigger)
	for r.scanner.Scan() {
		buffer := r.scanner.Text()
		log.Println(buffer)
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
