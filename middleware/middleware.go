package middleware

import (
	"github.com/gin-gonic/gin"
	featurerepositoryv1 "github.com/muhammadisa/nobita/repository/v1/feature/interface"
	"github.com/muhammadisa/nobita/util/jwt"
	"net/http"

	"strconv"
)

const (
	UserID = `user_id`
	RoleID = `role_id`
)

func RoleAccessibleMiddleware(featureRepo featurerepositoryv1.RW) gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.Request.Header.Get("Authorization")
		j := jwt.NewJWT(auth, "SECRET")
		keys, err := j.ExtractKeys([]string{"user_id", "role_id"})
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		id, err := strconv.Atoi(keys["user_id"])
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "user id not passed"})
			return
		}

		roleID, err := strconv.Atoi(keys["role_id"])
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "role id not passed"})
			return
		}

		c.Set(UserID, id)
		c.Set(UserID, id)
		c.Next()
	}
}
