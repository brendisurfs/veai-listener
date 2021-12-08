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
				// listen for the create, start doing the create stuff.
				if event.Op.String() == "CREATE" {
					enhance.ExecEnhance(config.InDir, "mp4", config.OutDir)
				}

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