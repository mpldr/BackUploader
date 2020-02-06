package display

import (
	"context"
	"sync"

	"github.com/wsxiaoys/terminal"
	"golang.org/x/sync/semaphore"
)

var (
	status    = make([]string, 0)
	name      = make([]string, 0)
	maxlength = -1
	ctx       = context.TODO()
	writing   = semaphore.NewWeighted(1)
	free      = true
	mapMutex  = sync.Mutex{}
)

func Add(initstatus string, givenname string) int {
	mapMutex.Lock()
	status = append(status, initstatus)
	name = append(name, givenname)
	mapMutex.Unlock()
	if maxlength < len(givenname) {
		maxlength = len(givenname)
	}
	if !free {
		terminal.Stdout.Nl(1)
	}
	return len(name) - 1
}

func Update(id int, newstatus string) {
	mapMutex.Lock()
	status[id] = newstatus
	mapMutex.Unlock()
	show()
}

func show() {
	if err := writing.Acquire(ctx, 1); err != nil {
		return
	}
	defer writing.Release(1)
	mapMutex.Lock()
	localName := name
	localStat := status
	mapMutex.Unlock()

	if !free {
		terminal.Stdout.Up(len(localName))
	} else {
		free = false
		terminal.Stdout.Nl(2)
	}

	for i := 0; i < len(localName); i++ {
		dotc := maxlength - len(localName[i])
		dots := "..."
		for j := 0; j < dotc; j++ {
			dots += "."
		}
		terminal.Stdout.ClearLine().Colorf(localName[i] + dots + localStat[i]).Nl().Left(len(localName[i] + dots + localStat[i]))
	}

}
