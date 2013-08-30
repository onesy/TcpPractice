package main

import (
	"fmt"
	"lib/user_error"
	"net"
)

func main() {

}

func client() {

	conn, conn_err := net.Dial(CONNECTION_CAT, SERVER_ADDR+":"+SERVER_PORT)
	user_error.Check_error(conn_err, user_error.FATAL)

	input = []byte("this is called from client")
	/**
	 * @TODO	 还没有写发信息和收信息的相关的东西
	 */
	// write msg from connection
	// read msg from connection
}

func writeMsg(conn net.Conn, msg []byte) {
	for {
		length, write_err := conn.Write(msg)
		user_error.Check_error(write_err, user_error.ERROR)
	}
}

func readMsg(conn net.Conn) {
	for {
		var buff [512]byte
		length, read_err := conn.Read(buff[0:])
		user_error.Check_error(write_err, user_error.ERROR)
		bufferStr := string(buff[0:length])
		fmt.Println("Client已经收到：", bufferStr)
	}

}
