package main

import (
	// "bufio"
	"fmt"
	"net"
	// "os"
	"io/ioutil"
)

// import "bufio"
// import "os"

func main() {

	// Подключаемся к сокету
	conn, _ := net.Dial("tcp", "127.0.0.1:8081")
	defer conn.Close()
	bs, _ := ioutil.ReadAll(conn)
	fmt.Println(string(bs))

	// // Подключаемся к сокету
	// conn, _ := net.Dial("tcp", "127.0.0.1:8081")
	// // defer conn.Close()
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
