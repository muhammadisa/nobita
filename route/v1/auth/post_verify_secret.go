package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/muhammadisa/nobita/model/v1/auth"
	"net/http"
)

func (r route) PostVerifySecret() (string, gin.HandlerFunc) {
	return "/auth/verify", func(c *gin.Context) {
		var json auth.Secret
		err := c.ShouldBindJSON(&json)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		verifyResult, err := r.UseCase.VerifySecret(c, json)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"auth": verifyResult})
	}
}
