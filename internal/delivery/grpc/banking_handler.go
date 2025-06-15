package grpcapi

import (
	"context"

	"github.com/LavaJover/shvark-banking-service/internal/domain"
	bankingpb "github.com/LavaJover/shvark-banking-service/proto/gen"
	"google.golang.org/protobuf/types/known/durationpb"
)

type BankingHandler struct {
	bankingpb.UnimplementedBankingServiceServer
	bankDetailUC domain.BankDetailUsecase
	bankUC 		 domain.BankUsecase
}

func NewBankingHandler(bankDetailUC domain.BankDetailUsecase, bankUC domain.BankUsecase) *BankingHandler {
	return &BankingHandler{bankDetailUC: bankDetailUC, bankUC: bankUC}
}

func (h *BankingHandler) CreateBankDetail(ctx context.Context, r *bankingpb.CreateBankDetailRequest) (*bankingpb.CreateBankDetailResponse, error) {
	bankDetail := domain.BankDetail{
		TraderID: r.TraderId,
		Country: r.Country,
		Currency: r.Currency,
		MinAmount: float32(r.MinAmount),
		MaxAmount: float32(r.MaxAmount),
		BankName: r.BankName,
		PaymentSystem: r.PaymentSystem,
		Delay: r.Delay.AsDuration(),
		Enabled: r.Enabled,
		CardNumber: r.CardNumber,
		Phone: r.Phone,
		Owner: r.Owner,
		MaxOrdersSimultaneosly: r.MaxOrdersSimultaneosly,
		MaxAmountDay: int32(r.MaxAmountDay),
		MaxAmountMonth: int32(r.MaxAmountMonth),
		MaxQuantityDay: int32(r.MaxQuantityDay),
		MaxQuantityMonth: int32(r.MaxQuantityMonth),
		DeviceID: r.DeviceId,
		InflowCurrency: r.InflowCurrency,
	}
	bankDetailID, err := h.bankDetailUC.CreateBankDetail(&bankDetail)
	if err != nil {
		return nil, err
	}

	return &bankingpb.CreateBankDetailResponse{
		BankDetailId: bankDetailID,
		Message: "Bank detail was successfully created!",
	}, nil
}

func (h *BankingHandler) GetBankDetailByID(ctx context.Context, r *bankingpb.GetBankDetailByIDRequest) (*bankingpb.GetBankDetailByIDResponse, error) {
	bankDetailID := r.BankDetailId
	response, err := h.bankDetailUC.GetBankDetailByID(bankDetailID)
	if err != nil {
		return nil, err
	}
	return &bankingpb.GetBankDetailByIDResponse{
		BankDetail: &bankingpb.BankDetail{
			BankDetailId: response.ID,
			TraderId: response.TraderID,
			Currency: response.Currency,
			Country: response.Country,
			MinAmount: float64(response.MinAmount),
			MaxAmount: float64(response.MaxAmount),
			BankName: response.BankName,
			PaymentSystem: response.PaymentSystem,
			Enabled: response.Enabled,
			Delay: durationpb.New(response.Delay),
			CardNumber: response.CardNumber,
			Phone: response.Phone,
			Owner: response.Owner,
			MaxOrdersSimultaneosly: response.MaxOrdersSimultaneosly,
			MaxAmountDay: float64(response.MaxAmountDay),
			MaxAmountMonth: float64(response.MaxAmountMonth),
			MaxQuantityDay: float64(response.MaxQuantityDay),
			MaxQuantityMonth: float64(response.MaxAmountMonth),
			DeviceId: response.DeviceID,
			InflowCurrency: response.InflowCurrency,
		},
	}, nil
}

func (h *BankingHandler) UpdateBankDetail(ctx context.Context, r *bankingpb.UpdateBankDetailRequest) (*bankingpb.UpdateBankDetailResponse, error) {
	bankDetailUpdate := domain.BankDetail{
		ID: r.BankDetail.BankDetailId,
		TraderID: r.BankDetail.TraderId,
		Country: r.BankDetail.Country,
		Currency: r.BankDetail.Currency,
		MinAmount: float32(r.BankDetail.MinAmount),
		MaxAmount: float32(r.BankDetail.MaxAmount),
		BankName: r.BankDetail.BankName,
		PaymentSystem: r.BankDetail.PaymentSystem,
		Delay: r.BankDetail.Delay.AsDuration(),
		Enabled: r.BankDetail.Enabled,
		CardNumber: r.BankDetail.CardNumber,
		Phone: r.BankDetail.Phone,
		Owner: r.BankDetail.Owner,
		MaxOrdersSimultaneosly: r.BankDetail.MaxOrdersSimultaneosly,
		MaxAmountDay: int32(r.BankDetail.MaxAmountDay),
		MaxAmountMonth: int32(r.BankDetail.MaxAmountMonth),
		MaxQuantityDay: int32(r.BankDetail.MaxQuantityDay),
		MaxQuantityMonth: int32(r.BankDetail.MaxQuantityMonth),
		DeviceID: r.BankDetail.DeviceId,
		InflowCurrency: r.BankDetail.InflowCurrency,
	}

	err := h.bankDetailUC.UpdateBankDetail(&bankDetailUpdate)
	if err != nil {
		return nil, err
	}

	return &bankingpb.UpdateBankDetailResponse{}, nil
}

