package main

import "fmt"

func main() {
 var intarr[6]int = [6]int{1,2,3,4,5,6}
 var slice []int  = intarr[2:5]
 fmt.Println(slice)
 fmt.Println(len(slice))
 slice2:=append(slice,11,22)
 fmt.Println(slice2)
	var a[]int = []int{1,2,3,4,5,6}
	var b[]int = make([]int,15)
	copy(b,a)
	fmt.Println(b)

}
