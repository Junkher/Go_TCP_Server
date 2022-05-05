package main

import (
	"TCP/backend"
	"fmt"
	"os"
)

func main() {

		if len(os.Args) !=2 {
			fmt.Println("请正确输入监听的端口号！")
			os.Exit(1)
		}   
		fmt.Println("TCP will listen on port", os.Args[1])
		s := backend.NewServer()
		s.Startup(os.Args[1])

}
