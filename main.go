package main

import (
	"fmt"
	"todo_app_go/app/controllers"
	"todo_app_go/app/models"
)

func main() {

	fmt.Println(models.Db)

	controllers.StartMainServer()
}
