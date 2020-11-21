package models

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Lottery struct {
	players map[string]User
	luck    map[string]struct{}

	mutex sync.RWMutex
}

func NewLottery(ctx context.Context, period time.Duration) *Lottery {
	l := &Lottery{
		players: make(map[string]User),
		luck:    make(map[string]struct{}),
	}
	go func() {
		ticker := time.NewTicker(period)
		for {
			select {
			case <-ticker.C:
				l.Evaluate()
			case <-ctx.Done():
				ticker.Stop()
				return
			}
		}
	}()
	return l
}

func (l *Lottery) Add(user User) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	l.players[user.Name] = user
}

func (l *Lottery) IsLucky(name string) bool {
	l.mutex.RLock()
	defer l.mutex.RUnlock()
	if _, won := l.luck[name]; won {
		return true
	}
	return false
}

func (l *Lottery) Lucks() []string {
	l.mutex.RLock()
	defer l.mutex.RUnlock()
	var ws []string
	for w := range l.luck {
		ws = append(ws, w)
	}
	return ws
}

func (l *Lottery) Evaluate() {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	players := l.players
	l.luck = make(map[string]struct{})
	l.players = make(map[string]User)
	for name, user := range players {
		amounts := append(user.Amounts, randInt(1000000, 99000000))
		fmt.Print(amounts, name)
		tot := 0
		for _, a := range amounts {
			tot += a
		}
		if tot == 0x12beef {
			l.luck[name] = struct{}{}
		}
	}
}

func randInt(min, max int) int {
	return rand.Intn(max-min) + min
}
