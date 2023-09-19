package controller

import (
	"github.com/PatipatCha/sgoc_account_service_demo.git/config/database"
	"github.com/PatipatCha/sgoc_account_service_demo.git/entities"
	"github.com/gofiber/fiber/v2"
)

func GetAccount(c *fiber.Ctx) error {
	db := database.DBConn
	var account []entities.Personal
	db.Find(&account)
	// json.Marshal(&data)
	return c.Status(200).JSON(&account)

}

func CreateAccount(c *fiber.Ctx) error {
	db := database.DBConn
	account := new(entities.CreateAccountRequest)
	err := c.BodyParser(account)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "massage": "Check your input", "data": err})
	}
	err = db.Create(&account).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "massage": "Could not create account", "data": err})
	}
	return c.Status(200).JSON(&account)
}

func UpdateAccount(c *fiber.Ctx) error {
	db := database.DBConn
	account := []entities.Account{}
	title := new(entities.UpdateAccountRequest)

	// err := c.BodyParser(account)
	// if err != nil {
	// 	return c.Status(500).JSON(fiber.Map{"status": "error", "massage": "Check your input", "data": err})
	// }

	db.Model(&account).Where("id = ?", title.ID).Update("firstname", title.Firstname)

	return c.Status(200).JSON("updated")
}
