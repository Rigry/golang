package main

import (
	"encoding/gob"
	"fmt"
	"net"
)

func server() {
	// слушать порт
	ln, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
	}
	for {
		// принятие соединения
		с, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		// обратботка соединения
		go handleServerConnection(с)
	}

}

func handleServerConnection(c net.Conn) {
	// получение сообщения
	var msg string
	err := gob.NewDecoder(c).Decode(&msg)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Received", msg)
	}

	c.Close()
}

func client() {
	// соединиться с сервером
	с, err := net.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	// послать сообщение
	msg := "Hello, World"
	fmt.Println("Sending", msg)
	err = gob.NewEncoder(с).Encode(msg)
	if err != nil {
		fmt.Println(err)
	}
	с.Close()
}

func main() {
	go server()
	go client()

	var input string
	fmt.Scanln(&input)
}
