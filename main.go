package main

import (
	"fmt"
	"todo_app_go/app/models"
)

func main() {
	// fmt.Println(config.Config.Port)
	// fmt.Println(config.Config.DbName)
	// fmt.Println(config.Config.LogFile)
	// fmt.Println(config.Config.SQLDriver)

	// log.Println("test")

	fmt.Println(models.Db)

	u := &models.User{}
	u.Name = "test"
	u.Email = "test@gmail.com"
	u.PassWord = "testtest"
	fmt.Println(u)

	u.CreateUser()
}
