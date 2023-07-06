package repository

import (
	"github.com/sharifsharifzoda/project-management-system/models"
	"gorm.io/gorm"
)

type Authorization interface {
	CreateUser(user *models.User) (int, error)
	GetUser(email string) (models.User, error)
	IsEmailUsed(email string) bool
}

type User interface {
	GetUser(id int) (models.User, error)
	UpdateUser(newUser models.User) error
	DeleteUser(id int) error
	RestoreUser(id int) error
	GetProjects(userId int) ([]models.ProjectParticipant, error)
	GetTasks(userId int) (models.Tasks, error)
}

type Project interface {
	CreateProject(project models.Project) (int, error)
	GetAllProjects(userId int) (models.Projects, error)
	GetProjectById(userId, projectId int) (models.Project, error)
	UpdateProject(project models.Project) error
	DeleteProject(userId, projectId int) error
	GetDeletedProjects(userId int) (models.Projects, error)
	RestoreProject(userId, projectId int) error
	AddUserToProject(propar models.ProjectParticipant) error
}

type Task interface {
	CreateTask(task models.Task) (int, error)
	GetTasks(userId int) (models.Tasks, error)
	GetTaskById(userId, taskId int) (models.Task, error)
	UpdateTask(task models.Task) error
	DeleteTask(userId, taskId int) error
	RestoreTask(userId, taskId int) error
}

type Repository struct {
	Authorization
	User
	Project
	Task
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		User:          NewUserRepository(db),
		Project:       NewProjectRepo(db),
		Task:          NewTaskRepo(db),
	}
}
