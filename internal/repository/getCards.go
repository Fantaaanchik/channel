package repository

import (
	"chans/models"
	"gorm.io/gorm"
	"sync"
	"time"
)

type Repository struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

type UserRepository struct {
	UserRep models.Card
}

func (r *Repository) GetCards(respch chan any, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 100)
	respch <- "100"
	wg.Done()
}
