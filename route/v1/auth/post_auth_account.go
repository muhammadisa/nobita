package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/muhammadisa/nobita/model/v1/auth"
	"net/http"
)

func (r route) PostAuthAccount() (string, gin.HandlerFunc) {
	return "/auth", func(c *gin.Context) {
		var json auth.Account
		err := c.ShouldBindJSON(&json)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		authResult, err := r.UseCase.AuthAccount(c, json)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"auth": authResult})
	}
}
