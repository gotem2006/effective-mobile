package car

import (
	handler "testtask/internal/hadnler"
	"github.com/gofiber/fiber/v2"
)

func InitCarRoutes(app *fiber.App, handler handler.CarHadnler){
	app.Post("/car", handler.PostCar)
	app.Delete("/car", handler.DeleteCar)
	app.Get("/car", handler.GetCars)
	app.Put("/car", handler.UpdateCar)
	
}