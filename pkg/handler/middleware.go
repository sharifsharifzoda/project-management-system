package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func (h *Handler) authMiddleware(c *gin.Context) {
	header := c.GetHeader("Authorization")

	if header == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"reason": "empty auth header",
		})
		return
	}

	split := strings.Split(header, " ")
	if len(split) != 2 || split[0] != "Alif" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"reason": "invalid auth header",
		})
		return
	}

	if len(split[1]) == 0 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"reason": "token is empty",
		})
		return
	}

	userId, userRole, err := h.Auth.ParseToken(split[1])
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Set("userId", userId)
	c.Set("userRole", userRole)
}

func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get("userId")
	if !ok {
		return 0, errors.New("userId not found")
	}

	idInt, ok := id.(int)
	if !ok {
		return 0, errors.New("invalid type of userId")
	}

	return idInt, nil
}

func GetUserRole(c *gin.Context) (string, error) {
	role, ok := c.Get("userRole")
	if !ok {
		return "", errors.New("userRole not found")
	}

	roleStr, ok := role.(string)
	if !ok {
		return "", errors.New("invalid type of userRole")
	}

	return roleStr, nil
}
