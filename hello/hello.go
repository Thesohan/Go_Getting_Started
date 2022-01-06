package main

import (
	"fmt"
	"github.com/thesohan/greetings"
"log")

func main(){
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	// message,err := greetings.Hello("Sohan")
	// fmt.Println(message)

	message,err := greetings.Hello("")
	if err!=nil{
		log.Fatal(err)

	}
	fmt.Println(message)
}