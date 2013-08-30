package main

import (
	"fmt"
	"lib/user_error"
	"net"
)

func server() {
	ln, err := net.Listen("tcp", ":8080")
	user_error.Check_error(err, user_error.FATAL)
	for {
		connection, err := ln.Accept()
		user_error.Check_error(err, user_error.ERROR)
		go worker(connection, contentProccessor)
	}
}

func contentProccessor(cntnt []byte) {
	Scntnt := string(cntnt)
	fmt.Printf("Scntn:%s", Scntnt)
}

func worker(conn net.Conn, delegation func(cntnt []byte)) {
	var buffer [512]byte
	length, err := conn.Read(buffer[0:])
	user_error.Check_error(err, user_error.ERROR)
	delegation(buffer[0:length])
	go wipeAss(conn)
}

func wipeAss(conn net.Conn) {
	defer conn.Close()
}

func main() {
	server()
}
