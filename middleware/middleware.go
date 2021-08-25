package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/muhammadisa/nobita/util/jwt"
	"net/http"
	"strconv"
)

const (
	UserID = `user_id`
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.Request.Header.Get("Authorization")
		j := jwt.NewJWT(auth, "SECRET")
		userID, err := j.ExtractKey("user_id")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		id, err := strconv.Atoi(userID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "user id not passed"})
			return
		}
		c.Set(UserID, id)
		c.Next()
	}
}
