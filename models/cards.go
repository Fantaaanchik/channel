package models

type Card struct {
	Id         int    `json:"id" gorm:"column:id"`
	Balance    string `json:"balance" gorm:"column:balance"`
	Number     string `json:"number" gorm:"column:number"`
	Bank       string `json:"bank" gorm:"column:bank"`
	ExpireDate string `json:"expire_date" gorm:"column:expire_date"`
}
