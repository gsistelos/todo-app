package db

import (
	"database/sql"
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
	if err := s.db.QueryRow("SELECT id, username, email, password FROM users WHERE id = ?", id).
		Scan(&user.ID, &user.Username, &user.Email, &user.Password); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (s *MysqlDB) GetUsers() (*[]models.User, error) {
	rows, err := s.db.Query("SELECT id, username, email, password FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
        return &users, err
    }

	return &users, nil
}

func (s *MysqlDB) createUsersTable() error {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id INT AUTO_INCREMENT PRIMARY KEY,
		username VARCHAR(255) NOT NULL CHECK (username <> ''),
		email VARCHAR(255) NOT NULL UNIQUE CHECK (email <> ''),
		password VARCHAR(255) NOT NULL CHECK (password <> '')
	)
	`

	_, err := s.db.Exec(query)
	return err
}
