package main

import (
	"fmt"
	"log"

	"github.com/PatipatCha/sgoc_account_service_demo.git/config/database"
	"github.com/PatipatCha/sgoc_account_service_demo.git/controller"
	"github.com/PatipatCha/sgoc_account_service_demo.git/entities"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDatabase() {
	var err error
	dsn := "host=pg-08082529-patipat-8829.aivencloud.com user=avnadmin password=AVNS_6ji60WZsOHzY6m3z88V dbname=account_service port=19122 sslmode=require TimeZone=Asia/Bangkok"
	database.DBConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	fmt.Println("Database connected!")
	database.DBConn.AutoMigrate(&entities.Account{})
	fmt.Println("Migrated DB")
}

func setupRoutes(app *fiber.App) {
	api := app.Group("/api") // /api

	v1 := api.Group("/v1")                    // /api/v1
	v1.Get("/account", controller.GetAccount) // /api/v1/list

}

func healthchecker(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "Welcome to Golang, Fiber, and GORM",
	})
}

func main() {
	app := fiber.New()

	initDatabase()
	app.Get("/", healthchecker)
	setupRoutes(app)
	log.Fatal(app.Listen(":3441"))
}
