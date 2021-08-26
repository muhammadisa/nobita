package middleware

import (
	"errors"
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
		roleID, ok := c.Get(RoleID)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": errors.New("role id not passed from middleware")})
		}

		err := featureRepo.ReadFeature(c, roleID.(int64), c.FullPath())
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		c.Next()
	}
}

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.Request.Header.Get("Authorization")
		j := jwt.NewJWT(auth, "SECRET")
		keys, err := j.ExtractKeys([]string{UserID, RoleID})
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		id, err := strconv.ParseInt(keys[UserID], 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "user id not passed"})
			return
		}

		roleID, err := strconv.ParseInt(keys[RoleID], 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "role id not passed"})
			return
		}

		c.Set(UserID, id)
		c.Set(RoleID, roleID)
		c.Next()
	}
}
