package auth

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/muhammadisa/nobita/middleware"
	"github.com/muhammadisa/nobita/model/v1/auth"
	"net/http"
)

func (r route) UpdateProfile() (string, gin.HandlerFunc) {
	return "/profile", func(c *gin.Context) {
		userID, ok := c.Get(middleware.UserID)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": errors.New("user id not passed from middleware")})
		}

		var json auth.Profile
		err := c.ShouldBindJSON(&json)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		json.AccountID = userID.(int64)

		err = r.UseCase.EditProfile(c, json)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		}
	}
}
