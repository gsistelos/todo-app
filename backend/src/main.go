package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/gsistelos/todo-app/controllers"
	_ "github.com/gsistelos/todo-app/models"
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
	dbUser := os.Getenv("MYSQL_USER")
	dbPassword := os.Getenv("MYSQL_PASSWORD")
	dbName := os.Getenv("MYSQL_DATABASE")

	connectionString := fmt.Sprintf("%s:%s@tcp(mysql:3306)/%s", dbUser, dbPassword, dbName)

	db, err := sql.Open("mysql", connectionString)
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

	router.HandleFunc("/users", controllers.GetAllUsers).Methods("GET")
	router.HandleFunc("/users/{id}", controllers.GetUser).Methods("GET")

	return router
}
