package main

import (
	"fmt"
	"net"
)

func main() {
fmt.Println("客户端启动==")
con,err:=net.Dial("tcp","127.0.0.1:8888")
	if err!=nil {
		fmt.Println("error not link",err)
	}
	fmt.Println("success",con)
}
