package handlers

import (
	"log"

	"github.com/erobx/tradeups/backend/pkg/db"
	"github.com/erobx/tradeups/backend/pkg/user"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func GetUser(c fiber.Ctx) error {
	return c.SendString("Got user")
}

// {"username":"","password":"","email":""}
func CreateUser(c fiber.Ctx) error {
	newUser := new(user.User)
	if err := c.Bind().Body(newUser); err != nil {
		return err
	}
	newUser.Uuid = uuid.New()

	hashed, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.SendStatus(500)
	}
	newUser.Hash = string(hashed)

	if err := db.CreateUser(newUser); err != nil {
		log.Printf("Error: %s", err.Error())
		return c.SendStatus(500)
	}

	return c.SendStatus(200)
}
