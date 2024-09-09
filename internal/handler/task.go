package handler

import (
	"strconv"
	"taskman/internal/models"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateTask(c *gin.Context) {
	var params models.CreateTaskParams
	if err := c.BindJSON(&params); err != nil {
		c.JSON(400, err.Error())
		return
	}

	task, err := h.services.CreateTask(c, params)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(201, task)
}

func (h *Handler) GetTaskByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}

	task, err := h.services.GetTaskByID(c, id)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, gin.H{"task": task})
}

func (h *Handler) GetTasks(c *gin.Context) {
	tasks, err := h.services.GetTasks(c)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, gin.H{"tasks": tasks})
}

func (h *Handler) UpdateTask(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}

	var params models.UpdateTaskParams
	err = c.BindJSON(&params)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}

	task, err := h.services.UpdateTask(c, id, params)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, gin.H{"task updeted": task})
}

func (h *Handler) DeleteTask(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}

	err = h.services.DeleteTask(c, id)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, "task deleted")
}

func (h *Handler) MarkTask(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}

	var params models.MarkTaskParams
	err = c.BindJSON(&params)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}

	err = h.services.MarkTask(c, id, params)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, gin.H{"message": "task marked succesfully"})
}
