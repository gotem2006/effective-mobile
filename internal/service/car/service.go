package car

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"testtask/internal/models"
	"testtask/internal/repository"
	"time"
)



type carService struct{
	repo repository.CarRepository
	apiUrl string
}


func NewCarService(repo repository.CarRepository, apiUrl string) *carService{

	return &carService{
		repo: repo,
		apiUrl: apiUrl,
	}
}


func (s *carService) AddCar(regNums models.RegNums) error{
	errors := make(chan error, len(regNums.RegNums))
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()
	for _, regNum := range regNums.RegNums{
		go func(regNum string) {
			car, err := s.fetchCarInfo(regNum)
			if err != nil{
				errors <- err
			}
			if err = s.repo.AddCar(car); err != nil{
				errors <- err
			}
		}(regNum)
	}
	for{
		select{
		case err := <- errors:
			return err
		case <- ctx.Done():
			return nil
		}
	}
}


func (s *carService) DeleteCar(regNum string) error{
	return s.repo.DeleteCar(regNum)
}



func (s *carService) GetCars(filter *models.Filter) (cars []models.Car, err error){
	// get cars by filters
	cars, err = s.repo.SelectCars(filter)

	if err != nil{
		return cars, err
	}


	return cars, nil
}

func (s *carService) UpdateCar(regNum string, car models.Car) error{
	return s.repo.UpdateCar(regNum, car)
}


func (s *carService) fetchCarInfo(number string) (car models.Car, err error){
	url := s.apiUrl + "/info?regNum=" + number
	log.Println(url)
	response, err := http.Get(url)
	if err != nil{
		return models.Car{}, err
	}
	defer response.Body.Close()
	if err = json.NewDecoder(response.Body).Decode(&car); err != nil{
		return models.Car{}, fmt.Errorf("can`t found car: %s", number)
	}
	if response.StatusCode != http.StatusOK{
		return models.Car{}, fmt.Errorf("сan`t find car `%s`", number)
	}
	return car, nil
}
