package controller

import (
	"fmt"
	"log"
	"storage-service-api/database"
	"storage-service-api/model/entity"
	"storage-service-api/model/request"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func ItemControllerCreate(c *fiber.Ctx) error {
	item := new(request.ItemCreateRequest)
	if err := c.BodyParser(item); err != nil {
		log.Println(err)
	}

	// VALIDATE REQUEST
	validate := validator.New()
	errValidate := validate.Struct(item)
	if errValidate != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "failed",
			"error":   errValidate.Error(),
		})
	}

	// FILE REQUIRED
	var filenameString string

	filename := c.Locals("filename")
	log.Println("filename = ", filename)
	if filename == nil {
		return c.Status(422).JSON(fiber.Map{
			"message": "file is required",
		})
	} else {
		filenameString = fmt.Sprintf("%v", filename)
	}

	fileOwner := c.Locals("username").(string)

	baseURL := "http://127.0.0.1:8000/public/"

	fileUrl := baseURL + filenameString

	newItem := entity.Item{
		Name:  item.Name,
		Item:  filenameString,
		Owner: fileOwner,
		Url:   fileUrl,
	}

	errCreateItem := database.DB.Create(&newItem).Error
	if errCreateItem != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "failed to store data",
		})
	}

	return c.JSON(fiber.Map{
		"message": "success",
		"data":    newItem,
	})
}

func ItemControllerGetAll(c *fiber.Ctx) error {
	var items []entity.Item
	result := database.DB.Find(&items)
	if result.Error != nil {
		log.Println(result.Error)
	}

	return c.JSON(items)
}
