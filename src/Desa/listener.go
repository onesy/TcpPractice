package main

import (
	"backgammon/protocal"
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
		go worker(connection, protocal.Enter)
	}
}

func contentProccessor(cntnt []byte) {
	Scntnt := string(cntnt)
	fmt.Printf("Scntn:%s", Scntnt)
}

func worker(conn net.Conn, delegation func(con net.Conn, cntnt []byte)) {
	var buffer [512]byte
	length, err := conn.Read(buffer[0:])
	user_error.Check_error(err, user_error.ERROR)
	delegation(conn, buffer[0:length])
	defer wipeAss(conn)
}

func wipeAss(conn net.Conn) {
	defer conn.Close()
}

func main() {
	server()
}
