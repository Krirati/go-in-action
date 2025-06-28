package controllers

import (
	"learning/app/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
)

func (s *Server) CreateTodo(c *gin.Context) {
	var todo Todo
	err := c.ShouldBindJSON(&todo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error bind": err.Error(),
		})
		return
	}

	createdTodo, err := s.Store.CreateTodo(c, createTodoParams(todo))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error create": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, todoRespone(createdTodo))
}

func createTodoParams(todo Todo) models.CreateTodoParams {
	return models.CreateTodoParams{
		Title:       todo.Title,
		Description: pgtype.Text{String: todo.Description},
		Completed:   pgtype.Bool{Bool: *todo.Completed, Valid: true},
		CreatedAt:   pgtype.Timestamp{Time: time.Now().UTC(), Valid: true},
		UpdatedAt:   pgtype.Timestamp{Time: time.Now().UTC(), Valid: true},
	}
}

func todoRespone(createdTodo models.Todo) Todo {
	return Todo{
		Title:       createdTodo.Title,
		Description: createdTodo.Description.String,
		CreatedAt:   createdTodo.CreatedAt.Time.Format("2006-01-02 15:04:05"),
		UpdatedAt:   createdTodo.UpdatedAt.Time.Format("2006-01-02 15:04:05"),
	}
}
