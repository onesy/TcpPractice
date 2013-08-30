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
	fmt.Println("this is callback say:", a)
	for ai := range a {
		fmt.Printf("%v", ai)
	}
}
