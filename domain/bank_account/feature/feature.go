package feature

import (
	"context"
	"go-test-1/domain/bank_account/model"
	"go-test-1/domain/bank_account/repository"
)

type BankAccountFeature interface {
	CheckBalanceFeature(ctx context.Context, bankNo string) (response model.CheckBalanceResponse, err error)
	DepositBalanceFeature(ctx context.Context, request *model.DepositBalanceRequest) (response model.DepositBalanceResponse, err error)
	WithDrawBalanceFeature(ctx context.Context, request *model.WithDrawBalanceRequest) (response model.WithDrawBalanceResponse, err error)
	SignUpFeature(ctx context.Context, request *model.SignUpRequest) (response model.SignUpResponse, err error)
}

type bankAccountFeature struct {
	bankAccountRepository repository.BankAccountRepository
}

func NewBankAccountFeature(bankAccountRepo repository.BankAccountRepository) BankAccountFeature {
	return &bankAccountFeature{
		bankAccountRepository: bankAccountRepo,
	}
}
