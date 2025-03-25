package lock

import "time"

type Lock struct {
    timer *time.Timer
    done chan struct{}
    countdown Countdown
}

type Countdown struct {
    minutes int
    seconds int
    quit chan struct{}
}

func NewLock() *Lock {
    return &Lock{
        timer: time.NewTimer(time.Second * 5),
        done: make(chan struct{}),
        countdown: NewCountdown(0, 5),
    }
}

func NewCountdown(minutes, seconds int) Countdown {
    return Countdown{
        minutes: minutes,
        seconds: seconds,
    }
}
