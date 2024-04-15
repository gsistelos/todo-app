package main

import (
	"log"

	"github.com/gsistelos/todo-app/api"
	"github.com/gsistelos/todo-app/db"
)

func main() {
	db, err := db.NewMysqlDB()
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Init(); err != nil {
		log.Fatal(err)
	}

	apiServer := api.NewAPIServer(db)

	if err := apiServer.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
