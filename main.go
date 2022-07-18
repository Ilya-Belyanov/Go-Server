package main

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
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
	if tp == REG {
		tryRegister(conn, packet)
	}
}

func tryRegister(conn net.Conn, packet map[string]interface{}) {
	keys := getClientKeys()
	isExist := isExistUser(packet[keys.NAME].(string))
	var answer RegisterStruct
	if isExist {
		answer = RegisterStruct{SERVER_ERR_REG_AUTH, "Имя уже занято", -1, false}
	} else if lastId, err := addNewUser(packet[keys.NAME].(string), packet[keys.PASS].(string)); err {
		answer = RegisterStruct{SERVER_ERR_REG_AUTH, "Ошибка со стороны бд", -1, false}
	} else {
		answer = RegisterStruct{SERVER_SUCCESS_REG, "Успешно", lastId, false}
	}
	b, _ := json.Marshal(answer)
	fmt.Println(string(b))
	conn.Write(b) // пишем в сокет
}
