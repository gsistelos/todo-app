package db

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/gsistelos/todo-app/models"
)

var (
	NotFound    = fmt.Errorf("Not found")
	NotModified = fmt.Errorf("Not modified")
)

func (s *MysqlDB) CreateUser(user *models.User) (int64, error) {
	result, err := s.db.Exec("INSERT INTO users (username, email, password, created_at) VALUES (?, ?, ?, ?)",
		user.Username, user.Email, user.Password, user.CreatedAt)
	if err != nil {
		return 0, err
	}

	id, _ := result.LastInsertId()
	return id, nil
}

func (s *MysqlDB) GetUserByID(id string) (*models.User, error) {
	var user models.User
	if err := s.db.QueryRow("SELECT * FROM users WHERE id = ?", id).
		Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, NotFound
		} else {
			return nil, err
		}
	}

	return &user, nil
}

func (s *MysqlDB) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := s.db.QueryRow("SELECT * FROM users WHERE email = ?", email).
		Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, NotFound
		} else {
			return nil, err
		}
	}

	return &user, nil
}

func (s *MysqlDB) GetUsers() (*[]models.User, error) {
	rows, err := s.db.Query("SELECT id, username, created_at FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err = rows.Scan(&user.ID, &user.Username, &user.CreatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		return &users, err
	}

	if users == nil {
		return nil, NotFound
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
		return NotFound
	}

	return nil
}

func (s *MysqlDB) UpdateUser(id string, user *models.User) error {
	result, err := s.db.Exec("UPDATE users SET username = ?, email = ?, password = ? WHERE id = ?",
		user.Username, user.Email, user.Password, id)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return NotModified
	}

	return nil
}

func (s *MysqlDB) UserExists(id string) (bool, error) {
	var exists bool
	if err := s.db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE id = ?)", id).Scan(&exists); err != nil {
		return false, err
	}

	return exists, nil
}

func (s *MysqlDB) UserEmailExists(email string) (bool, error) {
	var exists bool
	if err := s.db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE email = ?)", email).Scan(&exists); err != nil {
		return false, err
	}

	return exists, nil
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
		password VARCHAR(255) NOT NULL CHECK (password <> ''),
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)
	`

	_, err = s.db.Exec(query)
	if err != nil {
		return err
	}

	fmt.Println("Table 'users' created")

	return nil
}