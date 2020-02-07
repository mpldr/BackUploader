package display

import (
	"context"
	"sync"

	"github.com/wsxiaoys/terminal"
)

type DisplayController struct {
	status    []string
	name      []string
	maxlength int
	ctx       context.Context
	writing   sync.Mutex
	free      bool
	mapMutex  sync.Mutex
}

func (dc *DisplayController) Add(initstatus string, givenname string) int {
	dc.mapMutex.Lock()
	dc.status = append(dc.status, initstatus)
	dc.name = append(dc.name, givenname)
	dc.mapMutex.Unlock()
	if dc.maxlength < len(givenname) {
		dc.maxlength = len(givenname)
	}
	if !dc.free {
		terminal.Stdout.Nl(1)
	}
	return len(dc.name) - 1
}

func (dc *DisplayController) Update(id int, newstatus string) {
	dc.mapMutex.Lock()
	dc.status[id] = newstatus
	dc.mapMutex.Unlock()
	dc.show()
}

func (dc *DisplayController) show() {
	dc.writing.Lock()
	defer dc.writing.Unlock()

	dc.mapMutex.Lock()
	localName := dc.name
	localStat := dc.status
	maxl := dc.maxlength
	dc.mapMutex.Unlock()

	if !dc.free {
		terminal.Stdout.Up(len(localName))
	} else {
		dc.free = false
		terminal.Stdout.Nl(2)
	}

	for i := 0; i < len(localName); i++ {
		dotc := maxl - len(localName[i])
		dots := "..."
		for j := 0; j < dotc; j++ {
			dots += "."
		}
		terminal.Stdout.ClearLine().Colorf(localName[i] + dots + localStat[i]).Nl().Left(len(localName[i] + dots + localStat[i]))
	}

}
