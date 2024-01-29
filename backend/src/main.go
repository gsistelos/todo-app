package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/gsistelos/todo-app/controllers"
	"github.com/gsistelos/todo-app/models"
	"net/http"
	"os"
)

func main() {
	db := connectDB()
	defer db.Close()

	router := createRouter()

	http.ListenAndServe(":8080", router)
}

func connectDB() *sql.DB {
	dbPassword := os.Getenv("MYSQL_ROOT_PASSWORD")

	db, err := sql.Open("mysql", fmt.Sprintf("mysql:%s@tcp(mysql:3306)/mysql", dbPassword))
	if err != nil {
		panic(err.Error())
	}

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	return db
}

func createRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/users", controllers.getAllUsers).Methods("GET")
	router.HandleFunc("/users/{id}", controllers.getUser).Methods("GET")

	return router
}
