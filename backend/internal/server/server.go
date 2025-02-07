package server

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/erobx/tradeups/backend/internal/db"
	"github.com/erobx/tradeups/backend/pkg/handlers"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

type Server struct {
	addr string
	fiber *fiber.App
	db *db.PostgresDB
}

func NewServer(addr string, db *db.PostgresDB) *Server {
	return &Server{
		addr: addr,
		fiber: fiber.New(),
		db: db,
	}
}

func (s *Server) Run() error {
	if err := s.UseMiddleware(); err != nil {
		log.Fatalf("An error has occurred serving middleware: %v", err)
	}

	if err := s.MapHandlers(); err != nil {
		log.Fatalf("An error has occurred mapping the handlers: %v", err)
	}

	go func() {
		if err := s.fiber.Listen(":"+s.addr); err != nil {
			log.Fatalf("Could not start the server: %v", err)
		}
		
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT)

	<-quit

	log.Println("Server is stopping...")

	return nil
}

func (s *Server) UseMiddleware() error {
	s.fiber.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"Origin", "Content-Type", "Accept"},
		AllowMethods: []string{"OPTIONS"},
	}))
	return nil
}

func (s *Server) MapHandlers() error {
	auth := s.fiber.Group("/auth")
	auth.Post("/register", handlers.Register(s.db))
	// login
	//auth.Post("/login", )

	api := s.fiber.Group("/api")
	//api.Get("/user/:id", handlers.Login(s.db))
	api.Get("/user/inventory", handlers.GetSkins(s.db))

	return nil
}
