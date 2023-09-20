package controller

import (
	"fmt"

	"github.com/PatipatCha/sgoc_account_service_demo.git/config/database"
	"github.com/PatipatCha/sgoc_account_service_demo.git/entities"
	"github.com/gofiber/fiber/v2"
)

func GetAccount(c *fiber.Ctx) error {
	db := database.DBConn
	var account []entities.Account
	db.Find(&account).Order("created_at DESC, id DESC")
	return c.Status(200).JSON(&account)
}

func GetAccountById(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var account entities.Account
	err := db.Find(&account, id).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "massage": "Could not find Account", "data": err})
	}

	if account.ID == 0 {
		res := fiber.Map{
			"massage": "ID Not Found",
		}
		return c.Status(200).JSON(res)
	}

	return c.Status(200).JSON(&account)

}

func CreateAccount(c *fiber.Ctx) error {
	db := database.DBConn
	account := new(entities.Account)
	err := c.BodyParser(account)
	if err != nil {
		res := fiber.Map{
			"status":  "error",
			"massage": "Check your input",
			"data":    err,
		}
		fmt.Println(err)
		return c.Status(500).JSON(res)
	}
	err = db.Create(&account).Error
	if err != nil {
		resError := fiber.Map{"status": "error", "massage": "Could not create account", "data": err}
		return c.Status(500).JSON(resError)
	}
	return c.Status(200).JSON(&account)
}

func UpdateAccount(c *fiber.Ctx) error {
	db := database.DBConn
	account := []entities.Account{}
	accountReq := new(entities.Account)
	if err := c.BodyParser(accountReq); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	db.Model(&account).Where("id = ?", accountReq.ID).Updates(entities.Account{
		Firstname: accountReq.Firstname,
		Surname:   accountReq.Surname,
		Email:     accountReq.Email,
		Mobile:    accountReq.Mobile,
		Role:      accountReq.Role,
		IsActive:  accountReq.IsActive,
	})

	return c.Status(200).JSON("updated")
}
