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

	/*
		作成する
		u := &models.User{}
		u.Name = "tets"
		u.Email = "test@gmail.com"
		u.PassWord = "testtest"
		fmt.Println(u)

		u.CreateUser()
	*/

	// 取得する
	u, _ := models.GetUser(1)
	fmt.Println(u)

	// 更新する
	u.Name = "Test2"
	u.Email = "test2@example.com"
	u.UpdateUser()
	u, _ = models.GetUser(1)
	fmt.Println(u)

	// 削除する
	u.DeleteUser()
	u, _ = models.GetUser(1)
	fmt.Println(u)
}
