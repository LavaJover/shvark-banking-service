package usecase

import "github.com/LavaJover/shvark-banking-service/internal/domain"

type DefaultBankUsecase struct {
	BankRepo domain.BankRepo
}

func NewDefaultBankUsecase(bankRepo domain.BankRepo) *DefaultBankUsecase {
	return &DefaultBankUsecase{BankRepo: bankRepo}
}

func (uc *DefaultBankUsecase) CreateBank(bank *domain.Bank) error {
	return uc.BankRepo.CreateBank(bank)
}

func (uc *DefaultBankUsecase) GetBankByName(bankName string) (*domain.Bank, error) {
	return uc.BankRepo.GetBankByName(bankName)
}

func (uc *DefaultBankUsecase) GetBankByID(bankID string) (*domain.Bank, error) {
	return uc.BankRepo.GetBankByID(bankID)
}

func (uc *DefaultBankUsecase) GetBanksByCountry(country string) ([]*domain.Bank, error) {
	return uc.BankRepo.GetBanksByCountry(country)
}

func (uc *DefaultBankUsecase) DeleteBankByID(bankID string) error {
	return uc.BankRepo.DeleteBankByID(bankID)
}