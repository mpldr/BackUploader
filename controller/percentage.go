package controller

import (
	"bytes"
	"regexp"
	"time"

	"github.com/poldi1405/BackUploader/display"
)

var (
	percentageRegEx = regexp.MustCompile(`(?P<perc>[0-9.,]+) ?\%`)
)

func percentage(stdout, stderr *bytes.Buffer, displayId int, DC *display.DisplayController, action string, cont chan bool) {
	_, running := <-cont
	for running {

		stout := stdout.Bytes()
		sterr := stderr.Bytes()
		resultOut := percentageRegEx.FindAllStringSubmatch(string(stout), -1)
		resultErr := percentageRegEx.FindAllStringSubmatch(string(sterr), -1)

		perc := ""
		if len(resultOut) > 0 {
			perc = resultOut[len(resultOut)-1][1]
		}
		if len(resultErr) > 0 {
			perc = resultErr[len(resultErr)-1][1]
		}
		if perc != "" {
			DC.Update(displayId, action+" "+perc+"%")
		}
		time.Sleep(500 * time.Millisecond)
		select {
		case <-cont:
		default:
			running = true
		}
	}
}
