package main

import (
	"fmt"
	"log"

	"go_bmstu/internal/api"
)

func main(){
	log.Println("Application start!")
	api.StartServer()
	log.Println("Application terminated!")
	fmt.Println("Hello World!!")
}
