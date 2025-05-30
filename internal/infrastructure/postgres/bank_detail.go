package postgres

import "time"

type BankDetailModel struct {
	ID 				string	`gorm:"primaryKey;type:uuid"`
	TraderID 		string	`gorm:"type:uuid;not null"`
	Country 		string	
	Currency 		string	
	MinAmount 		float32
	MaxAmount 		float32
	BankName 		string
	PaymentSystem 	string
	Delay			time.Duration
	Enabled 		bool
}