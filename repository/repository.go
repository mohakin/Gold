package repository

import (
	"errors"
	"time"
)

var (
	errUpdateFailed = errors.New("update failed")
	errDeleteFailed = errors.New("delete failed")
)

// NOTE: Repository is the interface which must be satisfied in order to connect a database.

type Repository interface {
	Migrate() error
	InsertHolding(h Holdings) (*Holdings, error)
	AllHoldings() ([]Holdings, error)
	GetHoldingsByID(id int) (*Holdings, error)
	UpdateHolding(id int, updated Holdings) error
	DeleteHolding(id int64) error
}

// NOTE: Holdings is the type for the user's gold holdings.

type Holdings struct {
	ID            int64     `json:"id"`
	Amount        int       `json:"amount"`
	PurchaseDate  time.Time `json:"purchase_date"`
	PurchasePrice int       `json:"purchase_price"`
}
