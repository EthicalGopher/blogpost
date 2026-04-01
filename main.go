package main

import (
	"blog/db"
	"blog/utils"
)

func main() {
	DB := db.Connection()
	db.Migrate()
	utils.ViewAllUsers(DB)
}
