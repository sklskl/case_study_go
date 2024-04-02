package main

import (
	"fmt"
	"strconv"
	"time"
)

/*
协程上手案例，编写一个程序完成
1、主线程中开启一个goroutine，该goroutine每隔一秒输出hello golang
2、主线程上每隔一秒输出hello msb 输出10次后退出程序
3、主线程和goroutine同时执行
*/
func main() {//主线程
	 go test01()//开启协程
	for i := 0; i < 10; i++ {
		fmt.Println("hello msb+"+strconv.Itoa(i))
		//阻塞一秒
		time.Sleep(time.Second)
	}

}
func test01()  {
	for i := 0; i < 10; i++ {
		fmt.Println("hello golang"+strconv.Itoa(i))
		//阻塞一秒
		time.Sleep(time.Second)
	}
}