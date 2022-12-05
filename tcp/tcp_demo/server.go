package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {
	// 循环的接受客户端发送的数据
	defer conn.Close()
	for {
		// 创建一个新的切片
		buf := make([]byte, 1024)
		// 1 等待客户端通过conn发送信息
		// 2.如果客户端没有write【发送】，那么协程就阻塞在这里
		fmt.Println("服务器在等待客户发送信息", conn.RemoteAddr().String())
		n, err := conn.Read(buf) //从conn读取
		if err != nil {
			fmt.Println("服务器的Read err=", err)
			return //
		}
		// 3.显示客户端发送的内容服务器的终端
		fmt.Print(string(buf[:n]))
	}

}
func main() {
	fmt.Println("服务器开始监听。。。")
	// 1表示网络协议是tcp
	// 28888为端口
	listen, err := net.Listen("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("listen err = ", err)
		return
	}
	defer listen.Close() //延时关闭
	// fmt.Printf("listen suc = %v", listen)
	for {
		// 等待客户端来链接
		fmt.Println("等待客户端链接...")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Accept err = ", err)
		} else {
			fmt.Printf("Accept suc con=%v\n,客户端IP =%v\n", conn, conn.RemoteAddr().String())
		}
		go process(conn)
	}
}
