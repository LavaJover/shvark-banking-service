package grpcapi

import (
	"context"

	"github.com/LavaJover/shvark-banking-service/internal/domain"
	bankingpb "github.com/LavaJover/shvark-banking-service/proto/gen"
	"google.golang.org/protobuf/types/known/durationpb"
)

type BankDetailHandler struct {
	bankingpb.UnimplementedBankingServiceServer
	uc domain.BankDetailUsecase
}

func NewbankDetailHandler(uc domain.BankDetailUsecase) *BankDetailHandler {
	return &BankDetailHandler{uc: uc}
}

func (h *BankDetailHandler) CreateBankDetail(ctx context.Context, r *bankingpb.CreateBankDetailRequest) (*bankingpb.CreateBankDetailResponse, error) {
	bankDetail := domain.BankDetail{
		TraderID: r.TraderId,
		Country: r.Country,
		Currency: r.Currency,
		MinAmount: float32(r.MinAmount),
		MaxAmount: float32(r.MaxAmount),
		BankName: r.BankName,
		PaymentSystem: r.PaymentSystem,
		Delay: r.Delay.AsDuration(),
	}
	bankDetailID, err := h.uc.CreateBankDetail(&bankDetail)
	if err != nil {
		return nil, err
	}

	return &bankingpb.CreateBankDetailResponse{
		BankDetailId: bankDetailID,
		Message: "Bank detail was successfully created!",
	}, nil
}

func (h *BankDetailHandler) GetBankDetailByID(ctx context.Context, r *bankingpb.GetBankDetailByIDRequest) (*bankingpb.GetBankDetailByIDResponse, error) {
	bankDetailID := r.BankDetailId
	response, err := h.uc.GetBankDetailByID(bankDetailID)
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
		},
	}, nil
}

func (h *BankDetailHandler) UpdateBankDetail(ctx context.Context, r *bankingpb.UpdateBankDetailRequest) (*bankingpb.UpdateBankDetailResponse, error) {
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
	}

	err := h.uc.UpdateBankDetail(&bankDetailUpdate)
	if err != nil {
		return nil, err
	}

	return &bankingpb.UpdateBankDetailResponse{}, nil
}

func (h *BankDetailHandler) DeletebankDetail(ctx context.Context, r *bankingpb.DeleteBankDetailRequest) (*bankingpb.DeleteBankDetailResponse, error) {
	bankDetailID := r.BankDetailId
	response, err := h.uc.DeleteBankDetail(bankDetailID)
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
		},
	}, nil
}