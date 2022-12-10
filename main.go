package main

import (
	"fmt"
	"todo-app/app/models"
)

// go mod init todo-appで、todo-appとしたので、todo-app/configとする

func main() {
	/*
		fmt.Println(config.Config.Port)
		fmt.Println(config.Config.SQLDriver)
		fmt.Println(config.Config.DbName)
		fmt.Println(config.Config.LogFile)

		log.Println("test")
	*/

	fmt.Println(models.Db)
}
