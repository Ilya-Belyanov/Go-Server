package main

import (
	"encoding/json"
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:7777")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()
	fmt.Println("Server is listening...")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Err: ", err)
			return
		}
		go handleClient(conn)
	}

}

func handleClient(conn net.Conn) {
	defer conn.Close()
	fmt.Println("New connect")
	for {
		//conn.Write([]byte("Hello, what's your name?\n")) // пишем в сокет
		buf := make([]byte, 1024*4) // буфер для чтения клиентских данных

		readLen, err := conn.Read(buf) // читаем из сокета
		if err != nil {
			fmt.Println("Err: ", err)
			break
		}
		fmt.Println("JSON ", string(buf[:readLen]))
		var js interface{}
		err = json.Unmarshal(buf[:readLen], &js)
		if err != nil {
			fmt.Println("Err JSON: ", err)
			break
		}
		packet := js.(map[string]interface{})
		fmt.Println("JSON ", packet["type"])
		checkClientPacket(conn, packet)
	}
	fmt.Println("End connect")
}

func checkClientPacket(conn net.Conn, packet map[string]interface{}) {
	keys := getClientKeys()
	tp := int(packet[keys.TYPE].(float64))
	if tp == 1 {
		answer := RegisterStruct{1, "Success", 1, true}
		b, _ := json.Marshal(answer)
		fmt.Println(string(b))
		conn.Write(b) // пишем в сокет
	}
}
