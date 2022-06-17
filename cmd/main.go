package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/putcho01/gofiber/pkg/books"
	"github.com/putcho01/gofiber/pkg/common/config"
	"github.com/putcho01/gofiber/pkg/common/db"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	app := fiber.New()
	db := db.Init(c.DBUrl)

	books.RegisterRoutes(app, db)

	app.Listen(c.Port)
}
