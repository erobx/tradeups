package lock

import (
	"sync"
	"time"
)

// global state for active tradeups that become full
// counting down until timer finishes or a user takes
// a skin out of the tradeup
type LockManager struct {
    locks map[string]*Lock
    mu *sync.RWMutex
}

func NewLockManager() *LockManager {
    return &LockManager{
        locks: make(map[string]*Lock),
    }
}

func (lm *LockManager) StartTimer(tradeupId string) {
    lm.mu.Lock()
    defer lm.mu.Unlock()

    if lock, ok := lm.locks[tradeupId]; ok {
        lock.timer = time.NewTimer(time.Second * 5)
    } else {
        lm.locks[tradeupId] = &Lock{}
    }
}
