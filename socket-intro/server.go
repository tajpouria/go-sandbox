package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

const (
	connHost = "127.0.0.1"
	connPort = "8080"
	connType = "tcp"
)

func main() {
	fmt.Println("Starting " + connType + " server on " + "connHost " + connHost + ":" + connPort)
	l, err := net.Listen(connType, connHost+":"+connPort)
	if err != nil {
		fmt.Println("Error listening: " + err.Error())
		os.Exit(1)
	}
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println("Error connection: ", err.Error())
			return
		}
		log.Println("Client " + c.RemoteAddr().String() + " Connected")

		go handleConnection(c)
	}
}

func handleConnection(c net.Conn) {
	buffer, err := bufio.NewReader(c).ReadBytes('\n')
	if err != nil {
		fmt.Println("Client left")
		c.Close()
		return
	}

	log.Println("Client message: ", string(buffer[:len(buffer)-1]))

	c.Write(buffer)

	handleConnection(c)
}
