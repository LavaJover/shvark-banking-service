package postgres

import (
	"github.com/LavaJover/shvark-banking-service/internal/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DefaultBankRepository struct {
	DB *gorm.DB
}

func NewDefaultBankRepository(db *gorm.DB) *DefaultBankRepository {
	return &DefaultBankRepository{DB: db}
}

func (r *DefaultBankRepository) CreateBank(bank *domain.Bank) error {
	bankModel := BankModel{
		ID: uuid.New().String(),
		BankName: bank.BankName,
		UserfriendlyName: bank.UserfriendlyName,
		Country: bank.Country,
		SBP_ID: bank.SBP_ID,
		IconURL: bank.IconUrl,
	}
	if err := r.DB.Create(&bankModel).Error; err != nil {
		return err
	}
	bank.ID = bankModel.ID
	return nil
}

func (r *DefaultBankRepository) GetBankByName(bankName string) (*domain.Bank, error) {
	var bankModel BankModel
	if err := r.DB.Where("bank_name = ?", bankName).First(&bankModel).Error; err != nil {
		return nil, err
	}

	return &domain.Bank{
		ID: bankModel.ID,
		BankName: bankModel.BankName,
		UserfriendlyName: bankModel.UserfriendlyName,
		Country: bankModel.Country,
		SBP_ID: bankModel.SBP_ID,
		IconUrl: bankModel.IconURL,
	}, nil
}

func (r *DefaultBankRepository) GetBanksByCountry(country string) ([]*domain.Bank, error) {
	var bankModels []BankModel
	if err := r.DB.Where("country = ?", country).Find(&bankModels).Error; err != nil {
		return nil, err
	}
	banks := make([]*domain.Bank, len(bankModels))
	for i, bankModel := range bankModels {
		banks[i] = &domain.Bank{
			ID: bankModel.ID,
			BankName: bankModel.BankName,
			UserfriendlyName: bankModel.UserfriendlyName,
			Country: bankModel.Country,
			SBP_ID: bankModel.SBP_ID,
			IconUrl: bankModel.IconURL,
		}
	}

	return banks, nil
}

func (r *DefaultBankRepository) GetBankByID(bankID string) (*domain.Bank, error) {
	var bankModel BankModel
	if err := r.DB.Where("id = ?", bankID).First(&bankModel).Error; err != nil {
		return nil, err
	}

	return &domain.Bank{
		ID: bankModel.ID,
		BankName: bankModel.BankName,
		UserfriendlyName: bankModel.UserfriendlyName,
		Country: bankModel.Country,
		SBP_ID: bankModel.SBP_ID,
		IconUrl: bankModel.IconURL,
	}, nil
}

func (r *DefaultBankRepository) DeleteBankByID(bankID string) error {
	if err := r.DB.Delete(&BankModel{ID: bankID}).Error; err != nil {
		return err
	}
	return nil
}