package repository

import (
	"github.com/sharifsharifzoda/project-management-system/models"
	"gorm.io/gorm"
	"log"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (u *UserRepository) GetUser(id int) (models.User, error) {
	var user models.User
	err := u.db.Where("id = ? AND is_active = ?", id, true).First(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (u *UserRepository) UpdateUser(newUser models.User) error {
	if err := u.db.Save(&newUser).Error; err != nil {
		return err
	}

	return nil
}

func (u *UserRepository) DeleteUser(id int) error {
	tx := u.db.Model(&models.User{}).Where("id = ? AND is_active = true", id).
		Update("is_active", false)

	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (u *UserRepository) RestoreUser(id int) error {
	tx := u.db.Model(&models.User{}).Where("id = ? AND is_active = false", id).
		Update("is_active", true)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (u *UserRepository) GetProjects(userId int) ([]models.ProjectParticipant, error) {
	var projects []models.ProjectParticipant
	rows, err := u.db.Model(&models.ProjectParticipant{}).Joins("inner join users on project_participants.participant_id = users.id").
		Select([]string{"project_participants.id", "project_participants.role", "project_participants.project_id"}).
		Where("project_participants.participant_id = ?", userId).Rows()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var pro models.ProjectParticipant
		err := rows.Scan(&pro.ID, &pro.Role, &pro.ProjectId)
		if err != nil {
			log.Println("error while scanning from row")
			return nil, err
		}

		projects = append(projects, pro)
	}

	return projects, nil
}

func (u *UserRepository) GetTasks(userId int) (models.Tasks, error) {
	var tasks models.Tasks
	rows, err := u.db.Model(models.Task{}).Joins("inner join users on tasks.executor_id = users.id").
		Joins("inner join projects on tasks.project_id = projects.id").
		Select([]string{"tasks.id", "tasks.title", "tasks.description", "users.firstname", "tasks.status",
			"projects.name", "tasks.deadline"}).
		Where("tasks.executor_id = ? AND tasks.is_active = ?", userId, true).Rows()
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
