package _interface

import "github.com/gin-gonic/gin"

type Route interface {
	GetFoo() (string, gin.HandlerFunc)
}
