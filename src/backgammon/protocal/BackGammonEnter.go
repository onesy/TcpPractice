package protocal

import (
	"fmt"
)

/*
 服务器逻辑处理.入口函数.
 @author sunyuw <leentingOne@gmail.com>
*/
func Enter(tcpcontent []byte) {
	// BackGammonLevel1Protocal(BackGammon一层协议)解析封包.
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
	// test
	Scntnt := string(tcpcontent)
	fmt.Printf("Scntn:%s", Scntnt)
}
func BG1LevelProtocalProc(tcpget []byte) (contents map[string][]byte, err error) {
	contents = map[string][]byte{
		"test": tcpget,
	}
	err = nil
	return
}

func BG2LevelProtocalProc(content []byte) {
	// BackGammonLevel2Protocal(BackGammon二层协议)解析封包.
	/*
		这一层主要是提取出业务内容,比第一层协议更为特殊的是将一部分逻辑代码放在这里.
		这里的逻辑处理非常简单,就是通过magic字段去查找一个协议分析的格式
		被反射调用的函数返回2个map,第一个map用来表达解析格式,
		第二个map则为键和第一个map一致的接受解析内容的函数,类型byte[]为值.
		在返回的map中key'funcProc'为必选参数.
		byte[0:4] 长4byte的magic段.通过magic段在FuncMap常量中去查找方法名,反射调用.
	*/
}
