package controller

import (
	"github.com/PatipatCha/sgoc_account_service_demo.git/config/database"
	"github.com/PatipatCha/sgoc_account_service_demo.git/entities"
	"github.com/gofiber/fiber/v2"
)

func GetAccount(c *fiber.Ctx) error {
	db := database.DBConn
	var account []entities.Account
	db.Find(&account)
	return c.JSON(&account)
}
