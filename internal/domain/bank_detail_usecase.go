package domain

type BankDetailUsecase interface {
	CreateBankDetail(bankDetail *BankDetail) (string, error)
	GetBankDetailByID(bankDetailID string) (*BankDetail, error)
	DeleteBankDetail(bankDetailID string) (*BankDetail, error)
	UpdateBankDetail(bankDetail *BankDetail) error
	GetBankDetailsByTraderID(traderID string) ([]*BankDetail, error)
	GetEligibleBankDetails(query *BankDetailQuery) ([]*BankDetail, error)
}