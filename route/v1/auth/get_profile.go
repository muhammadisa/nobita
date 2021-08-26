package auth

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/muhammadisa/nobita/middleware"
	"net/http"
)

func (r route) GetProfile() (string, gin.HandlerFunc) {
	return "/profile", func(c *gin.Context) {
		userID, ok := c.Get(middleware.UserID)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": errors.New("user id not passed from middleware")})
		}
		profile, err := r.UseCase.MyProfile(c, userID.(int64))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		}
		c.JSON(http.StatusOK, gin.H{"profile": profile})
	}
}
