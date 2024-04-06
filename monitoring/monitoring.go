package monitoring

import (
    "github.com/fsnotify/fsnotify"
	"fmt"
)

func Watch(dir string, cb func(eventType fsnotify.Event)) {

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Println(err)
	}

    err = watcher.Add(dir)
	if err != nil {
		fmt.Println(err)
	}

	defer watcher.Close()

	done := make(chan bool)
    
	go func() {
		for {
			select {
                case event, ok := <- watcher.Events:
                    
                    fmt.Println("test")

                    if !ok {
                        return
                    }
                    fmt.Println("event:", event)
                    
                    cb(event)

                case err, ok := <- watcher.Errors:
                    if !ok {
                        return
                    }
                    fmt.Println("error:", err)
			}
		}
	}()

	<-done
}