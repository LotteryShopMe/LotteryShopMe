package models

import "github.com/pkg/errors"

const (
	MaxAmountsLen = 4
	MaxAmount     = 99
)

type User struct {
	Name    string `json:"name"`
	Amounts []int  `json:"amounts"`
}

func (a *User) Validate() error {
	if a.Name == "" {
		return errors.New("name cannot be empty")
	}
	return nil
}

func (a *User) AddAmount(amount int) error {
	if amount < 0 || amount > MaxAmount {
		return errors.Wrapf(ErrInvalidData, "别加太多了", MaxAmount+1, amount)
	}
	if len(a.Amounts) >= MaxAmountsLen {
		return errors.Wrapf(ErrInvalidData, "最多只能加三次... (%d)", MaxAmountsLen)
	}
	a.Amounts = append(a.Amounts, amount)
	return nil
}

func (a *User) AmountSum() int {
	sum := 0
	for _, a := range a.Amounts {
		sum += a
	}
	return sum
}
