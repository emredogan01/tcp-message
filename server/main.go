package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func handleConnection(conn net.Conn) {
    fmt.Println("Connected to", conn.RemoteAddr().String())
    defer conn.Close()
    for {
        message, err := bufio.NewReader(conn).ReadString('\n')
        if err != nil {
            fmt.Println("Connection closed by client", conn.RemoteAddr().String())
            return
        }
        fmt.Print("Message received:", string(message))
        conn.Write([]byte("Message received: " + message))
    }
}

func main() {
    listener, err := net.Listen("tcp", ":8080")
    if err != nil {
        fmt.Println("Error starting TCP server:", err)
        os.Exit(1)
    }
    defer listener.Close()
    fmt.Println("TCP server listening on port 8080")
    for {
        conn, err := listener.Accept()
        if err != nil {
            fmt.Println("Error accepting connection:", err)
            continue
        }
        go handleConnection(conn)
    }
}
