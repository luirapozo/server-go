package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	for {
		//creates a reader for terminal
		fmt.Printf("You can download files from this server\n")
		fmt.Printf("Write the option that you want:\n\n")
		fmt.Printf("1. List of files to download\n")
		fmt.Printf("2. Download File\n")
		fmt.Printf("0. Quit\n")
		fmt.Print(">> ")
		//reads the message
		text, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		if strings.TrimSpace(string(text)) == "1" {
			//connected to local host
			c, err := net.Dial("tcp", "127.0.0.1:8080")
			if err != nil {
				log.Fatal(err)
			}
			defer c.Close()
			Op1(c)
		} else if strings.TrimSpace(string(text)) == "2" {
			//connected to local host
			c, err := net.Dial("tcp", "127.0.0.1:8080")
			if err != nil {
				log.Fatal(err)
			}
			defer c.Close()
			Op2(c)
		} else if strings.TrimSpace(string(text)) == "0" {
			//connected to local host
			c, err := net.Dial("tcp", "127.0.0.1:8080")
			if err != nil {
				log.Fatal(err)
			}
			defer c.Close()

			fmt.Printf("Client leaving\n")
			c.Write([]byte("0\n"))
			os.Exit(0)
		} else {
			fmt.Printf("Invalid input, try again\n\n\n")
		}
	}
}

func Op1(c net.Conn) {

	//sends message to the server
	c.Write([]byte("1\n"))
	//reads messages from the server
	serverText, err := bufio.NewReader(c).ReadString('\t')
	if err != nil {
		log.Fatal(err)
	}
	//prints message from the server
	fmt.Println(serverText)
}
func Op2(c net.Conn) {

	c.Write([]byte("2\n"))

	fmt.Printf("Write the file you want to download: ")
	file, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	c.Write([]byte(file + "\n"))

	size, err := bufio.NewReader(c).ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	intSize, err := strconv.Atoi(strings.TrimSpace(size))
	if err != nil {
		log.Fatal(err)
	}

	var completeFile []byte
	for i := 0; i < intSize; i++ {
		serverFile, err := bufio.NewReader(c).ReadBytes('\n')
		if err != nil {
			log.Fatal(err)
		}
		c.Write([]byte("Recibido\n"))
		fmt.Printf("slice %v: %v\n\n", i+1, serverFile)
		completeFile = append(completeFile, serverFile...)
		time.Sleep(1 * time.Second)
	}
	err = ioutil.WriteFile(strings.TrimSpace(file), completeFile, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
