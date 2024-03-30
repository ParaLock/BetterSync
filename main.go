package main

import (
    // "io"
    // "os"
    // "os/exec"
    // "time"
    // "github.com/fsnotify/fsnotify"
    // "gopkg.in/yaml.v2"
    // "io/ioutil"
    // "log"
    // "net"
    // "net/http"
)

import "usync/synchronizers"

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

func main() {

	gitSynchronizer	:= &synchronizers.Git{}
    gitSynchronizer.Execute()





	// exePath, err := os.Executable()
	// if err != nil {
	// 	panic(err)
	// }

	// exeDir := filepath.Dir(exePath)
	// relativePath := "config/settings.json"
	// absolutePath := filepath.Join(exeDir, relativePath)
	// fmt.Printf("Absolute path: %s\n", absolutePath)






    // rcloneSynchronizer  := &synchronizers.RClone{}
    // var args Args
    // arg.MustParse(&args)
    // fmt.Println("Mode:", args.Mode)
    // fmt.Println("BootstrapMethod:", args.BootstrapMethod)
    // var args Args
    // arg.MustParse(&args)

    // fmt.Println(args.mode)
    // fmt.Println(args.bootstrapMethod)

	// config := loadConfig("config.yaml")

    // for _, sync := range config.Syncs {
    //     fmt.Printf("Setting up sync: %s\n", sync.Name)
    // }

    // fmt.Println("Setup complete. Press 'Enter' to exit.")
    // fmt.Scanln()
}

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

