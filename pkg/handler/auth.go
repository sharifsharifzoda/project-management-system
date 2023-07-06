package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sharifsharifzoda/project-management-system/models"
	"net/http"
)

type signUpData struct {
	FirstName string `json:"firstname" biding:"required"`
	LastName  string `json:"lastname" biding:"required"`
	Email     string `json:"email" biding:"required"`
	Password  string `json:"password" biding:"required"`
}

type signInData struct {
	Email    string `json:"email" biding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) signUp(c *gin.Context) {
	var request signUpData

	if err := c.BindJSON(&request); err != nil {
		c.JSON(400, map[string]any{
			"error": "invalid JSON provided",
		})
		return
	}

	var user = models.User{
		Firstname: request.FirstName,
		Lastname:  request.LastName,
		Email:     request.Email,
		Password:  request.Password,
	}

	err := h.Auth.ValidateUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "validate",
		})
		return
	}

	isUsed := h.Auth.IsEmailUsed(user.Email)
	if isUsed {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "email is already created",
		})
		return
	}

	id, err := h.Auth.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to create user",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id": id,
	})
}

func (h *Handler) signIn(c *gin.Context) {
	var request signInData
	if c.ShouldBindJSON(&request) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid JSON provided",
		})
		return
	}

	var user = models.User{
		Email:    request.Email,
		Password: request.Password,
	}

	err := h.Auth.ValidateUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "validate",
		})
		return
	}

	checkedUser, err := h.Auth.CheckUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	token, err := h.Auth.GenerateToken(checkedUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Writer.Header().Set("Authorization", token)

	c.JSON(200, map[string]any{
		"msg": "signed in",
	})
}
