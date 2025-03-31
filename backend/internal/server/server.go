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
	"time"

	"github.com/erobx/tradeups/backend/internal/db"
	"github.com/erobx/tradeups/backend/internal/middleware"
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
	if err := s.useMiddleware(); err != nil {
		log.Fatalf("An error has occurred serving middleware: %v", err)
	}

	if err := s.mapHandlers(); err != nil {
		log.Fatalf("An error has occurred mapping the handlers: %v", err)
	}

    // constantly check tradeups in progress, decide winner 
    go s.watchTradeups()
    go s.maintainTradeupCounts()
	//s.generateKeys()
	
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

func (s *Server) useMiddleware() error {
	s.fiber.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:5173"},
		AllowCredentials: true,
		AllowHeaders: []string{"Origin", "Authorization", "Content-Type", "Accept"},
        AllowMethods: []string{"PUT", "DELETE", "OPTIONS"},
	}))
	return nil
}

func (s *Server) mapHandlers() error {
    // Auth
	auth := s.fiber.Group("/auth")
	auth.Post("/register", handlers.Register(s.db))
	auth.Post("/login", handlers.Login(s.db))

	v1 := s.fiber.Group("/v1")

    // SSE Test
    events := v1.Group("/events")
    events.Get("/inventory/:token", handlers.InventoryUpdates(s.db), middleware.SSE())

    // User
    users := v1.Group("/users")

    users.Get("/", handlers.GetUser(s.db), middleware.Protected())
	users.Get("/:userId", handlers.GetUser(s.db), middleware.Protected())
	users.Get("/:userId/inventory", handlers.GetInventory(s.db), middleware.Protected())
    users.Get("/:userId/recent", handlers.GetRecentTradeups(s.db), middleware.Protected())
    users.Get("/:userId/stats", handlers.GetUserStats(s.db), middleware.Protected())

    users.Delete("/:userId/inventory/:invId", handlers.DeleteSkin(s.db), middleware.Protected())

    // Tradeups
    tradeups := v1.Group("/tradeups")

    tradeups.Get("/", handlers.GetActiveTradeupsSSE(s.db))
    tradeups.Get("/:tradeupId", handlers.GetTradeupSSE(s.db))

    tradeups.Post("/new", handlers.NewTradeup(s.db), middleware.Admin())
    tradeups.Put("/add", handlers.AddSkinToTradeup(s.db), middleware.Protected())
    tradeups.Delete("/remove", handlers.RemoveSkinFromTradeup(s.db), middleware.Protected())

    // Store
    store := v1.Group("/store")
    store.Post("/buy", handlers.BuyCrate(s.db), middleware.Protected())

	return nil
}

func (s *Server) watchTradeups() {
    ticker := time.NewTicker(500 * time.Millisecond)
    defer ticker.Stop()

    for {
        select {
        case <-ticker.C:
            // get tradeups who timers have expired
            activeTradeups, err := s.db.FindReadyActiveTradeups()
            if err != nil {
                log.Printf("Error finding ready active tradeups: %v\n", err)
                continue
            }

            if len(activeTradeups) > 0 {
                err = s.db.UpdateTradeupsToInProgress(activeTradeups)
                if err != nil {
                    log.Printf("Error updating tradeups to In Progress: %v\n", err)
                }
            }

            tradeups, err := s.db.GetTradeupsInProgress()
            if err != nil {
                log.Printf("Error fetching tradeups in progress: %v\n", err)
                continue
            }
            
            if len(tradeups) > 0 {
                err = s.db.ProcessTradeupWinners(tradeups)
                if err != nil {
                    log.Printf("Error processing tradeup winners: %v\n", err)
                }
            }
        }
    }
}

func (s *Server) maintainTradeupCounts() {
    ticker := time.NewTicker(1 * time.Minute)
    defer ticker.Stop()

    for {
        select {
        case <-ticker.C:
            err := s.db.MaintainTradeupCount()
            if err != nil {
                log.Printf("Error maintaining tradeup counts: %v\n", err)
                continue
            }
        }
    }
}

func (s *Server) generateKeys() {
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
