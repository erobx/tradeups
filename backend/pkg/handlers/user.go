package handlers

import (
	"log"

	"github.com/erobx/tradeups/backend/internal/db"
	"github.com/erobx/tradeups/backend/pkg/user"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// {"email":"","password":""}
func Login(p *db.PostgresDB) fiber.Handler {
	return func(c fiber.Ctx) error {
		creds := new(user.Creds)
		if err := c.Bind().Body(creds); err != nil {
			return err
		}

		return c.SendString("Got user")
	}
}

// {"username":"","email":"","password":""}
func Register(p *db.PostgresDB) fiber.Handler {
	return func(c fiber.Ctx) error {
		newUser := new(user.User)
		if err := c.Bind().Body(newUser); err != nil {
			return err
		}

		// check if email exists
		exists, err := p.FindEmail(newUser.Email)
		if err != nil {
			log.Printf("Error: %s", err.Error())
			return c.SendStatus(400)
		}

		if exists {
			log.Println("Email already exists")
			return c.SendStatus(400)
		}

		// check if username is taken
		exists, err = p.FindUsername(newUser.Username)
		if err != nil {
			log.Printf("Error: %s", err.Error())
			return c.SendStatus(400)
		}

		if exists {
			log.Println("Username already taken")
			return c.SendStatus(400)
		}

		newUser.Uuid = uuid.New()

		hashed, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
		if err != nil {
			return c.SendStatus(500)
		}
		newUser.Hash = string(hashed)

		if err := p.CreateUser(newUser); err != nil {
			log.Printf("Error: %s", err.Error())
			return c.SendStatus(500)
		}



		return c.SendStatus(200)
	}
}
