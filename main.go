package main

import (
"fmt"
"os"
"net"
"bufio"
"strings" // only needed below for sample processing
"encoding/json"
"os/signal"
"syscall"
)

func main() {

    ch := make(chan os.Signal)
    msgs := make(chan string)
    signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)    

    fmt.Println("Launching server...")

    // listen on all interfaces
    addr := fmt.Sprintf("%s:%s", os.Getenv("SYSLOG_HOST"), os.Getenv("SYSLOG_PORT"))
    ln, _ := net.Listen("tcp", addr)

    // accept connection on port
    conn, _ := ln.Accept()

    go func(c net.Conn, msgs chan string){
        // will listen for message to process ending in newline (\n)
        message, _ := bufio.NewReader(c).ReadString('\n')
        msgs <- message
    }(conn, msgs)
 
    // run loop forever (or until ctrl-c)
    for {
        select {
        case <- ch:
            os.Exit(0)
        case message := <- msgs:
            // sample process for string received
            tuple := strings.Split(message, "@cee:")
            if len(tuple) == 2 {
                var m map[string]interface{}
                err := json.Unmarshal([]byte(tuple[1]), &m)
                if err != nil {
                  fmt.Printf("Json Parsing Failed: %s \n", err.Error())
                  continue
                }
                fmt.Println(m["time"])
                if t, ok := m["time"];ok {
                    fmt.Println(t) 
                }
            }
        }
    }

}
