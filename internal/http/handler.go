package http

import (
        "net/http"
        "strconv"

        "github.com/gin-gonic/gin"

        "github.com/example/todo/internal/domain"
        "github.com/example/todo/internal/service"
)

// Handler exposes HTTP routes for todo service.
type Handler struct {
	Service *service.TodoService
}

// NewHandler creates an HTTP handler.
func NewHandler(s *service.TodoService) *Handler {
	return &Handler{Service: s}
}

func (h *Handler) Register(r *gin.Engine) {
        r.GET("/todos", h.listTodos)
        r.POST("/todos", h.createTodo)
        r.GET("/todos/:id", h.getTodo)
        r.PUT("/todos/:id", h.updateTodo)
        r.DELETE("/todos/:id", h.deleteTodo)
}

func (h *Handler) listTodos(c *gin.Context) {
        todos, err := h.Service.List()
        if err != nil {
                c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
                return
        }
        c.JSON(http.StatusOK, todos)
}

func (h *Handler) createTodo(c *gin.Context) {
        var todo domain.Todo
        if err := c.ShouldBindJSON(&todo); err != nil {
                c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
                return
        }
        if err := h.Service.Create(&todo); err != nil {
                c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
                return
        }
        c.JSON(http.StatusOK, todo)
}

func (h *Handler) getTodo(c *gin.Context) {
        id, err := strconv.ParseInt(c.Param("id"), 10, 64)
        if err != nil {
                c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
                return
        }
        todo, err := h.Service.Get(id)
        if err != nil {
                c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
                return
        }
        c.JSON(http.StatusOK, todo)
}

func (h *Handler) updateTodo(c *gin.Context) {
        id, err := strconv.ParseInt(c.Param("id"), 10, 64)
        if err != nil {
                c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
                return
        }
        var todo domain.Todo
        if err := c.ShouldBindJSON(&todo); err != nil {
                c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
                return
        }
        todo.ID = id
        if err := h.Service.Update(&todo); err != nil {
                c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
                return
        }
        c.JSON(http.StatusOK, todo)
}

func (h *Handler) deleteTodo(c *gin.Context) {
        id, err := strconv.ParseInt(c.Param("id"), 10, 64)
        if err != nil {
                c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
                return
        }
        if err := h.Service.Delete(id); err != nil {
                c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
                return
        }
        c.Status(http.StatusNoContent)
}
