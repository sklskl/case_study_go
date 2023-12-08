package main

import "fmt"

func main() {
  var a map[int] string
 a = make(map[int] string,10)
 a[2000]="zhangsan"
 a[2001]="lisi"
 a[2002]="wangwu"
 fmt.Println(a)
}
