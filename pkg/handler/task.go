package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sharifsharifzoda/project-management-system/models"
	"strconv"
	"strings"
)

type taskIn struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	ExecutorId  int    `json:"executor_id" binding:"required"`
	Status      string `json:"status"`
	ProjectId   int    `json:"project_id" binding:"required"`
	Deadline    string `json:"deadline" binding:"required"`
}

func (h *Handler) createTask(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		c.JSON(500, map[string]any{
			"error": err.Error(),
		})
		return
	}

	userRole, err := GetUserRole(c)
	if err != nil {
		c.JSON(500, map[string]any{
			"error": err.Error(),
		})
		return
	}

	if strings.ToLower(userRole) != "superuser" {
		c.JSON(400, map[string]any{
			"error": "You are not allowed to create a project",
		})
		return
	}

	var data taskIn
	if err := c.BindJSON(&data); err != nil {
		c.JSON(400, map[string]any{
			"error": "invalid JSON provided",
		})
		return
	}

	var task = models.Task{
		Title:        data.Title,
		Description:  data.Description,
		ControllerId: userId,
		ExecutorId:   data.ExecutorId,
		Status:       data.Status,
		ProjectId:    data.ProjectId,
		Deadline:     data.Deadline,
	}

	id, err := h.Task.CreateTask(task)
	if err != nil {
		c.JSON(500, map[string]any{
			"error": "failed to create anew task",
		})
		return
	}

	c.JSON(201, map[string]any{
		"id": id,
	})
}

func (h *Handler) getAllTasks(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		c.JSON(500, map[string]any{
			"error": err.Error(),
		})
		return
	}

	userRole, err := GetUserRole(c)
	if err != nil {
		c.JSON(500, map[string]any{
			"error": err.Error(),
		})
		return
	}

	if strings.ToLower(userRole) != "superuser" {
		c.JSON(400, map[string]any{
			"error": "You are not allowed to create a project",
		})
		return
	}

	tasks, err := h.Task.GetAllTasks(userId)
	if err != nil {
		c.JSON(500, map[string]any{
			"error": "failed to get the list of tasks",
		})
		return
	}

	c.JSON(200, map[string]any{
		"tasks": tasks,
	})
}

func (h *Handler) getTaskById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		c.JSON(500, map[string]any{
			"error": err.Error(),
		})
		return
	}

	userRole, err := GetUserRole(c)
	if err != nil {
		c.JSON(500, map[string]any{
			"error": err.Error(),
		})
		return
	}

	if strings.ToLower(userRole) != "superuser" {
		c.JSON(400, map[string]any{
			"error": "You are not allowed to create a project",
		})
		return
	}

	taskId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, map[string]any{
			"error": "invalid type of param",
		})
		return
	}

	task, err := h.Task.GetTaskById(userId, taskId)
	if err != nil {
		c.JSON(500, map[string]any{
			"error": "failed to get the task",
		})
		return
	}

	c.JSON(200, map[string]any{
		"task": task,
	})
}

func (h *Handler) updateTask(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		c.JSON(500, map[string]any{
			"error": err.Error(),
		})
		return
	}

	userRole, err := GetUserRole(c)
	if err != nil {
		c.JSON(500, map[string]any{
			"error": err.Error(),
		})
		return
	}

	if strings.ToLower(userRole) != "superuser" {
		c.JSON(400, map[string]any{
			"error": "You are not allowed to create a project",
		})
		return
	}

	taskId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, map[string]any{
			"error": "invalid type of param",
		})
		return
	}

	var in taskIn
	if err := c.BindJSON(&in); err != nil {
		c.JSON(400, map[string]any{
			"error": "invalid JSON provided",
		})
		return
	}

	var task = models.Task{
		ID:           taskId,
		Title:        in.Title,
		Description:  in.Description,
		ControllerId: userId,
		ExecutorId:   in.ExecutorId,
		Status:       in.Status,
		ProjectId:    in.ProjectId,
		Deadline:     in.Deadline,
		IsActive:     true,
	}

	if err := h.Task.UpdateTask(task); err != nil {
		c.JSON(500, map[string]any{
			"error": "failed to update the task",
		})
		return
	}

	c.JSON(200, map[string]any{
		"message": "task updated successfully",
	})
}

func (h *Handler) deleteTask(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		c.JSON(500, map[string]any{
			"error": err.Error(),
		})
		return
	}

	userRole, err := GetUserRole(c)
	if err != nil {
		c.JSON(500, map[string]any{
			"error": err.Error(),
		})
		return
	}

	if strings.ToLower(userRole) != "superuser" {
		c.JSON(400, map[string]any{
			"error": "You are not allowed to create a project",
		})
		return
	}

	taskId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, map[string]any{
			"error": "invalid type of param",
		})
		return
	}

	if err := h.Task.DeleteTask(userId, taskId); err != nil {
		c.JSON(500, map[string]any{
			"error": "failed to delete the task",
		})
		return
	}

	c.JSON(200, map[string]any{
		"message": "task deleted successfully",
	})
}

func (h *Handler) restoreTask(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		c.JSON(500, map[string]any{
			"error": err.Error(),
		})
		return
	}

	userRole, err := GetUserRole(c)
	if err != nil {
		c.JSON(500, map[string]any{
			"error": err.Error(),
		})
		return
	}

	if strings.ToLower(userRole) != "superuser" {
		c.JSON(400, map[string]any{
			"error": "You are not allowed to create a project",
		})
		return
	}

	taskId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, map[string]any{
			"error": "invalid type of param",
		})
		return
	}

	if err := h.Task.RestoreTask(userId, taskId); err != nil {
		c.JSON(500, map[string]any{
			"error": "failed to restore the task",
		})
		return
	}

	c.JSON(200, map[string]any{
		"message": "task restored successfully",
	})
}
