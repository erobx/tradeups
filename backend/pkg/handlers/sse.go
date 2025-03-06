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

func SSE(p *db.PostgresDB) fiber.Handler {
    return func(c fiber.Ctx) error {
        c.Set("Content-Type", "text/event-stream")
        c.Set("Cache-Control", "no-cache")
        c.Set("Connection", "keep-alive")
        c.Set("Transfer-Encoding", "chunked")

        log.Printf("New Request\n")

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
                time.Sleep(2 * time.Second)
            }
        })
    }
}
