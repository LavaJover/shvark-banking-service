package postgres

import (
	"github.com/LavaJover/shvark-banking-service/internal/domain"
	"gorm.io/gorm"
	"github.com/google/uuid"
)

type DefaultBankDetailRepository struct {
	DB *gorm.DB
}

func (r *DefaultBankDetailRepository) CreateBankDetail(bankDetail *domain.BankDetail) (string, error) {
	bankDetailModel := BankDetailModel{
		ID: uuid.New().String(),
		TraderID: bankDetail.TraderID,
		Country: bankDetail.Country,
		Currency: bankDetail.Currency,
		MinAmount: bankDetail.MinAmount,
		MaxAmount: bankDetail.MaxAmount,
		BankName: bankDetail.BankName,
		PaymentSystem: bankDetail.PaymentSystem,
		Delay: bankDetail.Delay,
		Enabled: bankDetail.Enabled,
	}
	if err := r.DB.Create(&bankDetailModel).Error; err != nil {
		return "", err
	}

	bankDetail.ID = bankDetailModel.ID
	return bankDetail.ID, nil
}

func (r *DefaultBankDetailRepository) DeleteBankDetail(bankDetailID string) (*domain.BankDetail, error) {
	var bankDetailModel BankDetailModel

	if err := r.DB.Where("id = ?", bankDetailID).First(&bankDetailModel).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, domain.ErrBankDetailNotFound
		}
		return nil, err
	}

	if err := r.DB.Delete(&bankDetailModel).Error; err != nil {
		return nil, err
	}

	return &domain.BankDetail{
		ID: bankDetailModel.ID,
		TraderID: bankDetailModel.TraderID,
		Country: bankDetailModel.Country,
		Currency: bankDetailModel.Currency,
		MinAmount: bankDetailModel.MinAmount,
		MaxAmount: bankDetailModel.MaxAmount,
		BankName: bankDetailModel.BankName,
		PaymentSystem: bankDetailModel.PaymentSystem,
		Delay: bankDetailModel.Delay,
		Enabled: bankDetailModel.Enabled,
	}, nil

}

func (r *DefaultBankDetailRepository) UpdateBankDetail(bankDetail *domain.BankDetail) error {
	modelToUpdate := BankDetailModel{
		TraderID: bankDetail.TraderID,
		Country: bankDetail.Country,
		Currency: bankDetail.Currency,
		MinAmount: bankDetail.MinAmount,
		MaxAmount: bankDetail.MaxAmount,
		BankName: bankDetail.BankName,
		PaymentSystem: bankDetail.PaymentSystem,
		Delay: bankDetail.Delay,
		Enabled: bankDetail.Enabled,
	}

	if err := r.DB.Save(&modelToUpdate).Error; err != nil {
		return err
	}

	return nil
}

func (r *DefaultBankDetailRepository) GetBankDetailByID(bankDetailID string) (*domain.BankDetail, error) {
	var bankDetailModel BankDetailModel

	if err := r.DB.Where("id = ?", bankDetailID).First(&bankDetailModel).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, domain.ErrBankDetailNotFound
		}
		return nil, err
	}

	return &domain.BankDetail{
		ID: bankDetailModel.ID,
		TraderID: bankDetailModel.TraderID,
		Country: bankDetailModel.Country,
		Currency: bankDetailModel.Currency,
		MinAmount: bankDetailModel.MinAmount,
		MaxAmount: bankDetailModel.MaxAmount,
		BankName: bankDetailModel.BankName,
		PaymentSystem: bankDetailModel.PaymentSystem,
		Delay: bankDetailModel.Delay,
		Enabled: bankDetailModel.Enabled,
	}, nil
}