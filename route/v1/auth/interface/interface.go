package _interface

import "github.com/gin-gonic/gin"

type Route interface {
	PostAuthAccount() (string, gin.HandlerFunc)
	PostVerifySecret() (string, gin.HandlerFunc)
	GetProfile() (string, gin.HandlerFunc)
	UpdateProfile() (string, gin.HandlerFunc)
}
