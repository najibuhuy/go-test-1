package repository

import (
	"context"
	"errors"
	"go-test-1/domain/bank_account/model"
	"go-test-1/domain/shared/constant"
	Error "go-test-1/domain/shared/error"
)

func (bar bankAccountRepository) CheckBalanceRepository(ctx context.Context, bankNo string) (response model.CheckBalanceResponse, err error) {
	tx, err := bar.Database.Pool.Begin(ctx)

	row := tx.QueryRow(ctx, "SELECT bank_no, balance FROM bank_accounts WHERE bank_no = $1", bankNo)
	var account model.BankAccount
	if err := row.Scan(&account.BankNo, &account.Balance); err != nil {
		err = Error.New(constant.ErrInvalidRequest, "error select query", errors.New("error query"))
		return response, err
	}
	if err := tx.Commit(ctx); err != nil {
		err = Error.New(constant.ErrInvalidRequest, "error committing transaction", errors.New("error commit"))
		return response, err
	}

	response = model.CheckBalanceResponse{
		Balance: int64(account.Balance),
		BankNo:  account.BankNo,
	}
	return
}
