package repository

import (
	"github.com/sharifsharifzoda/project-management-system/models"
	"gorm.io/gorm"
	"log"
)

type TaskRepo struct {
	db *gorm.DB
}

func NewTaskRepo(db *gorm.DB) *TaskRepo {
	return &TaskRepo{db: db}
}

func (t *TaskRepo) CreateTask(task models.Task) (int, error) {
	err := t.db.Model(&models.Task{}).Create(&task).Error
	if err != nil {
		return -1, err
	}

	return task.ID, nil
}

func (t *TaskRepo) GetTasks(userId int) (models.Tasks, error) {
	var tasks models.Tasks
	rows, err := t.db.Model(models.Task{}).Joins("inner join users on tasks.executor_id = users.id").
		Joins("inner join projects on tasks.project_id = projects.id").
		Select([]string{"tasks.id", "tasks.title", "tasks.description", "users.firstname", "tasks.status",
			"projects.name", "tasks.deadline"}).
		Where("tasks.controller_id = ? AND tasks.is_active = ?", userId, true).Rows()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var task models.Task
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.ExecutorName, &task.Status, &task.ProjectName,
			&task.Deadline)
		if err != nil {
			log.Println("error while scanning from row")
			return nil, err
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (t *TaskRepo) GetTaskById(userId, taskId int) (models.Task, error) {
	var task models.Task
	row := t.db.Model(models.Task{}).Joins("inner join users on tasks.executor_id = users.id").
		Joins("inner join projects on tasks.project_id = projects.id").
		Select([]string{"tasks.id", "tasks.title", "tasks.description", "users.firstname", "tasks.status",
			"projects.name", "tasks.deadline"}).
		Where("tasks.controller_id = ? AND tasks.id = ? AND tasks.is_active = ?", userId, taskId, true).Row()

	err := row.Scan(&task.ID, &task.Title, &task.Description, &task.ExecutorName, &task.Status, &task.ProjectName,
		&task.Deadline)
	if err != nil {
		return models.Task{}, err
	}

	return task, nil
}

func (t *TaskRepo) UpdateTask(task models.Task) error {
	err := t.db.Save(&task).Error
	if err != nil {
		return err
	}

	return nil
}

func (t *TaskRepo) DeleteTask(userId, taskId int) error {
	err := t.db.Model(&models.Task{}).Where("id = ? AND controller_id = ? AND is_active = ?", taskId, userId, true).
		Update("is_active", false).Error
	if err != nil {
		return err
	}

	return nil
}

func (t *TaskRepo) RestoreTask(userId, taskId int) error {
	err := t.db.Model(&models.Task{}).Where("id = ? AND controller_id = ? AND is_active = ?", taskId, userId, false).
		Update("is_active", true).Error
	if err != nil {
		return err
	}

	return nil
}
