package domain

type BankDetailRepository interface {
	CreateBankDetail(bankDetail *BankDetail) (string, error)
	DeleteBankDetail(bankDetailID string) (*BankDetail, error)
	UpdateBankDetail(bankDetail *BankDetail) error
	GetBankDetailByID(bankDetailID string) (*BankDetail, error)
}