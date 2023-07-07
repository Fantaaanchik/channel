package main

import (
	"chans/config"
	"chans/internal/db"
	"chans/models"
	"fmt"
	"gorm.io/gorm"
	"log"
	"sync"
)

//var wg sync.WaitGroup

func main() {
	config.ReadConfig("./config/config.json")
	fmt.Println("Connection to config success!!!")
	db.ConnectionToDB()
	fmt.Println("Connection to DB success!!!")
	defer db.CloseDB()

	var cards []models.Card

	cardChan := make(chan models.Card)

	rowsAffectedChan := make(chan int64, 100)
	go func() {
		defer close(cardChan)
		err := db.GetDB().FindInBatches(&cards, 100, func(tx *gorm.DB, batch int) error {
			fmt.Println("batch", batch)
			for _, result := range cards {
				if result.Id >= 1 && result.Id <= 500 {
					cardChan <- result
				}
			}

			rowsAffected := tx.RowsAffected
			fmt.Println("rowsAffected", rowsAffected)

			tx.Save(&cards)

			return nil
		})
		if err != nil {
			log.Println("ERROR: ", err)
		}
	}()
	fmt.Println("--------------------------------------------------------------------------------------")
	for card := range cardChan {
		// Обработка данных из заданного промежутка (от 100 до 200)
		// Добавьте соответствующую логику обработки сюда

		fmt.Println(card)
	}
	//close(cardChan)

	totalRowsAffected := int64(0)

	for rowsAffected := range rowsAffectedChan {
		totalRowsAffected += rowsAffected
	}

	fmt.Println("Total Rows Affected:", totalRowsAffected)
	//----------------------------------------------------------------------------------------
	//batchSize := 10
	//n := 5 // Количество повторений
	//
	//start := time.Now()
	//numOfAttempts := 0
	//
	//for i := 0; i < n; i++ {
	//	resp := make(chan []models.Card, 100)
	//	wg := &sync.WaitGroup{}
	//	wg.Add(1)
	//	go CallChan(resp, wg, batchSize)
	//	wg.Wait()
	//	close(resp)
	//
	//	data := <-resp
	//	fmt.Println("Data:", data)
	//	fmt.Println()
	//
	//	numOfAttempts++
	//
	//}
	//fmt.Println("Number of attempts:", numOfAttempts)
	//fmt.Println("Took:", time.Since(start))
}

//wg.Add(3)
//go say("Riki")
//go say("Invoker")
//go say("Lion")
//wg.Wait()
//}

//func say(s string) {
//	for i := 0; i < 3; i++ {
//		fmt.Println(s)
//		time.Sleep(time.Millisecond * 500)
//	}
//	wg.Done()
//}

//func ChannelExample(){
//message1 := make(chan string)
//message2 := make(chan string)
//
//go func() {
//	for {
//		time.Sleep(time.Millisecond * 500)
//		message1 <- "Прошло пол секунды"
//	}
//}()
//
//go func() {
//	for {
//		time.Sleep(time.Second * 2)
//		message2 <- "Прошло 2 секунды"
//	}
//}()
//
//for {
//	select {
//	case msg := <-message1:
//		fmt.Println(msg)
//	case msg := <-message2:
//		fmt.Println(msg)
//	}
//}
//}

func CallChan(channel chan []models.Card, wg *sync.WaitGroup, batchSize int) {
	defer wg.Done()

	type Card struct {
		Id         int    `json:"id" gorm:"column:id"`
		Balance    string `json:"balance" gorm:"column:balance"`
		Number     string `json:"number" gorm:"column:number"`
		Bank       string `json:"bank" gorm:"column:bank"`
		ExpireDate string `json:"expire_date" gorm:"column:expire_date"`
	}

	defer db.CloseDB()

	err := db.GetDB().FindInBatches(&[]models.Card{}, batchSize, func(tx *gorm.DB, batchCount int) error {
		var cards []models.Card
		if err := tx.Find(&cards).Error; err != nil {
			log.Println(err.Error())
		}
		err := tx.AutoMigrate(&Card{})
		if err != nil {
			log.Println(err.Error())
		}

		channel <- cards

		return nil
	}).Error

	if err != nil {
		log.Println(err.Error())
	}
	for cards := range channel {
		for _, card := range cards {
			fmt.Println(card)
		}
	}

}
