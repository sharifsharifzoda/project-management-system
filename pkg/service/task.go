package service

import (
	"github.com/sharifsharifzoda/project-management-system/models"
	"github.com/sharifsharifzoda/project-management-system/pkg/repository"
	"log"
)

type TaskService struct {
	repo repository.Task
}

func NewTaskService(repo repository.Task) *TaskService {
	return &TaskService{repo: repo}
}

func (t *TaskService) CreateTask(task models.Task) (int, error) {
	id, err := t.repo.CreateTask(task)
	if err != nil {
		log.Println("failed to create a new task. Error is: ", err.Error())
		return -1, err
	}

	return id, nil
}

func (t *TaskService) GetAllTasks(userId int) (models.Tasks, error) {
	tasks, err := t.repo.GetTasks(userId)
	if err != nil {
		log.Println("failed to get the list of tasks. Error is: ", err.Error())
		return nil, err
	}

	return tasks, nil
}

func (t *TaskService) GetTaskById(userId, taskId int) (models.Task, error) {
	task, err := t.repo.GetTaskById(userId, taskId)
	if err != nil {
		log.Println("failed to get the task. Error is: ", err.Error())
		return models.Task{}, err
	}

	return task, nil
}

func (t *TaskService) UpdateTask(task models.Task) error {
	err := t.repo.UpdateTask(task)
	if err != nil {
		log.Println("failed to update the task. Error is: ", err.Error())
		return err
	}

	return nil
}

func (t *TaskService) DeleteTask(userId, taskId int) error {
	err := t.repo.DeleteTask(userId, taskId)
	if err != nil {
		log.Println("failed to delete the task. Error is: ", err.Error())
		return err
	}

	return nil
}

func (t *TaskService) RestoreTask(userId, taskId int) error {
	err := t.repo.RestoreTask(userId, taskId)
	if err != nil {
		log.Println("failed to restore the task. Error is: ", err.Error())
		return err
	}

	return nil
}
