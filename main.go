package main

import (
	"blog/db"
	"blog/server"
	"log"
)

func main() {
	DB := db.Connection()
	db.Migrate()
	app := server.Server(DB)
	if err := app.Listen(":8143"); err != nil {
		log.Panic(err)
	}
}
