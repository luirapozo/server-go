package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
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
		fmt.Printf("You can download files from this server\n")
		fmt.Printf("Write the option that you want:\n\n")
		fmt.Printf("1. List of files to download\n")
		fmt.Printf("2. Download File\n")
		fmt.Printf("0. Quit\n")
		fmt.Print(">> ")
		//reads the message
		text, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		if strings.TrimSpace(string(text)) == "1" {

			//sends message to the server
			c.Write([]byte("1\n"))
			//reads messages from the server
			serverText, err := bufio.NewReader(c).ReadString('\t')
			if err != nil {
				log.Fatal(err)
			}
			//prints message from the server
			fmt.Println(serverText)
		} else if strings.TrimSpace(string(text)) == "2" {
			c.Write([]byte("2\n"))

			fmt.Printf("Write the file you want to download: ")
			file, err := reader.ReadString('\n')
			if err != nil {
				log.Fatal(err)
			}
			c.Write([]byte(file + "\n"))

			serverFile, err := bufio.NewReader(c).ReadBytes(' ')
			if err != nil {
				fmt.Printf("Error aqui\n")
				log.Fatal(err)
			}

			err = ioutil.WriteFile(strings.TrimSpace(file), serverFile, 0644)
			if err != nil {
				fmt.Printf("Error aqua\n")
				log.Fatal(err)
			}

		} else if strings.TrimSpace(string(text)) == "0" {
			fmt.Printf("Client leaving\n")
			c.Write([]byte("0\n"))
			os.Exit(0)
		} else {
			fmt.Printf("Invalid input, try again\n\n\n")
		}
	}
}
