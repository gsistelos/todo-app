package db

import (
	"fmt"
	"github.com/gsistelos/todo-app/models"
)

func (s *MysqlDB) CreateTask(task *models.Task) (int64, error) {
	res, err := s.db.Exec("INSERT INTO tasks (user_id, description, done, created_at) VALUES (?, ?, ?, ?)", task.UserID, task.Description, task.Done, task.CreatedAt)
	if err != nil {
		return 0, err
	}

	id, _ := res.LastInsertId()
	return id, nil
}

func (s *MysqlDB) GetTasks(id string) (*[]models.Task, error) {
	rows, err := s.db.Query("SELECT * FROM tasks WHERE user_id = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var task models.Task
		if err = rows.Scan(&task.ID, &task.UserID, &task.Description, &task.Done, &task.CreatedAt); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	if err = rows.Err(); err != nil {
		return &tasks, err
	}

	if tasks == nil {
		return nil, NotFound
	}

	return &tasks, nil
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
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)
	`

	_, err = s.db.Exec(query)
	if err != nil {
		return err
	}

	fmt.Println("Table 'tasks' created")

	return nil
}
