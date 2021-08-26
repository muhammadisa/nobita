package foo

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (r route) GetFoo() (string, gin.HandlerFunc) {
	return "/foo", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"foo": "ok"})
	}
}
