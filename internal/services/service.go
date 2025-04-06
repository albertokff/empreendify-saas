package services

import (
	"time"
)

type Service struct {
	ID 			uint 	`gorm:"primaryKey"`
	ClientName  string  `gorm:"size:100"`
	Description string  `gorm:"type:text"`
	Price 		float64
	Status 		string  `gorm:"size:20`
	Date 		time.Time 
}