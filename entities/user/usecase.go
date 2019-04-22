package user

import (
    "github.com/gin-gonic/gin"
    "github.com/jinzhu/gorm"
)

type InDB struct {
    DB *gorm.DB
}

// Usecase represent the user's usecases
type Usecase interface {
	Fetch(c *gin.Context)
	Store(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}