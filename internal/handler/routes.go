package handler

import (
	"taskman/internal/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.TaskService
}

func NewHandler(services *service.TaskService) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	task := router.Group("/task")
	{
		task.POST("/create", h.CreateTask)
		task.GET("/:id", h.GetTaskByID)
		task.GET("/all", h.GetTasks)
		task.PUT("/update/:id", h.UpdateTask)
		task.DELETE("/delete/:id", h.DeleteTask)
		task.PATCH("/:id/done", h.MarkTask)
	}

	return router
}
