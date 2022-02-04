package main

import (
	"fmt"
	"log"
)
import "hongjun.com/greetings"

func main() {

	//设置日志预定义的属性
	log.SetPrefix("greetings:")
	log.SetFlags(0)

	//message, err := greetings.Hello("hongjun500")
	message, err := greetings.Hello("hongjun500")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
