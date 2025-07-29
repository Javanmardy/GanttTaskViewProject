package api

import (
	"net/http"

	"gantt-task-view-project/backend/internal/db"
	model "gantt-task-view-project/backend/internal/models"

	"github.com/gin-gonic/gin"
)

func GetAllTasks(c *gin.Context) {
	rows, err := db.DB.Query("SELECT id, title_fa, assignee_fa, start, end, status, progress FROM tasks")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "DB error"})
		return
	}
	defer rows.Close()

	var tasks []model.Task
	for rows.Next() {
		var t model.Task
		err := rows.Scan(&t.ID, &t.TitleFa, &t.AssigneeFa, &t.Start, &t.End, &t.Status, &t.Progress)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Scan error"})
			return
		}
		tasks = append(tasks, t)
	}
	c.Writer.Header().Set("Content-Type", "application/json; charset=utf-8")

	c.JSON(http.StatusOK, tasks)
}

func CreateTask(c *gin.Context) {
	var task model.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := db.InsertTask(task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	task.ID = id
	c.JSON(http.StatusOK, task)
}
func GetTaskByID(c *gin.Context) {
	id := c.Param("id")
	task, err := db.GetTaskByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	c.JSON(http.StatusOK, task)
}
func UpdateTask(c *gin.Context) {
	id := c.Param("id")
	var task model.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := db.UpdateTask(id, task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Update failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "updated"})
}
func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	err := db.DeleteTask(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Delete failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "deleted"})
}
