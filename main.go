package main

import (
    // "io"
    "os"
    // "os/exec"
    // "time"
    // "github.com/rjeczalik/notify"
    // "gopkg.in/yaml.v2"
    // "io/ioutil"
    // "log"
    // "net"
    // "net/http"
    "strings"
	"log"
    "github.com/fsnotify/fsnotify"
    "path/filepath"
	// "os"
	// "path/filepath"
)

// import "usync/synchronizers"
// import "usync/monitoring"

type Config struct {
    Syncs []SyncConfig `yaml:"syncs"`
}

type SyncConfig struct {
    Name      string      `yaml:"name"`
    Type      string      `yaml:"type"`
    OnFailure string      `yaml:"on_failure"`
    RunWhen   struct {
        DirChanged struct {
            Events string `yaml:"events"`
        } `yaml:"dir_changed"`
        Timer struct {
            Interval int `yaml:"interval"`
        } `yaml:"timer"`
    } `yaml:"run_when"`
    Remotes map[string]string `yaml:"remotes"`
    Local  string `yaml:"local"`
}

type BootstrapArgs struct {
    
}

type Args struct {
    Mode             string `arg:"positional"`
	BootstrapMethod  string `arg:"--bootstrap-method"`
}

// func watch(dir string, allowedEvents []string, watcher fsnotify.Watcher) {

//     err := watcher.Add(sync.Local)
//     watcher, err := fsnotify.NewWatcher()

//     if err != nil {
//         panic(fmt.Sprintf("error initializing watcher: %v", err))
//     }
//     if err != nil {
//         panic(fmt.Sprintf("error adding directory to watcher: %v", err))
//     }
//     defer watcher.Close()

//     go func() {
//         for {
//             select {
//                 case event, ok := <-watcher.Events:

//                     if !ok {
//                         return
//                     }
//                     fmt.Printf("Event: %s\n", event)

//                 case err, ok := <-watcher.Errors:
                    
//                     if !ok {
//                         return
//                     }

//                     fmt.Printf("Watcher error: %v\n", err)
//             }
//         }
//     }()
    
// }

// func callback(event fsnotify.Event) {
    
//     fmt.Println("My event")
//     fmt.Println(event)
// }   

func main() {

    watcher, err := fsnotify.NewWatcher()
    if err != nil {
        log.Fatal(err)
    }
    defer watcher.Close()

    go func() {
        for {
            select {
            case event, ok := <-watcher.Events:
                if !ok {
                    return
                }
                if event.Op&fsnotify.Create == fsnotify.Create {
					fi, err := os.Stat(event.Name)
					if err == nil && fi.IsDir() {
						addWatcher(event.Name, watcher)
					}
				}
                log.Println("event:", event)
     
            case err, ok := <-watcher.Errors:
                if !ok {
                    return
                }
                log.Println("error:", err)
            }
        }
    }()

    err = watcher.Add("/home/linuxdev/SyncTest")
    if err != nil {
        log.Fatal(err)
    }

    <-make(chan struct{})
}

func addWatcher(path string, watcher *fsnotify.Watcher) error {
	
    //Skip .git or there will be an infinite loop 
    if(strings.Contains(path, ".git")) {
        return nil
    }
    
    err := watcher.Add(path)
	if err != nil {
		return err
	}
	return filepath.Walk(path, func(path string, fi os.FileInfo, err error) error {
		if fi.IsDir() {
			return watcher.Add(path)
		}
		return nil
	})
}

// func main() {

//     c := make(chan notify.EventInfo, 1)

// 	if err := notify.Watch("/home/linuxdev/dev/SyncTest", c, notify.Remove); err != nil {
// 		fmt.Println(err)
// 	}
// 	defer notify.Stop(c)

// 	ei := <-c
// 	fmt.Println("Got event:", ei)
// }

// func loadConfig(path string) Config {
//     var config Config
//     data, err := ioutil.ReadFile(path)
//     if err != nil {
//         panic(fmt.Sprintf("cannot read config file: %v", err))
//     }
//     err = yaml.Unmarshal(data, &config)
//     if err != nil {
//         panic(fmt.Sprintf("cannot parse config file: %v", err))
//     }
//     return config
// }

