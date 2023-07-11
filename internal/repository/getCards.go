package repository

import (
	"chans/internal/db"
	"chans/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"sync"
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

func (r *Repository) GetCardFromDB(c *gin.Context) ([]models.Card, error) {
	var cards []models.Card
	var (
		batchSize int
		start     int
		end       int
	)

	fmt.Println("Введите размер batch-а")
	_, err := fmt.Scan(&batchSize)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad type of value, need integer"})
		return nil, err
	}
	fmt.Println("Введите размер batch-а")
	_, err = fmt.Scan(&start)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad type of value, need integer"})
		return nil, err
	}
	endNum := start
	if start == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Число 0 не подходит, вам нужно написать число больше, либо равным 1"})
		return nil, err
	} else if start > 1 {
		start = (start - 1) * batchSize
		start++
	}

	end = endNum * batchSize

	var wg sync.WaitGroup
	cardChan := make(chan models.Card)
	rowsAffectedChan := make(chan int64)

	wg.Add(1)
	go func() {
		defer close(cardChan)
		db.GetDB().FindInBatches(&cards, batchSize, func(tx *gorm.DB, batch int) error {
			fmt.Println("batch", batch)
			for _, result := range cards {
				if result.Id >= start && result.Id <= end {
					cardChan <- result
				}
			}

			rowsAffected := tx.RowsAffected
			c.JSON(http.StatusOK, gin.H{"rowsAffected": rowsAffected})

			tx.Save(&cards)

			return nil
		})
		wg.Done()
	}()

	go func() {
		for card := range cardChan {
			c.JSON(http.StatusOK, gin.H{"cards": card})
		}
	}()

	totalRowsAffected := int64(0)

	go func() {
		for rowsAffected := range rowsAffectedChan {
			totalRowsAffected += rowsAffected
		}
	}()

	wg.Wait()

	close(rowsAffectedChan)

	c.JSON(http.StatusOK, gin.H{
		"cards": cards,
	})
	return cards, nil
}
