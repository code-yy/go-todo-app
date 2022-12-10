package main

import (
	"fmt"
	"log"

	// go mod init todo-appで、todo-appとしたので、todo-app/configとする
	"todo-app/config"
)

func main() {
	fmt.Println(config.Config.Port)
	fmt.Println(config.Config.SQLDriver)
	fmt.Println(config.Config.DbName)
	fmt.Println(config.Config.LogFile)

	log.Println("test")
}
