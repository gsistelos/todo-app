package db

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/gsistelos/todo-app/models"
)

func (s *MysqlDB) CreateTask(task *models.Task) (int64, error) {
	res, err := s.db.Exec("INSERT INTO tasks (user_id, description, done, term, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)",
		task.UserID, task.Description, task.Done, task.Term, task.CreatedAt, task.UpdatedAt)
	if err != nil {
		return 0, err
	}

	id, _ := res.LastInsertId()
	return id, nil
}

func (s *MysqlDB) GetTasks(userID string) (*[]models.Task, error) {
	rows, err := s.db.Query("SELECT * FROM tasks WHERE user_id = ?", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var task models.Task
		if err = rows.Scan(&task.ID, &task.UserID, &task.Description, &task.Done, &task.Term, &task.CreatedAt, &task.UpdatedAt); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	if err = rows.Err(); err != nil {
		return &tasks, err
	}

	if tasks == nil {
		return nil, ErrNotFound
	}

	return &tasks, nil
}

func (s *MysqlDB) GetTaskByID(userID, taskID string) (*models.Task, error) {
	var task models.Task
	err := s.db.QueryRow("SELECT * FROM tasks WHERE user_id = ? AND id = ?", userID, taskID).
		Scan(&task.ID, &task.UserID, &task.Description, &task.Done, &task.Term, &task.CreatedAt, &task.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotFound
		} else {
			return nil, err
		}
	}

	return &task, nil
}

func (s *MysqlDB) UpdateTask(userID, taskID string, task *models.Task) error {
	result, err := s.db.Exec("UPDATE tasks SET description = ?, done = ?, term = ?, updated_at = ? WHERE user_id = ? AND id = ?",
		task.Description, task.Done, task.Term, task.UpdatedAt, userID, taskID)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return ErrNotFound
	}

	return nil
}

func (s *MysqlDB) DeleteTask(userID, taskID string) error {
	result, err := s.db.Exec("DELETE FROM tasks WHERE user_id = ? AND id = ?", userID, taskID)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return ErrNotFound
	}

	return nil
}

func (s *MysqlDB) createTasksTable() error {
	exists, err := s.doTableExists("tasks")
	if err != nil {
		return err
	}

	if exists {
		fmt.Println("Table 'tasks' already exists")
		return nil
	}

	query := `
	CREATE TABLE tasks (
		id INT AUTO_INCREMENT PRIMARY KEY,
		user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
		description VARCHAR(255) NOT NULL CHECK (description <> ''),
		done BOOLEAN NOT NULL DEFAULT FALSE,
		term TIMESTAMP NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
	)
	`

	_, err = s.db.Exec(query)
	if err != nil {
		return err
	}

	fmt.Println("Table 'tasks' created")

	return nil
}
