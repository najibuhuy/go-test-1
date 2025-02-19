package repository

import (
	"context"
	"go-test-1/domain/bank_account/model"
	"go-test-1/infrastructure/database"
)

type BankAccountRepository interface {
	CheckBalanceRepository(ctx context.Context, bankNo string) (response model.CheckBalanceResponse, err error)
	DepositBalanceRepository(ctx context.Context, request *model.DepositBalanceRequest) (response model.DepositBalanceResponse, err error)
	WithDrawBalanceRepository(ctx context.Context, request *model.WithDrawBalanceRequest) (response model.WithDrawBalanceResponse, err error)
	SignUpRepository(ctx context.Context, request *model.SignUpRequest) (response model.SignUpResponse, err error)
}

type bankAccountRepository struct {
	Database *database.Database
}

func NewBankAccountRepository(db *database.Database) BankAccountRepository {
	return &bankAccountRepository{
		Database: db,
	}
}
