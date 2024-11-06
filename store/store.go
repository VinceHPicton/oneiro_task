package store

import (
	"math/rand"
	"oneiro_task/loan"
)

type Store map[int64]loan.Loan

func InitStore() Store {
	return Store{}
}

// This isn't great, I would use a UUID dependency for IDs preferably and the map would be map[UUID]loan.Loan
// But I want to keep it simple and with zero dependencies for this small task
func generateID() int64 {
	return rand.Int63()
}

func (store *Store) Add(loan loan.Loan) int64 {
	newID := generateID()

	(*store)[newID] = loan

	return newID
}
