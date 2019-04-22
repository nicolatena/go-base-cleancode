package config


import (
	"github.com/jinzhu/gorm"
    . "go-base-cleancode/models"
)


type InDB struct {
	DB *gorm.DB
}

func Migrate(idb *gorm.DB) {

	idb.Debug().AutoMigrate(&User{})

}