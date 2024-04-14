package main

import (
	"log"
	"testtask/config"
	handler "testtask/internal/hadnler"
	"testtask/internal/repository"
	"testtask/internal/routes"
	"testtask/internal/service"
	postgres "testtask/pkg/database"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// @title Car API
// @version 1.0
// @description Swagger for car catalog api
// @host 0.0.0.0:8080
// @BasePath /
func main(){
	app := fiber.New(fiber.Config{
		EnableSplittingOnParsers: true,
	})
	app.Use(cors.New(), logger.New())

	config, err := config.LoadConfig(".")
	if err != nil{
		log.Fatal(err)
	}

	time.Sleep(10 * time.Second)

	pg := postgres.NewPostgres(config)
	conn, err := pg.Connection()
	if err != nil{
		log.Fatal(err)
	}
	if err := pg.MigrationUp(); err != nil{
		log.Println("Faile migrate:",err)
	}

	repositories := repository.NewRepository(conn)
	services := service.NewServices(repositories, config.APIUrl)
	hadnlers := handler.NewHadnlers(services)
	routes.InitRoutes(app, hadnlers)

	err = app.Listen(":8080")
	if err != nil{
		log.Fatal(err)
	}
}
