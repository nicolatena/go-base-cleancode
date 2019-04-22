package main

import (
	"fmt"

    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/viper"
    "github.com/gin-gonic/gin"


	_userHttpDeliver "go-base-cleancode/entities/user/delivery/http"
	// _articleRepo "github.com/bxcodec/go-clean-arch/article/repository"
	// _articleUcase "github.com/bxcodec/go-clean-arch/article/usecase"
	// _authorRepo "github.com/bxcodec/go-clean-arch/author/repository"
)

func init() {
	viper.SetConfigFile(`go-base-cleancode/config.json`)
	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		fmt.Println("Service RUN on DEBUG mode")
	}

}

func main() {

	dbHost := viper.GetString(`database.host`)
	dbPort := viper.GetString(`database.port`)
	dbUser := viper.GetString(`database.user`)
	dbPass := viper.GetString(`database.pass`)
	dbName := viper.GetString(`database.name`)

	connection := fmt.Sprintf("host=%s port=%s user=%s dbname=%s  password=%s", dbHost, dbPort, dbUser, dbName, dbPass)
	
	dbConn, err := gorm.Open(`postgres`, connection)
	defer dbConn.Close()
	if err != nil && viper.GetBool("debug") {
		fmt.Println(err)
	}

	
	// add struct to add table in database
	// dbConn.Debug().AutoMigrate(&User{})



	route := gin.Default()


	_userHttpDeliver.NewUserHttpHandler(route, dbConn)


	route.Run(viper.GetString("server.address"))

}