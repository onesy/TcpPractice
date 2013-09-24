package main

import (
	"fmt"
	"lib/user_error"
	"net"
)

func main() {

}

func client() {
	// 协议描述.
	/*
		这一层主要处理安全和登录等等.
		byte[0:4] 长4byte的magic段.
		byte[4:7] 通信协议版本号,版本为3部分,通信协议主版本号byte[0],通信协议次版本号byte[1],通信协议子版本号byte[2].
		byte[7:10] 客户端版本号,版本为3部分,客户端主版本号byte[3],客户端次版本号[4],客户端子版本号[5].
		byte[10:13] 客户端对应服务器的版本号,格式如上.byte[6],byte[7],byte[8].
		byte[13:17] 用户名加密串长度UserNameLen.
		byte[17:20] 密码加密串长度PasswdLen.
		byte[20:24] 凭证长度CertificateLen.
		byte[24:32] 内容长度ContentLen.
		byte[32:32 + UserNameLen] 用户名加密信息.
		byte[32+UserNameLen:32+UserNameLen + PasswdLen] 密码加密内容.
		byte[32+ UserNameLen + PasswdLen:32+UserNameLen + PasswdLen + CertificateLen] 凭证内容.
		设:baseInfo = 32+UserNameLen + PasswdLen + CertificateLen;
		byte[baseInfo:baseInfo + ContentLen] 内容.此内容交由二层协议处理.
	*/

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
