package main

import (
	"fmt"
	"log"
	"net"
)

func server() {
	ln, err := net.Listen("tcp", ":8080")
	check_error(err, FATAL)
	for {
		connection, err := ln.Accept()
		check_error(err, ERROR)
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
	check_error(err, ERROR)
	delegation(buffer[0:length])
	go wipeAss(conn)
}

func check_error(err error, level int) {
	if err != nil {
		switch level {
		case FATAL:
			log.Println("Event Class:FATAL!", err)
		case ERROR:
			log.Println("Event Class:ERROR!", err)
		case WARNING:
			log.Println("Event Class:WAUNING!", err)
		case NOTICE:
			log.Println("Event Class:NOTICE!", err)
		default:
			log.Println("Event Class:UNKNOWN!", err)
		}
	}
}

func wipeAss(conn net.Conn) {
	defer conn.Close()
}

func main() {
	server()
}
