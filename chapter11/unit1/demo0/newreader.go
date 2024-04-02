package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file,err :=os.OpenFile("/Users/skl/Desktop/skl.docx",os.O_RDWR|os.O_APPEND|os.O_CREATE,0666)

	if err!=nil{
		fmt.Println("文件打开失败，err=",err)
	}

	defer file.Close()
	wirter :=bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		wirter.WriteString("hello winter\t")
	}

	wirter.Flush()


}
