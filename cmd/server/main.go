// go get - u github.com/gin-gonic/gin gorm.io/gorm gorm.io/driver/postgres github.com/spf13/viper github.com/google/uuid github.com/golang-jwt/jwt/v5
package main

import (
	"fmt"
	"log"

	"example.com/restful-api-example/configs"
	"example.com/restful-api-example/internal/models"
	"example.com/restful-api-example/internal/routes"
)

func main() {
	config, err := configs.LoadConfig("./")
	if err != nil {
		fmt.Println("enviroment load fail")
		return
	}

	configs.ConnectDatabase(&config)

	err = configs.Db_todos.AutoMigrate(&models.Todo{}, &models.User{})
	if err != nil {
		return
	}

	route := routes.InitRouter()
	log.Fatal(route.Run(":" + config.ServerPort))
}