func (h *BankingHandler) DeleteBankDetail(ctx context.Context, r *bankingpb.DeleteBankDetailRequest) (*bankingpb.DeleteBankDetailResponse, error) {
	bankDetailID := r.BankDetailId
	response, err := h.bankDetailUC.DeleteBankDetail(bankDetailID)
	if err != nil {
		return nil, err
	}

	return &bankingpb.DeleteBankDetailResponse{
		BankDetail: &bankingpb.BankDetail{
			BankDetailId: response.ID,
			TraderId: response.TraderID,
			Currency: response.Currency,
			Country: response.Country,
			MinAmount: float64(response.MinAmount),
			MaxAmount: float64(response.MaxAmount),
			BankName: response.BankName,
			PaymentSystem: response.PaymentSystem,
			Enabled: response.Enabled,
			Delay: durationpb.New(response.Delay),
			CardNumber: response.CardNumber,
			Phone: response.Phone,
			Owner: response.Owner,
			MaxOrdersSimultaneosly: response.MaxOrdersSimultaneosly,
			MaxAmountDay: float64(response.MaxAmountDay),
			MaxAmountMonth: float64(response.MaxAmountMonth),
			MaxQuantityDay: float64(response.MaxQuantityDay),
			MaxQuantityMonth: float64(response.MaxQuantityMonth),
			DeviceId: response.DeviceID,
			InflowCurrency: response.InflowCurrency,
		},
	}, nil
}

func (h *BankingHandler) GetBankDetailsByTraderID(ctx context.Context, r *bankingpb.GetBankDetailsByTraderIDRequest) (*bankingpb.GetBankDetailsByTraderIDResponse, error) {
	traderID := r.TraderId
	bankDetails, err := h.bankDetailUC.GetBankDetailsByTraderID(traderID)
	if err != nil {
		return nil, err
	}

	bankDetailsResponse := make([]*bankingpb.BankDetail, len(bankDetails))

	for i, bankDetail := range bankDetails {
		bankDetailsResponse[i] = &bankingpb.BankDetail{
			BankDetailId: bankDetail.ID,
			TraderId: bankDetail.TraderID,
			Currency: bankDetail.Currency,
			Country: bankDetail.Country,
			MinAmount: float64(bankDetail.MinAmount),
			MaxAmount: float64(bankDetail.MaxAmount),
			BankName: bankDetail.BankName,
			PaymentSystem: bankDetail.PaymentSystem,
			Enabled: bankDetail.Enabled,
			Delay: durationpb.New(bankDetail.Delay),
			CardNumber: bankDetail.CardNumber,
			Phone: bankDetail.Phone,
			Owner: bankDetail.Owner,
			MaxOrdersSimultaneosly: bankDetail.MaxOrdersSimultaneosly,
			MaxAmountDay: float64(bankDetail.MaxAmountDay),
			MaxAmountMonth: float64(bankDetail.MaxAmountMonth),
			MaxQuantityDay: float64(bankDetail.MaxQuantityDay),
			MaxQuantityMonth: float64(bankDetail.MaxQuantityMonth),
			DeviceId: bankDetail.DeviceID,
			InflowCurrency: bankDetail.InflowCurrency,
		}
	}

	return &bankingpb.GetBankDetailsByTraderIDResponse{
		BankDetails: bankDetailsResponse,
	}, nil
}

