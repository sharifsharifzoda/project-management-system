package service

import (
	"github.com/sharifsharifzoda/project-management-system/logging"
	"github.com/sharifsharifzoda/project-management-system/models"
	"github.com/sharifsharifzoda/project-management-system/pkg/repository"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type Authorization interface {
	ValidateUser(user models.User) error
	IsEmailUsed(email string) bool
	CreateUser(user *models.User) (int, error)
	CheckUser(user models.User) (models.User, error)
	GenerateToken(user models.User) (string, error)
	ParseToken(token string) (int, string, error)
}

type User interface {
	GetUser(id int) (models.User, error)
	UpdateUser(newUser models.User) error
	DeleteUser(id int) error
	Restore(id int) error
	GetProjects(userId int) ([]models.ProjectParticipant, error)
	GetTasks(userId int) (models.Tasks, error)
	UploadUserPicture(id int, filepath string) (models.User, error)
	UpdatePictureUser(id int, filepath string) (models.User, error)
}

type Project interface {
	CreateProject(project models.Project) (int, error)
	GetAllProjects(userId int) (models.Projects, error)
	GetProjectById(userId, projectId int) (models.Project, error)
	UpdateProject(project models.Project) error
	DeleteProject(userId, projectId int) error
	GetDeletedProjects(userId int) (models.Projects, error)
	Restore(userId, projectId int) error
	AddUserToProject(managerId int, propar models.ProjectParticipant) error
}

type Task interface {
	CreateTask(task models.Task) (int, error)
	GetAllTasks(userId int) (models.Tasks, error)
	GetTaskById(userId, taskId int) (models.Task, error)
	UpdateTask(task models.Task) error
	DeleteTask(userId, taskId int) error
	RestoreTask(userId, taskId int) error
}

type Service struct {
	Auth    Authorization
	User    User
	Project Project
	Task    Task
	Logger  *logging.Logger
}

func NewService(repository *repository.Repository, log *logging.Logger) *Service {
	return &Service{
		Auth:    NewAuthService(repository.Authorization, log),
		User:    NewUserService(repository.User),
		Project: NewProjectService(repository.Project),
		Task:    NewTaskService(repository.Task),
		Logger:  log,
	}
}
