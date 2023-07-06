package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sharifsharifzoda/project-management-system/models"
	"strconv"
	"strings"
)

type dataIn struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Department  string `json:"department" binding:"required"`
	Status      string `json:"status"`
	StartDate   string `json:"start_date,omitempty"`
	Deadline    string `json:"deadline" binding:"required"`
}

func (h *Handler) createProject(c *gin.Context) {
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

	var data dataIn
	if err := c.BindJSON(&data); err != nil {
		c.JSON(400, map[string]any{
			"error": "invalid JSON provided",
		})
		return
	}

	var project = models.Project{
		Name:        data.Name,
		Description: data.Description,
		Department:  data.Department,
		ManagerID:   userId,
		Status:      data.Status,
		StartDate:   data.StartDate,
		Deadline:    data.Deadline,
	}

	id, err := h.Project.CreateProject(project)
	if err != nil {
		c.JSON(500, map[string]any{
			"error": "failed to create a new project",
		})
		return
	}

	c.JSON(201, map[string]any{
		"id": id,
	})
}

func (h *Handler) getAllProjects(c *gin.Context) {
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
			"error": "You are not allowed to get all projects",
		})
		return
	}

	projects, err := h.Project.GetAllProjects(userId)
	if err != nil {
		c.JSON(400, map[string]any{
			"error": "failed to get the list of projects",
		})
		return
	}

	if projects == nil {
		c.JSON(200, map[string]any{
			"message": "there is no any project",
		})
		return
	}

	c.JSON(200, map[string]any{
		"projects": projects,
	})
}

func (h *Handler) getProjectById(c *gin.Context) {
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
			"error": "You are not allowed to get a project",
		})
		return
	}

	projectId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, map[string]any{
			"error": "invalid type of param",
		})
		return
	}

	project, err := h.Project.GetProjectById(userId, projectId)
	if err != nil {
		c.JSON(400, map[string]any{
			"error": "project doesn't exist",
		})
		return
	}

	c.JSON(200, map[string]any{
		"project": project,
	})
}

func (h *Handler) updateProject(c *gin.Context) {
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
			"error": "You are not allowed to update a project",
		})
		return
	}

	projectId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, map[string]any{
			"error": "invalid type of param",
		})
		return
	}

	var pro dataIn
	if err = c.BindJSON(&pro); err != nil {
		c.JSON(400, map[string]any{
			"error": "invalid JSON provided",
		})
		return
	}

	var project = models.Project{
		ID:          projectId,
		Name:        pro.Name,
		Description: pro.Description,
		Department:  pro.Department,
		ManagerID:   userId,
		Status:      pro.Status,
		StartDate:   pro.StartDate,
		Deadline:    pro.Deadline,
		IsActive:    true,
	}

	if err := h.Project.UpdateProject(project); err != nil {
		c.JSON(400, map[string]any{
			"error": "failed to update the project",
		})
		return
	}

	c.JSON(200, map[string]any{
		"message": "project updated successfully",
	})
}

func (h *Handler) deleteProject(c *gin.Context) {
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
			"error": "You are not allowed to delete a project",
		})
		return
	}

	projectId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, map[string]any{
			"error": "invalid type of param",
		})
		return
	}

	if err := h.Project.DeleteProject(userId, projectId); err != nil {
		c.JSON(500, map[string]any{
			"error": "failed to delete the project",
		})
		return
	}

	c.JSON(200, map[string]any{
		"message": "project deleted successfully",
	})
}

func (h *Handler) getDeletedProjects(c *gin.Context) {
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
			"error": "You are not allowed to get a list of deleted projects",
		})
		return
	}

	projects, err := h.Project.GetDeletedProjects(userId)
	if err != nil {
		c.JSON(400, map[string]any{
			"error": "failed to get the list of deleted projects",
		})
		return
	}

	if projects == nil {
		c.JSON(200, map[string]any{
			"message": "there is no any deleted project",
		})
		return
	}

	c.JSON(200, map[string]any{
		"projects": projects,
	})
}

func (h *Handler) restoreProject(c *gin.Context) {
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
			"error": "You are not allowed to restore a project",
		})
		return
	}

	projectId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, map[string]any{
			"error": "invalid type of param",
		})
		return
	}

	if err := h.Project.Restore(userId, projectId); err != nil {
		c.JSON(500, map[string]any{
			"error": "failed to restore the project",
		})
		return
	}

	c.JSON(200, map[string]any{
		"message": "project restored successfully",
	})
}

func (h *Handler) addUserToProject(c *gin.Context) {
	managerId, err := getUserId(c)
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
			"error": "You are not allowed to add participants to a project",
		})
		return
	}

	var propar models.ProjectParticipant
	if err := c.BindJSON(&propar); err != nil {
		c.JSON(400, map[string]any{
			"error": "invalid JSON provided",
		})
		return
	}

	if err := h.Project.AddUserToProject(managerId, propar); err != nil {
		c.JSON(400, map[string]any{
			"error": "failed to add a new participant to the project",
		})
		return
	}

	c.JSON(200, map[string]any{
		"message": "added a new participant to the project",
	})
}
