package models

import (
	"time"
)

type Service struct {
	ID 			uint 		`gorm:"primaryKey" json:"id"`
	ClientName  string  	`gorm:"size:100" json:"client_name"`
	Description string  	`gorm:"type:text" json:"description"`
	Price 		float64 	`json:"price"`
	Status 		string  	`gorm:"size:20" json:"status"`
	Date 		time.Time   `json:"date"`
}