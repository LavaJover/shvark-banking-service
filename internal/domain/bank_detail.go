package domain

import "time"

type BankDetail struct {
	ID 						string
	TraderID 				string
	Country 				string
	Currency 				string
	InflowCurrency			string
	MinAmount 				float32
	MaxAmount 				float32
	BankName 				string
	PaymentSystem 			string
	Delay					time.Duration
	Enabled 				bool
	CardNumber 				string
	Phone 					string
	Owner 					string
	MaxOrdersSimultaneosly  int32
	MaxAmountDay			int32
	MaxAmountMonth  		int32
	MaxQuantityDay			int32
	MaxQuantityMonth		int32
	DeviceID				string
}

type BankDetailQuery struct {
	Amount 			float32
	Currency 		string
	PaymentSystem 	string
	Country 		string
}

type Statistics struct {
	ActiveOrders 	int64
	ActiveOrdersSum int64
	TodayStatistics TodayStatistics
	MonthStatistics MonthStatistics
}

type TodayStatistics struct {
	SucceedOrders 	 	int64
	CanceledOrders   	int64
	SucceedOrdersSum 	float64
	CanceledOrdersSum 	float64
}

type MonthStatistics struct {
	SucceedOrders 	 	int64
	CanceledOrders   	int64
	SucceedOrdersSum 	float64
	CanceledOrdersSum 	float64
}

type BankDetailStatistics struct {
	BankDetail BankDetail
	Statistics Statistics
}