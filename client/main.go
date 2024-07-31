package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
    conn, err := net.Dial("tcp", "localhost:8080")
    if err != nil {
        fmt.Println("Error connecting to server:", err)
        os.Exit(1)
    }
    defer conn.Close()

    fmt.Println("Connected to TCP server")
    for {
        reader := bufio.NewReader(os.Stdin)
        fmt.Print("Enter message: ")
        text, _ := reader.ReadString('\n')
        conn.Write([]byte(text))

        message, _ := bufio.NewReader(conn).ReadString('\n')
        fmt.Print("Server response: " + message)
    }
}
