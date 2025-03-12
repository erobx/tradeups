package handlers

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
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

        id := c.Params("tradeupId")
        // get specific tradeup
        if id != "" {
            return c.SendStreamWriter(func(w *bufio.Writer) {
                for {
                    tradeup, err := p.GetTradeup(id)
                    if err != nil {
                        log.Println(err)
                        return
                    }

                    b, _ := json.Marshal(tradeup)
                    fmt.Fprintf(w, "data: %s\n\n", string(b))

                    if err := w.Flush(); err != nil {
                        log.Printf("Client disconnected!")
                        return
                    }
                    time.Sleep(2 * time.Second)
                }
            })
        }
        return c.SendStatus(500)       
    }

}
