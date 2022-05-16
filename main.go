package main

import (
	"fmt"
	"todo_app_go/config"
)

func main() {
	fmt.Println(config.Config.Port)
	fmt.Println(config.Config.DbName)
	fmt.Println(config.Config.LogFile)
	fmt.Println(config.Config.SQLDriver)
}
