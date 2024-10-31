package handlers

import (
	"log"
	"net/http"

	"github.com/codepnw/go-mongo-todos/models"
	"github.com/codepnw/go-mongo-todos/services"
	"github.com/gin-gonic/gin"
)

type todoHandler struct {
	service services.ITodos
}

func NewTodoHandler(service services.ITodos) *todoHandler {
	return &todoHandler{service: service}
}

func (h *todoHandler) CreateTodo(c *gin.Context) {
	req := &models.Todo{}

	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"code":    http.StatusBadRequest,
		})
		return
	}

	res, err := h.service.InsertTodo(req)
	if err != nil {
		log.Fatal(err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "created", "data": res})
}

func (h *todoHandler) GetTodos(c *gin.Context) {
	todos, err := h.service.FindAllTodos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, todos)
}

func (h *todoHandler) GetTodo(c *gin.Context) {
	id := c.Param("id")

	todo, err := h.service.FindTodoById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, todo)
}

func (h *todoHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var req *models.Todo

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.service.UpdateTodo(id, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *todoHandler) DeleteTodo(c *gin.Context) {
	id := c.Param("id")

	if err := h.service.DeleteTodo(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"message": "todo deleted"})
}
