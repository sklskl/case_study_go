package main

import "fmt"

func main() {
	v :=GenericDef
	fmt.Println(v)
}
func GenericDef(){
	////范性类型的定义
	//type mySlice [P int|string] []P
	//type myMap [K int|string,V float64|float32]map[K]V
	//type myList [l int|float32]struct{
	//	data []int
	//	l int
	//}

}
