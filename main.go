/*
TODO:
somehow set up a pubsub system with the RunCmd only firing when the chan has "CREATE"
createChan := make(chan string)

if event.Op.String() == "CREATE" {
	createChan <- "CREATE"
}
then run veai in a seperate thread to allow it to fully exec.


**/
package main

import (
	"brendisurfs/dropbox-sr-watcher/enhance"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/fsnotify/fsnotify"
)

// Config - struct to define the config file.
type Config struct {
	InDir  string `json:"in_dir"`
	OutDir string `json:"out_dir"`
}

var config Config

func readConfig() {

	file, err := os.ReadFile("./config.json")
	if err != nil {
		log.Fatal("could not find config file", err)
	}

	if err = json.Unmarshal(file, &config); err != nil {
		log.Fatal("could not unmarshal config")
	}

	fmt.Println(config.InDir)

}

func main() {
	readConfig()

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer watcher.Close()

	done := make(chan bool)

	go func() {
		for {
			select {
			// check for events.
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				enhance.RunCmd(config.InDir, "mp4", config.OutDir, event)

				// listen for the create, start doing the create stuff.

			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error: ", err)
			}
		}
	}()

	err = watcher.Add(config.InDir)
	if err != nil {
		log.Fatal(err)
	}
	<-done
}
