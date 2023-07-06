package service

import (
	"fmt"
	"github.com/sharifsharifzoda/project-management-system/models"
	"github.com/sharifsharifzoda/project-management-system/pkg/repository"
	"log"
	"os"
	"time"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (u *UserService) GetUser(id int) (models.User, error) {
	user, err := u.repo.GetUser(id)
	if err != nil {
		log.Println("failed to get the user. error is: ", err.Error())
		return models.User{}, err
	}
	return user, nil
}

func (u *UserService) UpdateUser(newUser models.User) error {
	err := u.repo.UpdateUser(newUser)
	if err != nil {
		log.Println("failed to update the user. error is: ", err.Error())
		return err
	}

	return nil
}

func (u *UserService) DeleteUser(id int) error {
	err := u.repo.DeleteUser(id)
	if err != nil {
		log.Println("failed to delete the user. error is: ", err.Error())
		return err
	}

	return nil
}

func (u *UserService) Restore(id int) error {
	err := u.repo.RestoreUser(id)
	if err != nil {
		log.Println("failed to restore the user. error is: ", err.Error())
		return err
	}

	return nil
}

func (u *UserService) GetProjects(userId int) ([]models.ProjectParticipant, error) {
	projects, err := u.repo.GetProjects(userId)
	if err != nil {
		log.Println("failed to get the list of projects. Error is: ", err.Error())
		return nil, err
	}

	return projects, nil
}

func (u *UserService) GetTasks(userId int) (models.Tasks, error) {
	tasks, err := u.repo.GetTasks(userId)
	if err != nil {
		log.Println("failed to get the list of tasks. Error is: ", err.Error())
		return nil, err
	}

	return tasks, nil
}

func (u *UserService) UploadUserPicture(id int, filepath string) (models.User, error) {
	user, err := u.GetUser(id)
	if err != nil {
		return models.User{}, err
	}

	user.Photo = filepath
	user.UpdatedAt = time.Now()

	if err := u.repo.UpdateUser(user); err != nil {
		log.Println("failed to update user while uploading photo. Error is: ", err.Error())
		return models.User{}, err
	}

	return user, nil
}

func (u *UserService) UpdatePictureUser(id int, filepath string) (models.User, error) {
	user, err := u.GetUser(id)
	if err != nil {
		return user, err
	}

	if err := os.Remove(fmt.Sprintf("./files/layouts/%s", user.Photo)); err != nil {
		return user, err
	}

	user.Photo = filepath
	user.UpdatedAt = time.Now()

	if err := u.repo.UpdateUser(user); err != nil {
		log.Println("failed to update user while changing the profile photo. Error is: ", err.Error())
		return models.User{}, err
	}

	return user, nil
}
