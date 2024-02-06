package db

import (
	"database/sql"
	"fmt"
	"github.com/gsistelos/todo-app/models"
)

func (s *MysqlDB) CreateUser(userReq *models.CreateUserReq) (*models.User, error) {
	result, err := s.db.Exec("INSERT INTO users (username, email, password) VALUES (?, ?, ?)",
		userReq.Username, userReq.Email, userReq.Password)
	if err != nil {
		return &models.User{}, err
	}

	user := models.User{
		Username: userReq.Username,
		Email:    userReq.Email,
		Password: userReq.Password,
	}

	id, _ := result.LastInsertId()
	user.ID = int(id)

	return &user, nil
}

func (s *MysqlDB) GetUser(id string) (*models.User, error) {
	var user models.User
	if err := s.db.QueryRow("SELECT * FROM users WHERE id = ?", id).
		Scan(&user.ID, &user.Username, &user.Email, &user.Password); err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *MysqlDB) GetUsers() (*[]models.User, error) {
	rows, err := s.db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err = rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		return &users, err
	}

	if users == nil {
		return nil, sql.ErrNoRows
	}

	return &users, nil
}

func (s *MysqlDB) DeleteUser(id string) error {
	result, err := s.db.Exec("DELETE FROM users WHERE id = ?", id)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (s *MysqlDB) createUsersTable() error {
	exists, err := s.doTableExists("users")
	if err != nil {
		return err
	}

	if exists {
		fmt.Println("Table 'users' already exists")
		return nil
	}

	query := `
	CREATE TABLE users (
		id INT AUTO_INCREMENT PRIMARY KEY,
		username VARCHAR(255) NOT NULL CHECK (username <> ''),
		email VARCHAR(255) NOT NULL UNIQUE CHECK (email <> ''),
		password VARCHAR(255) NOT NULL CHECK (password <> '')
	)
	`

	_, err = s.db.Exec(query)
	if err != nil {
		return err
	}

	fmt.Println("Table 'users' created")

	return nil
}
