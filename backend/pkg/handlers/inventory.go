package handlers

import (
    "bufio"
    "encoding/json"
    "fmt"
	"log"
    "time"

	"github.com/erobx/tradeups/backend/internal/db"
	"github.com/erobx/tradeups/backend/pkg/common"
	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
)

// {id: 0, name: "M4A4 | Howl", wear: "Factory New", rarity: "Contraband", float: 0.01, isStatTrak: true, imgSrc: "/m4a4-howl.png"},
func GetInventory(p *db.PostgresDB) fiber.Handler {
	return func(c fiber.Ctx) error {
		urlUserId := c.Params("userId")
		// for now using Bearer token instead of jwt in the cookie bc of localhost
        token := c.Locals("jwt").(*jwt.Token)

        jwtUserId, err := common.ValidateAndReturnUserId(token, urlUserId)
        if err != nil {
            log.Println(err)
            return c.SendStatus(500)
        }

		inv, err := p.GetInventory(jwtUserId)
		if err != nil {
            log.Println(err)
			return c.JSON(fiber.Map{
                "skins": "empty",
            })
		}

		return c.JSON(inv)
	}
}

func DeleteSkin(p *db.PostgresDB) fiber.Handler {
    return func(c fiber.Ctx) error {
        urlUserId := c.Params("userId")
        urlInvId := c.Params("invId")

        token := c.Locals("jwt").(*jwt.Token)

        jwtUserId, err := common.ValidateAndReturnUserId(token, urlUserId)
        if err != nil {
            log.Println(err)
            return c.SendStatus(500)
        }

        err = p.DeleteSkin(jwtUserId, urlInvId)
        if err != nil {
            log.Println(err)
            return c.SendStatus(500)
        }

        log.Printf("User: %s deleted item %s from their inventory\n", urlUserId, urlInvId)
        return c.SendStatus(204)
    }
}

func InventoryUpdates(p *db.PostgresDB) fiber.Handler {
    return func(c fiber.Ctx) error {
        token := c.Locals("jwt").(*jwt.Token)

        var jwtUserId string
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			jwtUserId, _ = claims.GetSubject()
		} else {
			return c.SendStatus(500)
		}

        c.Set("Content-Type", "text/event-stream")
        c.Set("Cache-Control", "no-cache")
        c.Set("Connection", "keep-alive")
        c.Set("Transfer-Encoding", "chunked")

        log.Printf("Starting inventory updates for user %s...\n", jwtUserId)

        // send updates to user inventory
        return c.SendStreamWriter(func(w *bufio.Writer) {
            ticker := time.NewTicker(time.Second)
            defer ticker.Stop()

            lastKnownInventory, err := p.GetInventory(jwtUserId)
            if err != nil {
                log.Printf("Error getting initial inventory: %v\n", err)
                fmt.Fprintf(w, "data: {\"error\": \"Failed to fetch inventory\"}\n\n")
                w.Flush()
                return
            }

            // track deleted items for notifications
            lastInvIds := make(map[int]bool)
            for _, skin := range lastKnownInventory.Skins {
                lastInvIds[skin.Id] = true
            }

            for {
                select {
                case <-ticker.C:
                    currInv, err := p.GetInventory(jwtUserId)
                    if err != nil {
                        log.Printf("Error checking inventory: %v\n", err)
                        fmt.Fprintf(w, "data: {\"error\": \"Failed to check for updates\"}\n\n")
                        if err := w.Flush(); err != nil {
                            log.Printf("Client disconnected: %v", err)
                            return
                        }
                        continue
                    }

                    for _, skin := range currInv.Skins {
                        found := false
                        for _, oldSkin := range lastKnownInventory.Skins {
                            if skin.Id == oldSkin.Id {
                                found = true
                                break
                            }
                        }

                        if !found {
                            update := map[string]any{
                                "type": "inventory_update",
                                "action": "add",
                                "item": skin,
                            }

                            jsonData, _ := json.Marshal(update)
                            fmt.Fprintf(w, "data: %s\n\n", string(jsonData))
                            if err := w.Flush(); err != nil {
                                log.Printf("Client disconnected: %v", err)
                                return
                            }

                            log.Printf("Sent inventory update to user %s: New item %d added",
                                jwtUserId, skin.Id)
                        }
                    }

                    // find deleted items
                    currentInvIds := make(map[int]bool)
                    for _, skin := range currInv.Skins {
                        currentInvIds[skin.Id] = true
                    }
                    
                    for invId := range lastInvIds {
                        if !currentInvIds[invId] {
                            // this item was deleted
                            update := map[string]any{
                                "type": "inventory_update",
                                "action": "remove",
                                "invId": invId,
                            }

                            jsonData, _ := json.Marshal(update)
                            fmt.Fprintf(w, "data: %s\n\n", string(jsonData))
                            if err := w.Flush(); err != nil {
                                log.Printf("Client disconnected: %v\n", err)
                                return
                            }

                            log.Printf("Sent inventory update to user %s: Item %d removed",
                                jwtUserId, invId)
                        }
                    }

                    // update reference state
                    lastKnownInventory = currInv
                    lastInvIds = currentInvIds
                }
            }
        })
    }
}
