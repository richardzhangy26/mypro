package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:7777")
	defer conn.Close()
	if err != nil {
		fmt.Println("conn err = ", err)
		return

	}
	reader := bufio.NewReader(os.Stdin) //stdin代表标准输入【终端】
	// fmt.Printf("conn成功=%v\n 客户端的ip = %v\n ", conn)
	for {
		// 抓取终端的输入，换行结束
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("readString err=", err)
		}
		line = strings.Trim(line, "\r\n")
		if line == "exit" {
			fmt.Println("客户端退出")
			break
		}
		// 发送给服务器数据
		n, err := conn.Write([]byte(line + "\n"))
		if err != nil {
			fmt.Println("conn.Write err=", err)
		}
		fmt.Println("客户端发送了 %d 字节的数据", n)
	}

}
