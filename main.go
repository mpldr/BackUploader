package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sync"

	"./controller"
	"./display"
	"github.com/bigkevmcd/go-configparser"
	"github.com/wsxiaoys/terminal/color"
	"golang.org/x/sync/semaphore"
)

var (
	Wg           sync.WaitGroup
	buildVersion = "not built using Makefile, version unknown!"
)

func init() {
	log.SetOutput(os.Stderr)
}

func main() {
	// read and parse configuration
	fmt.Print("Reading Config........")
	conf, err := configparser.NewConfigParserFromFile("config.ini")
	if err != nil {
		color.Println("@{^r*!}ERROR")
		log.Fatal(err)
	}

	// get limit for concurrent jobs
	setclimit := int64(10)
	setclimitc, err := conf.GetInt64("Settings", "Concurrent")
	if err != nil {
		log.Println("no concurrency limit set. using default (10).")
	} else {
		setclimit = setclimitc
	}
	// get limit for concurrent packing
	setpclimit := int64(10)
	setpclimitc, err := conf.GetInt64("Settings", "Packing")
	if err != nil {
		log.Println("no packing limit set. using default (2).")
	} else {
		setpclimit = setpclimitc
	}
	// get limit for concurrent paring
	setprlimit := int64(10)
	setprlimitc, err := conf.GetInt64("Settings", "Paring")
	if err != nil {
		log.Println("no paring limit set. using default (3).")
	} else {
		setprlimit = setprlimitc
	}
	// get limit for concurrent uploading
	setulimit := int64(10)
	setulimitc, err := conf.GetInt64("Settings", "Uploading")
	if err != nil {
		log.Println("no upload limit set. using default (5).")
	} else {
		setulimit = setulimitc
	}

	// get commands...
	// ...for packing
	controller.PackCmd, err = conf.Get("Commands", "Packing")
	if err != nil {
		log.Fatal(err)
	}
	// ...for paring
	controller.ParCmd, err = conf.Get("Commands", "Paring")
	if err != nil {
		log.Fatal(err)
	}
	// ...for uploading
	controller.UpldCmd, err = conf.Get("Commands", "Uploading")
	if err != nil {
		log.Fatal(err)
	}
	// ...executed...
	controller.Executor, err = conf.Get("Commands", "Executor")
	if err != nil {
		log.Fatal(err)
	}
	// ...and its options
	controller.ExecOpt, err = conf.Get("Commands", "ExecOpt")
	if err != nil {
		log.Fatal(err)
	}
	// get folders
	// ...for uploads
	controller.Path, err = conf.Get("Directories", "Upload")
	if err != nil {
		log.Fatal(err)
	}
	// ...for uploading
	controller.FailPath, err = conf.Get("Directories", "Failed")
	if err != nil {
		log.Fatal(err)
	}
	// ...for uploading
	controller.SuccPath, err = conf.Get("Directories", "Finished")
	if err != nil {
		log.Fatal(err)
	}

	color.Println("@{g}DONE@{|}")

	// read value from config
	fmt.Print("Listing Directories...")
	path, err := conf.Get("Directories", "Upload")
	if err != nil {
		color.Println("@{^r*!}ERROR")
		log.Fatal(err)
	}
	// try to read subdirectories
	subdirs, err := ioutil.ReadDir(path)
	if err != nil {
		color.Println("@{^r*!}ERROR")
		log.Fatal(err)
	}
	color.Println("@{g}DONE@{|}")

	var dirs = make([]string, 0)
	for _, dir := range subdirs {
		if dir.IsDir() {
			dirs = append(dirs, dir.Name())
		}
	}

	color.Println("Found a total of@{!}", len(dirs), "@{|}directories.")
	if len(dirs) == 0 {
		fmt.Println("Nothing to do.")
		return
	}

	//TODO: Context controller
	fmt.Println("a maximum of", setclimit, "jobs is performed at a time")
	controller.Running = semaphore.NewWeighted(setclimit)
	fmt.Println("a maximum of", setpclimit, "jobs is packed at a time")
	controller.Packing = semaphore.NewWeighted(setpclimit)
	fmt.Println("a maximum of", setprlimit, "jobs is pared at a time")
	controller.Paring = semaphore.NewWeighted(setprlimit)
	fmt.Println("a maximum of", setulimit, "jobs is uploaded at a time")
	controller.Uploading = semaphore.NewWeighted(setulimit)
	controller.Path = path

	for _, nextdir := range dirs {
		controller.Running.Acquire(controller.Contxt, 1)
		displayId := display.Add("@{y}idle", nextdir)
		go controller.Start(nextdir, displayId, &Wg)
		Wg.Add(1)
	}
	Wg.Wait()
	fmt.Println("Group completed")
}
