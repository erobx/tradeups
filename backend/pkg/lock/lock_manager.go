package lock

import (
	"log"
	"time"
)

// global state for active tradeups that become full
// counting down until timer finishes or a user takes
// a skin out of the tradeup
type LockManager struct {
    start time.Time
    lockMap map[int]time.Timer
}

type Lock struct {
    status string // could be active or locked
    counter time.Ticker
}

func NewLockManager() *LockManager {
    return &LockManager{
        start: time.Now(),
        lockMap: make(map[int]time.Timer),
    }
}

func (lm *LockManager) TestTimer() {
    log.Println(lm.start)
    time.Sleep(2 * time.Second)
    log.Println(time.Since(lm.start))
}
