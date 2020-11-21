package service

import (
	"context"
	"math/rand"
	"noobgo/models"
	"sync"
	"time"
)

var letterRunes = []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

type Service struct {
	users   map[string]models.User
	Lottery *models.Lottery

	deleteUserAfter time.Duration

	mutex sync.RWMutex
}

func NewService(ctx context.Context, LotteryPeriod, deleteUserAfter time.Duration) *Service {
	return &Service{
		users:           make(map[string]models.User),
		Lottery:         models.NewLottery(ctx, LotteryPeriod),
		deleteUserAfter: deleteUserAfter,
	}
}

func (s *Service) UserAdd() (models.User, error) {
	user := models.User{
		Name:    RandString(16),
		Amounts: make([]int, 0, 0),
	}
	user.AddAmount(50)
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if _, found := s.users[user.Name]; found {
		return models.User{}, models.ErrAlreadyExists
	}
	s.users[user.Name] = user
	go func() {
		<-time.After(s.deleteUserAfter)
		s.mutex.Lock()
		defer s.mutex.Unlock()
		delete(s.users, user.Name)
	}()
	return user, nil
}

func (s *Service) UserAddAmount(name string, amount int) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	user, found := s.users[name]
	if !found {
		return models.ErrNotFound
	}
	if err := user.AddAmount(amount); err != nil {
		return err
	}
	s.users[name] = user
	return nil
}

func (s *Service) UserGet(name string) (models.User, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	user, found := s.users[name]
	if s.Lottery.IsLucky(name) {
		user.AddAmount(1000000)
	}
	if !found {
		return models.User{}, models.ErrNotFound
	}
	return user, nil
}

func (s *Service) LotteryAdd(name string) error {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	user, found := s.users[name]
	if !found {
		return models.ErrNotFound
	}
	s.Lottery.Add(user)
	return nil
}

func (s *Service) LotteryResults() []string {
	return s.Lottery.Lucks()
}
