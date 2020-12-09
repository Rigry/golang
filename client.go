package main

import (
	// "bufio"
	"io"
	"fmt"
	"net"
	"time"
	// "os"
	// "io/ioutil"
)

// import "bufio"
// import "os"

func handle(conn net.Conn) {
	for {
		input := make([]byte, 20)
		n, _ := conn.Read(input)
		input = input[0:n]
		fmt.Println(input)
	}
}

func main() {

	// Подключаемся к сокету
	// conn, _ := net.Dial("tcp", "127.0.0.1:8081")
	// defer conn.Close()
	// bs, _ := ioutil.ReadAll(conn)
	// fmt.Println(string(bs))

	// Подключаемся к сокету
	// conn, _ := net.Dial("tcp", "10.121.3.34:8081")
	conn, _ := net.Dial("tcp", "127.0.0.1:8081")
	

	bytes := make([]byte, 20)
	bytes[0] = 0
	bytes[1] = 1
	bytes[2] = 0
	bytes[3] = 0
	bytes[4] = 0
	bytes[5] = 0xB
	bytes[6] = 1
	bytes[7] = 0x10
	bytes[8] = 0
	bytes[9] = 0
	bytes[10] = 0
	bytes[11] = 2
	bytes[12] = 4
	bytes[13] = 0
	bytes[14] = 8
	bytes[15] = 0
	bytes[16] = 10

	request := make([]byte, 20)
	request[0] = 0
	request[1] = 1
	request[2] = 0
	request[3] = 0
	request[4] = 0
	request[5] = 6
	request[6] = 1
	request[7] = 0x03
	request[8] = 0
	request[9] = 0
	request[10] = 0
	request[11] = 2

	var b bool 

	for {
		defer conn.Close()
		
		if b {
			conn.Write(request)
			b = false
			fmt.Println(10)
		} else {
			conn.Write(bytes)
			b = true
			fmt.Println(3)
		}

		var data [12]byte
		if _, err := io.ReadFull(conn, data[:12]); err != nil {
			fmt.Println(err)
		}
		// fmt.Println(len)
		// fmt.Println(io.ReadFull(conn, data[:13]))
		// input = input[0:n]
		fmt.Println(data)

		time.Sleep(time.Second)
	}


	// defer conn.Close()
	// for {
	// 	// Чтение входных данных от stdin
	// 	reader := bufio.NewReader(os.Stdin)
	// 	fmt.Print("Text to send: ")
	// 	text, _ := reader.ReadString('\n')
	// 	// Отправляем в socket
	// 	fmt.Fprintf(conn, text+"\n")
		
	// 	// Прослушиваем ответ
	// 	message, _ := bufio.NewReader(conn).ReadString('\n')
	// 	fmt.Print("Message from server: " + string(message))
	// }
}
