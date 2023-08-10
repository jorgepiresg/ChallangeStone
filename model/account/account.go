package model_account

import (
	"fmt"
	"strconv"
	"time"
)

type Account struct {
	ID         string    `json:"id,omitempty" db:"id"`
	Name       string    `json:"name,omitempty" db:"name"`
	CPF        string    `json:"cpf,omitempty" db:"cpf"`
	Secret     string    `json:"secret,omitempty" db:"secret"`
	Balance    int       `json:"balance,omitempty" db:"balance"`
	Created_at time.Time `json:"created_at,omitempty" db:"created_at"`
}

type Create struct {
	Name     string `json:"name" db:"name"`
	CPF      string `json:"cpf" db:"cpf"`
	Password string `json:"password"`
	Secret   string `json:"-" db:"secret"`
}

type CreateResponse struct {
	ID string `json:"id"`
}

type BalanceResponse struct {
	Balance int `json:"balance"`
}

func (c Create) Valid() error {
	switch true {
	case len(c.Name) <= 0:
		return fmt.Errorf("name invalid")

	case len(c.CPF) != 11:
		return fmt.Errorf("cpf invalid")

	case len(c.Password) <= 0:
		return fmt.Errorf("password invalid")
	}

	if _, err := strconv.Atoi(c.CPF); err != nil {
		return fmt.Errorf("cpf invalid")
	}

	return nil
}
