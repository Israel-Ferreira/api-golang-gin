package models

import "time"

type Account struct {
	Id        int
	Status    string
	Archived  bool
	Timestamp
}

func (a *Account) DeleteAccount() {
	a.Archived = true
	a.DeletedAt = time.Now()
}
