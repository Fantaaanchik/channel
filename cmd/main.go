package main

import (
	"chans/config"
	"chans/internal/db"
	"chans/internal/handler"
	"chans/internal/repository"
	"chans/internal/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	config.ReadConfig("./config/config.json")
	fmt.Println("Connection to config success!!!")
	dbc := db.ConnectionToDB()
	fmt.Println("Connection to DB success!!!")
	defer db.CloseDB()

	repo := repository.NewRepo(dbc)

	serviceCon := service.NewService(repo)

	r := gin.Default()

	handlers := handler.NewH(serviceCon, r)

	handlers.AllRoutes()

	err := r.Run(config.Configure.PortRun)
	if err != nil {
		log.Fatal("router failed to start")
	}
}
