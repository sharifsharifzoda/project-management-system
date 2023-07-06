package service

import (
	"github.com/sharifsharifzoda/project-management-system/models"
	"github.com/sharifsharifzoda/project-management-system/pkg/repository"
	"log"
)

type ProjectService struct {
	repo repository.Project
}

func NewProjectService(repo repository.Project) *ProjectService {
	return &ProjectService{repo: repo}
}

func (p *ProjectService) CreateProject(project models.Project) (int, error) {
	id, err := p.repo.CreateProject(project)
	if err != nil {
		log.Println("failed to create a new project. Error is: ", err.Error())
		return -1, err
	}

	return id, nil
}

func (p *ProjectService) GetAllProjects(userId int) (models.Projects, error) {
	projects, err := p.repo.GetAllProjects(userId)
	if err != nil {
		log.Println("failed to get the list of projects. Error is: ", err.Error())
		return nil, err
	}

	return projects, nil
}

func (p *ProjectService) GetProjectById(userId, projectId int) (models.Project, error) {
	project, err := p.repo.GetProjectById(userId, projectId)
	if err != nil {
		log.Println("failed to get the project by id. Error is: ", err.Error())
		return models.Project{}, err
	}

	return project, nil
}

func (p *ProjectService) UpdateProject(project models.Project) error {
	_, err := p.repo.GetProjectById(project.ManagerID, project.ID)
	if err != nil {
		log.Println("you don't have any project. Error is: ", err.Error())
		return err
	}

	if err := p.repo.UpdateProject(project); err != nil {
		log.Println("failed to update the project. Error is: ", err.Error())
		return err
	}

	return nil
}

func (p *ProjectService) DeleteProject(userId, projectId int) error {
	err := p.repo.DeleteProject(userId, projectId)
	if err != nil {
		log.Println("failed to delete the project. Error is: ", err.Error())
		return err
	}

	return nil
}

func (p *ProjectService) GetDeletedProjects(userId int) (models.Projects, error) {
	projects, err := p.repo.GetDeletedProjects(userId)
	if err != nil {
		log.Println("failed to get the list of deleted projects. Error is: ", err.Error())
		return nil, err
	}

	return projects, nil
}

func (p *ProjectService) Restore(userId, projectId int) error {
	err := p.repo.RestoreProject(userId, projectId)
	if err != nil {
		log.Println("failed to restore the project. Error is: ", err.Error())
		return err
	}

	return nil
}

func (p *ProjectService) AddUserToProject(managerId int, propar models.ProjectParticipant) error {
	_, err := p.repo.GetProjectById(managerId, propar.ProjectId)
	if err != nil {
		log.Println("failed to add a new participant to the project. Error is: ", err.Error())
		return err
	}

	if err := p.repo.AddUserToProject(propar); err != nil {
		log.Println("failed to add a new participant to the project. Error is: ", err.Error())
		return err
	}

	return nil
}
