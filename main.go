package main

import (
	"fmt"
	"todo-app/app/controllers"
	"todo-app/app/models"
	"todo-app/config"
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
		// 作成する
		u := &models.User{}
		u.Name = "test2"
		u.Email = "test2@gmail.com"
		u.PassWord = "testtest"
		fmt.Println(u)

		u.CreateUser()
	*/

	/*
		取得する
		u, _ := models.GetUser(1)
		fmt.Println(u)

		更新する
		u.Name = "Test2"
		u.Email = "test2@example.com"
		u.UpdateUser()
		u, _ = models.GetUser(1)
		fmt.Println(u)

		削除する
		u.DeleteUser()
		u, _ = models.GetUser(1)
		fmt.Println(u)
	*/

	/*
		// todoを追加する
		user, _ := models.GetUser(2)
		user.CreateTodo("First User")
	*/

	/*
		t, _ := models.GetTodo(1)
		fmt.Println(t)
	*/

	/*
		user, _ := models.GetUser(3)
		user.CreateTodo("Third User")
	*/

	/*
		todos, _ := models.GetTodos()
		for _, v := range todos {
			fmt.Println(v)
		}
	*/

	/*
		user2, _ := models.GetUser(3)
		todos, _ := user2.GetTodosByUser()
		for _, v := range todos {
			fmt.Println(v)
		}
	*/

	/*
		t, _ := models.GetTodo(1)
		t.Content = "Update Todo"
		t.UpdateTodo()
	*/

	/*
		t, _ := models.GetTodo(3)
		t.DeleteTodo()
	*/

	controllers.StartMainServer()
	fmt.Println(config.Config.Port)

}