func (h *BankingHandler) GetEligibleBankDetails(ctx context.Context, r *bankingpb.GetEligibleBankDetailsRequest) (*bankingpb.GetEligibleBankDetailsResponse, error) {
	query := &domain.BankDetailQuery{
		Amount: float32(r.Amount),
		Currency: r.Currency,
		PaymentSystem: r.PaymentSystem,
		Country: r.Country,
	}

	bankDetails, err := h.bankDetailUC.GetEligibleBankDetails(query)
	if err != nil {
		return nil, err
	}

	bankDetailsResponse := make([]*bankingpb.BankDetail, len(bankDetails))

	for i, bankDetail := range bankDetails {
		bankDetailsResponse[i] = &bankingpb.BankDetail{
			BankDetailId: bankDetail.ID,
			TraderId: bankDetail.TraderID,
			Currency: bankDetail.Currency,
			Country: bankDetail.Country,
			MinAmount: float64(bankDetail.MinAmount),
			MaxAmount: float64(bankDetail.MaxAmount),
			BankName: bankDetail.BankName,
			PaymentSystem: bankDetail.PaymentSystem,
			Enabled: bankDetail.Enabled,
			Delay: durationpb.New(bankDetail.Delay),
			CardNumber: bankDetail.CardNumber,
			Phone: bankDetail.Phone,
			Owner: bankDetail.Owner,
			MaxOrdersSimultaneosly: bankDetail.MaxOrdersSimultaneosly,
			MaxAmountDay: float64(bankDetail.MaxAmountDay),
			MaxAmountMonth: float64(bankDetail.MaxAmountMonth),
			MaxQuantityDay: float64(bankDetail.MaxQuantityDay),
			MaxQuantityMonth: float64(bankDetail.MaxQuantityMonth),
			DeviceId: bankDetail.DeviceID,
			InflowCurrency: bankDetail.InflowCurrency,
		}
	}

	return &bankingpb.GetEligibleBankDetailsResponse{
		BankDetails: bankDetailsResponse,
	}, nil
}

func (h *BankingHandler) CreateBank(ctx context.Context, r *bankingpb.CreateBankRequest) (*bankingpb.CreateBankResponse, error) {
	bank := domain.Bank{
		BankName: r.BankName,
		UserfriendlyName: r.UserfriendlyName,
		Country: r.Country,
		SBP_ID: r.SbpId,
		IconUrl: r.IconUrl,
	}
	if err := h.bankUC.CreateBank(&bank); err != nil {
		return nil, err
	}

	return &bankingpb.CreateBankResponse{
		Id: bank.ID,
	}, nil
}

func (h *BankingHandler) GetBankByName(ctx context.Context, r *bankingpb.GetBankByNameRequest) (*bankingpb.GetBankByNameResponse, error) {
	bankName := r.BankName
	bank, err := h.bankUC.GetBankByName(bankName)
	if err != nil {
		return nil, err
	}
	return &bankingpb.GetBankByNameResponse{
		Bank: &bankingpb.Bank{
			Id: bank.ID,
			BankName: bank.BankName,
			UserfriendlyName: bank.UserfriendlyName,
			Country: bank.Country,
			SbpId: bank.SBP_ID,
			IconUrl: bank.IconUrl,
		},
	}, nil
}

func (h *BankingHandler) GetBankByID(ctx context.Context, r *bankingpb.GetBankByIDRequest) (*bankingpb.GetBankByIDResponse, error) {
	bankID := r.BankId
	bank, err := h.bankUC.GetBankByID(bankID)
	if err != nil {
		return nil, err
	}
	return &bankingpb.GetBankByIDResponse{
		Bank: &bankingpb.Bank{
			Id: bank.ID,
			BankName: bank.BankName,
			UserfriendlyName: bank.UserfriendlyName,
			Country: bank.Country,
			SbpId: bank.SBP_ID,
			IconUrl: bank.IconUrl,
		},
	}, nil
}

func (h *BankingHandler) DeleteBankByID(ctx context.Context, r *bankingpb.DeleteBankByIDRequest) (*bankingpb.DeleteBankByIDResponse, error) {
	bankID := r.BankId
	if err := h.bankUC.DeleteBankByID(bankID); err != nil {
		return nil, err
	}
	return &bankingpb.DeleteBankByIDResponse{}, nil
}

func (h *BankingHandler) GetBanksByCountry(ctx context.Context, r *bankingpb.GetBanksByCountryRequest) (*bankingpb.GetBanksByCountryResponse, error) {
	country := r.Country
	banks, err := h.bankUC.GetBanksByCountry(country)
	if err != nil {
		return nil, err
	}
	respBanks := make([]*bankingpb.Bank, len(banks))
	for i, bank := range banks {
		respBanks[i] = &bankingpb.Bank{
			Id: bank.ID,
			BankName: bank.BankName,
			UserfriendlyName: bank.UserfriendlyName,
			Country: bank.Country,
			SbpId: bank.SBP_ID,
			IconUrl: bank.IconUrl,
		}
	}

	return &bankingpb.GetBanksByCountryResponse{
		Banks: respBanks,
	}, nil
}