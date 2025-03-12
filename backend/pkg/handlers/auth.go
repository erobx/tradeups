package handlers

import (
	"log"
	"time"

	"github.com/erobx/tradeups/backend/internal/db"
	"github.com/erobx/tradeups/backend/pkg/common"
	"github.com/erobx/tradeups/backend/pkg/user"
	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func GetUser(p *db.PostgresDB) fiber.Handler {
	return func(c fiber.Ctx) error {
		//accessToken := c.Cookies("JWT")
		//if accessToken == "" {
		//	log.Println("No access token")
		//	return c.SendStatus(401)
		//}

        _, err := common.ValidateHeaders(c)
        if err != nil {
            return c.SendStatus(400)
        }

		return c.JSON(fiber.Map{
			"loggedIn": true,
		})
	}
}

// {"email":"","password":""}
func Login(p *db.PostgresDB) fiber.Handler {
	return func(c fiber.Ctx) error {
		creds := new(user.RegisteredUserPayload)
		if err := c.Bind().Body(creds); err != nil {
			return err
		}

		id, hash, err := p.GetHash(creds.Email)
		if err != nil {
			log.Printf("Email doesn't exist: %v\n", err)
			return c.SendStatus(400)
		}

		// check if password and hash match
		err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(creds.Password))
		if err != nil {
			return c.SendStatus(400)
		}

		//createAndSetJWT(c)
		jwt, err := newJWT(id)
		if err != nil {
			c.SendStatus(500)
		}

		log.Printf("User %s logged in\n", creds.Email)
		return c.JSON(fiber.Map{
			"jwt": jwt,
			"userId": id,
		})
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
			log.Printf("Error: %s\n", err.Error())
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

		if err := createAndSetJWT(c, newUser.Uuid.String()); err != nil {
			return c.SendStatus(500)
		}

		if err := p.CreateUser(newUser); err != nil {
			log.Printf("Error: %s\n", err.Error())
			return c.SendStatus(500)
		}
		
		log.Printf("New user %s registered\n", newUser.Username)
		return c.SendStatus(200)
	}
}

func createAndSetJWT(c fiber.Ctx, id string) error {
	// new user created, make new jwt
	jwt, err := newJWT(id)
	if err != nil {
		log.Printf("JWT not signed: %v\n", err)
		return err
	}

	// set cookie for jwt
	cookie := createJWTCookie(jwt)
	c.Cookie(cookie)
	return nil
}

func newJWT(id string) (string, error) {
	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(99999 * time.Hour)),
		IssuedAt: jwt.NewNumericDate(time.Now()),
		Subject: id,
		Issuer: "tradeups",
	}

	keyBytes, err := common.ReadPrivKey()
	if err != nil {
		return "", err
	}

	signingKey, err := jwt.ParseECPrivateKeyFromPEM(keyBytes)
	if err != nil {
		log.Printf("Error parsing private key: %v\n", err)
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES512, claims)
	s, err := token.SignedString(signingKey)
	return s, err
}

func createJWTCookie(jwt string) *fiber.Cookie {
	cookie := new(fiber.Cookie)

	cookie.Name = "JWT"
	cookie.Value = jwt
	//cookie.Domain = ""
	//cookie.MaxAge = int(time.Second)*60*60*24*365*10 // 10 year
	cookie.Path = "/"
	cookie.HTTPOnly = true
	cookie.Secure = false
	cookie.SameSite = fiber.CookieSameSiteLaxMode

	return cookie
}



