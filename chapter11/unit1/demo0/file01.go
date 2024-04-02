package main

import (
	"fmt"
	"os"
)

func main() {
//打开文件
	file,err := os.Open("/Users/skl/Desktop/skl.docx")
	if err!=nil {
		fmt.Println("文件打开出错，对应错误为：",nil)
	}else {
		fmt.Printf("文件：%v", file)
	}
	err2 := file.Close()
	if err2!=nil {
		fmt.Println("关闭失败")
	}
}
