package handler

import (
	"testtask/internal/hadnler/car"
	"testtask/internal/service"

	"github.com/gofiber/fiber/v2"
	_ "testtask/docs"
)


type CarHadnler interface{
	PostCar(c *fiber.Ctx) error
	DeleteCar(c *fiber.Ctx) error
	GetCars(c *fiber.Ctx) error
	UpdateCar(c *fiber.Ctx) error
}



type Handler struct{
	Car CarHadnler
}


func NewHadnlers(services *service.Service) *Handler{
	carHadnler := car.NewCarHandler(services.Car)
	return &Handler{
		Car: carHadnler,
	}
}


