package repository

import (
	"testtask/internal/models"
	"testtask/internal/repository/car"

	"github.com/jackc/pgx/v5/pgxpool"
)


type CarRepository interface{
	DeleteCar(regNum string) error
	UpdateCar(regNum string, car models.Car) error
	AddCar(Car models.Car) error
	SelectCars(filter *models.Filter) (cars []models.Car, err error)

}

type Repository struct{
	Car CarRepository
}


func NewRepository(db *pgxpool.Pool) *Repository{
	carRepository := car.NewCarRepository(db)
	return &Repository{
		Car: carRepository,
	}
}