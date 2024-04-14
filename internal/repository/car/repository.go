package car

import (
	"context"
	"fmt"
	"testtask/internal/models"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)



type carRepository struct{
	db *pgxpool.Pool
}


func NewCarRepository(db *pgxpool.Pool) *carRepository{
	return &carRepository{db}
}



func (c *carRepository) AddCar(car models.Car) error{
	query := 
	`
		INSERT INTO cars (regNum, mark, model, year, owner_name, owner_surname, owner_patronymic)
		VALUES($1, $2, $3, $4, $5, $6, $7)
	`
	_, err := c.db.Exec(context.Background(), query, car.RegNum, car.Mark, car.Model, car.Year, car.Owner.Name, car.Owner.Surname, car.Owner.Patronymic)
	if err != nil{
		return fmt.Errorf("car is exist %s", car.RegNum)
	}
	return nil
}


 func (c *carRepository) DeleteCar(number string) error{
	query := 
	`
		DELETE FROM cars WHERE regNum = $1 RETURNING regnum
	`
	var regnum string
	c.db.QueryRow(context.Background(), query, number).Scan(&regnum)
	if regnum == ""{
		return fmt.Errorf("car is not exist %s", number)
	}
	return nil
 }


 func (c *carRepository) SelectCars(filter *models.Filter) (cars []models.Car, err error){
 	query :=
	`
		SELECT * FROM cars WHERE 1=1
	`

	//building query by filters
	query, args := filter.ParseToSql(query)

	// get rows
	rows, err := c.db.Query(context.Background(), query, args...)

	if err != nil{
		return nil, err
	}

	// parse cars from db to struct
	cars, err = pgx.CollectRows(rows, func(row pgx.CollectableRow) (models.Car, error) {
		var car models.Car
		err := row.Scan(&car.RegNum, &car.Mark, &car.Model, &car.Year, &car.Owner.Name, &car.Owner.Surname, &car.Owner.Patronymic)
		if err != nil{
			return models.Car{}, err
		}
		return car, nil
	})
	if err != nil{
		return cars, err
	}
 	return cars, nil
 }

 func (c *carRepository) UpdateCar(number string, car models.Car) error{
	query :=
	`
		UPDATE cars 
		SET mark=$1, model=$2, year=$3, owner_name=$4, owner_surname =$5, owner_patronymic =$6 
		WHERE regnum=$7
		RETURNING regnum
	`
	var regnum string
	c.db.QueryRow(context.Background(), query,  car.Mark, car.Model, car.Year, car.Owner.Name, car.Owner.Surname, car.Owner.Patronymic, number).Scan(&regnum)
	if regnum == ""{
		return fmt.Errorf("car is not exist %s", number)
	}
 	return nil
 }