package main

import (
	"gopkg.in/fsnotify.v1"
	"log"
)

func main() {
	log.SetFlags(0)
	log.SetPrefix("🍣  ")
	log.Println("started")

	// create watcher
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
	    log.Fatal(err)
	}
	defer watcher.Close()

	doneCh := make(chan bool)

	// configure dirs to watch
	err = watcher.Add("_work/foo")
	if err != nil {
	    log.Fatal(err)
	}
	err = watcher.Add("_work/bar")
	if err != nil {
	    log.Fatal(err)
	}

	// start event receiver
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				log.Println("event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Println("modified file:", event.Name)
				}
			case err := <-watcher.Errors:
				log.Println("error:", err)
			}
		}
	}()

	<-doneCh
}
