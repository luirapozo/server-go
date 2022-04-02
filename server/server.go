package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	//net.Listen opens the server in the assigned port
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	//listener will close at the end
	defer listener.Close()

	//accepts connection with the client
	c, err := listener.Accept()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("client connected: %v\n", listener.Addr())

	for {
		//creates a reader for the client
		reader := bufio.NewReader(c)
		//reads the message
		clientText, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		//reads the message from the client
		if strings.TrimSpace(string(clientText)) == "0" {
			fmt.Printf("The server has closed\n")
			os.Exit(0)
		}
		fmt.Printf("Client: %v\n", string(clientText))
		//sends message to the clients
		c.Write([]byte("Server has recieved the message\n"))
	}
}
