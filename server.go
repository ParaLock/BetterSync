package main

import (
    "log"
    "net"
    "net/http"
)

func start_server() {
    
    fs := http.FileServer(http.Dir("static"))

    http.Handle("/data", fs)
    listener, err := net.Listen("tcp", ":0")
    if err != nil {
        log.Fatal(err)
    }
    defer listener.Close()

    port := listener.Addr().(*net.TCPAddr).Port
    log.Printf("Listening on :%d...", port)

    if err := http.Serve(listener, nil); err != nil {
        log.Fatal(err)
    }
}