package handlers

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/erobx/tradeups/backend/internal/db"
	"github.com/gofiber/fiber/v3"
)

func GetActiveTradeupsSSE(p *db.PostgresDB) fiber.Handler {
    return func(c fiber.Ctx) error {
        c.Set("Content-Type", "text/event-stream")
        c.Set("Cache-Control", "no-cache")
        c.Set("Connection", "keep-alive")
        c.Set("Transfer-Encoding", "chunked")

        log.Printf("Requesting all tradeups...\n")

        return c.SendStreamWriter(func(w *bufio.Writer) {
            for {
                tradeups, err := p.GetActiveTradeups()
                if err != nil {
                    log.Println(err)
                    return
                }

                b, _ := json.Marshal(tradeups)
                fmt.Fprintf(w, "data: %s\n\n", string(b))

                if err := w.Flush(); err != nil {
                    log.Printf("Client disconnected!")
                    return
                }
                time.Sleep(200 * time.Millisecond)
            }
        })
    }
}

func GetTradeupSSE(p *db.PostgresDB) fiber.Handler {
    return func(c fiber.Ctx) error {
        c.Set("Content-Type", "text/event-stream")
        c.Set("Cache-Control", "no-cache")
        c.Set("Connection", "keep-alive")
        c.Set("Transfer-Encoding", "chunked")

        log.Printf("Requesting tradeup...\n")

        id, err := strconv.Atoi(c.Params("tradeupId"))
        if err != nil {
            return c.SendStatus(500)
        }
        // get specific tradeup
        return c.SendStreamWriter(func(w *bufio.Writer) {
            ticker := time.NewTicker(250 * time.Millisecond)
            defer ticker.Stop()

            tradeup, err := p.GetTradeup(id)
            if err != nil {
                log.Printf("Error getting tradeup: %v\n", err)
                return
            }

            b, _ := json.Marshal(tradeup)
            fmt.Fprintf(w, "data: %s\n\n", string(b))

            if err := w.Flush(); err != nil {
                log.Printf("Client disconnected!")
                return
            }

            for {
                select {
                case <-ticker.C:
                    tradeup, err := p.GetTradeup(id)
                    if err != nil {
                        log.Printf("Error getting tradeup: %v\n", err)
                        return
                    }

                    b, _ := json.Marshal(tradeup)
                    fmt.Fprintf(w, "data: %s\n\n", string(b))

                    if err := w.Flush(); err != nil {
                        log.Printf("Client disconnected!")
                        return
                    }
                }
            }
        })
    }
}
