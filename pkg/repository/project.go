package repository

import (
	"github.com/sharifsharifzoda/project-management-system/models"
	"gorm.io/gorm"
	"log"
)

type ProjectRepo struct {
	db *gorm.DB
}

func NewProjectRepo(db *gorm.DB) *ProjectRepo {
	return &ProjectRepo{db: db}
}

func (p *ProjectRepo) CreateProject(project models.Project) (int, error) {
	err := p.db.Create(&project).Error
	if err != nil {
		return -1, err
	}

	return project.ID, nil
}

func (p *ProjectRepo) GetAllProjects(userId int) (models.Projects, error) {
	var projects models.Projects
	rows, err := p.db.Model(&models.Project{}).Joins("inner join users on projects.manager_id = users.id").
		Select([]string{"projects.id", "projects.name", "projects.description", "projects.department",
			"projects.status", "projects.start_date", "projects.deadline", "users.firstname"}).
		Where("projects.is_active = ? AND projects.manager_id = ?", true, userId).Rows()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var pro models.Project
		err := rows.Scan(&pro.ID, &pro.Name, &pro.Description, &pro.Department, &pro.Status, &pro.StartDate,
			&pro.Deadline, &pro.ManagerName)
		if err != nil {
			log.Println("error while scanning from row")
			return nil, err
		}

		projects = append(projects, pro)
	}

	return projects, nil
}

func (p *ProjectRepo) GetProjectById(userId, projectId int) (models.Project, error) {
	var pro models.Project
	row := p.db.Model(&models.Project{}).Joins("inner join users on projects.manager_id = users.id").
		Select([]string{"projects.id", "projects.name", "projects.description", "projects.department",
			"projects.status", "projects.start_date", "projects.deadline", "users.firstname"}).
		Where("projects.is_active = ? AND projects.manager_id = ? AND projects.id = ?",
			true, userId, projectId).Row()

	if err := row.Scan(&pro.ID, &pro.Name, &pro.Description, &pro.Department, &pro.Status, &pro.StartDate,
		&pro.Deadline, &pro.ManagerName); err != nil {
		return models.Project{}, err
	}

	return pro, nil
}

func (p *ProjectRepo) UpdateProject(project models.Project) error {
	err := p.db.Where("projects.id = ? AND projects.manager_id = ?", project.ID, project.ManagerID).
		Save(&project).Error
	if err != nil {
		return err
	}

	return nil
}

func (p *ProjectRepo) DeleteProject(userId, projectId int) error {
	err := p.db.Model(&models.Project{}).Where("id = ? AND manager_id = ? AND is_active = ?",
		projectId, userId, true).
		Update("is_active", false).Error

	if err != nil {
		return err
	}

	return nil
}

func (p *ProjectRepo) GetDeletedProjects(userId int) (models.Projects, error) {
	var projects models.Projects
	rows, err := p.db.Model(&models.Project{}).Joins("inner join users on projects.manager_id = users.id").
		Select([]string{"projects.id", "projects.name", "projects.description", "projects.department",
			"projects.status", "projects.start_date", "projects.deadline", "users.firstname"}).
		Where("projects.is_active = ? AND projects.manager_id = ?", false, userId).Rows()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var pro models.Project
		err := rows.Scan(&pro.ID, &pro.Name, &pro.Description, &pro.Department, &pro.Status, &pro.StartDate,
			&pro.Deadline, &pro.ManagerName)
		if err != nil {
			log.Println("error while scanning from row")
			return nil, err
		}

		projects = append(projects, pro)
	}

	return projects, nil
}

func (p *ProjectRepo) RestoreProject(userId, projectId int) error {
	err := p.db.Model(&models.Project{}).Where("id = ? AND manager_id = ?", projectId, userId).
		Update("is_active", true).Error
	if err != nil {
		return err
	}

	return nil
}

func (p *ProjectRepo) AddUserToProject(propar models.ProjectParticipant) error {
	err := p.db.Create(&propar).Error
	if err != nil {
		return err
	}

	return nil
}
