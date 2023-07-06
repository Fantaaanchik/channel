package service

import (
	"chans/internal/db"
	"chans/models"
	"fmt"
	"sync"
	"time"
)

func CallChan(respch chan any, wg *sync.WaitGroup) {
	var cards models.Card
	err := db.GetDB().Find(&cards).Error
	if err != nil {
		fmt.Println("cannot find data from table cards, err: ", err.Error())
	}
	time.Sleep(time.Millisecond * 100)
	respch <- cards
	wg.Done()
}
