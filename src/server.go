package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

func ListFiles() string {
	fmt.Printf("The client has asked for the file list\n")
	files, err := ioutil.ReadDir("./data")
	if err != nil {
		log.Fatal(err)
	}
	list := "List of downloadable files:\n-------------------\n"
	for _, file := range files {

		list = list + file.Name() + "\n"
	}
	return list
}

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

		fmt.Printf("Client: %v\n", string(clientText))

		if strings.TrimSpace(string(clientText)) == "0" {
			fmt.Printf("The server has closed\n")
			os.Exit(0)
		} else if strings.TrimSpace(string(clientText)) == "1" {
			list := ListFiles()
			c.Write([]byte(list + "-------------------\t"))
		} else {

			fmt.Printf("The client has asked to download a file\n")
			file, err := reader.ReadString('\n')
			if err != nil {
				log.Fatal(err)
			}

			direccion := "./data/" + file
			fmt.Printf("File: %v", direccion)
			input, err := ioutil.ReadFile(strings.TrimSpace(direccion))
			if err != nil {
				log.Fatal(err)
			}

			res := bytes.Split(input, []byte("\n"))
			fmt.Printf("tamanio %v\n", len(res))
			c.Write([]byte(strconv.Itoa(len(res)) + "\n"))
			//fmt.Printf("%v\n", res)
			var partialFile []byte
			for i := 0; i < len(res); i++ {
				partialFile = append(res[i], byte('\n'))
				fmt.Printf("VALOR PARCIAL#%v %v\n\n", i+1, partialFile)
				c.Write(append(partialFile))
				time.Sleep(2 * time.Second)
				validacion, err := reader.ReadString('\n')
				if err != nil {
					log.Fatal(err)
				}
				fmt.Printf("validacionm%v\n\n", validacion)

			}
		}
	}
}
