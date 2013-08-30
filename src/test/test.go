package main

import (
	"fmt"
	"lib/user_error"
	"net"
)

func main() {
	conn, _ := net.Listen("tcp", "127.0.0.1:8089")
	user_error.WipeAss(nil, user_error.NOTICE, callbacktest, conn)

}

func callbacktest(a ...interface{}) {
	fmt.Printf("callback params:%v", (a[0]).(net.Conn))

}
