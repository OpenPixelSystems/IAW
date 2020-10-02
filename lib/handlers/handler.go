package handler

import (
	"sync"
)

type Handler interface {
	RunHandler(wg *sync.WaitGroup, outdir string)
}
