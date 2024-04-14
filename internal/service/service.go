package service

import (
	"testtask/internal/models"
	_ "testtask/internal/models"
	"testtask/internal/repository"
	"testtask/internal/service/car"
)



type CarService interface{
	AddCar(regNums models.RegNums) error
	DeleteCar(regNum string) error
	GetCars(filter *models.Filter) (cars []models.Car, err error)
	UpdateCar(regNum string, car models.Car) error
}


type Service struct{
	Car CarService
}



func NewServices(repositories *repository.Repository, apiUrl string) *Service{
	carService := car.NewCarService(repositories.Car, apiUrl)
	return &Service{
		Car: carService,
	}
}