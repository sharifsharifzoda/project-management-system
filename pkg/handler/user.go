package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sharifsharifzoda/project-management-system/models"
	"github.com/sharifsharifzoda/project-management-system/utils"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func (h *Handler) getUser(c *gin.Context) {
	id, err := getUserId(c)
	if err != nil {
		c.JSON(500, map[string]any{
			"error": err.Error(),
		})
		return
	}

	user, err := h.User.GetUser(id)
	if err != nil {
		c.JSON(500, map[string]any{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func (h *Handler) updateUser(c *gin.Context) {
	id, err := getUserId(c)
	if err != nil {
		c.JSON(500, map[string]any{
			"error": err.Error(),
		})
		return
	}

	var request signUpData
	if err := c.BindJSON(&request); err != nil {
		c.JSON(400, map[string]any{
			"error": "invalid JSON provided",
		})
		return
	}

	var newUser = models.User{
		ID:        id,
		Firstname: request.FirstName,
		Lastname:  request.LastName,
		Email:     request.Email,
		Password:  request.Password,
		Role:      "user",
		IsActive:  true,
	}

	if err = h.Auth.ValidateUser(newUser); err != nil {
		c.JSON(400, map[string]any{
			"error": "validate",
		})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(500, map[string]any{
			"error": "failed to hash the password",
		})
		return
	}

	newUser.Password = string(hash)

	if err := h.User.UpdateUser(newUser); err != nil {
		c.JSON(500, map[string]any{
			"error": "failed to update the user",
		})
		return
	}

	c.JSON(200, map[string]any{
		"message": "user successfully updated",
	})
}

func (h *Handler) deleteUser(c *gin.Context) {
	id, err := getUserId(c)
	if err != nil {
		c.JSON(500, map[string]any{
			"error": err.Error(),
		})
		return
	}

	if err := h.User.DeleteUser(id); err != nil {
		c.JSON(500, map[string]any{
			"error": "failed to delete the user",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "user successfully deleted",
	})
}

func (h *Handler) restoreUser(c *gin.Context) {
	id, err := getUserId(c)
	if err != nil {
		c.JSON(500, map[string]any{
			"error": err.Error(),
		})
		return
	}

	if err := h.User.Restore(id); err != nil {
		c.JSON(500, map[string]any{
			"error": "failed to restore the user",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "successfully restored. Now you can sign in again.",
	})
}

func (h *Handler) getProjects(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		c.JSON(500, map[string]any{
			"error": err.Error(),
		})
		return
	}

	projects, err := h.User.GetProjects(userId)
	if err != nil {
		c.JSON(500, map[string]any{
			"error": "failed to get the list of projects",
		})
		return
	}

	if projects == nil {
		c.JSON(200, map[string]any{
			"message": "you don't have any project",
		})
		return
	}

	c.JSON(200, map[string]any{
		"projects": projects,
	})
}

func (h *Handler) getTasks(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		c.JSON(500, map[string]any{
			"error": err.Error(),
		})
		return
	}

	tasks, err := h.User.GetTasks(userId)
	if err != nil {
		c.JSON(400, map[string]any{
			"error": "failed ti get the list of tasks",
		})
		return
	}

	if tasks == nil {
		c.JSON(200, map[string]any{
			"message": "there is no any tasks",
		})
		return
	}

	c.JSON(200, map[string]any{
		"tasks": tasks,
	})
}

func (h *Handler) setProfilePhoto(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		c.JSON(400, map[string]any{
			"error": err.Error(),
		})
		return
	}

	image, err := c.FormFile("photo")
	if err != nil {
		c.JSON(400, map[string]any{
			"error": err.Error(),
		})
		return
	}

	filePath, err := utils.GenFilenameWithDir(image.Filename)
	if err != nil {
		c.JSON(http.StatusUnsupportedMediaType, map[string]any{
			"error": err.Error(),
		})
		return
	}

	user, err := h.User.UploadUserPicture(userId, filePath)
	if err != nil {
		c.JSON(400, map[string]any{
			"error": err.Error(),
		})
		return
	}

	if err := c.SaveUploadedFile(image, fmt.Sprintf("./files/%s", filePath)); err != nil {
		c.JSON(500, map[string]any{
			"error": "failed to upload file",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "profile photo uploaded successfully",
		"data":    user,
	})
}

func (h *Handler) changeProfilePhoto(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		c.JSON(500, map[string]any{
			"error": err.Error(),
		})
		return
	}

	image, err := c.FormFile("photo")
	if err != nil {
		c.JSON(400, map[string]any{
			"error": err.Error(),
		})
		return
	}

	filePath, err := utils.GenFilenameWithDir(image.Filename)
	if err != nil {
		c.JSON(http.StatusUnsupportedMediaType, map[string]any{
			"error": err.Error(),
		})
		return
	}

	user, err := h.User.UpdatePictureUser(userId, filePath)
	if err != nil {
		c.JSON(500, map[string]any{
			"error": err.Error(),
		})
		return
	}

	if err := c.SaveUploadedFile(image, fmt.Sprintf("./file/layouts/%s", filePath)); err != nil {
		c.JSON(500, map[string]any{
			"error": "failed to upload file",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    user,
	})

}
