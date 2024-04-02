package main

import (
	"fmt"
)

func main() {
 //创建结构体实力
	p :=NewPerson("lili")
	p.SetAge(18)
	fmt.Println(*p)
	fmt.Println(p.GetAge())
}
