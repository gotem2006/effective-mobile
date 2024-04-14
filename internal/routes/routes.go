package routes

import (
	handler "testtask/internal/hadnler"
	"testtask/internal/routes/car"
	"github.com/gofiber/swagger"
	"github.com/gofiber/fiber/v2"
	_ "testtask/docs"
)




func InitRoutes(app *fiber.App, handlers *handler.Handler){	
	car.InitCarRoutes(app, handlers.Car)
	app.Get("/swagger/*", swagger.HandlerDefault)
}