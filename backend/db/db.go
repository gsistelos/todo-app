package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var (
	ErrNotFound    = fmt.Errorf("not found")
	ErrNotModified = fmt.Errorf("not modified")
)

type MysqlDB struct {
	db *sql.DB
}

func NewMysqlDB() (*MysqlDB, error) {
	dbPassword := os.Getenv("MYSQL_ROOT_PASSWORD")
	dbAddress := os.Getenv("MYSQL_ADDRESS")
	dbName := os.Getenv("MYSQL_DATABASE")

	db, err := sql.Open("mysql", fmt.Sprintf("root:%s@tcp(%s)/%s?parseTime=true", dbPassword, dbAddress, dbName))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to MySQL database")

	return &MysqlDB{
		db: db,
	}, nil
}

func (s *MysqlDB) Init() error {
	if err := s.createUsersTable(); err != nil {
		return err
	}

	if err := s.createTasksTable(); err != nil {
		return err
	}

	return nil
}

func (s *MysqlDB) doTableExists(table string) (bool, error) {
	query := fmt.Sprintf("SHOW TABLES LIKE '%s'", table)

	rows, err := s.db.Query(query)
	if err != nil {
		return false, err
	}
	defer rows.Close()

	return rows.Next(), nil
}
