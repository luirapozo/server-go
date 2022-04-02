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
	//connected to local host
	c, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}

	for {
		//creates a reader for terminal
		reader := bufio.NewReader(os.Stdin)
		fmt.Printf("Write something to the server:\n")
		//reads the message
		text, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		//sends message to the server
		c.Write([]byte(text + "\n"))
		//reads messages from the server
		serverText, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		//prints message from the server
		//%!v(MISSING) appears when using Printf
		fmt.Println("Server: " + serverText)
		//Closes the client
		if strings.TrimSpace(string(text)) == "0" {
			fmt.Printf("Client leaving\n")
			os.Exit(0)
		}
	}
}
