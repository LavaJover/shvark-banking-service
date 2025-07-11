package postgres

import (
	"github.com/LavaJover/shvark-banking-service/internal/domain"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type DefaultBankDetailRepository struct {
	DB *gorm.DB
}

func NewDefaultbankDetailRepository(db *gorm.DB) (*DefaultBankDetailRepository, error) {
	return &DefaultBankDetailRepository{DB: db}, nil
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
		CardNumber: bankDetail.CardNumber,
		Phone: bankDetail.Phone,
		Owner: bankDetail.Owner,
		MaxOrdersSimultaneosly: bankDetail.MaxOrdersSimultaneosly,
		MaxAmountDay: bankDetail.MaxAmountDay,
		MaxAmountMonth: bankDetail.MaxAmountMonth,
		MaxQuantityDay: bankDetail.MaxQuantityDay,
		MaxQuantityMonth: bankDetail.MaxAmountMonth,
		DeviceID: bankDetail.DeviceID,
		InflowCurrency: bankDetail.InflowCurrency,
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
		CardNumber: bankDetailModel.CardNumber,
		Phone: bankDetailModel.Phone,
		Owner: bankDetailModel.Owner,
		MaxOrdersSimultaneosly: bankDetailModel.MaxOrdersSimultaneosly,
		MaxAmountDay: bankDetailModel.MaxAmountDay,
		MaxAmountMonth: bankDetailModel.MaxAmountMonth,
		MaxQuantityDay: bankDetailModel.MaxQuantityDay,
		MaxQuantityMonth: bankDetailModel.MaxQuantityMonth,
		DeviceID: bankDetailModel.DeviceID,
		InflowCurrency: bankDetailModel.InflowCurrency,
	}, nil

}

func (r *DefaultBankDetailRepository) UpdateBankDetail(bankDetail *domain.BankDetail) error {
	modelToUpdate := BankDetailModel{
		ID: bankDetail.ID,
		TraderID: bankDetail.TraderID,
		Country: bankDetail.Country,
		Currency: bankDetail.Currency,
		MinAmount: bankDetail.MinAmount,
		MaxAmount: bankDetail.MaxAmount,
		BankName: bankDetail.BankName,
		PaymentSystem: bankDetail.PaymentSystem,
		Delay: bankDetail.Delay,
		Enabled: bankDetail.Enabled,
		CardNumber: bankDetail.CardNumber,
		Phone: bankDetail.Phone,
		Owner: bankDetail.Owner,
		MaxOrdersSimultaneosly: bankDetail.MaxOrdersSimultaneosly,
		MaxAmountDay: bankDetail.MaxAmountDay,
		MaxAmountMonth: bankDetail.MaxAmountMonth,
		MaxQuantityDay: bankDetail.MaxQuantityDay,
		MaxQuantityMonth: bankDetail.MaxAmountMonth,
		DeviceID: bankDetail.DeviceID,
		InflowCurrency: bankDetail.InflowCurrency,
	}

	if err := r.DB.Updates(&modelToUpdate).Error; err != nil {
		return err
	}

	if err := r.DB.Model(&modelToUpdate).Updates(map[string]interface{}{
		"enabled": bankDetail.Enabled,
		"delay": bankDetail.Delay,
		}).Error; err != nil {
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
		CardNumber: bankDetailModel.CardNumber,
		Phone: bankDetailModel.Phone,
		Owner: bankDetailModel.Owner,
		MaxOrdersSimultaneosly: bankDetailModel.MaxOrdersSimultaneosly,
		MaxAmountDay: bankDetailModel.MaxAmountDay,
		MaxAmountMonth: bankDetailModel.MaxAmountMonth,
		MaxQuantityDay: bankDetailModel.MaxQuantityDay,
		MaxQuantityMonth: bankDetailModel.MaxQuantityMonth,
		DeviceID: bankDetailModel.DeviceID,
		InflowCurrency: bankDetailModel.InflowCurrency,
	}, nil
}

func (r *DefaultBankDetailRepository) GetBankDetailsByTraderID(traderID string) ([]*domain.BankDetail, error) {
	var bankDetailModels []BankDetailModel
	if err := r.DB.Where("trader_id = ?", traderID).Find(&bankDetailModels).Error; err != nil {
		return nil, err
	}

	bankDetailsResponse := make([]*domain.BankDetail, len(bankDetailModels))

	for i, bankDetailModel := range bankDetailModels {
		bankDetailsResponse[i] = &domain.BankDetail{
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
			CardNumber: bankDetailModel.CardNumber,
			Phone: bankDetailModel.Phone,
			Owner: bankDetailModel.Owner,
			MaxOrdersSimultaneosly: bankDetailModel.MaxOrdersSimultaneosly,
			MaxAmountDay: bankDetailModel.MaxAmountDay,
			MaxAmountMonth: bankDetailModel.MaxAmountMonth,
			MaxQuantityDay: bankDetailModel.MaxQuantityDay,
			MaxQuantityMonth: bankDetailModel.MaxQuantityMonth,
			DeviceID: bankDetailModel.DeviceID,
			InflowCurrency: bankDetailModel.InflowCurrency,
		}
	}

	return bankDetailsResponse, nil
}

func (r *DefaultBankDetailRepository) GetEligibleBankDetails(query *domain.BankDetailQuery) ([]*domain.BankDetail, error) {
	var bankDetailModels []BankDetailModel
	err := r.DB.
		Where("enabled = ?", true).
		Where("min_amount <= ? AND max_amount >= ?", query.Amount, query.Amount).
		Where("payment_system = ?", query.PaymentSystem).
		Find(&bankDetailModels).Error
	
	if err != nil {
		return nil, status.Errorf(codes.Internal, "db query failed: %v\n", err.Error())
	}

	bankDetails := make([]*domain.BankDetail, len(bankDetailModels))
	for i, bankDetailModel := range bankDetailModels {
		bankDetails[i] = &domain.BankDetail{
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
			CardNumber: bankDetailModel.CardNumber,
			Phone: bankDetailModel.Phone,
			Owner: bankDetailModel.Owner,
			MaxOrdersSimultaneosly: bankDetailModel.MaxOrdersSimultaneosly,
			MaxAmountDay: bankDetailModel.MaxAmountDay,
			MaxAmountMonth: bankDetailModel.MaxAmountMonth,
			MaxQuantityDay: bankDetailModel.MaxQuantityDay,
			MaxQuantityMonth: bankDetailModel.MaxQuantityMonth,
			DeviceID: bankDetailModel.DeviceID,
			InflowCurrency: bankDetailModel.InflowCurrency,
		}
	}

	return bankDetails, nil
}

// func (r *DefaultBankDetailRepository) GetBankDetailsStatisticsByTraderID(traderID string, page, limit int64) ([]*domain.BankDetailStatistics, error) {

// }