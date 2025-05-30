package grpc

import (
	"context"

	"github.com/LavaJover/shvark-banking-service/internal/domain"
	bankingpb "github.com/LavaJover/shvark-banking-service/proto/gen"
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
		return &bankingpb.CreateBankDetailResponse{
			BankDetailId: bankDetailID,
			Message: "Failed to create bank detail!",
		}, err
	}

	return &bankingpb.CreateBankDetailResponse{
		BankDetailId: bankDetailID,
		Message: "Bank detail was successfully created!",
	}, nil
}