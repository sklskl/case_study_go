package main

import "fmt"

func main() {
	//创建chinese
	c:=Chinese{}
	//创建american
	a:=Amerincan{}
 	greet(c)
	greet(a)

}

type SayHello interface {
	sayHello()
}
type Chinese struct {

}
type Amerincan struct {

}
//实现接口的方法
func (person Chinese) sayHello(){
	fmt.Println("你好")
}
func (person Amerincan) sayHello(){
	fmt.Println("hi")
}
//定义一个函数，用来打招呼，具备接受sayhello接口的能力变量
func greet(s SayHello)  {
	s.sayHello()
}