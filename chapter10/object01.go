package main

import "fmt"

func main() {
	var t *Teacher = new (Teacher)
	//t是指针，指向的是地址,(*t)作用是根据地址取值
	(*t).Name="tome"
	(*t).Age =18
	(*t).School ="bj"
	//{
	//	Name:   "tome",
	//	Age:    18,
	//	School: "sdu",
	//}
	fmt.Println(*t)
}
 //定义一个老师的结构体，统一放入结构体中管理
 type Teacher struct {
	 //变量名字大写，外界可以访问这个属性
    Name string
	Age int
	School string
 }