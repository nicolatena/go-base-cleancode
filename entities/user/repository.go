package user

import (
    "github.com/gin-gonic/gin"
)

// Repository represent the user's repository contract
type Repository interface {
	Fetch(c *gin.Context)
	Store(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}