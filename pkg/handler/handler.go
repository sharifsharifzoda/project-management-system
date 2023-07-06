package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sharifsharifzoda/project-management-system/pkg/service"
)

type Handler struct {
	Auth    service.Authorization
	User    service.User
	Project service.Project
	Task    service.Task
}

func NewHandler(auth service.Authorization, user service.User, project service.Project, task service.Task) *Handler {
	return &Handler{
		Auth:    auth,
		User:    user,
		Project: project,
		Task:    task,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, map[string]any{
			"message": "pong",
		})
	})

	api := r.Group("/v1")
	{
		api.POST("/restore", h.restoreUser)

		auth := api.Group("/auth")
		{
			auth.POST("/sign-up", h.signUp)
			auth.POST("/sign-in", h.signIn)
		}

		user := api.Group("/user", h.authMiddleware)
		{
			user.GET("/", h.getUser)
			user.PUT("/", h.updateUser)
			user.DELETE("/", h.deleteUser)
			user.GET("/projects", h.getProjects)
			user.GET("/tasks", h.getTasks)
			user.POST("/photo", h.setProfilePhoto)
			user.PUT("/photo", h.changeProfilePhoto)
		}

		project := api.Group("/project", h.authMiddleware)
		{
			project.POST("/", h.createProject)
			project.GET("/", h.getAllProjects)
			project.GET("/:id", h.getProjectById)
			project.POST("/users", h.addUserToProject)
			project.PUT("/:id", h.updateProject)
			project.DELETE("/:id", h.deleteProject)
			project.GET("/deleted", h.getDeletedProjects)
			project.POST("/:id/restore", h.restoreProject)
			//project.GET("/:id/users", h.getParticipants)
		}

		task := api.Group("/task", h.authMiddleware)
		{
			task.POST("/", h.createTask)
			task.GET("/", h.getAllTasks)
			task.GET("/:id", h.getTaskById)
			task.PUT("/:id", h.updateTask)
			task.DELETE("/:id", h.deleteTask)
			task.POST("/:id/restore", h.restoreTask)
		}
	}

	return r
}
