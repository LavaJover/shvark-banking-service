package postgres

import "time"

type BankModel struct {
	ID 		 		 string `gorm:"primaryKey;type:uuid"`
	BankName 		 string
	UserfriendlyName string
	Country 		 string
	SBP_ID			 string
	IconURL 		 string
	CreatedAt 		 time.Time
	UpdatedAt 		 time.Time
}