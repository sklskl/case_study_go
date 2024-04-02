package main

import "fmt"

func main() {
	//方式一
	var s2 Student1 = Student1{"zhangsan",16}
	fmt.Println("s2=",s2)
}

type Student1 struct {
	Name string
	Age int
}