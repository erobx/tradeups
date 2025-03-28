package db

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand/v2"

	"github.com/erobx/tradeups/backend/pkg/skins"
	"github.com/erobx/tradeups/backend/pkg/tradeups"
)

func (p *PostgresDB) GetActiveTradeups() ([]tradeups.Tradeup, error) {
	var activeTradeups []tradeups.Tradeup
	q := `
	select t.id tradeup_id, t.rarity, t.current_status, t.stop_time,
	coalesce(
		jsonb_agg(
			distinct jsonb_build_object(
				'username', u.username,
				'avatar', u.avatar_key
			) 
		) filter (where u.id is not null), '[]'
	) as players,
	coalesce(
		json_agg(
			json_build_object(
				'inventoryId', ts.inv_id,
				'price', (select round(cast(i.price as numeric),2)),
				'imageSrc', s.image_key
			)
		) filter (where ts.inv_id is not null), '[]'
	) as skins
	from tradeups t
	left join tradeups_skins ts on t.id = ts.tradeup_id
	left join inventory i on ts.inv_id = i.id
	left join skins s on i.skin_id = s.id
	left join users u on i.user_id = u.id
	where t.current_status = 'Active'
	group by t.id, t.rarity, t.current_status
	`
	rows, err := p.conn.Query(context.Background(), q)
	if err != nil {
		return activeTradeups, err
	}
	defer rows.Close()

    var tempTradeups []tradeups.Tradeup
    var imageKeys []string

	for rows.Next() {
		var t tradeups.Tradeup
		var playersJson, skinsJson []byte
		err := rows.Scan(&t.Id, &t.Rarity, &t.Status, &t.StopTime, &playersJson, &skinsJson)
		if err != nil {
			return activeTradeups, err
		}

		err = json.Unmarshal(playersJson, &t.Players)
		if err != nil {
			return activeTradeups, err
		}

        var tempSkins []skins.TradeupSkin
		err = json.Unmarshal(skinsJson, &tempSkins)
		if err != nil {
			return activeTradeups, err
		}

        for _, skin := range tempSkins {
            if skin.ImageSrc != "" {
                imageKeys = append(imageKeys, skin.ImageSrc)
            }
        }

        t.Skins = tempSkins

        tempTradeups = append(tempTradeups, t)
	}

    if err := rows.Err(); err != nil {
        return activeTradeups, err
    }

    urlMap := p.urlManager.GetUrls(imageKeys)

    for i := range tempTradeups {
        for j := range tempTradeups[i].Skins {
            if url, exists := urlMap[tempTradeups[i].Skins[j].ImageSrc]; exists {
                tempTradeups[i].Skins[j].ImageSrc = url
            }
        }
        activeTradeups = append(activeTradeups, tempTradeups[i])
    }

	return activeTradeups, nil 
}

func (p *PostgresDB) GetTradeup(id string) (tradeups.Tradeup, error) {
    t, err := p.getTradeup(id)
    if err != nil {
        return t, err
    }

    if t.Status == "Completed" {
        return t, nil
    }

    return t, nil
}

func (p *PostgresDB) FindReadyActiveTradeups() ([]string, error) {
    var readyTradeupIds []string
    q := "select id from tradeups where current_status='Active' and now() >= stop_time"
    rows, err := p.conn.Query(context.Background(), q)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    for rows.Next() {
        var id string
        if err := rows.Scan(&id); err != nil {
            return nil, err
        }
        readyTradeupIds = append(readyTradeupIds, id)
    }

    return readyTradeupIds, rows.Err()
}

func (p *PostgresDB) UpdateTradeupsToInProgress(tradeupIds []string) error {
    q := `
    update tradeups set current_status = 'In Progress'
    where id = any($1)
    `
    _, err := p.conn.Exec(context.Background(), q, tradeupIds)
    return err
}

func (p *PostgresDB) TradeupIsFull(tradeupId string) error {
    numSkins, err := p.getSkinCount(tradeupId)
    if err != nil {
        return err
    }

    if numSkins > 10 {
        return fmt.Errorf("Tradeup full")
    }
    return nil
}

func (p *PostgresDB) AddSkinToTradeup(userId, tradeupId string, invId int) error {
    // if the user actually owns the skin to add
    var exists bool
    q := "select exists(select 1 from inventory where user_id=$1 and id=$2)"
    row := p.conn.QueryRow(context.Background(), q, userId, invId)
    err := row.Scan(&exists)
    if err != nil {
        return err
    }
    
    if !exists {
        return fmt.Errorf("User does not own that item")
    }

    // finally add the skin
    q = "insert into tradeups_skins values($1,$2)"
    _, err = p.conn.Exec(context.Background(), q, tradeupId, invId)
    if err != nil {
        return err
    }

    // have to check if tradeup has 10 skins in order to start timer
    numSkins, err := p.getSkinCount(tradeupId)
    if err != nil {
        return err
    }

    if numSkins == 10 {
        q = "update tradeups set stop_time=now() + interval '10 sec' where id=$1"
        _, err = p.conn.Exec(context.Background(), q, tradeupId)
        log.Printf("Started timer for tradeup %s\n", tradeupId)
        if err != nil {
            return err
        }
    }

    return nil
}

