package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

const (
	connType = "tcp"
	connHost = "127.0.0.1"
	connPort = "8080"
)

func main() {
	fmt.Println("Establish connection to server " + connHost + ":" + connPort)
	c, err := net.Dial(connType, connHost+":"+connPort)
	if err != nil {
		fmt.Println("Error connecting to server: " + err.Error())
		os.Exit(1)
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Text for send: ")

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Erro reading from stdin: " + err.Error())
			return
		}

		c.Write([]byte(input))

		msg, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println("Error reading server response: " + err.Error())
			return
		}

		fmt.Println("Server respond: " + msg)
	}
}
