package car

import (
	"log"
	"log/slog"
	"net/http"
	"testtask/internal/models"
	"testtask/internal/service"

	"github.com/gofiber/fiber/v2"
)


type carHadnler struct{
	service service.CarService
}

func NewCarHandler(service service.CarService) *carHadnler{
	return &carHadnler{service: service}
}


//@Summary PostCar
//@Tags CarAPI
//@Description Post car by regNum
//@Accept json
//@Param input body models.RegNums true "RegNums"
//@Router /car [post]
func (h *carHadnler) PostCar(c *fiber.Ctx) error{
	var regNums models.RegNums


	// parse numbers
	if err := c.BodyParser(&regNums); err != nil{
		slog.Error("Can`t parse data", string(c.Body()))
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}


	// add numbers
	if err := h.service.AddCar(regNums); err != nil{
		slog.Error(err.Error())
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// result
	slog.Info("Success request with data: %v", regNums)
	return c.Status(http.StatusOK).SendString("Success")
}


//@Summary DeleteCar
//@Tags CarAPI
//@Description Delete car by regNum
//@Accept json
//@Param regNum query string true "RegNum"
//@Router /car [delete]
func (h *carHadnler) DeleteCar(c *fiber.Ctx) error{
	// parse number
	regNum := c.Query("regNum")

	// delete number
	err := h.service.DeleteCar(regNum)


	if err != nil{
		log.Println("Error: ", err.Error())
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}


	return c.Status(http.StatusOK).SendString("Success")
}

//@Tags CarAPI
//@Summary GetCars
//@Description Get slice of car
//@Param limit query int false "limit"
//@Param offset query int false "offset"
//@Param mark query []string false "mark"
//@Param model query []string false "model"
//@Param year query []string false "year"
//@Param owner_name query []string false "owner_name"
//@Param owner_surname query []string false "owner_surname"
//@Param owner_patronymic query []string false "owner_patronymic"
// @Router /car [get]
func (h *carHadnler) GetCars(c *fiber.Ctx) error{
	filter := &models.Filter{}

	// parse filter
	if err := c.QueryParser(filter); err != nil{
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// get cars
	cars, err := h.service.GetCars(filter)
	if err != nil{
		return err
	}

	return c.Status(http.StatusOK).JSON(cars)
}



//@Summary UpdateCar
//@Tags CarAPI
//@Description update car
//@Accept json
//@Param regNum query string true "regNum"
//@Param input body models.Car true "Car"
//@Router /car [put]
func (h *carHadnler) UpdateCar(c *fiber.Ctx) error{
	var car models.Car

	if err := c.BodyParser(&car); err != nil{
		log.Println("Can`t parse data: ", string(c.Body()))
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	err := h.service.UpdateCar(c.Query("regNum"), car)

	if err != nil{
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusOK).SendString("Success")
}
