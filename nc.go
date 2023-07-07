package main

import (
    "flag"
    "io"
    "log"
    "net"
    "os"
)

func main() {
    var (
        listen       bool
        target       string
        listenPort   int
        targetPort   int
    )

    flag.BoolVar(&listen, "l", false, "Listen mode")
    flag.StringVar(&target, "t", "", "Target hostname or IP address")
    flag.IntVar(&listenPort, "lp", 0, "Listen port")
    flag.IntVar(&targetPort, "tp", 0, "Target port")

    flag.Parse()

    if listen {
        if listenPort == 0 {
            log.Fatal("Please specify the listen port using -lp")
        }

        listenMode(listenPort)
    } else {
        if target == "" || targetPort == 0 {
            log.Fatal("Please specify the target and target port using -t and -tp")
        }

        connectMode(target, targetPort)
    }
}

func listenMode(port int) {
    listener, err := net.Listen("tcp", ":"+strconv.Itoa(port))
    if err != nil {
        log.Fatalf("Failed to start listener: %v", err)
    }
    defer listener.Close()

    log.Printf("Listening on port %d\n", port)

    conn, err := listener.Accept()
    if err != nil {
        log.Fatalf("Failed to accept connection: %v", err)
    }
    defer conn.Close()

    log.Printf("Connection accepted from %s\n", conn.RemoteAddr())

    io.Copy(os.Stdout, conn)
}

func connectMode(target string, port int) {
    conn, err := net.Dial("tcp", target+":"+strconv.Itoa(port))
    if err != nil {
        log.Fatalf("Failed to connect to target: %v", err)
    }
    defer conn.Close()

    log.Printf("Connected to %s\n", conn.RemoteAddr())

    io.Copy(conn, os.Stdin)
}
