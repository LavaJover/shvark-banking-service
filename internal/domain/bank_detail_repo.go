package domain

type BankDetailRepository interface {
	CreateBankDetail(bankDetail *BankDetail) (string, error)
	DeleteBankDetail(bankDetailID string) (*BankDetail, error)
	UpdateBankDetail(bankDetail *BankDetail) error
	GetBankDetailByID(bankDetailID string) (*BankDetail, error)
	GetBankDetailsByTraderID(traderID string) ([]*BankDetail, error)
	GetEligibleBankDetails(query *BankDetailQuery) ([]*BankDetail, error)

	// GetBankDetailsStatisticsByTraderID(traderID string, page, limit int64) ([]*BankDetailStatistics, error)
}