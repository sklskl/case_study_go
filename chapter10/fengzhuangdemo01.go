package main

import (
	"fmt"

)

type person struct {
	Name string
	age int//其他包不能直接访问

}
//定义工厂模式的函数，相当于构造器
func NewPerson(name string)  *person{
	return &person{
		Name: name,
	}
}
func (p *person)  SetAge(age int){
	if age>0 && age<150 {
		p.age = age
	}else {
		fmt.Println("sorry，传入年龄范围不正确")
	}
}
func (p *person) GetAge( ) int{
	return p.age
}