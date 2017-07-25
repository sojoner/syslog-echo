package main

import (
"fmt"
"os"
"net"
"bufio"
"strings" // only needed below for sample processing
"encoding/json"
)

func handleRequest(conn net.Conn){
    for{
        message, err := bufio.NewReader(conn).ReadString('\n')
        if err != nil {
            fmt.Printf("Error:: %s \n", err.Error())
            break
        }
        // sample process for string received
        tuple := strings.Split(message, "@cee:")
        if len(tuple) == 2 {
            var m map[string]interface{}
            err := json.Unmarshal([]byte(tuple[1]), &m)
            if err != nil {
              fmt.Printf("Json Parsing Failed: %s \n", err.Error())
              continue
            }
            if t, ok := m["time"];ok {
                fmt.Println(t) 
            }   
        }   
    }
    
}

func main() {

    fmt.Println("Launching server...")

    // listen on all interfaces
    addr := fmt.Sprintf("%s:%s", os.Getenv("SYSLOG_HOST"), os.Getenv("SYSLOG_PORT"))
    ln, err := net.Listen("tcp", addr)
    
    if err != nil {
        fmt.Println("Error listening:", err.Error())
        os.Exit(1)
    }
    // Close the listener when the application closes.
    defer ln.Close()

    // run loop forever (or until ctrl-c)
    for {
        // Listen for an incoming connection.
        conn, err := ln.Accept()
        if err != nil {
            fmt.Println("Error accepting: ", err.Error())
            os.Exit(1)
        }
        // Handle connections in a new goroutine.
        go handleRequest(conn)
    }
}
