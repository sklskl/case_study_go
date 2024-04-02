package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	content,err := ioutil.ReadFile("/Users/skl/Desktop/skl.docx")
	if err!=nil {
		fmt.Println("读取出错，对应错误为：",nil)
	}
		//读取成功返回内容
		fmt.Printf("文件：%v", string(content))

}
