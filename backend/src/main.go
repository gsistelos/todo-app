package main

import (
	"github.com/gsistelos/todo-app/api"
	"github.com/gsistelos/todo-app/db"
	"log"
)

func main() {
	db, err := db.NewMysqlDB()
	if err != nil {
		log.Fatal(err)
	}

	err = db.Init()
	if err != nil {
		log.Fatal(err)
	}

	apiServer := api.NewAPIServer(":8080", db)

	apiServer.Run()
}
