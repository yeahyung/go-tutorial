package main

import (
	"github.com/gin-gonic/gin"
	"go-tutorial/gin-tutorial/config"
	"go-tutorial/gin-tutorial/handler"
	"go-tutorial/gin-tutorial/repository"
	"go-tutorial/gin-tutorial/router"
	"log"
)

func init() {

}

func main() {
	gin.SetMode(config.GinMode)
	r := gin.Default()
	server := r.Group("/gin-tutorial")

	db, err := config.SetUpDB()
	if err != nil {
		log.Fatalln(err.Error())
	}

	personRepo := repository.NewGormPersonRepo(db)
	serverHandler := handler.NewServerHandler(personRepo)
	router.NewServerRouter(server, serverHandler)

	r.Run()
}
