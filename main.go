package main

import (
    "fmt"
    "io"
    "os"
    "os/exec"
    "time"
    "github.com/fsnotify/fsnotify"
    "gopkg.in/yaml.v2"
    "io/ioutil"
	"synchronizsers/RClone"
)

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
    Remotes []string `yaml:"remotes"`
    Local  string `yaml:"local"`
}

func watch(sync SyncConfig, watcher) {

    err := watcher.Add(sync.Local)
    watcher, err := fsnotify.NewWatcher()

    if err != nil {
        panic(fmt.Sprintf("error initializing watcher: %v", err))
    }
    if err != nil {
        panic(fmt.Sprintf("error adding directory to watcher: %v", err))
    }
    defer watcher.Close()

    go func() {
        for {
            select {
                case event, ok := <-watcher.Events:

                    if !ok {
                        return
                    }
                    fmt.Printf("Event: %s\n", event)

                case err, ok := <-watcher.Errors:
                    
                    if !ok {
                        return
                    }

                    fmt.Printf("Watcher error: %v\n", err)
            }
        }
    }()
    
}

func main() {

	gitSynchronizer 	:= &synchronizers.Git{}
    rcloneSynchronizer  := &synchronizers.RClone{}

	config := loadConfig("config.yaml")


    for _, sync := range config.Syncs {
        fmt.Printf("Setting up sync: %s\n", sync.Name)

  

        if sync.RunWhen.Timer.Interval > 0 {
            ticker := time.NewTicker(time.Duration(sync.RunWhen.Timer.Interval) * time.Second)
            go func() {
                for range ticker.C {
                    if err := executeRcloneSync(sync); err != nil {
                        handleSyncError(sync, err)
                    }
                }
            }()
        }
    }

    fmt.Println("Setup complete. Press 'Enter' to exit.")
    fmt.Scanln()
}

func loadConfig(path string) Config {
    var config Config
    data, err := ioutil.ReadFile(path)
    if err != nil {
        panic(fmt.Sprintf("cannot read config file: %v", err))
    }
    err = yaml.Unmarshal(data, &config)
    if err != nil {
        panic(fmt.Sprintf("cannot parse config file: %v", err))
    }
    return config
}

func executeSync() {

}

func executeRcloneSync(sync SyncConfig) error {
    fmt.Printf("Executing rclone sync from %s to %s\n", sync.Local, sync.Remote)
    cmd := exec.Command("rclone", "bisync", sync.Local, sync.Remote)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    err := cmd.Run()
    if err != nil {
        return fmt.Errorf("rclone sync error: %w", err)
    }
    fmt.Println("Sync completed successfully")
    return nil
}

func handleSyncError(sync SyncConfig, err error) {
    fmt.Printf("Error during sync '%s': %v\n", sync.Name, err)
    if sync.OnFailure == "exit" {
        os.Exit(1)
    }
}