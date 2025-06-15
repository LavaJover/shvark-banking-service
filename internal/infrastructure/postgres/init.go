package postgres

import (
	"log"

	"github.com/LavaJover/shvark-banking-service/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func MustInitDB(cfg *config.BankingConfig) *gorm.DB {
	dsn := cfg.BankingDB.Dsn
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to load db: %v\n", err.Error())
	}

	db.AutoMigrate(&BankModel{}, &BankDetailModel{})

	return db
}