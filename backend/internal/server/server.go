package server

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
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

	//s.GenerateKeys()
	
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
		AllowOrigins: []string{"http://localhost:5173"},
		AllowCredentials: true,
		AllowHeaders: []string{"Origin", "Authorization", "Content-Type", "Accept"},
		AllowMethods: []string{"OPTIONS"},
	}))
	return nil
}

func (s *Server) MapHandlers() error {
	auth := s.fiber.Group("/auth")
	auth.Post("/register", handlers.Register(s.db))
	auth.Post("/login", handlers.Login(s.db))
	auth.Get("/users/:id", handlers.GetUser(s.db))

	api := s.fiber.Group("/api")
	api.Get("/users/:id/inventory", handlers.GetInventory(s.db))
	api.Get("/tradeups", handlers.GetActiveTradeups(s.db))
    api.Get("/sse", handlers.SSE(s.db))

	return nil
}

func (s *Server) GenerateKeys() {
	privateKey, _ := ecdsa.GenerateKey(elliptic.P521(), rand.Reader)
    publicKey := &privateKey.PublicKey
	encPriv, encPub := encode(privateKey, publicKey)

	writeToFiles(encPriv, encPub)
}

func writeToFiles(privKey, pubKey string) {
	f, _ := os.Create("jwt-priv-key.pem")
	f.WriteString(privKey)
	f.Close()

	f, _ = os.Create("jwt-pub-key.pem")
	f.WriteString(pubKey)
	f.Close()
}

func encode(privateKey *ecdsa.PrivateKey, publicKey *ecdsa.PublicKey) (string, string) {
    x509Encoded, _ := x509.MarshalECPrivateKey(privateKey)
    pemEncoded := pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: x509Encoded})

    x509EncodedPub, _ := x509.MarshalPKIXPublicKey(publicKey)
    pemEncodedPub := pem.EncodeToMemory(&pem.Block{Type: "PUBLIC KEY", Bytes: x509EncodedPub})

    return string(pemEncoded), string(pemEncodedPub)
}
