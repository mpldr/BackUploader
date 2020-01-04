package display

import (
	"context"

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
)

func Add(initstatus string, givenname string) int {
	status = append(status, initstatus)
	name = append(name, givenname)
	if maxlength < len(givenname) {
		maxlength = len(givenname)
	}
	if !free {
		terminal.Stdout.Nl(1)
	}
	return len(name) - 1
}

func Update(id int, newstatus string) {
	status[id] = newstatus
	show()
}

func show() {
	if err := writing.Acquire(ctx, 1); err != nil {
		return
	}
	defer writing.Release(1)

	if !free {
		terminal.Stdout.Up(len(name))
	} else {
		free = false
		terminal.Stdout.Nl(2)
	}

	for i := 0; i < len(name); i++ {
		dotc := maxlength - len(name[i])
		dots := "..."
		for j := 0; j < dotc; j++ {
			dots += "."
		}
		terminal.Stdout.ClearLine().Colorf(name[i] + dots + status[i]).Nl().Left(len(name[i] + dots + status[i]))
	}

}
