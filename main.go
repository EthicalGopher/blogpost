package main

import (
	"blog/db"
	"blog/utils"
	"fmt"
)

func main() {
	DB := db.Connection()
	db.Migrate()
	utils.ViewAllUsers(DB)
	user, _ := utils.MyData(DB, 1)
	for _, u := range user {
		fmt.Println(u)
	}
}
