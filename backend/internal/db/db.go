package db

import (
	"database/sql"
	"fmt"
	"gantt-task-view-project/backend/internal/models"
	model "gantt-task-view-project/backend/internal/models"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect(user, password, dbname string) error {
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=true&loc=Local", user, password, dbname)
	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	return DB.Ping()
}

func InsertTask(task model.Task) (int, error) {
	query := `INSERT INTO tasks (title_fa, assignee_fa, start, end, status, progress)
              VALUES (?, ?, ?, ?, ?, ?)`
	result, err := DB.Exec(query, task.TitleFa, task.AssigneeFa, task.Start, task.End, task.Status, task.Progress)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	return int(id), err
}
func GetTaskByID(id string) (models.Task, error) {
	var task models.Task
	err := DB.QueryRow("SELECT id, title_fa, assignee_fa, start, end, status, progress FROM tasks WHERE id = ?", id).
		Scan(&task.ID, &task.TitleFa, &task.AssigneeFa, &task.Start, &task.End, &task.Status, &task.Progress)
	return task, err
}
func UpdateTask(id string, task models.Task) error {
	_, err := DB.Exec(`UPDATE tasks SET title_fa=?, assignee_fa=?, start=?, end=?, status=?, progress=? WHERE id=?`,
		task.TitleFa, task.AssigneeFa, task.Start, task.End, task.Status, task.Progress, id)
	return err
}
func DeleteTask(id string) error {
	_, err := DB.Exec("DELETE FROM tasks WHERE id = ?", id)
	return err
}