func (p *PostgresDB) RemoveSkinFromTradeup(tradeupId string, invId int) (skins.InventorySkin, error) {
    var status string
    var invSkin skins.InventorySkin
    var imageKey string

    // check if status is 'Active'
    q := "select current_status from tradeups where id=$1"
    row := p.conn.QueryRow(context.Background(), q, tradeupId)
    err := row.Scan(&status)
    if err != nil {
        return invSkin, err
    }

    if status != "Active" {
        return invSkin, fmt.Errorf("Cannot remove a skin from non-active tradeup")
    }

    numSkins, err := p.getSkinCount(tradeupId)
    if err != nil {
        return invSkin, err
    }

    // tradeup full before removal, check if stop_time < now()
    if numSkins == 10 {
        q = "update tradeups set stop_time=now() + interval '5 year' where id=$1"
        _, err = p.conn.Exec(context.Background(), q, tradeupId)
        if err != nil {
            return invSkin, err
        }
        log.Printf("Stopped timer for tradeup %s\n", tradeupId)
    }

    q = `
    with deleted_skin as (
        delete from tradeups_skins ts
        where tradeup_id=$1 and inv_id=$2
        returning inv_id
    )
    select i.id, i.wear_str, i.wear_num, round(cast(i.price as numeric),2), i.is_stattrak, i.created_at,
		s.name, s.rarity, s.collection, s.image_key
    from inventory i
	join skins s on s.id = i.skin_id
	where i.id=$3
    order by s.image_key, i.wear_str
    `
    row = p.conn.QueryRow(context.Background(), q, tradeupId, invId, invId)
    err = row.Scan(&invSkin.Id, &invSkin.Wear, &invSkin.SkinFloat, &invSkin.Price, &invSkin.IsStatTrak, &invSkin.CreatedAt,
        &invSkin.Name, &invSkin.Rarity, &invSkin.Collection, &imageKey)
    
    if err != nil {
        return invSkin, err
    }

    urlMap := p.urlManager.GetUrls([]string{imageKey})
    invSkin.ImageSrc = urlMap[imageKey]
    return invSkin, err
}

func (p *PostgresDB) CreateTradeup(rarity string) error {
    q := "insert into tradeups values (nextval('tradeups_id_seq'),$1)"
    _, err := p.conn.Exec(context.Background(), q, rarity)
    if err != nil {
        return err
    }

    return nil
}

// Loop through all rarities, ensuring there's always 3 tradeups of each rarity
func (p *PostgresDB) MaintainTradeupCount() error {
    rarities := []string{"Consumer", "Industrial", "Mil-Spec", "Restricted", "Classified"}
    
    for _, r := range rarities {
        var count int
        q := "select count(*) from tradeups where current_status = 'Active' and rarity = $1"
        if err := p.conn.QueryRow(context.Background(), q, r).Scan(&count); err != nil {
            return err
        }
        
        if count < 3 {
            if err := p.CreateTradeup(r); err != nil {
                return err
            }
        }
    }
    return nil
}

// Main winner logic
func (p *PostgresDB) decideWinner(tradeup tradeups.Tradeup) error {
    // TODO: come up with algo
    // for now percentage split
    // ex: 8/10, 2/10 => user1 has 80%, user2 has 20%
    
    usersSkins := make(map[string]int)
    // group the skins based on the userId
    for _, skin := range tradeup.Skins {
        _, ok := usersSkins[skin.UserId]
        if ok {
            usersSkins[skin.UserId] += 1
        } else {
            usersSkins[skin.UserId] = 1
        }
    }

    var winner string
    randomNum := rand.IntN(100)
    currWeight := 0
    
    for player, weight := range usersSkins {
        currWeight += weight * 10
        if randomNum < currWeight {
            winner = player
            break
        }
    }

    log.Printf("User %s won!\n", winner)

    q := "update tradeups set current_status='Completed', winner=$1 where id=$2"
    _, err := p.conn.Exec(context.Background(), q, winner, tradeup.Id)
    if err != nil {
        return err
    }

    return nil
}

func (p *PostgresDB) GetTradeupsInProgress() ([]tradeups.Tradeup, error) {
    var tradeupsInProgress []tradeups.Tradeup
    // get all tradeups with status In Progress
    q := "select id from tradeups where current_status='In Progress'"
    rows, err := p.conn.Query(context.Background(), q)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var tradeupIds []string
    for rows.Next() {
        var id string
        err := rows.Scan(&id)
        if err != nil {
            return nil, err
        }
        tradeupIds = append(tradeupIds, id)
    }

    if err = rows.Err(); err != nil {
        return nil, err
    }

    for _, id := range tradeupIds {
        tradeup, err := p.getTradeup(id)
        if err != nil {
            continue
        }
        tradeupsInProgress = append(tradeupsInProgress, tradeup)
    }

    return tradeupsInProgress, nil
}

func (p *PostgresDB) ProcessTradeupWinners(toProcess []tradeups.Tradeup) error {
    tx, err := p.conn.Begin(context.Background())
    if err != nil {
        return err
    }
    defer func() {
        if err != nil {
            tx.Rollback(context.Background())
            return
        }
        err = tx.Commit(context.Background())
    }()

    for _, t := range toProcess {
        p.decideWinner(t)
    }

    return nil
}

