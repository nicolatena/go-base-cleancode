package main

import (
    _ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/viper"
    "github.com/gin-gonic/gin"

	"go-base-cleancode/infrastructure/config"
	_userHttpDeliver "go-base-cleancode/entities/user/delivery/http"
	_productHttpDeliver "go-base-cleancode/entities/product/delivery/http"
	// _articleRepo "github.com/bxcodec/go-clean-arch/article/repository"
	// _articleUcase "github.com/bxcodec/go-clean-arch/article/usecase"
	// _authorRepo "github.com/bxcodec/go-clean-arch/author/repository"
)


func main() {
	db := config.Init()

	
	route := gin.Default()


	_userHttpDeliver.NewUserHttpHandler(route, db)
	_productHttpDeliver.NewProductHttpHandler(route, db)


	route.Run(viper.GetString("server.address"))
}