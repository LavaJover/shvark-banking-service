package domain

type BankRepo interface {
	CreateBank(bank *Bank) error
	GetBankByName(bankName string) (*Bank, error)
	GetBanksByCountry(country string) ([]*Bank, error)
	DeleteBankByID(bankID string) error
	GetBankByID(bankID string) (*Bank, error)
}