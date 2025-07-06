package usecase

import (
	"github.com/LavaJover/shvark-banking-service/internal/domain"
)

type DefaultBankDetailUsecase struct {
	repo domain.BankDetailRepository
}

func NewDefaultBankDetailUsecase(repo domain.BankDetailRepository) *DefaultBankDetailUsecase {
	return &DefaultBankDetailUsecase{repo: repo}
}

func (uc *DefaultBankDetailUsecase) CreateBankDetail(bankDetail *domain.BankDetail) (string, error) {
	return uc.repo.CreateBankDetail(bankDetail)
}

func (uc *DefaultBankDetailUsecase) GetBankDetailByID(bankDetailID string) (*domain.BankDetail, error) {
	return uc.repo.GetBankDetailByID(bankDetailID)
}

func (uc *DefaultBankDetailUsecase) UpdateBankDetail(bankDetail *domain.BankDetail) error {
	return uc.repo.UpdateBankDetail(bankDetail)
}

func (uc *DefaultBankDetailUsecase) DeleteBankDetail(bankDetailID string) (*domain.BankDetail, error) {
	return uc.repo.DeleteBankDetail(bankDetailID)
}

func (uc *DefaultBankDetailUsecase) GetBankDetailsByTraderID(traderID string) ([]*domain.BankDetail, error) {
	return uc.repo.GetBankDetailsByTraderID(traderID)
}

func (uc *DefaultBankDetailUsecase) GetEligibleBankDetails(query *domain.BankDetailQuery) ([]*domain.BankDetail, error) {
	return uc.repo.GetEligibleBankDetails(query)
}

// func (uc *DefaultBankDetailUsecase) GetBankDetailsStatisticsByTraderID(traderID string, page, limit int64) ([]*domain.BankDetailStatistics, error) {
// 	bankDetails, err := uc.GetBankDetailsByTraderID(traderID)
// 	if err != nil {
// 		return nil, err
// 	}

// }