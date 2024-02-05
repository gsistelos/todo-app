package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

type MysqlDB struct {
	db *sql.DB
}

func NewMysqlDB() (*MysqlDB, error) {
	dbPassword := os.Getenv("MYSQL_ROOT_PASSWORD")
	dbName := os.Getenv("MYSQL_DATABASE")

	db, err := sql.Open("mysql", fmt.Sprintf("root:%s@tcp(mysql:3306)/%s", dbPassword, dbName))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to MySQL")

	return &MysqlDB{
		db: db,
	}, nil
}

func (s *MysqlDB) Init() error {
	if err := s.createUsersTable(); err != nil {
		return err
	}

	return nil
}
